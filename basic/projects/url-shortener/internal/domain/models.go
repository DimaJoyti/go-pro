package domain

import (
	"errors"
	"net/url"
	"time"
)

// Errors
var (
	ErrInvalidURL        = errors.New("invalid URL")
	ErrURLNotFound       = errors.New("URL not found")
	ErrCodeExists        = errors.New("short code already exists")
	ErrInvalidCode       = errors.New("invalid short code")
	ErrEmptyURL          = errors.New("URL cannot be empty")
	ErrCodeTooShort      = errors.New("code must be at least 3 characters")
	ErrCodeTooLong       = errors.New("code must be at most 20 characters")
	ErrInvalidCharacters = errors.New("code contains invalid characters")
)

// URL represents a shortened URL with its metadata
type URL struct {
	ShortCode    string    `json:"short_code"`
	OriginalURL  string    `json:"original_url"`
	CreatedAt    time.Time `json:"created_at"`
	LastAccessed time.Time `json:"last_accessed,omitempty"`
	Clicks       int64     `json:"clicks"`
	Analytics    Analytics `json:"analytics,omitempty"`
}

// Analytics contains detailed tracking information
type Analytics struct {
	TotalClicks  int64            `json:"total_clicks"`
	Referrers    map[string]int64 `json:"referrers,omitempty"`
	UserAgents   map[string]int64 `json:"user_agents,omitempty"`
	ClickHistory []Click          `json:"click_history,omitempty"`
}

// Click represents a single click event
type Click struct {
	Timestamp time.Time `json:"timestamp"`
	Referrer  string    `json:"referrer,omitempty"`
	UserAgent string    `json:"user_agent,omitempty"`
	IPAddress string    `json:"ip_address,omitempty"`
}

// ShortenRequest represents a request to shorten a URL
type ShortenRequest struct {
	URL        string `json:"url"`
	CustomCode string `json:"custom_code,omitempty"`
}

// ShortenResponse represents the response after shortening a URL
type ShortenResponse struct {
	ShortCode   string    `json:"short_code"`
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

// StatsResponse represents URL statistics
type StatsResponse struct {
	ShortCode    string    `json:"short_code"`
	OriginalURL  string    `json:"original_url"`
	Clicks       int64     `json:"clicks"`
	CreatedAt    time.Time `json:"created_at"`
	LastAccessed time.Time `json:"last_accessed,omitempty"`
	Analytics    Analytics `json:"analytics,omitempty"`
}

// HealthResponse represents health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// Validate validates the ShortenRequest
func (r *ShortenRequest) Validate() error {
	if r.URL == "" {
		return ErrEmptyURL
	}

	// Validate URL format
	parsedURL, err := url.ParseRequestURI(r.URL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		return ErrInvalidURL
	}

	// Validate custom code if provided
	if r.CustomCode != "" {
		if len(r.CustomCode) < 3 {
			return ErrCodeTooShort
		}
		if len(r.CustomCode) > 20 {
			return ErrCodeTooLong
		}
		if !isValidCode(r.CustomCode) {
			return ErrInvalidCharacters
		}
	}

	return nil
}

// isValidCode checks if a code contains only valid characters
func isValidCode(code string) bool {
	for _, char := range code {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '-' || char == '_') {
			return false
		}
	}
	return true
}

// NewURL creates a new URL instance
func NewURL(shortCode, originalURL string) *URL {
	now := time.Now()
	return &URL{
		ShortCode:   shortCode,
		OriginalURL: originalURL,
		CreatedAt:   now,
		Clicks:      0,
		Analytics: Analytics{
			TotalClicks:  0,
			Referrers:    make(map[string]int64),
			UserAgents:   make(map[string]int64),
			ClickHistory: []Click{},
		},
	}
}

// IncrementClicks increments the click count and updates analytics
func (u *URL) IncrementClicks(referrer, userAgent, ipAddress string) {
	u.Clicks++
	u.LastAccessed = time.Now()
	u.Analytics.TotalClicks++

	// Track referrer
	if referrer != "" {
		u.Analytics.Referrers[referrer]++
	}

	// Track user agent
	if userAgent != "" {
		u.Analytics.UserAgents[userAgent]++
	}

	// Add to click history (limit to last 100 clicks)
	click := Click{
		Timestamp: time.Now(),
		Referrer:  referrer,
		UserAgent: userAgent,
		IPAddress: ipAddress,
	}

	u.Analytics.ClickHistory = append(u.Analytics.ClickHistory, click)
	if len(u.Analytics.ClickHistory) > 100 {
		u.Analytics.ClickHistory = u.Analytics.ClickHistory[1:]
	}
}

// ToStatsResponse converts URL to StatsResponse
func (u *URL) ToStatsResponse() *StatsResponse {
	return &StatsResponse{
		ShortCode:    u.ShortCode,
		OriginalURL:  u.OriginalURL,
		Clicks:       u.Clicks,
		CreatedAt:    u.CreatedAt,
		LastAccessed: u.LastAccessed,
		Analytics:    u.Analytics,
	}
}
