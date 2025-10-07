# 📘 Lesson 13: HTTP Servers and REST APIs

Welcome to building web services with Go! This lesson covers creating robust HTTP servers, designing REST APIs, and implementing web service best practices.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Build HTTP servers using Go's standard library
- Design and implement RESTful APIs
- Handle HTTP requests, responses, and routing
- Implement middleware for cross-cutting concerns
- Handle JSON serialization and validation
- Implement authentication and authorization
- Add logging, monitoring, and error handling
- Deploy and scale HTTP services

## 📚 Theory

### Basic HTTP Server

Go's `net/http` package provides everything needed for HTTP servers:

```go
func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/users", usersHandler)
    
    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Welcome to our API",
        "version": "1.0.0",
    })
}
```

### Modern Routing with ServeMux (Go 1.22+)

Use the enhanced ServeMux for better routing:

```go
func setupRoutes() *http.ServeMux {
    mux := http.NewServeMux()
    
    // Method-specific routes
    mux.HandleFunc("GET /users", listUsers)
    mux.HandleFunc("POST /users", createUser)
    mux.HandleFunc("GET /users/{id}", getUser)
    mux.HandleFunc("PUT /users/{id}", updateUser)
    mux.HandleFunc("DELETE /users/{id}", deleteUser)
    
    return mux
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id") // Extract path parameter
    
    user, err := userService.GetByID(id)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

### REST API Design

Follow REST principles for API design:

```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

type UserHandler struct {
    service UserService
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        h.handleGet(w, r)
    case http.MethodPost:
        h.handlePost(w, r)
    case http.MethodPut:
        h.handlePut(w, r)
    case http.MethodDelete:
        h.handleDelete(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func (h *UserHandler) handlePost(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    if err := h.service.Create(&user); err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
```

### Middleware Pattern

Implement cross-cutting concerns with middleware:

```go
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Wrap ResponseWriter to capture status code
        wrapped := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        
        next.ServeHTTP(wrapped, r)
        
        log.Printf("%s %s %d %v", r.Method, r.URL.Path, wrapped.statusCode, time.Since(start))
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

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Authorization required", http.StatusUnauthorized)
            return
        }
        
        user, err := validateToken(token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        // Add user to request context
        ctx := context.WithValue(r.Context(), "user", user)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

### JSON Handling and Validation

Handle JSON data with validation:

```go
type CreateUserRequest struct {
    Name  string `json:"name" validate:"required,min=2,max=50"`
    Email string `json:"email" validate:"required,email"`
}

func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    
    // Decode JSON
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        writeErrorResponse(w, "Invalid JSON", http.StatusBadRequest)
        return
    }
    
    // Validate request
    if err := validator.Struct(&req); err != nil {
        writeValidationError(w, err)
        return
    }
    
    // Create user
    user := &User{
        Name:      req.Name,
        Email:     req.Email,
        CreatedAt: time.Now(),
    }
    
    if err := h.service.Create(user); err != nil {
        writeErrorResponse(w, "Failed to create user", http.StatusInternalServerError)
        return
    }
    
    writeJSONResponse(w, user, http.StatusCreated)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

func writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(map[string]string{
        "error": message,
    })
}
```

### Server Configuration

Configure servers for production:

```go
func main() {
    mux := setupRoutes()
    
    // Apply middleware
    handler := loggingMiddleware(corsMiddleware(mux))
    
    server := &http.Server{
        Addr:         ":8080",
        Handler:      handler,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    
    // Graceful shutdown
    go func() {
        log.Println("Server starting on :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()
    
    // Wait for interrupt signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    log.Println("Shutting down server...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }
    
    log.Println("Server exited")
}
```

## 💻 Hands-On Examples

Navigate to `../../code/lesson-13/` to see and run these examples.

### Example 1: Complete REST API
```go
type BookAPI struct {
    books []Book
    mu    sync.RWMutex
}

func (api *BookAPI) setupRoutes() *http.ServeMux {
    mux := http.NewServeMux()
    
    mux.HandleFunc("GET /books", api.listBooks)
    mux.HandleFunc("POST /books", api.createBook)
    mux.HandleFunc("GET /books/{id}", api.getBook)
    mux.HandleFunc("PUT /books/{id}", api.updateBook)
    mux.HandleFunc("DELETE /books/{id}", api.deleteBook)
    
    return mux
}
```

### Example 2: File Upload Handler
```go
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    // Parse multipart form (32MB max)
    if err := r.ParseMultipartForm(32 << 20); err != nil {
        http.Error(w, "Failed to parse form", http.StatusBadRequest)
        return
    }
    
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Failed to get file", http.StatusBadRequest)
        return
    }
    defer file.Close()
    
    // Save file
    dst, err := os.Create("uploads/" + header.Filename)
    if err != nil {
        http.Error(w, "Failed to create file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()
    
    if _, err := io.Copy(dst, file); err != nil {
        http.Error(w, "Failed to save file", http.StatusInternalServerError)
        return
    }
    
    writeJSONResponse(w, map[string]string{
        "message":  "File uploaded successfully",
        "filename": header.Filename,
    }, http.StatusOK)
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-13/exercises/`:

1. **Basic HTTP Server**: Build a simple web server
2. **REST API**: Implement a complete CRUD API
3. **Middleware Chain**: Create authentication and logging middleware
4. **File Server**: Build a static file server with directory listing
5. **WebSocket Chat**: Implement real-time chat with WebSockets
6. **API Gateway**: Build a simple API gateway with routing

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-13
go test -v ./exercises/...
go run main.go  # Start the server and test endpoints
```

## 🔍 Key Takeaways

- Go's standard library provides powerful HTTP server capabilities
- REST APIs should follow standard HTTP methods and status codes
- Middleware enables clean separation of cross-cutting concerns
- Proper JSON handling and validation are essential
- Graceful shutdown ensures clean server termination
- Testing HTTP handlers requires special techniques

## 📖 Additional Resources

- [net/http Package](https://pkg.go.dev/net/http)
- [HTTP Server Examples](https://go.dev/doc/articles/wiki/)
- [REST API Design](https://restfulapi.net/)
- [HTTP Testing](https://pkg.go.dev/net/http/httptest)

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 14: Database Integration](../lesson-14/README.md)**

---

**Build amazing APIs!** 🌐

*Remember: Good APIs are intuitive, well-documented, and handle errors gracefully!*
