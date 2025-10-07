# 📘 Lesson 11: Advanced Concurrency Patterns

Welcome to advanced Go concurrency! This lesson explores sophisticated patterns for building concurrent applications using goroutines, channels, and the context package.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Implement worker pool patterns for concurrent processing
- Use fan-out/fan-in patterns for parallel data processing
- Build pipeline patterns for streaming data transformation
- Apply the context package for cancellation and timeouts
- Use sync package primitives for advanced synchronization
- Detect and prevent race conditions
- Design scalable concurrent architectures

## 📚 Theory

### Worker Pool Pattern

Worker pools manage a fixed number of goroutines to process work items:

```go
func workerPool(jobs <-chan Job, results chan<- Result, numWorkers int) {
    var wg sync.WaitGroup
    
    // Start workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobs {
                result := processJob(job)
                results <- result
            }
        }(i)
    }
    
    // Close results when all workers are done
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

### Fan-Out/Fan-In Pattern

Distribute work across multiple goroutines and collect results:

```go
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    
    for i := 0; i < workers; i++ {
        output := make(chan int)
        outputs[i] = output
        
        go func() {
            defer close(output)
            for n := range input {
                output <- process(n)
            }
        }()
    }
    
    return outputs
}

func fanIn(inputs ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    
    for _, input := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for n := range ch {
                output <- n
            }
        }(input)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}
```

### Pipeline Pattern

Chain processing stages together:

```go
func pipeline() <-chan Result {
    // Stage 1: Generate data
    numbers := make(chan int)
    go func() {
        defer close(numbers)
        for i := 1; i <= 100; i++ {
            numbers <- i
        }
    }()
    
    // Stage 2: Square numbers
    squares := make(chan int)
    go func() {
        defer close(squares)
        for n := range numbers {
            squares <- n * n
        }
    }()
    
    // Stage 3: Filter even squares
    results := make(chan Result)
    go func() {
        defer close(results)
        for sq := range squares {
            if sq%2 == 0 {
                results <- Result{Value: sq}
            }
        }
    }()
    
    return results
}
```

### Context Package

Use context for cancellation, timeouts, and request-scoped values:

```go
func processWithTimeout(ctx context.Context, data []int) ([]int, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    results := make(chan []int, 1)
    errors := make(chan error, 1)
    
    go func() {
        // Simulate long-running work
        processed := make([]int, len(data))
        for i, v := range data {
            select {
            case <-ctx.Done():
                errors <- ctx.Err()
                return
            default:
                processed[i] = v * 2
                time.Sleep(100 * time.Millisecond)
            }
        }
        results <- processed
    }()
    
    select {
    case result := <-results:
        return result, nil
    case err := <-errors:
        return nil, err
    case <-ctx.Done():
        return nil, ctx.Err()
    }
}
```

### Sync Package Primitives

#### **Mutex for Exclusive Access**
```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

#### **RWMutex for Read/Write Access**
```go
type SafeMap struct {
    mu   sync.RWMutex
    data map[string]int
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    val, ok := sm.data[key]
    return val, ok
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.data[key] = value
}
```

#### **Once for One-Time Initialization**
```go
var (
    instance *Database
    once     sync.Once
)

func GetDatabase() *Database {
    once.Do(func() {
        instance = &Database{
            // expensive initialization
        }
    })
    return instance
}
```

### Race Condition Detection

Use Go's race detector to find data races:

```bash
go run -race main.go
go test -race ./...
```

## 💻 Hands-On Examples

Navigate to `../../code/lesson-11/` to see and run these examples.

### Example 1: Worker Pool for Image Processing
```go
type ImageJob struct {
    ID       int
    Filename string
}

type ImageResult struct {
    ID    int
    Size  int64
    Error error
}

func processImages(jobs <-chan ImageJob, results chan<- ImageResult) {
    for job := range jobs {
        // Simulate image processing
        time.Sleep(100 * time.Millisecond)
        
        result := ImageResult{
            ID:   job.ID,
            Size: int64(len(job.Filename) * 1024), // Mock size
        }
        
        results <- result
    }
}
```

### Example 2: Rate-Limited API Client
```go
type RateLimitedClient struct {
    limiter *time.Ticker
    client  *http.Client
}

func NewRateLimitedClient(requestsPerSecond int) *RateLimitedClient {
    return &RateLimitedClient{
        limiter: time.NewTicker(time.Second / time.Duration(requestsPerSecond)),
        client:  &http.Client{Timeout: 10 * time.Second},
    }
}

func (c *RateLimitedClient) Get(ctx context.Context, url string) (*http.Response, error) {
    select {
    case <-c.limiter.C:
        // Rate limit satisfied
    case <-ctx.Done():
        return nil, ctx.Err()
    }
    
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    return c.client.Do(req)
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-11/exercises/`:

1. **Worker Pool Implementation**: Build a configurable worker pool
2. **Pipeline Processing**: Create a multi-stage data pipeline
3. **Context Cancellation**: Implement timeout and cancellation handling
4. **Concurrent Map**: Build a thread-safe map with RWMutex
5. **Rate Limiter**: Implement a token bucket rate limiter
6. **Producer-Consumer**: Build a bounded buffer pattern

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-11
go test -v ./exercises/...
go test -race ./exercises/...  # Check for race conditions
```

## 🔍 Key Takeaways

- Worker pools provide controlled concurrency
- Fan-out/fan-in patterns enable parallel processing
- Pipelines create efficient streaming architectures
- Context package enables proper cancellation and timeouts
- Sync primitives provide fine-grained synchronization control
- Race detection is essential for concurrent code quality

## 📖 Additional Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Context Package](https://go.dev/blog/context)
- [Sync Package Documentation](https://pkg.go.dev/sync)
- [Race Detector](https://go.dev/blog/race-detector)

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 12: Testing and Benchmarking](../lesson-12/README.md)**

---

**Master concurrency!** 🚀

*Remember: Concurrent programming is powerful but requires careful design. Always test with the race detector!*
