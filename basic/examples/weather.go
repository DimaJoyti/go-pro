package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
)

const apiKey = "a3b952be51caa80eb0f6c20331b9f332" // Replace with your OpenWeatherMap API key
const baseURL = "https://api.openweathermap.org/data/2.5/weather"

type WeatherResponse struct {
    Name    string    `json:"name"`
    Main    Main      `json:"main"`
    Weather []Weather `json:"weather"`
}

type Main struct {
    Temp     float64 `json:"temp"`
    Pressure float64 `json:"pressure"`
    Humidity float64 `json:"humidity"`
}

type Weather struct {
    Description string `json:"description"`
}

func getWeather(city string) (*WeatherResponse, error) {
    url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", baseURL, city, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error: %s", resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var weatherResponse WeatherResponse
    if err := json.Unmarshal(body, &weatherResponse); err != nil {
        return nil, err
    }

    return &weatherResponse, nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Please provide a city name.")
        return
    }
    city := os.Args[1]

    weather, err := getWeather(city)
    if err != nil {
        fmt.Printf("Error fetching weather: %v\n", err)
        return
    }

    fmt.Printf("Weather in %s:\n", weather.Name)
    fmt.Printf("Temperature: %.2f °C\n", weather.Main.Temp)
    fmt.Printf("Pressure: %.2f hPa\n", weather.Main.Pressure)
    fmt.Printf("Humidity: %.2f%%\n", weather.Main.Humidity)
    if len(weather.Weather) > 0 {
        fmt.Printf("Description: %s\n", weather.Weather[0].Description)
    }
}

// go run weather.go London