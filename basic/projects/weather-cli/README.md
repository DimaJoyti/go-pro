# ☀️ Weather CLI Application

A beautiful command-line weather application built with Go, featuring colorful terminal output, multiple data sources, and intelligent caching.

## 📋 Project Overview

This project demonstrates building a production-ready CLI application in Go with:
- Integration with external weather APIs
- Beautiful terminal UI with colors and formatting
- Intelligent caching for performance
- Multiple output formats (table, JSON, chart)
- Location-based forecasts
- Historical data tracking
- Error handling and retry logic

## 🎯 Learning Objectives

By completing this project, you will learn:

- **CLI Development**: Build user-friendly command-line tools
- **API Integration**: Work with external REST APIs
- **Terminal UI**: Create colorful, formatted terminal output
- **Caching**: Implement efficient data caching
- **Configuration**: Manage API keys and settings
- **Error Handling**: Robust error handling and retries
- **Testing**: Test CLI applications effectively

## 🏗️ Architecture

```
┌──────────────┐
│     CLI      │
│   (Cobra)    │
└──────┬───────┘
       │
       ▼
┌──────────────────────────────┐
│    Weather Service           │
│  (Business Logic)            │
└──────┬───────────────────────┘
       │
       ├─────────────┬──────────────┐
       ▼             ▼              ▼
┌──────────┐  ┌──────────┐  ┌──────────┐
│   API    │  │  Cache   │  │Formatter │
│ Client   │  │  Layer   │  │  (UI)    │
└──────────┘  └──────────┘  └──────────┘
```

## 🚀 Features

### Core Features
- ✅ Current weather conditions
- ✅ 5-day forecast
- ✅ Hourly forecast
- ✅ Location search
- ✅ Multiple units (metric/imperial)
- ✅ Beautiful terminal output

### Display Options
- ✅ Colorful ASCII art weather icons
- ✅ Table format
- ✅ JSON output
- ✅ Chart visualization
- ✅ Detailed/summary views

### Advanced Features
- ✅ Intelligent caching (TTL-based)
- ✅ Multiple weather providers
- ✅ Offline mode with cached data
- ✅ Configuration file support
- ✅ Historical data tracking

## 📁 Project Structure

```
weather-cli/
├── cmd/
│   └── main.go                 # CLI entry point
├── internal/
│   ├── api/
│   │   ├── client.go           # API client interface
│   │   ├── openweather.go      # OpenWeather implementation
│   │   └── weatherapi.go       # WeatherAPI implementation
│   ├── formatter/
│   │   ├── table.go            # Table formatter
│   │   ├── json.go             # JSON formatter
│   │   └── chart.go            # Chart formatter
│   └── cache/
│       └── cache.go            # Caching layer
├── pkg/
│   └── weather/
│       └── models.go           # Weather data models
├── configs/
│   └── config.yaml             # Configuration file
├── docs/
│   ├── API.md                  # API documentation
│   └── USAGE.md                # Usage examples
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 🔧 Installation

### Prerequisites
- Go 1.21 or higher
- Weather API key (free from [OpenWeatherMap](https://openweathermap.org/api))

### Quick Start

```bash
# Clone and navigate
cd basic/projects/weather-cli

# Install dependencies
go mod download

# Build the application
go build -o weather ./cmd/main.go

# Set your API key
export WEATHER_API_KEY="your-api-key-here"

# Run the application
./weather current --city "London"
```

## 📖 Usage

### Get Current Weather

```bash
# By city name
weather current --city "New York"

# By coordinates
weather current --lat 40.7128 --lon -74.0060

# With specific units
weather current --city "Tokyo" --units metric

# Detailed output
weather current --city "Paris" --detailed
```

### Get Forecast

```bash
# 5-day forecast
weather forecast --city "London"

# Hourly forecast
weather forecast --city "Berlin" --hourly

# With chart visualization
weather forecast --city "Sydney" --chart
```

### Output Formats

```bash
# Table format (default)
weather current --city "Moscow"

# JSON format
weather current --city "Dubai" --format json

# Compact format
weather current --city "Singapore" --compact
```

### Configuration

```bash
# Set default city
weather config set city "San Francisco"

# Set default units
weather config set units imperial

# View configuration
weather config show
```

## 🎨 Example Output

### Current Weather (Table Format)

```
╔══════════════════════════════════════════════════════════╗
║              Weather in New York, US                     ║
╠══════════════════════════════════════════════════════════╣
║                                                          ║
║                    ☀️  Clear Sky                         ║
║                                                          ║
║  Temperature:     72°F (22°C)                           ║
║  Feels Like:      70°F (21°C)                           ║
║  Humidity:        65%                                    ║
║  Wind Speed:      8 mph (13 km/h) NW                    ║
║  Pressure:        1013 hPa                              ║
║  Visibility:      10 km                                  ║
║  UV Index:        6 (High)                              ║
║                                                          ║
║  Sunrise:         06:24 AM                              ║
║  Sunset:          07:45 PM                              ║
║                                                          ║
║  Last Updated:    2024-01-15 14:30:00                   ║
╚══════════════════════════════════════════════════════════╝
```

### 5-Day Forecast

```
┌─────────────────────────────────────────────────────────┐
│              5-Day Forecast for London                  │
├─────────────────────────────────────────────────────────┤
│                                                         │
│  Mon 15 Jan    ☀️   18°C / 12°C    Clear               │
│  Tue 16 Jan    ⛅   16°C / 10°C    Partly Cloudy       │
│  Wed 17 Jan    🌧️   14°C / 9°C     Rain                │
│  Thu 18 Jan    ☁️   15°C / 11°C    Cloudy              │
│  Fri 19 Jan    ☀️   17°C / 13°C    Clear               │
│                                                         │
└─────────────────────────────────────────────────────────┘
```

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Test specific package
go test -v ./internal/api/

# Run integration tests (requires API key)
WEATHER_API_KEY=your-key go test -tags=integration ./...
```

## 🔐 Configuration

### Environment Variables

```bash
WEATHER_API_KEY=your-api-key-here
WEATHER_CACHE_TTL=300              # Cache TTL in seconds
WEATHER_DEFAULT_CITY=London
WEATHER_DEFAULT_UNITS=metric       # metric or imperial
```

### Config File (`~/.weather-cli/config.yaml`)

```yaml
api:
  key: your-api-key-here
  provider: openweather
  timeout: 10s

cache:
  enabled: true
  ttl: 300s
  directory: ~/.weather-cli/cache

defaults:
  city: London
  units: metric
  format: table

display:
  colors: true
  icons: true
  detailed: false
```

## 📊 API Providers

### OpenWeatherMap (Default)

```bash
# Sign up: https://openweathermap.org/api
# Free tier: 1,000 calls/day
export WEATHER_API_KEY="your-openweather-key"
```

### WeatherAPI.com

```bash
# Sign up: https://www.weatherapi.com/
# Free tier: 1,000,000 calls/month
weather config set provider weatherapi
export WEATHER_API_KEY="your-weatherapi-key"
```

## 🎓 Learning Path

1. **Start Here**: Read `cmd/main.go` to understand CLI structure
2. **API Integration**: Study `internal/api/client.go`
3. **Data Models**: Review `pkg/weather/models.go`
4. **Formatting**: Examine `internal/formatter/`
5. **Caching**: Understand `internal/cache/cache.go`
6. **Testing**: Look at test files

## 🚀 Advanced Usage

### Scripting

```bash
#!/bin/bash
# Get weather and send notification
TEMP=$(weather current --city "London" --format json | jq '.temperature')
if [ $TEMP -gt 25 ]; then
    notify-send "Hot day!" "Temperature is ${TEMP}°C"
fi
```

### Cron Jobs

```bash
# Check weather every hour
0 * * * * /usr/local/bin/weather current --city "NYC" >> /var/log/weather.log
```

### Integration

```go
package main

import "github.com/DimaJoyti/go-pro/basic/projects/weather-cli/pkg/weather"

func main() {
    client := weather.NewClient("your-api-key")
    data, _ := client.GetCurrent("London")
    fmt.Printf("Temperature: %.1f°C\n", data.Temperature)
}
```

## 📚 Additional Resources

- [Usage Guide](docs/USAGE.md)
- [API Documentation](docs/API.md)
- [Cobra CLI Framework](https://github.com/spf13/cobra)
- [OpenWeatherMap API](https://openweathermap.org/api)

## 🎯 Next Steps

After completing this project, try:
1. Add weather alerts and warnings
2. Implement location autocomplete
3. Add weather maps (ASCII art)
4. Create a TUI (Terminal UI) version
5. Add air quality data
6. Implement weather comparisons
7. Add historical weather data
8. Create weather-based recommendations

---

**Happy Coding! ☀️🌧️⛅**

