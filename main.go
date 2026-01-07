package main

import (
	"fmt"
	"net/http"
	handler "weather-api/handlers"

	"weather-api/database"

	"github.com/joho/godotenv"
)

func main() {

	database.ConnectRedis()
	fmt.Println("Redis connected...")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("cannot load env")
	}
	fmt.Println("env loaded..")

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.GetWelcome)
	mux.HandleFunc("/api/v1/weather/{city}", handler.GetCityWeather)

	fmt.Println("server run on http://localhost:8383")
	http.ListenAndServe(":8383", mux)
	//server run on http://localhost:8383

}
