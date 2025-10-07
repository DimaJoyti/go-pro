# 📘 Lesson 16: Performance Optimization and Profiling

Welcome to Lesson 16! This advanced lesson covers performance optimization techniques, profiling tools, and building high-performance Go applications.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Profile Go applications using built-in tools
- Identify and fix performance bottlenecks
- Optimize memory usage and reduce allocations
- Apply CPU optimization techniques
- Use benchmarking for performance validation
- Implement caching strategies
- Build performance-critical applications

## 📚 Theory

### Go Profiling Tools

**CPU Profiling:**
```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()
    
    // Your application code
}
```

**Memory Profiling:**
```bash
go tool pprof http://localhost:6060/debug/pprof/heap
go tool pprof http://localhost:6060/debug/pprof/profile
```

### Memory Optimization

**Reduce Allocations:**
```go
// Bad: Creates new slice on each call
func processData(data []int) []int {
    result := make([]int, 0)
    for _, v := range data {
        if v > 0 {
            result = append(result, v*2)
        }
    }
    return result
}

// Good: Pre-allocate with capacity
func processDataOptimized(data []int) []int {
    result := make([]int, 0, len(data))
    for _, v := range data {
        if v > 0 {
            result = append(result, v*2)
        }
    }
    return result
}
```

**Object Pooling:**
```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

func processWithPool() {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf)
    
    // Use buffer
}
```

### CPU Optimization

**Avoid Unnecessary Work:**
```go
// Bad: Repeated expensive operations
func findUsers(names []string, users []User) []User {
    var result []User
    for _, name := range names {
        for _, user := range users {
            if user.Name == name {
                result = append(result, user)
                break
            }
        }
    }
    return result
}

// Good: Use map for O(1) lookup
func findUsersOptimized(names []string, users []User) []User {
    userMap := make(map[string]User, len(users))
    for _, user := range users {
        userMap[user.Name] = user
    }
    
    var result []User
    for _, name := range names {
        if user, exists := userMap[name]; exists {
            result = append(result, user)
        }
    }
    return result
}
```

## 💻 Hands-On Examples

### Example 1: Benchmarking
```go
func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var result string
        for j := 0; j < 100; j++ {
            result += "hello"
        }
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var builder strings.Builder
        for j := 0; j < 100; j++ {
            builder.WriteString("hello")
        }
        _ = builder.String()
    }
}
```

### Example 2: Memory Pool
```go
type RequestPool struct {
    pool sync.Pool
}

func NewRequestPool() *RequestPool {
    return &RequestPool{
        pool: sync.Pool{
            New: func() interface{} {
                return &Request{
                    Headers: make(map[string]string),
                    Body:    make([]byte, 0, 1024),
                }
            },
        },
    }
}

func (p *RequestPool) Get() *Request {
    return p.pool.Get().(*Request)
}

func (p *RequestPool) Put(req *Request) {
    req.Reset()
    p.pool.Put(req)
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-16/exercises/`:

1. **Profiling Practice**: Profile and optimize slow functions
2. **Memory Optimization**: Reduce memory allocations
3. **CPU Optimization**: Optimize CPU-intensive operations
4. **Caching Implementation**: Build efficient caching systems
5. **Benchmark Comparison**: Compare different implementations
6. **Real-World Optimization**: Optimize a complete application

## ✅ Validation

Run the tests and benchmarks:

```bash
cd ../../code/lesson-16
go test -v ./exercises/...
go test -bench=. ./exercises/...
go test -benchmem ./exercises/...
```

## 🔍 Key Takeaways

- Profile before optimizing
- Focus on the biggest bottlenecks first
- Reduce memory allocations where possible
- Use object pooling for frequently allocated objects
- Benchmark to validate optimizations
- Consider algorithmic improvements
- Balance readability with performance

## 🛠️ Profiling Commands

```bash
# CPU profiling
go tool pprof cpu.prof

# Memory profiling
go tool pprof mem.prof

# Live profiling
go tool pprof http://localhost:6060/debug/pprof/profile

# Trace analysis
go tool trace trace.out
```

## 🎯 Real-World Applications

### How This is Used in GO-PRO Backend

**Repository Layer Optimization:**
```go
// Optimized query with proper indexing
func (r *LessonRepository) GetByCourseID(ctx context.Context, courseID string, pagination *domain.PaginationRequest) ([]*domain.Lesson, int64, error) {
    // Use prepared statement for better performance
    query := `
        SELECT l.id, l.course_id, l.title, l.description, l.content, l.lesson_order
        FROM gopro.lessons l
        WHERE l.course_id = $1 AND l.status = 'published'
        ORDER BY l.lesson_order ASC
        LIMIT $2 OFFSET $3
    `

    limit := pagination.PageSize
    offset := (pagination.Page - 1) * pagination.PageSize

    rows, err := r.db.QueryContext(ctx, query, courseID, limit, offset)
    if err != nil {
        return nil, 0, fmt.Errorf("failed to query lessons: %w", err)
    }
    defer rows.Close()

    // Pre-allocate slice with estimated capacity
    lessons := make([]*domain.Lesson, 0, limit)
    for rows.Next() {
        var lesson domain.Lesson
        if err := rows.Scan(&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Description, &lesson.Content, &lesson.Order); err != nil {
            return nil, 0, err
        }
        lessons = append(lessons, &lesson)
    }

    return lessons, total, nil
}
```

**Caching Strategy:**
```go
func (s *lessonService) GetLessonByID(ctx context.Context, id string) (*domain.Lesson, error) {
    // Check cache first
    cacheKey := fmt.Sprintf("lesson:%s", id)

    if cached, err := s.cache.Get(ctx, cacheKey); err == nil {
        return cached.(*domain.Lesson), nil
    }

    // Cache miss - query database
    lesson, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    // Cache for 5 minutes
    _ = s.cache.Set(ctx, cacheKey, lesson, 5*time.Minute)

    return lesson, nil
}
```

## 🔒 Security Considerations

**Prevent Timing Attacks:**
```go
// Use constant-time comparison for sensitive data
import "crypto/subtle"

func compareSecrets(a, b string) bool {
    return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
```

**Resource Limits:**
```go
// Limit concurrent operations to prevent resource exhaustion
type RateLimiter struct {
    sem chan struct{}
}

func NewRateLimiter(maxConcurrent int) *RateLimiter {
    return &RateLimiter{
        sem: make(chan struct{}, maxConcurrent),
    }
}

func (r *RateLimiter) Do(ctx context.Context, fn func() error) error {
    select {
    case r.sem <- struct{}{}:
        defer func() { <-r.sem }()
        return fn()
    case <-ctx.Done():
        return ctx.Err()
    }
}
```

## 📊 Observability Insights

**Performance Metrics:**
```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/metric"
)

var (
    requestDuration metric.Float64Histogram
    cacheHits       metric.Int64Counter
    cacheMisses     metric.Int64Counter
)

func init() {
    meter := otel.Meter("lesson-service")

    requestDuration, _ = meter.Float64Histogram(
        "lesson.request.duration",
        metric.WithDescription("Duration of lesson requests"),
        metric.WithUnit("ms"),
    )

    cacheHits, _ = meter.Int64Counter(
        "lesson.cache.hits",
        metric.WithDescription("Number of cache hits"),
    )

    cacheMisses, _ = meter.Int64Counter(
        "lesson.cache.misses",
        metric.WithDescription("Number of cache misses"),
    )
}

func (s *lessonService) GetLessonByID(ctx context.Context, id string) (*domain.Lesson, error) {
    start := time.Now()
    defer func() {
        duration := time.Since(start).Milliseconds()
        requestDuration.Record(ctx, float64(duration))
    }()

    // Check cache
    if lesson, err := s.cache.Get(ctx, "lesson:"+id); err == nil {
        cacheHits.Add(ctx, 1)
        return lesson.(*domain.Lesson), nil
    }

    cacheMisses.Add(ctx, 1)
    return s.repo.GetByID(ctx, id)
}
```

**Profiling in Production:**
```go
import _ "net/http/pprof"

func main() {
    // Enable pprof endpoints
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // Your application
}

// Access profiles:
// http://localhost:6060/debug/pprof/
// http://localhost:6060/debug/pprof/heap
// http://localhost:6060/debug/pprof/profile?seconds=30
```

## 🧪 Advanced Testing

**Benchmark with Different Scenarios:**
```go
func BenchmarkGetLesson(b *testing.B) {
    scenarios := []struct {
        name     string
        cacheHit bool
    }{
        {"cache_hit", true},
        {"cache_miss", false},
    }

    for _, sc := range scenarios {
        b.Run(sc.name, func(b *testing.B) {
            // Setup
            service := setupService(sc.cacheHit)
            ctx := context.Background()

            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                _, _ = service.GetLessonByID(ctx, "lesson-1")
            }
        })
    }
}
```

**Memory Allocation Tracking:**
```go
func BenchmarkProcessData(b *testing.B) {
    data := generateTestData(1000)

    b.ReportAllocs()
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = processData(data)
    }
}

// Run with: go test -bench=. -benchmem
// Output shows allocations per operation
```

## 📖 Additional Resources

- [Go Blog - Profiling Go Programs](https://go.dev/blog/pprof)
- [Effective Go - Performance](https://go.dev/doc/effective_go)
- [GO-PRO Backend Performance Patterns](../../backend/internal/service/)
- [pprof Documentation](https://pkg.go.dev/net/http/pprof)

## 🎓 Key Takeaways Summary

✅ **Profile First**: Always measure before optimizing
✅ **Reduce Allocations**: Pre-allocate slices and use object pools
✅ **Cache Wisely**: Cache expensive operations with proper TTL
✅ **Benchmark Everything**: Validate optimizations with benchmarks
✅ **Monitor Production**: Use pprof and metrics in production

## ➡️ Next Steps

Once you've completed all exercises and benchmarks show improvements, move on to:
**[Lesson 17: Security Best Practices](../lesson-17/README.md)**

---

**Optimize for performance!** ⚡
