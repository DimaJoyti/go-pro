// Package handler provides HTTP handlers for the API Gateway
package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/DimaJoyti/go-pro/services/api-gateway/internal/auth"
	"github.com/DimaJoyti/go-pro/services/api-gateway/internal/config"
	"github.com/DimaJoyti/go-pro/services/api-gateway/internal/proxy"
	"github.com/DimaJoyti/go-pro/services/shared/middleware"
)

// Handler holds the HTTP handlers
type Handler struct {
	config     *config.Config
	router     *proxy.Router
	jwtManager *auth.JWTManager
}

// New creates a new handler
func New(cfg *config.Config) *Handler {
	return &Handler{
		config: cfg,
		router: proxy.NewRouter(
			cfg.Services.UserServiceURL,
			cfg.Services.CourseServiceURL,
			cfg.Services.ProgressServiceURL,
			30*time.Second,
		),
		jwtManager: auth.NewJWTManager(cfg.Auth.JWTSecret, cfg.Auth.TokenDuration),
	}
}

// SetupRoutes sets up the HTTP routes
func (h *Handler) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("GET /health", h.handleHealth)
	mux.HandleFunc("GET /ready", h.handleReady)

	// API routes - these will be proxied to backend services
	// Public routes (no authentication required)
	publicRoutes := http.NewServeMux()
	publicRoutes.HandleFunc("/api/v1/auth/", h.router.Route)
	publicRoutes.HandleFunc("/api/v1/courses", h.router.Route) // List courses (public)

	// Protected routes (authentication required)
	protectedRoutes := http.NewServeMux()
	protectedRoutes.HandleFunc("/api/v1/users/", h.router.Route)
	protectedRoutes.HandleFunc("/api/v1/courses/", h.router.Route) // Course details, create, update
	protectedRoutes.HandleFunc("/api/v1/lessons/", h.router.Route)
	protectedRoutes.HandleFunc("/api/v1/progress/", h.router.Route)
	protectedRoutes.HandleFunc("/api/v1/achievements/", h.router.Route)

	// Apply middleware
	mux.Handle("/api/v1/auth/", publicRoutes)
	mux.Handle("/api/v1/courses", h.jwtManager.OptionalAuthMiddleware(publicRoutes))
	mux.Handle("/api/v1/", h.jwtManager.AuthMiddleware(protectedRoutes))

	// Apply global middleware
	handler := middleware.Chain(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recovery,
		middleware.CORS([]string{"*"}),
		middleware.Timeout(30*time.Second),
		middleware.ServiceInfoMiddleware(middleware.ServiceInfo{
			Name:    "api-gateway",
			Version: "1.0.0",
			Env:     "development",
		}),
	)(mux)

	return handler
}

// handleHealth handles health check requests
func (h *Handler) handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"service":   "api-gateway",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// handleReady handles readiness check requests
func (h *Handler) handleReady(w http.ResponseWriter, r *http.Request) {
	// Check if backend services are reachable
	services := map[string]string{
		"user-service":     h.config.Services.UserServiceURL,
		"course-service":   h.config.Services.CourseServiceURL,
		"progress-service": h.config.Services.ProgressServiceURL,
	}

	allReady := true
	serviceStatus := make(map[string]string)

	for name, url := range services {
		if err := h.checkService(url + "/health"); err != nil {
			serviceStatus[name] = "unhealthy"
			allReady = false
		} else {
			serviceStatus[name] = "healthy"
		}
	}

	response := map[string]interface{}{
		"status":    "ready",
		"services":  serviceStatus,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	statusCode := http.StatusOK
	if !allReady {
		response["status"] = "not ready"
		statusCode = http.StatusServiceUnavailable
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// checkService checks if a service is healthy
func (h *Handler) checkService(url string) error {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return http.ErrNotSupported
	}

	return nil
}

// RateLimiter implements simple in-memory rate limiting
type RateLimiter struct {
	requests map[string][]time.Time
	limit    int
	window   time.Duration
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

// Allow checks if a request is allowed
func (rl *RateLimiter) Allow(key string) bool {
	now := time.Now()

	// Get existing requests for this key
	requests, exists := rl.requests[key]
	if !exists {
		rl.requests[key] = []time.Time{now}
		return true
	}

	// Remove old requests outside the window
	var validRequests []time.Time
	for _, req := range requests {
		if now.Sub(req) < rl.window {
			validRequests = append(validRequests, req)
		}
	}

	// Check if limit is exceeded
	if len(validRequests) >= rl.limit {
		return false
	}

	// Add new request
	validRequests = append(validRequests, now)
	rl.requests[key] = validRequests

	return true
}

// Cleanup removes old entries from the rate limiter
func (rl *RateLimiter) Cleanup() {
	now := time.Now()
	for key, requests := range rl.requests {
		var validRequests []time.Time
		for _, req := range requests {
			if now.Sub(req) < rl.window {
				validRequests = append(validRequests, req)
			}
		}
		if len(validRequests) == 0 {
			delete(rl.requests, key)
		} else {
			rl.requests[key] = validRequests
		}
	}
}

