# 🔄 Concurrency Deep Dive: Mastering Goroutines and Channels

This advanced tutorial provides an in-depth exploration of Go's concurrency primitives, patterns, and best practices.

## 🎯 Learning Objectives

By the end of this tutorial, you will:
- Understand goroutine lifecycle and scheduling
- Master channel patterns and idioms
- Prevent and debug deadlocks
- Detect and fix race conditions
- Implement advanced concurrency patterns
- Apply Go's memory model principles

---

## 📚 Table of Contents

1. [Goroutine Fundamentals](#goroutine-fundamentals)
2. [Channel Patterns](#channel-patterns)
3. [Deadlock Prevention](#deadlock-prevention)
4. [Race Condition Detection](#race-condition-detection)
5. [Advanced Patterns](#advanced-patterns)
6. [Memory Model](#memory-model)
7. [Best Practices](#best-practices)

---

## 1. Goroutine Fundamentals

### What is a Goroutine?

A goroutine is a lightweight thread managed by the Go runtime. Unlike OS threads, goroutines:
- Start with a small stack (2KB) that grows as needed
- Are multiplexed onto OS threads by the Go scheduler
- Have minimal context-switching overhead
- Can number in the hundreds of thousands

### Goroutine Lifecycle

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func demonstrateGoroutineLifecycle() {
    var wg sync.WaitGroup
    
    // 1. Creation
    wg.Add(1)
    go func() {
        defer wg.Done()
        
        // 2. Execution
        fmt.Println("Goroutine executing")
        time.Sleep(100 * time.Millisecond)
        
        // 3. Completion (automatic cleanup)
    }()
    
    // 4. Wait for completion
    wg.Wait()
    
    fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
}
```

### The Go Scheduler

Go uses an M:N scheduler:
- **M**: OS threads (machines)
- **P**: Processors (logical CPUs)
- **G**: Goroutines

```go
func demonstrateScheduler() {
    // Get number of logical CPUs
    numCPU := runtime.NumCPU()
    fmt.Printf("Number of CPUs: %d\n", numCPU)
    
    // Set max OS threads
    runtime.GOMAXPROCS(numCPU)
    
    // Launch many goroutines
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            // Simulate work
            time.Sleep(time.Millisecond)
        }(i)
    }
    
    wg.Wait()
}
```

---

## 2. Channel Patterns

### Unbuffered Channels

Unbuffered channels provide synchronization - send blocks until receive is ready.

```go
func unbufferedChannelExample() {
    ch := make(chan int)
    
    go func() {
        fmt.Println("Sending value...")
        ch <- 42  // Blocks until received
        fmt.Println("Value sent!")
    }()
    
    time.Sleep(time.Second) // Simulate delay
    value := <-ch           // Receive
    fmt.Printf("Received: %d\n", value)
}
```

### Buffered Channels

Buffered channels allow asynchronous communication up to buffer capacity.

```go
func bufferedChannelExample() {
    ch := make(chan int, 3) // Buffer size 3
    
    // Can send 3 values without blocking
    ch <- 1
    ch <- 2
    ch <- 3
    
    fmt.Printf("Channel length: %d, capacity: %d\n", len(ch), cap(ch))
    
    // Receive values
    fmt.Println(<-ch, <-ch, <-ch)
}
```

### Channel Direction

Specify channel direction for type safety.

```go
// Send-only channel
func sendOnly(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func receiveOnly(ch <-chan int) {
    value := <-ch
    fmt.Println(value)
}

func channelDirectionExample() {
    ch := make(chan int)
    
    go sendOnly(ch)
    receiveOnly(ch)
}
```

### Select Statement

Handle multiple channel operations.

```go
func selectExample() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(100 * time.Millisecond)
        ch1 <- "from ch1"
    }()
    
    go func() {
        time.Sleep(200 * time.Millisecond)
        ch2 <- "from ch2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        case <-time.After(500 * time.Millisecond):
            fmt.Println("timeout")
        }
    }
}
```

### Range Over Channels

```go
func rangeOverChannel() {
    ch := make(chan int, 5)
    
    // Send values
    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch) // Important: close when done sending
    }()
    
    // Receive until channel is closed
    for value := range ch {
        fmt.Println(value)
    }
}
```

---

## 3. Deadlock Prevention

### Understanding Deadlocks

A deadlock occurs when goroutines are waiting for each other indefinitely.

### Common Deadlock Scenario

```go
// ❌ BAD: This will deadlock
func deadlockExample() {
    ch := make(chan int)
    ch <- 42  // Blocks forever - no receiver!
    fmt.Println(<-ch)
}
```

### Solution: Use Goroutine

```go
// ✅ GOOD: Use goroutine for send
func noDeadlock() {
    ch := make(chan int)
    
    go func() {
        ch <- 42
    }()
    
    fmt.Println(<-ch)
}
```

### Real-World Example: Avoiding Deadlock

This is based on your `basic/deadlock.go` file, enhanced with best practices:

```go
package main

import (
    "fmt"
    "sync"
)

// ✅ GOOD: Proper channel usage with WaitGroup
func properChannelUsage() {
    var wg sync.WaitGroup
    nums := make(chan int)
    
    // Launch goroutines to send values
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go func(val int) {
            defer wg.Done()
            fmt.Printf("Sending: %d\n", val)
            nums <- val
            fmt.Printf("Sent: %d\n", val)
        }(i)
    }
    
    // Close channel after all sends complete
    go func() {
        fmt.Println("Waiting for all sends to complete...")
        wg.Wait()
        fmt.Println("Closing channel...")
        close(nums)
    }()
    
    // Read until channel is closed
    fmt.Println("Starting to receive values:")
    for num := range nums {
        fmt.Printf("Received: %d\n", num)
    }
    fmt.Println("Channel closed, program ending")
}

func main() {
    properChannelUsage()
}
```

### Deadlock Detection Patterns

```go
// Pattern 1: Timeout with select
func withTimeout() {
    ch := make(chan int)
    
    select {
    case val := <-ch:
        fmt.Println(val)
    case <-time.After(1 * time.Second):
        fmt.Println("Operation timed out - potential deadlock avoided")
    }
}

// Pattern 2: Non-blocking send/receive
func nonBlocking() {
    ch := make(chan int, 1)
    
    select {
    case ch <- 42:
        fmt.Println("Sent successfully")
    default:
        fmt.Println("Channel full, skipping send")
    }
    
    select {
    case val := <-ch:
        fmt.Printf("Received: %d\n", val)
    default:
        fmt.Println("No value available")
    }
}
```

---

## 4. Race Condition Detection

### What is a Race Condition?

A race condition occurs when multiple goroutines access shared data concurrently, and at least one modifies it.

### Example of Race Condition

```go
// ❌ BAD: Race condition
func raceConditionExample() {
    counter := 0
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++ // Race condition!
        }()
    }
    
    wg.Wait()
    fmt.Printf("Counter: %d\n", counter) // Unpredictable result
}
```

### Solution 1: Mutex

```go
// ✅ GOOD: Use mutex
func mutexSolution() {
    var (
        counter int
        mu      sync.Mutex
        wg      sync.WaitGroup
    )
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Counter: %d\n", counter) // Always 1000
}
```

### Solution 2: Atomic Operations

```go
import "sync/atomic"

// ✅ GOOD: Use atomic operations
func atomicSolution() {
    var (
        counter int64
        wg      sync.WaitGroup
    )
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            atomic.AddInt64(&counter, 1)
        }()
    }
    
    wg.Wait()
    fmt.Printf("Counter: %d\n", counter) // Always 1000
}
```

### Solution 3: Channel-Based

```go
// ✅ GOOD: Use channels
func channelSolution() {
    counter := 0
    ch := make(chan int)
    done := make(chan bool)
    
    // Counter goroutine
    go func() {
        for range ch {
            counter++
        }
        done <- true
    }()
    
    // Send increments
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            ch <- 1
        }()
    }
    
    wg.Wait()
    close(ch)
    <-done
    
    fmt.Printf("Counter: %d\n", counter) // Always 1000
}
```

### Detecting Races

Run your program with the race detector:

```bash
go run -race main.go
go test -race ./...
go build -race
```

---

## 5. Advanced Patterns

### Worker Pool Pattern

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type Job struct {
    ID     int
    Data   string
}

type Result struct {
    Job    Job
    Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()

    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job.ID)
        time.Sleep(100 * time.Millisecond) // Simulate work

        results <- Result{
            Job:    job,
            Output: fmt.Sprintf("Processed by worker %d", id),
        }
    }
}

func workerPoolExample() {
    const numWorkers = 3
    const numJobs = 10

    jobs := make(chan Job, numJobs)
    results := make(chan Result, numJobs)

    // Start workers
    var wg sync.WaitGroup
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- Job{ID: j, Data: fmt.Sprintf("data-%d", j)}
    }
    close(jobs)

    // Close results channel after all workers finish
    go func() {
        wg.Wait()
        close(results)
    }()

    // Collect results
    for result := range results {
        fmt.Printf("Job %d: %s\n", result.Job.ID, result.Output)
    }
}
```

### Fan-Out/Fan-In Pattern

```go
func fanOutFanIn() {
    // Input channel
    input := make(chan int)

    // Fan-out: multiple workers
    const numWorkers = 3
    workers := make([]<-chan int, numWorkers)

    for i := 0; i < numWorkers; i++ {
        workers[i] = worker(input)
    }

    // Fan-in: merge results
    output := merge(workers...)

    // Send data
    go func() {
        for i := 0; i < 10; i++ {
            input <- i
        }
        close(input)
    }()

    // Receive results
    for result := range output {
        fmt.Println(result)
    }
}

func worker(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for n := range input {
            output <- n * 2 // Process
        }
    }()
    return output
}

func merge(channels ...<-chan int) <-chan int {
    var wg sync.WaitGroup
    output := make(chan int)

    // Start goroutine for each input channel
    multiplex := func(c <-chan int) {
        defer wg.Done()
        for n := range c {
            output <- n
        }
    }

    wg.Add(len(channels))
    for _, c := range channels {
        go multiplex(c)
    }

    // Close output when all inputs are done
    go func() {
        wg.Wait()
        close(output)
    }()

    return output
}
```

### Pipeline Pattern

```go
func pipelineExample() {
    // Stage 1: Generate numbers
    gen := func(nums ...int) <-chan int {
        out := make(chan int)
        go func() {
            defer close(out)
            for _, n := range nums {
                out <- n
            }
        }()
        return out
    }

    // Stage 2: Square numbers
    sq := func(in <-chan int) <-chan int {
        out := make(chan int)
        go func() {
            defer close(out)
            for n := range in {
                out <- n * n
            }
        }()
        return out
    }

    // Stage 3: Filter even numbers
    filter := func(in <-chan int) <-chan int {
        out := make(chan int)
        go func() {
            defer close(out)
            for n := range in {
                if n%2 == 0 {
                    out <- n
                }
            }
        }()
        return out
    }

    // Build pipeline
    numbers := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    squared := sq(numbers)
    filtered := filter(squared)

    // Consume results
    for n := range filtered {
        fmt.Println(n)
    }
}
```

### Context for Cancellation

```go
import "context"

func contextCancellationExample() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    results := make(chan int)

    go func() {
        for i := 0; ; i++ {
            select {
            case <-ctx.Done():
                fmt.Println("Worker cancelled:", ctx.Err())
                close(results)
                return
            case results <- i:
                time.Sleep(500 * time.Millisecond)
            }
        }
    }()

    for result := range results {
        fmt.Println("Result:", result)
    }
}
```

### Semaphore Pattern

```go
func semaphoreExample() {
    const maxConcurrent = 3
    sem := make(chan struct{}, maxConcurrent)

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()

            // Acquire semaphore
            sem <- struct{}{}
            defer func() { <-sem }() // Release

            fmt.Printf("Task %d running\n", id)
            time.Sleep(time.Second)
            fmt.Printf("Task %d done\n", id)
        }(i)
    }

    wg.Wait()
}
```

---

## 6. Memory Model

### Happens-Before Relationship

Go's memory model specifies when reads of a variable in one goroutine can observe writes in another.

```go
var a, b int

// Goroutine 1
func goroutine1() {
    a = 1
    b = 2
}

// Goroutine 2
func goroutine2() {
    // May see b = 2 before a = 1!
    fmt.Println(a, b)
}
```

### Synchronization Guarantees

```go
// Channel send happens-before channel receive
func channelSync() {
    ch := make(chan int)
    var a string

    go func() {
        a = "hello"
        ch <- 0 // Send happens-before receive
    }()

    <-ch
    fmt.Println(a) // Guaranteed to see "hello"
}

// Mutex unlock happens-before next lock
func mutexSync() {
    var mu sync.Mutex
    var a string

    go func() {
        mu.Lock()
        a = "hello"
        mu.Unlock() // Unlock happens-before next lock
    }()

    time.Sleep(time.Millisecond) // Give goroutine time to run

    mu.Lock()
    fmt.Println(a) // Guaranteed to see "hello"
    mu.Unlock()
}

// WaitGroup.Done happens-before Wait returns
func waitGroupSync() {
    var wg sync.WaitGroup
    var a string

    wg.Add(1)
    go func() {
        a = "hello"
        wg.Done() // Done happens-before Wait returns
    }()

    wg.Wait()
    fmt.Println(a) // Guaranteed to see "hello"
}
```

---

## 7. Best Practices

### 1. Always Close Channels (When Appropriate)

```go
// ✅ GOOD: Sender closes channel
func goodChannelClose() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        close(ch) // Sender closes
    }()

    for val := range ch {
        fmt.Println(val)
    }
}

// ❌ BAD: Never close or receiver closes
func badChannelClose() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 5; i++ {
            ch <- i
        }
        // Forgot to close!
    }()

    // This will deadlock after receiving 5 values
    for val := range ch {
        fmt.Println(val)
    }
}
```

### 2. Use Context for Cancellation

```go
// ✅ GOOD: Respect context cancellation
func goodContextUsage(ctx context.Context) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            // Do work
            time.Sleep(100 * time.Millisecond)
        }
    }
}
```

### 3. Avoid Goroutine Leaks

```go
// ❌ BAD: Goroutine leak
func goroutineLeak() {
    ch := make(chan int)

    go func() {
        val := <-ch // Blocks forever if nothing sent
        fmt.Println(val)
    }()

    // Goroutine never exits!
}

// ✅ GOOD: Use context or timeout
func noGoroutineLeak() {
    ch := make(chan int)
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    go func() {
        select {
        case val := <-ch:
            fmt.Println(val)
        case <-ctx.Done():
            fmt.Println("Timeout, exiting")
            return
        }
    }()
}
```

### 4. Prefer Channels Over Shared Memory

```go
// ❌ Less idiomatic: Shared memory with mutex
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}

// ✅ More idiomatic: Channel-based
func channelCounter() {
    ops := make(chan int)
    done := make(chan bool)

    go func() {
        counter := 0
        for range ops {
            counter++
        }
        done <- true
    }()

    // Send operations
    for i := 0; i < 1000; i++ {
        ops <- 1
    }
    close(ops)
    <-done
}
```

### 5. Use Buffered Channels Wisely

```go
// ✅ GOOD: Buffer size matches expected load
func goodBuffering() {
    const numJobs = 100
    jobs := make(chan int, numJobs) // Buffer for all jobs

    // Send all jobs without blocking
    for i := 0; i < numJobs; i++ {
        jobs <- i
    }
    close(jobs)

    // Process
    for job := range jobs {
        fmt.Println(job)
    }
}
```

---

## 🧪 Exercises

### Exercise 1: Fix the Deadlock

```go
// Fix this deadlock
func exercise1() {
    ch := make(chan int)
    ch <- 42
    fmt.Println(<-ch)
}
```

### Exercise 2: Implement Worker Pool

Create a worker pool that processes 100 jobs with 5 workers.

### Exercise 3: Build a Pipeline

Create a 3-stage pipeline:
1. Generate numbers 1-100
2. Filter even numbers
3. Square the results

### Exercise 4: Race Condition Fix

Fix the race condition in this code:

```go
func exercise4() {
    var counter int
    var wg sync.WaitGroup

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter++
        }()
    }

    wg.Wait()
    fmt.Println(counter)
}
```

---

## 📚 Additional Resources

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Advanced Go Concurrency Patterns](https://go.dev/blog/io2013-talk-concurrency)
- [Go Memory Model](https://go.dev/ref/mem)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)

---

## 🎓 Key Takeaways

✅ **Goroutines are cheap** - use them liberally but responsibly
✅ **Channels provide synchronization** - prefer them over shared memory
✅ **Always close channels** when done sending (sender's responsibility)
✅ **Use context for cancellation** - respect cancellation signals
✅ **Detect races early** - use `-race` flag during development
✅ **Avoid goroutine leaks** - ensure all goroutines can exit
✅ **Understand happens-before** - know Go's memory model guarantees

---

**Master concurrency, master Go!** 🚀

Next: [Tutorial 10: Packages and Modules](../../course/lessons/lesson-10/README.md)

