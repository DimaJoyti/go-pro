# URL Shortener Architecture

## Overview

This document describes the architecture and design decisions for the URL Shortener service.

## Architecture Pattern

The application follows **Clean Architecture** principles with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────┐
│                    HTTP Layer                           │
│              (handlers/url_handler.go)                  │
│  - Request validation                                   │
│  - Response formatting                                  │
│  - HTTP routing                                         │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                  Service Layer                          │
│              (service/url_service.go)                   │
│  - Business logic                                       │
│  - Code generation                                      │
│  - Analytics tracking                                   │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│                Repository Layer                         │
│         (repository/memory.go, redis.go)                │
│  - Data persistence                                     │
│  - CRUD operations                                      │
│  - Storage abstraction                                  │
└─────────────────────────────────────────────────────────┘
```

## Layers

### 1. Domain Layer (`internal/domain/`)

Contains business entities and rules:
- **models.go**: Core domain models (URL, Analytics, Click)
- **Validation**: Business rule validation
- **No dependencies**: Pure business logic

**Key Types:**
- `URL`: Represents a shortened URL with metadata
- `Analytics`: Tracks usage statistics
- `ShortenRequest/Response`: DTOs for API communication

### 2. Repository Layer (`internal/repository/`)

Handles data persistence:
- **repository.go**: Interface definition
- **memory.go**: In-memory implementation
- **redis.go**: Redis implementation (future)

**Interface:**
```go
type URLRepository interface {
    Save(ctx context.Context, url *domain.URL) error
    FindByCode(ctx context.Context, code string) (*domain.URL, error)
    Update(ctx context.Context, url *domain.URL) error
    Delete(ctx context.Context, code string) error
    Exists(ctx context.Context, code string) (bool, error)
    GetAll(ctx context.Context) ([]*domain.URL, error)
}
```

**Benefits:**
- Easy to swap implementations
- Testable with mocks
- Concurrent-safe operations

### 3. Service Layer (`internal/service/`)

Contains business logic:
- URL shortening logic
- Code generation
- Analytics tracking
- Validation orchestration

**Responsibilities:**
- Coordinate between repository and handlers
- Implement business rules
- Generate unique codes
- Track analytics

### 4. Handler Layer (`internal/handlers/`)

HTTP request handling:
- Route registration
- Request parsing
- Response formatting
- Error handling
- Middleware integration

**Endpoints:**
- `POST /api/shorten` - Create short URL
- `GET /:code` - Redirect to original
- `GET /api/stats/:code` - Get statistics
- `GET /api/urls` - List all URLs
- `GET /health` - Health check

### 5. Package Layer (`pkg/`)

Reusable utilities:
- **shortener/generator.go**: Code generation algorithms

## Data Flow

### Shortening a URL

```
1. Client sends POST /api/shorten
   ↓
2. Handler validates request
   ↓
3. Service checks if custom code exists
   ↓
4. Service generates unique code (if needed)
   ↓
5. Repository saves URL
   ↓
6. Handler returns response
```

### Redirecting

```
1. Client sends GET /:code
   ↓
2. Handler extracts code
   ↓
3. Service finds URL in repository
   ↓
4. Service updates analytics
   ↓
5. Handler redirects to original URL
```

## Code Generation

### Algorithm

1. **Random Generation**: Uses crypto/rand for secure randomness
2. **Character Set**: `a-zA-Z0-9` (62 characters)
3. **Default Length**: 6 characters
4. **Collision Handling**: Retry with longer code if needed

### Uniqueness

- Check repository before saving
- Retry up to 10 times
- Increase length if collisions persist
- Probability of collision: ~1 in 56 billion (62^6)

## Analytics Tracking

### Tracked Metrics

1. **Total Clicks**: Incremented on each access
2. **Referrers**: Source of traffic
3. **User Agents**: Browser/client information
4. **IP Addresses**: Geographic tracking (optional)
5. **Click History**: Last 100 clicks

### Implementation

```go
func (u *URL) IncrementClicks(referrer, userAgent, ipAddress string) {
    u.Clicks++
    u.LastAccessed = time.Now()
    u.Analytics.TotalClicks++
    
    // Track referrer
    u.Analytics.Referrers[referrer]++
    
    // Track user agent
    u.Analytics.UserAgents[userAgent]++
    
    // Add to history (limit 100)
    u.Analytics.ClickHistory = append(...)
}
```

## Concurrency

### Thread Safety

- **Repository**: Uses `sync.RWMutex` for concurrent access
- **Read Operations**: Multiple readers allowed
- **Write Operations**: Exclusive lock
- **Context**: Propagated for cancellation

### Goroutine Safety

All operations are goroutine-safe:
```go
type MemoryRepository struct {
    urls map[string]*domain.URL
    mu   sync.RWMutex  // Protects urls map
}
```

## Error Handling

### Error Types

```go
var (
    ErrInvalidURL      = errors.New("invalid URL")
    ErrURLNotFound     = errors.New("URL not found")
    ErrCodeExists      = errors.New("short code already exists")
    ErrInvalidCode     = errors.New("invalid short code")
)
```

### Error Propagation

1. Domain layer: Validation errors
2. Repository layer: Storage errors
3. Service layer: Business logic errors
4. Handler layer: HTTP status codes

## Performance Considerations

### Memory Repository

- **Lookup**: O(1) - Hash map access
- **Insert**: O(1) - Hash map insert
- **Memory**: ~200 bytes per URL
- **Capacity**: Limited by available RAM

### Optimization Opportunities

1. **Caching**: Add Redis for distributed caching
2. **Database**: PostgreSQL for persistence
3. **Indexing**: B-tree indexes on short_code
4. **Sharding**: Distribute across multiple nodes
5. **CDN**: Cache redirects at edge

## Scalability

### Horizontal Scaling

```
┌──────────┐     ┌──────────┐     ┌──────────┐
│ Instance │     │ Instance │     │ Instance │
│    1     │     │    2     │     │    3     │
└────┬─────┘     └────┬─────┘     └────┬─────┘
     │                │                │
     └────────────────┼────────────────┘
                      │
                ┌─────▼─────┐
                │   Redis   │
                │  Cluster  │
                └───────────┘
```

### Load Balancing

- Round-robin distribution
- Health check endpoints
- Session affinity (optional)

## Security

### Current Implementation

- URL validation
- Input sanitization
- CORS headers
- No SQL injection (no SQL)

### Future Enhancements

- [ ] Rate limiting per IP
- [ ] API key authentication
- [ ] HTTPS enforcement
- [ ] DDoS protection
- [ ] Input length limits
- [ ] Malicious URL detection

## Testing Strategy

### Unit Tests

- Domain model validation
- Code generation
- Repository operations
- Service logic

### Integration Tests

- End-to-end API flows
- Database interactions
- Error scenarios

### Benchmarks

```bash
go test -bench=. -benchmem ./...
```

## Deployment

### Docker

```dockerfile
# Multi-stage build
FROM golang:1.21-alpine AS builder
# Build application
FROM alpine:latest
# Run application
```

### Environment Variables

```bash
PORT=8080
BASE_URL=http://localhost:8080
REDIS_URL=localhost:6379
REDIS_ENABLED=false
```

## Monitoring

### Metrics to Track

1. **Request Rate**: Requests per second
2. **Response Time**: P50, P95, P99 latencies
3. **Error Rate**: 4xx and 5xx responses
4. **Storage**: Number of URLs stored
5. **Cache Hit Rate**: Redis cache efficiency

### Logging

```go
log.Printf("Created short URL: %s -> %s", shortCode, originalURL)
log.Printf("%s %s %d %v", method, path, status, duration)
```

## Future Enhancements

### Phase 1: Core Features
- [x] URL shortening
- [x] Analytics tracking
- [x] In-memory storage
- [ ] Redis caching
- [ ] PostgreSQL persistence

### Phase 2: Advanced Features
- [ ] QR code generation
- [ ] Link expiration
- [ ] Password protection
- [ ] Custom domains
- [ ] Bulk operations

### Phase 3: Enterprise Features
- [ ] User authentication
- [ ] Team workspaces
- [ ] API rate limiting
- [ ] Advanced analytics
- [ ] Webhook notifications

## Design Decisions

### Why Clean Architecture?

- **Testability**: Easy to mock dependencies
- **Maintainability**: Clear separation of concerns
- **Flexibility**: Easy to swap implementations
- **Scalability**: Independent layer scaling

### Why In-Memory First?

- **Simplicity**: No external dependencies
- **Performance**: Fastest possible lookups
- **Learning**: Focus on Go concepts
- **Extensibility**: Easy to add persistence later

### Why Standard Library?

- **Stability**: No breaking changes
- **Performance**: Optimized by Go team
- **Learning**: Understand fundamentals
- **Deployment**: No dependency management

## References

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go Project Layout](https://github.com/golang-standards/project-layout)
- [Effective Go](https://golang.org/doc/effective_go)

