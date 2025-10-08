package repository

import (
	"context"
	"sync"

	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/domain"
)

// MemoryRepository implements URLRepository using in-memory storage
type MemoryRepository struct {
	urls map[string]*domain.URL
	mu   sync.RWMutex
}

// NewMemoryRepository creates a new in-memory repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		urls: make(map[string]*domain.URL),
	}
}

// Save stores a new URL
func (r *MemoryRepository) Save(ctx context.Context, url *domain.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.urls[url.ShortCode]; exists {
		return domain.ErrCodeExists
	}

	r.urls[url.ShortCode] = url
	return nil
}

// FindByCode retrieves a URL by its short code
func (r *MemoryRepository) FindByCode(ctx context.Context, code string) (*domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, exists := r.urls[code]
	if !exists {
		return nil, domain.ErrURLNotFound
	}

	return url, nil
}

// Update updates an existing URL
func (r *MemoryRepository) Update(ctx context.Context, url *domain.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.urls[url.ShortCode]; !exists {
		return domain.ErrURLNotFound
	}

	r.urls[url.ShortCode] = url
	return nil
}

// Delete removes a URL by its short code
func (r *MemoryRepository) Delete(ctx context.Context, code string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.urls[code]; !exists {
		return domain.ErrURLNotFound
	}

	delete(r.urls, code)
	return nil
}

// Exists checks if a short code exists
func (r *MemoryRepository) Exists(ctx context.Context, code string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, exists := r.urls[code]
	return exists, nil
}

// GetAll retrieves all URLs
func (r *MemoryRepository) GetAll(ctx context.Context) ([]*domain.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	urls := make([]*domain.URL, 0, len(r.urls))
	for _, url := range r.urls {
		urls = append(urls, url)
	}

	return urls, nil
}

// GetStats retrieves statistics for a URL
func (r *MemoryRepository) GetStats(ctx context.Context, code string) (*domain.StatsResponse, error) {
	url, err := r.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	return url.ToStatsResponse(), nil
}

// Close closes the repository (no-op for memory)
func (r *MemoryRepository) Close() error {
	return nil
}

// Count returns the number of URLs stored
func (r *MemoryRepository) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.urls)
}

// Clear removes all URLs (useful for testing)
func (r *MemoryRepository) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.urls = make(map[string]*domain.URL)
}
