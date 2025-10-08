package repository

import (
	"context"

	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/domain"
)

// URLRepository defines the interface for URL storage
type URLRepository interface {
	// Save stores a new URL
	Save(ctx context.Context, url *domain.URL) error

	// FindByCode retrieves a URL by its short code
	FindByCode(ctx context.Context, code string) (*domain.URL, error)

	// Update updates an existing URL
	Update(ctx context.Context, url *domain.URL) error

	// Delete removes a URL by its short code
	Delete(ctx context.Context, code string) error

	// Exists checks if a short code exists
	Exists(ctx context.Context, code string) (bool, error)

	// GetAll retrieves all URLs (for admin purposes)
	GetAll(ctx context.Context) ([]*domain.URL, error)

	// GetStats retrieves statistics for a URL
	GetStats(ctx context.Context, code string) (*domain.StatsResponse, error)

	// Close closes the repository connection
	Close() error
}
