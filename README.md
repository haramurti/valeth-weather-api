

# 🌤️ Valeth Weather API

![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-blue) ![Redis](https://img.shields.io/badge/redis-v9-red) ![License](https://img.shields.io/badge/license-MIT-green) ![Status](https://img.shields.io/badge/status-Native%20Mode-orange)

> "Built with sweat, tears, and pure `net/http`. No frameworks, because life is too short for shortcuts."

## 🧐 What is this?

This is a robust RESTful API to check real-time weather in any city around the globe.

This project is built **WITHOUT ANY FRAMEWORKS** (No Fiber, No Gin, No Echo). It utilizes Go's Standard Library (`net/http`) and the new routing capabilities of Go 1.22+.

**Why?** Because I wanted to understand how HTTP requests work under the hood, how to manually parse JSON, and how to structure a scalable application before jumping into "magic" frameworks.

## ✨ Key Features

* **⚡ Native Performance:** Extremely lightweight without the overhead of heavy dependencies.
* **🧠 Smart Caching (Redis):** Implements a cache-first strategy. We check Redis (local cache) before hitting the external API. This saves costs and reduces latency to sub-milliseconds.
* **🛡️ Error Handling:** Graceful error responses (JSON) instead of server panics.
* **🕵️ Transparency:** Response headers (`X-Cache-Source`) indicate whether data was served from `Redis` or `VisualCrossingAPI`.

## 🛠️ Tech Stack

* **[Go (Golang)](https://go.dev/)**: Main language (Version 1.22+ required for routing).
* **[Redis](https://redis.io/)**: In-memory data store for caching.
* **[Visual Crossing API](https://www.visualcrossing.com/)**: Third-party weather data provider.
* **Standard Library**: `net/http`, `encoding/json`, `context`.

## 🚀 Installation & Setup

It's easier than it looks. Follow these steps:

### 1. Prerequisites
Ensure you have the following installed:
* Go (Min version 1.22)
* Redis (Running locally or via Docker/OrbStack)

### 2. Clone Repository
```bash
git clone https://github.com/YOUR_USERNAME/weather-api.git
cd weather-api
```

### 3. Environment Variables

Create a `.env` file in the root directory and add your API Key:

```env
WEATHER_API_KEY=YOUR_VISUAL_CROSSING_API_KEY

```

### 4. Install Dependencies & Run

Tidy up the modules (downloads Redis driver & Godotenv):

```bash
go mod tidy

```

Start the server:

```bash
go run main.go

```

If you see `Server run on http://localhost:8383` and `Redis connected...`, you are good to go!

## 🎮 API Usage

You can use **Postman**, **Insomnia**, or your browser.

**Endpoint:**
`GET http://localhost:8383/api/v1/weather/{city_name}`

**Examples:**

```
GET http://localhost:8383/api/v1/weather/jakarta
GET http://localhost:8383/api/v1/weather/london

```

### 🎩 Witness the Magic (Caching Strategy)

1. **First Hit:** Check the `Time` (approx. 500ms - 1s). The Header `X-Cache-Source` will be `VisualCrossingAPI`.
2. **Second Hit:** Check the `Time` again. **BOOM!** It drops to 5ms - 20ms. The Header `X-Cache-Source` becomes `Redis`.

## 📂 Project Structure

```
.
├── database/
│   └── redis.go       # Redis connection logic
├── handlers/
│   └── controller.go  # Business logic (Check Redis -> Fetch API -> Response)
├── .env               # Secrets (Ignored by Git)
├── go.mod             # Dependencies
├── main.go            # Entry point & Routing
└── README.md          # Documentation

```

## 🤝 Contribution

Found a bug? Or found a way to make the code cleaner? Feel free to open a **Pull Request**. But remember, keep it Native! No frameworks allowed in this branch! 😛

## 📝 License

[MIT](https://choosealicense.com/licenses/mit/)

