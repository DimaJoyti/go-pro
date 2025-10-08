package cache

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/weather-cli/pkg/weather"
)

// Cache provides caching functionality for weather data
type Cache struct {
	dir string
	ttl time.Duration
	mu  sync.RWMutex
}

// CacheEntry represents a cached item
type CacheEntry struct {
	Data      *weather.WeatherData `json:"data"`
	ExpiresAt time.Time            `json:"expires_at"`
}

// NewCache creates a new cache instance
func NewCache(dir string, ttl time.Duration) (*Cache, error) {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	return &Cache{
		dir: dir,
		ttl: ttl,
	}, nil
}

// Get retrieves data from cache
func (c *Cache) Get(key string) (*weather.WeatherData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	path := c.getPath(key)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}

	var entry CacheEntry
	if err := json.Unmarshal(data, &entry); err != nil {
		return nil, false
	}

	// Check if expired
	if time.Now().After(entry.ExpiresAt) {
		os.Remove(path)
		return nil, false
	}

	return entry.Data, true
}

// Set stores data in cache
func (c *Cache) Set(key string, data *weather.WeatherData) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := CacheEntry{
		Data:      data,
		ExpiresAt: time.Now().Add(c.ttl),
	}

	jsonData, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	path := c.getPath(key)
	return os.WriteFile(path, jsonData, 0644)
}

// Clear removes all cached data
func (c *Cache) Clear() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return os.RemoveAll(c.dir)
}

// getPath returns the file path for a cache key
func (c *Cache) getPath(key string) string {
	return filepath.Join(c.dir, key+".json")
}
