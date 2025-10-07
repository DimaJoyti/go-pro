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

## ➡️ Next Steps

Once you've completed all exercises and benchmarks show improvements, move on to:
**[Lesson 17: Security Best Practices](../lesson-17/README.md)**

---

**Optimize for performance!** ⚡
