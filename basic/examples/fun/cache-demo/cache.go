package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Task: Implement a thread-safe in-memory cache with expiration support.
// The cache should support Get, Set, Delete operations and automatic cleanup of expired items.

// CacheItem represents a single cache entry
type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

// IsExpired checks if the cache item has expired
func (item *CacheItem) IsExpired() bool {
	return time.Now().After(item.Expiration)
}

// Cache is a thread-safe in-memory cache with expiration
type Cache struct {
	items map[string]*CacheItem
	mu    sync.RWMutex
	ttl   time.Duration
}

// NewCache creates a new cache with default TTL
func NewCache(ttl time.Duration) *Cache {
	cache := &Cache{
		items: make(map[string]*CacheItem),
		ttl:   ttl,
	}

	// Start cleanup goroutine
	go cache.cleanupExpired()

	return cache
}

// Set adds or updates an item in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.SetWithTTL(key, value, c.ttl)
}

// SetWithTTL adds or updates an item with custom TTL
func (c *Cache) SetWithTTL(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = &CacheItem{
		Value:      value,
		Expiration: time.Now().Add(ttl),
	}
}

// Get retrieves an item from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	// Check if expired
	if item.IsExpired() {
		return nil, false
	}

	return item.Value, true
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// Clear removes all items from the cache
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*CacheItem)
}

// Size returns the number of items in the cache
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return len(c.items)
}

// cleanupExpired periodically removes expired items
func (c *Cache) cleanupExpired() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.items {
			if item.IsExpired() {
				delete(c.items, key)
			}
		}
		c.mu.Unlock()
	}
}

// GetAll returns all non-expired items
func (c *Cache) GetAll() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := make(map[string]interface{})
	for key, item := range c.items {
		if !item.IsExpired() {
			result[key] = item.Value
		}
	}

	return result
}

// User represents a user object for demo
type User struct {
	ID   int
	Name string
}

func main() {
	fmt.Println("In-Memory Cache Demo")
	fmt.Println(strings.Repeat("=", 60))

	// Create cache with 3 second default TTL
	cache := NewCache(3 * time.Second)

	// Demo 1: Basic Set and Get
	fmt.Println("\n1. Basic Set and Get Operations")
	fmt.Println(strings.Repeat("-", 60))

	cache.Set("name", "Alice")
	cache.Set("age", 30)
	cache.Set("city", "New York")

	if value, found := cache.Get("name"); found {
		fmt.Printf("Found 'name': %v\n", value)
	}

	if value, found := cache.Get("age"); found {
		fmt.Printf("Found 'age': %v\n", value)
	}

	fmt.Printf("Cache size: %d items\n", cache.Size())

	// Demo 2: Storing complex objects
	fmt.Println("\n2. Storing Complex Objects")
	fmt.Println(strings.Repeat("-", 60))

	user := User{ID: 1, Name: "Bob Smith"}
	cache.Set("user:1", user)

	if value, found := cache.Get("user:1"); found {
		if u, ok := value.(User); ok {
			fmt.Printf("Found user: ID=%d, Name=%s\n", u.ID, u.Name)
		}
	}

	// Demo 3: Custom TTL
	fmt.Println("\n3. Custom TTL (1 second)")
	fmt.Println(strings.Repeat("-", 60))

	cache.SetWithTTL("temp", "This expires in 1 second", 1*time.Second)
	fmt.Println("Set 'temp' with 1 second TTL")

	if value, found := cache.Get("temp"); found {
		fmt.Printf("Immediately after set: %v\n", value)
	}

	time.Sleep(1500 * time.Millisecond)

	if _, found := cache.Get("temp"); !found {
		fmt.Println("After 1.5 seconds: 'temp' has expired")
	}

	// Demo 4: Expiration
	fmt.Println("\n4. Automatic Expiration (3 second default TTL)")
	fmt.Println(strings.Repeat("-", 60))

	cache.Set("expires", "This will expire")
	fmt.Printf("Set 'expires', cache size: %d\n", cache.Size())

	fmt.Println("Waiting 2 seconds...")
	time.Sleep(2 * time.Second)

	if value, found := cache.Get("expires"); found {
		fmt.Printf("After 2 seconds, still found: %v\n", value)
	}

	fmt.Println("Waiting 2 more seconds...")
	time.Sleep(2 * time.Second)

	if _, found := cache.Get("expires"); !found {
		fmt.Println("After 4 seconds total: 'expires' has expired")
	}

	// Demo 5: GetAll
	fmt.Println("\n5. Get All Items")
	fmt.Println(strings.Repeat("-", 60))

	cache.Clear()
	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	allItems := cache.GetAll()
	fmt.Printf("All items in cache (%d):\n", len(allItems))
	for key, value := range allItems {
		fmt.Printf("  %s: %v\n", key, value)
	}

	// Demo 6: Delete
	fmt.Println("\n6. Delete Operation")
	fmt.Println(strings.Repeat("-", 60))

	cache.Delete("key2")
	fmt.Println("Deleted 'key2'")
	fmt.Printf("Cache size: %d items\n", cache.Size())

	if _, found := cache.Get("key2"); !found {
		fmt.Println("'key2' not found (successfully deleted)")
	}

	fmt.Println("\nCache demo completed!")
}
