package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// Generics (Go 1.18+)
func Min[T int | float64](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Generic slice function
func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Reflection
func inspectType(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)
	fmt.Printf("   Type: %v, Kind: %v, Value: %v\n", t, t.Kind(), val)
}

// JSON handling
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Context usage
func doWork(ctx context.Context, id int) {
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("   Worker %d completed\n", id)
	case <-ctx.Done():
		fmt.Printf("   Worker %d cancelled: %v\n", id, ctx.Err())
	}
}

// Method chaining
type Calculator struct {
	value int
}

func (c *Calculator) Add(n int) *Calculator {
	c.value += n
	return c
}

func (c *Calculator) Multiply(n int) *Calculator {
	c.value *= n
	return c
}

func (c *Calculator) Result() int {
	return c.value
}

// Functional options pattern
type Server struct {
	host string
	port int
}

type ServerOption func(*Server)

func WithHost(host string) ServerOption {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func NewServer(opts ...ServerOption) *Server {
	s := &Server{
		host: "localhost",
		port: 8080,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	fmt.Println("=== Advanced Topics ===\n")

	// Generics
	fmt.Println("1. Generics:")
	fmt.Printf("   Min(10, 20) = %d\n", Min(10, 20))
	fmt.Printf("   Min(3.14, 2.71) = %.2f\n", Min(3.14, 2.71))

	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(n int) int { return n * 2 })
	fmt.Printf("   Doubled: %v\n\n", doubled)

	// Reflection
	fmt.Println("2. Reflection:")
	inspectType(42)
	inspectType("hello")
	inspectType(true)
	inspectType([]int{1, 2, 3})
	fmt.Println()

	// JSON encoding/decoding
	fmt.Println("3. JSON Handling:")
	user := User{Name: "Alice", Email: "alice@example.com", Age: 25}

	// Marshal (encode)
	jsonData, _ := json.Marshal(user)
	fmt.Printf("   JSON: %s\n", string(jsonData))

	// Unmarshal (decode)
	var decoded User
	json.Unmarshal(jsonData, &decoded)
	fmt.Printf("   Decoded: %+v\n\n", decoded)

	// Context with timeout
	fmt.Println("4. Context with Timeout:")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go doWork(ctx, 1)
	time.Sleep(150 * time.Millisecond)
	fmt.Println()

	// Context with cancellation
	fmt.Println("5. Context with Cancellation:")
	ctx2, cancel2 := context.WithCancel(context.Background())

	go doWork(ctx2, 2)
	time.Sleep(50 * time.Millisecond)
	cancel2() // Cancel the context
	time.Sleep(50 * time.Millisecond)
	fmt.Println()

	// Method chaining
	fmt.Println("6. Method Chaining:")
	calc := &Calculator{value: 10}
	result := calc.Add(5).Multiply(2).Add(3).Result()
	fmt.Printf("   (10 + 5) * 2 + 3 = %d\n\n", result)

	// Functional options pattern
	fmt.Println("7. Functional Options Pattern:")
	server1 := NewServer()
	fmt.Printf("   Default server: %s:%d\n", server1.host, server1.port)

	server2 := NewServer(WithHost("0.0.0.0"), WithPort(3000))
	fmt.Printf("   Custom server: %s:%d\n\n", server2.host, server2.port)

	// Type aliases
	fmt.Println("8. Type Aliases:")
	type UserID int
	type Username string

	var id UserID = 123
	var name Username = "alice"
	fmt.Printf("   UserID: %d, Username: %s\n\n", id, name)

	// Iota for constants
	fmt.Println("9. Iota for Constants:")
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf("   Sunday=%d, Monday=%d, Friday=%d\n\n", Sunday, Monday, Friday)

	// Variadic functions with spread
	fmt.Println("10. Variadic Functions:")
	nums := []int{1, 2, 3, 4, 5}
	total := sum(nums...)
	fmt.Printf("   Sum of %v = %d\n", nums, total)
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
