package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/weather-cli/internal/api"
	"github.com/DimaJoyti/go-pro/basic/projects/weather-cli/internal/cache"
	"github.com/DimaJoyti/go-pro/basic/projects/weather-cli/internal/formatter"
	"github.com/DimaJoyti/go-pro/basic/projects/weather-cli/pkg/weather"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "current":
		handleCurrent()
	case "forecast":
		handleForecast()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func handleCurrent() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: WEATHER_API_KEY environment variable not set")
		fmt.Println("Get your free API key from: https://openweathermap.org/api")
		os.Exit(1)
	}

	city := getFlag("--city", "London")
	format := getFlag("--format", "table")
	detailed := hasFlag("--detailed")

	// Initialize cache
	cacheDir := os.ExpandEnv("$HOME/.weather-cli/cache")
	weatherCache, err := cache.NewCache(cacheDir, 5*time.Minute)
	if err != nil {
		fmt.Printf("Warning: Failed to initialize cache: %v\n", err)
	}

	// Check cache first
	cacheKey := fmt.Sprintf("current_%s", city)
	if weatherCache != nil {
		if data, found := weatherCache.Get(cacheKey); found {
			displayWeather(data, format, detailed, false)
			return
		}
	}

	// Fetch from API
	client := api.NewOpenWeatherClient(apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := client.GetCurrent(ctx, city, weather.Imperial)
	if err != nil {
		fmt.Printf("Error fetching weather: %v\n", err)
		os.Exit(1)
	}

	// Cache the result
	if weatherCache != nil {
		weatherCache.Set(cacheKey, data)
	}

	displayWeather(data, format, detailed, false)
}

func handleForecast() {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		fmt.Println("Error: WEATHER_API_KEY environment variable not set")
		os.Exit(1)
	}

	city := getFlag("--city", "London")
	format := getFlag("--format", "table")

	client := api.NewOpenWeatherClient(apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	data, err := client.GetForecast(ctx, city, 5, weather.Imperial)
	if err != nil {
		fmt.Printf("Error fetching forecast: %v\n", err)
		os.Exit(1)
	}

	displayWeather(data, format, false, true)
}

func displayWeather(data *weather.WeatherData, format string, detailed, isForecast bool) {
	switch format {
	case "json":
		jsonData, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(jsonData))
	case "table":
		formatter := formatter.NewTableFormatter(true)
		if isForecast {
			fmt.Println(formatter.FormatForecast(data))
		} else {
			fmt.Println(formatter.FormatCurrent(data, detailed))
		}
	default:
		fmt.Printf("Unknown format: %s\n", format)
	}
}

func getFlag(flag, defaultValue string) string {
	for i, arg := range os.Args {
		if arg == flag && i+1 < len(os.Args) {
			return os.Args[i+1]
		}
	}
	return defaultValue
}

func hasFlag(flag string) bool {
	for _, arg := range os.Args {
		if arg == flag {
			return true
		}
	}
	return false
}

func printUsage() {
	fmt.Println("Weather CLI - Beautiful weather forecasts in your terminal")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  weather <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  current     Get current weather")
	fmt.Println("  forecast    Get weather forecast")
	fmt.Println("  help        Show this help message")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --city <name>      City name (default: London)")
	fmt.Println("  --format <type>    Output format: table, json (default: table)")
	fmt.Println("  --detailed         Show detailed information")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  weather current --city \"New York\"")
	fmt.Println("  weather current --city \"Tokyo\" --detailed")
	fmt.Println("  weather forecast --city \"Paris\"")
	fmt.Println("  weather current --city \"London\" --format json")
	fmt.Println()
	fmt.Println("Environment Variables:")
	fmt.Println("  WEATHER_API_KEY    Your OpenWeatherMap API key (required)")
	fmt.Println()
	fmt.Println("Get your free API key: https://openweathermap.org/api")
}
