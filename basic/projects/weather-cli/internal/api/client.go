package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/weather-cli/pkg/weather"
)

// Client defines the weather API client interface
type Client interface {
	GetCurrent(ctx context.Context, city string, units weather.Units) (*weather.WeatherData, error)
	GetForecast(ctx context.Context, city string, days int, units weather.Units) (*weather.WeatherData, error)
	GetCurrentByCoords(ctx context.Context, lat, lon float64, units weather.Units) (*weather.WeatherData, error)
}

// Config holds API client configuration
type Config struct {
	APIKey     string
	BaseURL    string
	Timeout    time.Duration
	MaxRetries int
}

// BaseClient provides common HTTP client functionality
type BaseClient struct {
	config     Config
	httpClient *http.Client
}

// NewBaseClient creates a new base client
func NewBaseClient(config Config) *BaseClient {
	if config.Timeout == 0 {
		config.Timeout = 10 * time.Second
	}
	if config.MaxRetries == 0 {
		config.MaxRetries = 3
	}

	return &BaseClient{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}
}

// doRequest performs an HTTP request with retry logic
func (c *BaseClient) doRequest(ctx context.Context, url string) ([]byte, error) {
	var lastErr error

	for attempt := 0; attempt < c.config.MaxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("creating request: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(time.Second * time.Duration(attempt+1))
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			lastErr = fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
			if resp.StatusCode >= 500 {
				time.Sleep(time.Second * time.Duration(attempt+1))
				continue
			}
			return nil, lastErr
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("reading response: %w", err)
		}

		return body, nil
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

// parseJSON parses JSON response into target struct
func parseJSON(data []byte, target interface{}) error {
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("parsing JSON: %w", err)
	}
	return nil
}
