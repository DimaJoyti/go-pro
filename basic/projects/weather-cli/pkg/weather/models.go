package weather

import "time"

// WeatherData represents current weather information
type WeatherData struct {
	Location    Location   `json:"location"`
	Current     Current    `json:"current"`
	Forecast    []Forecast `json:"forecast,omitempty"`
	LastUpdated time.Time  `json:"last_updated"`
}

// Location represents geographic location
type Location struct {
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone,omitempty"`
}

// Current represents current weather conditions
type Current struct {
	Temperature   float64   `json:"temperature"`
	FeelsLike     float64   `json:"feels_like"`
	Condition     string    `json:"condition"`
	Description   string    `json:"description"`
	Humidity      int       `json:"humidity"`
	Pressure      int       `json:"pressure"`
	WindSpeed     float64   `json:"wind_speed"`
	WindDirection string    `json:"wind_direction"`
	Visibility    int       `json:"visibility"`
	UVIndex       float64   `json:"uv_index"`
	CloudCover    int       `json:"cloud_cover"`
	Precipitation float64   `json:"precipitation,omitempty"`
	Sunrise       time.Time `json:"sunrise,omitempty"`
	Sunset        time.Time `json:"sunset,omitempty"`
}

// Forecast represents weather forecast for a specific time
type Forecast struct {
	Date          time.Time `json:"date"`
	TempMax       float64   `json:"temp_max"`
	TempMin       float64   `json:"temp_min"`
	Condition     string    `json:"condition"`
	Description   string    `json:"description"`
	Humidity      int       `json:"humidity"`
	WindSpeed     float64   `json:"wind_speed"`
	Precipitation float64   `json:"precipitation"`
	ChanceOfRain  int       `json:"chance_of_rain"`
}

// Units represents measurement units
type Units string

const (
	Metric   Units = "metric"
	Imperial Units = "imperial"
)

// WeatherIcon returns an emoji icon for weather condition
func WeatherIcon(condition string) string {
	icons := map[string]string{
		"clear":         "☀️",
		"sunny":         "☀️",
		"partly cloudy": "⛅",
		"cloudy":        "☁️",
		"overcast":      "☁️",
		"rain":          "🌧️",
		"drizzle":       "🌦️",
		"thunderstorm":  "⛈️",
		"snow":          "❄️",
		"mist":          "🌫️",
		"fog":           "🌫️",
		"wind":          "💨",
	}

	for key, icon := range icons {
		if contains(condition, key) {
			return icon
		}
	}
	return "🌡️"
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) &&
		(s == substr ||
			len(s) > len(substr) &&
				(s[:len(substr)] == substr || s[len(s)-len(substr):] == substr))
}
