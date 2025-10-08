package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/api/handlers"
	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/api/middleware"
	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/auth"
	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/repository"
	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration from environment
	dbURL := getEnv("DATABASE_URL", "postgres://localhost/blogdb?sslmode=disable")
	jwtSecret := getEnv("JWT_SECRET", "your-secret-key-change-in-production")
	port := getEnv("PORT", "8080")

	// Connect to database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("✓ Connected to database")

	// Initialize JWT manager
	jwtManager := auth.NewJWTManager(jwtSecret, 24*time.Hour)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	postRepo := repository.NewPostRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, jwtManager)
	postService := service.NewPostService(postRepo, userRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	postHandler := handlers.NewPostHandler(postService)

	// Setup router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/health", healthHandler).Methods("GET")
	router.HandleFunc("/api/auth/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/api/auth/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/api/auth/refresh", authHandler.RefreshToken).Methods("POST")

	// Public post routes
	router.HandleFunc("/api/posts", postHandler.ListPosts).Methods("GET")
	router.HandleFunc("/api/posts/{id}", postHandler.GetPost).Methods("GET")
	router.HandleFunc("/api/posts/slug/{slug}", postHandler.GetPostBySlug).Methods("GET")

	// Protected routes (require authentication)
	authRouter := router.PathPrefix("/api").Subrouter()
	authRouter.Use(middleware.AuthMiddleware(jwtManager))

	// Auth routes
	authRouter.HandleFunc("/auth/change-password", authHandler.ChangePassword).Methods("POST")

	// Post routes (authenticated)
	authRouter.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	authRouter.HandleFunc("/posts/{id}", postHandler.UpdatePost).Methods("PUT")
	authRouter.HandleFunc("/posts/{id}", postHandler.DeletePost).Methods("DELETE")
	authRouter.HandleFunc("/posts/{id}/publish", postHandler.PublishPost).Methods("POST")

	// Start server
	addr := ":" + port
	log.Printf("🚀 Blog Engine server starting on %s\n", addr)
	log.Printf("📚 API Documentation: http://localhost:%s/health\n", port)

	server := &http.Server{
		Addr:         addr,
		Handler:      corsMiddleware(loggingMiddleware(router)),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed:", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"healthy","service":"blog-engine","version":"1.0.0"}`)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
