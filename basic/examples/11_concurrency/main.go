package main

import (
	"fmt"
	"sync"
	"time"
)

// Simple goroutine
func sayHello(id int) {
	fmt.Printf("   Hello from goroutine %d\n", id)
}

// Goroutine with channel
func sendData(ch chan int, value int) {
	ch <- value
}

// Worker pattern
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("   Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

// Producer-Consumer
func producer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int, done chan<- bool) {
	for num := range ch {
		fmt.Printf("   Consumed: %d\n", num)
	}
	done <- true
}

func main() {
	fmt.Println("=== Concurrency ===\n")

	// Basic goroutine
	fmt.Println("1. Basic Goroutine:")
	go sayHello(1)
	go sayHello(2)
	time.Sleep(100 * time.Millisecond) // Wait for goroutines
	fmt.Println()

	// WaitGroup
	fmt.Println("2. WaitGroup:")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("   Goroutine %d started\n", id)
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("   Goroutine %d finished\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("   All goroutines completed\n")

	// Channels - basic
	fmt.Println("3. Channels (Basic):")
	ch := make(chan int)
	go sendData(ch, 42)
	value := <-ch
	fmt.Printf("   Received: %d\n\n", value)

	// Buffered channels
	fmt.Println("4. Buffered Channels:")
	buffered := make(chan string, 2)
	buffered <- "first"
	buffered <- "second"
	fmt.Printf("   Received: %s\n", <-buffered)
	fmt.Printf("   Received: %s\n\n", <-buffered)

	// Channel direction
	fmt.Println("5. Channel Direction:")
	messages := make(chan string)
	go func(ch chan<- string) { // Send-only
		ch <- "Hello"
	}(messages)
	msg := <-messages
	fmt.Printf("   Message: %s\n\n", msg)

	// Select statement
	fmt.Println("6. Select Statement:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from channel 1"
	}()
	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("   Received %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("   Received %s\n", msg2)
		}
	}
	fmt.Println()

	// Worker pool
	fmt.Println("7. Worker Pool:")
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var workerWg sync.WaitGroup

	// Start workers
	for w := 1; w <= 3; w++ {
		workerWg.Add(1)
		go worker(w, jobs, results, &workerWg)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers
	go func() {
		workerWg.Wait()
		close(results)
	}()

	// Collect results
	fmt.Println("   Results:")
	for result := range results {
		fmt.Printf("   Result: %d\n", result)
	}
	fmt.Println()

	// Producer-Consumer
	fmt.Println("8. Producer-Consumer:")
	dataCh := make(chan int, 3)
	done := make(chan bool)

	go producer(dataCh, 5)
	go consumer(dataCh, done)

	<-done
	fmt.Println()

	// Mutex
	fmt.Println("9. Mutex:")
	var mutex sync.Mutex
	counter := 0
	var mutexWg sync.WaitGroup

	for i := 0; i < 5; i++ {
		mutexWg.Add(1)
		go func(id int) {
			defer mutexWg.Done()
			mutex.Lock()
			counter++
			fmt.Printf("   Goroutine %d: counter = %d\n", id, counter)
			mutex.Unlock()
		}(i)
	}
	mutexWg.Wait()
	fmt.Printf("   Final counter: %d\n\n", counter)

	// Timeout with select
	fmt.Println("10. Timeout with Select:")
	timeoutCh := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		timeoutCh <- "result"
	}()

	select {
	case res := <-timeoutCh:
		fmt.Printf("   Received: %s\n", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("   Timeout!")
	}
}
