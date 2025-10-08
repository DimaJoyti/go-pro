package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/handlers"
	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/repository"
	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/service"
)

const (
	defaultPort    = "8080"
	defaultBaseURL = "http://localhost:8080"
)

func main() {
	// Get configuration from environment
	port := getEnv("PORT", defaultPort)
	baseURL := getEnv("BASE_URL", defaultBaseURL)

	// Initialize repository
	repo := repository.NewMemoryRepository()
	log.Println("✓ Initialized in-memory repository")

	// Initialize service
	urlService := service.NewURLService(repo, baseURL)
	log.Println("✓ Initialized URL service")

	// Initialize handlers
	handler := handlers.NewURLHandler(urlService)
	log.Println("✓ Initialized HTTP handlers")

	// Setup routes
	mux := http.NewServeMux()
	handler.SetupRoutes(mux)
	log.Println("✓ Configured routes")

	// Create server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      loggingMiddleware(corsMiddleware(mux)),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("🚀 Server starting on http://localhost:%s", port)
		log.Printf("📝 Base URL: %s", baseURL)
		log.Println("📚 API Endpoints:")
		log.Println("   POST   /api/shorten     - Shorten a URL")
		log.Println("   GET    /:code           - Redirect to original URL")
		log.Println("   GET    /api/stats/:code - Get URL statistics")
		log.Println("   GET    /api/urls        - List all URLs")
		log.Println("   GET    /health          - Health check")
		log.Println()
		log.Println("Press Ctrl+C to stop the server")
		log.Println(repeatString("=", 60))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("\n🛑 Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Close service
	if err := urlService.Close(); err != nil {
		log.Printf("Error closing service: %v", err)
	}

	log.Println("✓ Server stopped gracefully")
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a response writer wrapper to capture status code
		wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)
		log.Printf(
			"%s %s %d %v",
			r.Method,
			r.URL.Path,
			wrapped.statusCode,
			duration,
		)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// corsMiddleware adds CORS headers
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// repeatString repeats a string n times
func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
