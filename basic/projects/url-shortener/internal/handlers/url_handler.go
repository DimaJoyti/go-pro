package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/domain"
	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/service"
)

// URLHandler handles HTTP requests for URL operations
type URLHandler struct {
	service *service.URLService
}

// NewURLHandler creates a new URL handler
func NewURLHandler(service *service.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

// ShortenURL handles POST /api/shorten
func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.respondError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response, err := h.service.ShortenURL(r.Context(), &req)
	if err != nil {
		switch err {
		case domain.ErrInvalidURL, domain.ErrEmptyURL, domain.ErrCodeTooShort, domain.ErrCodeTooLong, domain.ErrInvalidCharacters:
			h.respondError(w, err.Error(), http.StatusBadRequest)
		case domain.ErrCodeExists:
			h.respondError(w, "Short code already exists", http.StatusConflict)
		default:
			h.respondError(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	h.respondJSON(w, response, http.StatusCreated)
}

// Redirect handles GET /:code
func (h *URLHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	// Extract short code from path
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" || code == "/" {
		h.respondError(w, "Short code required", http.StatusBadRequest)
		return
	}

	// Get analytics data
	referrer := r.Header.Get("Referer")
	userAgent := r.Header.Get("User-Agent")
	ipAddress := getIPAddress(r)

	// Get original URL
	originalURL, err := h.service.GetOriginalURL(r.Context(), code, referrer, userAgent, ipAddress)
	if err != nil {
		if err == domain.ErrURLNotFound {
			h.respondError(w, "Short URL not found", http.StatusNotFound)
		} else {
			h.respondError(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Redirect to original URL
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}

// GetStats handles GET /api/stats/:code
func (h *URLHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.respondError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract code from path
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/api/stats/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		h.respondError(w, "Short code required", http.StatusBadRequest)
		return
	}
	code := parts[0]

	stats, err := h.service.GetStats(r.Context(), code)
	if err != nil {
		if err == domain.ErrURLNotFound {
			h.respondError(w, "Short URL not found", http.StatusNotFound)
		} else {
			h.respondError(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	h.respondJSON(w, stats, http.StatusOK)
}

// GetAllURLs handles GET /api/urls (admin endpoint)
func (h *URLHandler) GetAllURLs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.respondError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	urls, err := h.service.GetAllURLs(r.Context())
	if err != nil {
		h.respondError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.respondJSON(w, urls, http.StatusOK)
}

// Health handles GET /health
func (h *URLHandler) Health(w http.ResponseWriter, r *http.Request) {
	response := domain.HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}
	h.respondJSON(w, response, http.StatusOK)
}

// respondJSON sends a JSON response
func (h *URLHandler) respondJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}

// respondError sends an error response
func (h *URLHandler) respondError(w http.ResponseWriter, message string, status int) {
	response := domain.ErrorResponse{
		Error:   http.StatusText(status),
		Message: message,
		Code:    status,
	}
	h.respondJSON(w, response, status)
}

// getIPAddress extracts the client IP address from the request
func getIPAddress(r *http.Request) string {
	// Check X-Forwarded-For header first
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Take the first IP if multiple are present
		ips := strings.Split(forwarded, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP header
	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	// Fall back to RemoteAddr
	return r.RemoteAddr
}

// SetupRoutes configures all HTTP routes
func (h *URLHandler) SetupRoutes(mux *http.ServeMux) {
	// API routes
	mux.HandleFunc("/api/shorten", h.ShortenURL)
	mux.HandleFunc("/api/stats/", h.GetStats)
	mux.HandleFunc("/api/urls", h.GetAllURLs)
	mux.HandleFunc("/health", h.Health)

	// Redirect route (catch-all)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Skip API routes
		if strings.HasPrefix(r.URL.Path, "/api/") || r.URL.Path == "/health" {
			http.NotFound(w, r)
			return
		}
		h.Redirect(w, r)
	})
}
