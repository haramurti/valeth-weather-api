package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"weather-api/database"

	"github.com/redis/go-redis/v9"
)

type VisualCrossingResponse struct {
	Days []struct {
		Temp       float64 `json:"temp"`
		Humidity   float64 `json:"humidity"`
		WindSpeed  float64 `json:"windspeed"`
		UVIndex    float64 `json:"uvindex"`
		Conditions string  `json:"conditions"`
		Datetime   string  `json:"datetime"`
		Sunrise    string  `json:"sunrise"`
		Sunset     string  `json:"sunset"`
	} `json:"days"`
}

type Weather struct {
	ID        int     `json:"id"`
	City      string  `json:"city"`
	Weather   string  `json:"weather"`
	Time      string  `json:"time"`
	AvgTemp   float64 `json:"temp_celcius"`
	Humidity  int     `json:"humidity"`
	WindSpeed float64 `json:"wind_speed"`
	UVindex   int     `json:"uv_index"`
	Sunrise   string  `json:"sunrise"`
	Sunset    string  `json:"sunset"`
}

func GetWelcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to weather api"))
}

func GetCityWeather(w http.ResponseWriter, r *http.Request) {

	cityParam := r.PathValue("city")

	RedisKey := "weather" + cityParam
	val, err := database.Rdb.Get(database.Ctx, RedisKey).Result()

	if err == nil {
		fmt.Println("cahce hit : getting data from redis for citydata : " + cityParam)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Cache-Source", "Redis")
		w.Write([]byte(val))
		return
	} else if err != redis.Nil {
		// Kalau  BUKAN karena kosong (misal Redis mati), laporin aja tapi lanjut
		fmt.Println("Error Redis:", err)
	}

	fmt.Println("Cache MISS: shooting to API : for city data " + cityParam)

	apiKey := os.Getenv("WEATHER_API_KEY")
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", cityParam, apiKey)
	fmt.Println("shooting url :" + url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "failed to shoot url : "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		http.Error(w, "error fetching url api", resp.StatusCode)
		return
	}

	var vcResponse VisualCrossingResponse
	if err := json.NewDecoder(resp.Body).Decode(&vcResponse); err != nil {
		http.Error(w, "failed to read weather", http.StatusInternalServerError)
		return
	}
	today := vcResponse.Days[0]

	finalData := Weather{
		ID:        1,
		City:      cityParam,
		Weather:   today.Conditions,
		Time:      today.Datetime,
		AvgTemp:   today.Temp,
		Humidity:  int(today.Humidity),
		WindSpeed: today.WindSpeed,
		UVindex:   int(today.UVIndex),
		Sunrise:   today.Sunrise,
		Sunset:    today.Sunset,
	}

	jsonData, err := json.Marshal(finalData)
	if err == nil {
		err := database.Rdb.Set(database.Ctx, RedisKey, jsonData, 12*time.Hour).Err()
		if err != nil {
			fmt.Println("failed to save to redis")
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Cache-Source", "VisualCrossingAPI")
	w.Write(jsonData)

}
