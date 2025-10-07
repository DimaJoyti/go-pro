# 📘 Lesson 20: Building Production Systems

Welcome to the final lesson! This capstone lesson covers building production-ready Go systems, combining all previous concepts into real-world applications.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Design and build complete production systems
- Implement comprehensive observability
- Apply all Go concepts in real-world scenarios
- Build scalable and maintainable applications
- Handle production challenges and requirements
- Create systems that can evolve and grow
- Apply industry best practices

## 📚 Theory

### Production System Architecture

**Layered Architecture:**
```go
// Domain Layer
type User struct {
    ID       string
    Email    string
    Name     string
    Active   bool
    Created  time.Time
}

type UserRepository interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}

// Application Layer
type UserService struct {
    repo   UserRepository
    logger Logger
    events EventBus
}

func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
    // Validation
    if err := req.Validate(); err != nil {
        return nil, fmt.Errorf("invalid request: %w", err)
    }
    
    // Business logic
    user := &User{
        ID:      generateID(),
        Email:   req.Email,
        Name:    req.Name,
        Active:  true,
        Created: time.Now(),
    }
    
    // Persistence
    if err := s.repo.Create(ctx, user); err != nil {
        s.logger.Error("failed to create user", "error", err)
        return nil, err
    }
    
    // Event publishing
    s.events.Publish(UserCreatedEvent{UserID: user.ID})
    
    return user, nil
}
```

### Comprehensive Observability

**Structured Logging:**
```go
type Logger interface {
    Debug(msg string, fields ...Field)
    Info(msg string, fields ...Field)
    Warn(msg string, fields ...Field)
    Error(msg string, fields ...Field)
    With(fields ...Field) Logger
}

type Field struct {
    Key   string
    Value interface{}
}

func String(key, value string) Field {
    return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
    return Field{Key: key, Value: value}
}

func Error(err error) Field {
    return Field{Key: "error", Value: err.Error()}
}
```

**Metrics Collection:**
```go
type Metrics interface {
    Counter(name string, tags map[string]string) Counter
    Gauge(name string, tags map[string]string) Gauge
    Histogram(name string, tags map[string]string) Histogram
}

type Counter interface {
    Inc()
    Add(delta float64)
}

type RequestMetrics struct {
    requestsTotal    Counter
    requestDuration  Histogram
    activeRequests   Gauge
}

func (m *RequestMetrics) RecordRequest(method, path string, duration time.Duration, status int) {
    tags := map[string]string{
        "method": method,
        "path":   path,
        "status": strconv.Itoa(status),
    }
    
    m.requestsTotal.Add(1)
    m.requestDuration.Observe(duration.Seconds())
}
```

### Configuration Management

**Environment-based Configuration:**
```go
type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
    Redis    RedisConfig    `yaml:"redis"`
    Logging  LoggingConfig  `yaml:"logging"`
    Metrics  MetricsConfig  `yaml:"metrics"`
}

type ServerConfig struct {
    Port         int           `yaml:"port" env:"SERVER_PORT" envDefault:"8080"`
    ReadTimeout  time.Duration `yaml:"read_timeout" env:"SERVER_READ_TIMEOUT" envDefault:"30s"`
    WriteTimeout time.Duration `yaml:"write_timeout" env:"SERVER_WRITE_TIMEOUT" envDefault:"30s"`
    IdleTimeout  time.Duration `yaml:"idle_timeout" env:"SERVER_IDLE_TIMEOUT" envDefault:"120s"`
}

func LoadConfig() (*Config, error) {
    var cfg Config
    
    // Load from file
    if data, err := os.ReadFile("config.yaml"); err == nil {
        if err := yaml.Unmarshal(data, &cfg); err != nil {
            return nil, fmt.Errorf("failed to parse config file: %w", err)
        }
    }
    
    // Override with environment variables
    if err := env.Parse(&cfg); err != nil {
        return nil, fmt.Errorf("failed to parse environment variables: %w", err)
    }
    
    return &cfg, nil
}
```

## 💻 Hands-On Examples

### Example 1: Complete HTTP Service
```go
type Server struct {
    config   *Config
    logger   Logger
    metrics  Metrics
    services *Services
    router   *mux.Router
}

func NewServer(cfg *Config, logger Logger, metrics Metrics, services *Services) *Server {
    s := &Server{
        config:   cfg,
        logger:   logger,
        metrics:  metrics,
        services: services,
        router:   mux.NewRouter(),
    }
    
    s.setupRoutes()
    s.setupMiddleware()
    
    return s
}

func (s *Server) setupMiddleware() {
    s.router.Use(s.loggingMiddleware)
    s.router.Use(s.metricsMiddleware)
    s.router.Use(s.recoveryMiddleware)
    s.router.Use(s.corsMiddleware)
}

func (s *Server) setupRoutes() {
    api := s.router.PathPrefix("/api/v1").Subrouter()
    
    // Health checks
    s.router.HandleFunc("/health", s.healthCheck).Methods("GET")
    s.router.HandleFunc("/ready", s.readinessCheck).Methods("GET")
    
    // User routes
    users := api.PathPrefix("/users").Subrouter()
    users.HandleFunc("", s.createUser).Methods("POST")
    users.HandleFunc("/{id}", s.getUser).Methods("GET")
    users.HandleFunc("/{id}", s.updateUser).Methods("PUT")
    users.HandleFunc("/{id}", s.deleteUser).Methods("DELETE")
}

func (s *Server) Start(ctx context.Context) error {
    server := &http.Server{
        Addr:         fmt.Sprintf(":%d", s.config.Server.Port),
        Handler:      s.router,
        ReadTimeout:  s.config.Server.ReadTimeout,
        WriteTimeout: s.config.Server.WriteTimeout,
        IdleTimeout:  s.config.Server.IdleTimeout,
    }
    
    go func() {
        <-ctx.Done()
        shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()
        server.Shutdown(shutdownCtx)
    }()
    
    s.logger.Info("starting server", String("port", strconv.Itoa(s.config.Server.Port)))
    return server.ListenAndServe()
}
```

### Example 2: Graceful Shutdown
```go
func main() {
    // Load configuration
    cfg, err := LoadConfig()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }
    
    // Initialize dependencies
    logger := NewLogger(cfg.Logging)
    metrics := NewMetrics(cfg.Metrics)
    db := NewDatabase(cfg.Database)
    
    // Initialize services
    services := &Services{
        UserService: NewUserService(db, logger, metrics),
    }
    
    // Create server
    server := NewServer(cfg, logger, metrics, services)
    
    // Setup graceful shutdown
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // Handle shutdown signals
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    go func() {
        <-sigChan
        logger.Info("shutdown signal received")
        cancel()
    }()
    
    // Start server
    if err := server.Start(ctx); err != nil && err != http.ErrServerClosed {
        logger.Error("server error", Error(err))
        os.Exit(1)
    }
    
    logger.Info("server stopped gracefully")
}
```

## 🧪 Final Project

Build a complete production system in `../../code/lesson-20/project/`:

**Requirements:**
- RESTful API with full CRUD operations
- Database integration with migrations
- Authentication and authorization
- Comprehensive logging and metrics
- Health checks and monitoring
- Docker containerization
- CI/CD pipeline
- Documentation and tests

## ✅ Validation

Validate your complete system:

```bash
cd ../../code/lesson-20
go test -v ./...
go test -race ./...
go test -cover ./...
docker build -t production-app .
docker run -p 8080:8080 production-app
```

## 🔍 Key Takeaways

- Production systems require comprehensive planning
- Observability is crucial for operations
- Configuration management enables flexibility
- Graceful shutdown prevents data loss
- Testing at all levels ensures reliability
- Documentation facilitates maintenance
- Monitoring enables proactive issue resolution

## 🏆 Production Checklist

- [ ] Comprehensive error handling
- [ ] Structured logging with correlation IDs
- [ ] Metrics collection and monitoring
- [ ] Health checks and readiness probes
- [ ] Graceful shutdown handling
- [ ] Configuration management
- [ ] Security best practices
- [ ] Performance optimization
- [ ] Comprehensive testing
- [ ] Documentation and runbooks

## 🎓 Congratulations!

You've completed the GO-PRO course! You now have the skills to:
- Build robust Go applications
- Apply best practices and design patterns
- Create production-ready systems
- Handle real-world challenges
- Continue learning and growing as a Go developer

## 🚀 Next Steps

- Build your own projects
- Contribute to open source
- Join the Go community
- Keep learning and experimenting
- Share your knowledge with others

---

**You're now a Go Pro!** 🎉
