package service

import (
	"context"
	"log"

	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/domain"
	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/repository"
	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/pkg/shortener"
)

// URLService handles business logic for URL shortening
type URLService struct {
	repo      repository.URLRepository
	generator *shortener.Generator
	baseURL   string
}

// NewURLService creates a new URL service
func NewURLService(repo repository.URLRepository, baseURL string) *URLService {
	return &URLService{
		repo:      repo,
		generator: shortener.NewGenerator(6),
		baseURL:   baseURL,
	}
}

// ShortenURL creates a shortened URL
func (s *URLService) ShortenURL(ctx context.Context, req *domain.ShortenRequest) (*domain.ShortenResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var shortCode string
	var err error

	// Use custom code if provided, otherwise generate one
	if req.CustomCode != "" {
		shortCode = req.CustomCode

		// Check if custom code already exists
		exists, err := s.repo.Exists(ctx, shortCode)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, domain.ErrCodeExists
		}
	} else {
		// Generate a unique short code
		shortCode, err = s.generateUniqueCode(ctx)
		if err != nil {
			return nil, err
		}
	}

	// Create URL entity
	url := domain.NewURL(shortCode, req.URL)

	// Save to repository
	if err := s.repo.Save(ctx, url); err != nil {
		return nil, err
	}

	// Build response
	response := &domain.ShortenResponse{
		ShortCode:   shortCode,
		ShortURL:    s.buildShortURL(shortCode),
		OriginalURL: req.URL,
		CreatedAt:   url.CreatedAt,
	}

	log.Printf("Created short URL: %s -> %s", shortCode, req.URL)
	return response, nil
}

// GetOriginalURL retrieves the original URL and tracks analytics
func (s *URLService) GetOriginalURL(ctx context.Context, code string, referrer, userAgent, ipAddress string) (string, error) {
	// Find URL
	url, err := s.repo.FindByCode(ctx, code)
	if err != nil {
		return "", err
	}

	// Increment clicks and update analytics
	url.IncrementClicks(referrer, userAgent, ipAddress)

	// Update in repository
	if err := s.repo.Update(ctx, url); err != nil {
		log.Printf("Failed to update analytics for %s: %v", code, err)
		// Don't fail the redirect if analytics update fails
	}

	return url.OriginalURL, nil
}

// GetStats retrieves statistics for a shortened URL
func (s *URLService) GetStats(ctx context.Context, code string) (*domain.StatsResponse, error) {
	return s.repo.GetStats(ctx, code)
}

// DeleteURL removes a shortened URL
func (s *URLService) DeleteURL(ctx context.Context, code string) error {
	return s.repo.Delete(ctx, code)
}

// GetAllURLs retrieves all shortened URLs (admin function)
func (s *URLService) GetAllURLs(ctx context.Context) ([]*domain.URL, error) {
	return s.repo.GetAll(ctx)
}

// generateUniqueCode generates a unique short code
func (s *URLService) generateUniqueCode(ctx context.Context) (string, error) {
	maxAttempts := 10

	for i := 0; i < maxAttempts; i++ {
		code, err := s.generator.Generate()
		if err != nil {
			return "", err
		}

		// Check if code already exists
		exists, err := s.repo.Exists(ctx, code)
		if err != nil {
			return "", err
		}

		if !exists {
			return code, nil
		}
	}

	// If we couldn't generate a unique code after max attempts,
	// try with a longer code
	return s.generator.GenerateWithLength(8)
}

// buildShortURL constructs the full short URL
func (s *URLService) buildShortURL(code string) string {
	return s.baseURL + "/" + code
}

// Close closes the service and its dependencies
func (s *URLService) Close() error {
	return s.repo.Close()
}
