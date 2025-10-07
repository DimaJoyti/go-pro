# 📘 Lesson 9: Goroutines and Channels

Welcome to Lesson 9! Concurrency is one of Go's standout features. This lesson introduces goroutines and channels - Go's primitives for concurrent programming.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Launch and manage goroutines
- Use channels for communication between goroutines
- Understand buffered vs unbuffered channels
- Apply select statements for channel operations
- Implement basic concurrency patterns
- Avoid common concurrency pitfalls
- Use sync package primitives when needed

## 📚 Theory

### Goroutines

Goroutines are lightweight threads managed by Go runtime:

```go
func main() {
    // Launch a goroutine
    go sayHello()
    
    // Launch anonymous goroutine
    go func() {
        fmt.Println("Anonymous goroutine")
    }()
    
    time.Sleep(time.Second) // Wait for goroutines
}

func sayHello() {
    fmt.Println("Hello from goroutine!")
}
```

### Channels

Channels enable communication between goroutines:

```go
// Unbuffered channel
ch := make(chan int)

// Buffered channel
buffered := make(chan string, 3)

// Send and receive
go func() {
    ch <- 42 // Send
}()

value := <-ch // Receive
```

### Channel Operations

```go
// Close a channel
close(ch)

// Check if channel is closed
value, ok := <-ch
if !ok {
    fmt.Println("Channel is closed")
}

// Range over channel
for value := range ch {
    fmt.Println(value)
}
```

### Select Statement

Handle multiple channel operations:

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("Timeout!")
default:
    fmt.Println("No channels ready")
}
```

## 💻 Hands-On Examples

### Example 1: Basic Goroutines
```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Second)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for r := 1; r <= 5; r++ {
        <-results
    }
}
```

### Example 2: Producer-Consumer Pattern
```go
func producer(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i
        time.Sleep(100 * time.Millisecond)
    }
    close(ch)
}

func consumer(ch <-chan int) {
    for value := range ch {
        fmt.Printf("Consumed: %d\n", value)
    }
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-09/exercises/`:

1. **Basic Goroutines**: Launch and coordinate goroutines
2. **Channel Communication**: Use channels for data passing
3. **Buffered Channels**: Work with buffered channels
4. **Select Statements**: Handle multiple channels
5. **Worker Pools**: Implement concurrent worker patterns
6. **Pipeline Pattern**: Build data processing pipelines

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-09
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Goroutines are cheap and lightweight
- Channels provide safe communication between goroutines
- Unbuffered channels are synchronous
- Buffered channels allow asynchronous communication
- Select enables non-blocking channel operations
- Always close channels when done sending
- Use sync package for low-level synchronization

## ⚠️ Common Pitfalls

- Forgetting to close channels
- Deadlocks from blocking operations
- Race conditions on shared data
- Goroutine leaks from infinite loops

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 10: Packages and Modules](../lesson-10/README.md)**

---

**Embrace concurrency!** 🚀
