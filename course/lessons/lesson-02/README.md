# 📘 Lesson 2: Variables, Constants, and Functions

Welcome to Lesson 2! Now that you understand Go's basic types, let's dive deeper into variables, constants, and functions - the building blocks of Go programs.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Declare variables using different methods (`var`, `:=`, multiple declarations)
- Understand variable scope and visibility rules
- Create and use functions with various parameter and return patterns
- Work with multiple return values and named returns
- Implement variadic functions (functions with variable arguments)
- Apply best practices for variable and function naming

## 📚 Theory

### Variable Declarations

Go provides several ways to declare variables:

#### **1. Explicit Declaration with `var`**
```go
var name string = "Go"
var age int = 13
var isActive bool = true

// Type inference
var language = "Go"        // string inferred
var year = 2009           // int inferred
```

#### **2. Short Variable Declaration with `:=`**
```go
// Only inside functions
name := "Go"
age := 13
isActive := true
```

#### **3. Multiple Variable Declarations**
```go
// Multiple variables of same type
var x, y, z int

// Multiple variables with initialization
var name, language string = "Go", "Programming"

// Multiple variables with different types
var (
    name     string = "Go"
    version  float64 = 1.21
    isActive bool = true
)

// Short declaration for multiple variables
name, age := "Alice", 25
```

### Variable Scope

Variables have different scopes in Go:

#### **Package Level (Global)**
```go
package main

var globalVar = "I'm global"  // Accessible throughout the package

func main() {
    // Can access globalVar here
}
```

#### **Function Level (Local)**
```go
func example() {
    localVar := "I'm local"  // Only accessible within this function
    // localVar is not accessible outside this function
}
```

#### **Block Level**
```go
func example() {
    if true {
        blockVar := "I'm in a block"  // Only accessible within this block
    }
    // blockVar is not accessible here
}
```

### Functions

Functions are first-class citizens in Go and follow this syntax:

```go
func functionName(parameter1 type1, parameter2 type2) returnType {
    // function body
    return value
}
```

#### **Basic Function Examples**
```go
// Simple function
func greet(name string) string {
    return "Hello, " + name
}

// Function with multiple parameters
func add(a, b int) int {
    return a + b
}

// Function with multiple parameters of same type (shorthand)
func multiply(a, b, c int) int {
    return a * b * c
}
```

#### **Multiple Return Values**
```go
// Function returning multiple values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Using the function
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

#### **Named Return Values**
```go
// Named return values
func rectangle(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = 2 * (length + width)
    return  // naked return - returns named values
}
```

#### **Variadic Functions**
```go
// Function that accepts variable number of arguments
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Usage
result1 := sum(1, 2, 3)        // 6
result2 := sum(1, 2, 3, 4, 5)  // 15
```

### Function Types and Variables

Functions can be assigned to variables and passed as parameters:

```go
// Function type
type Calculator func(int, int) int

// Function as variable
var add Calculator = func(a, b int) int {
    return a + b
}

// Anonymous function
multiply := func(a, b int) int {
    return a * b
}
```

## 💻 Hands-On Examples

Navigate to `../../code/lesson-02/` to see and run these examples.

### Example 1: Variable Declarations
```go
func demonstrateVariables() {
    // Different ways to declare variables
    var name string = "Go Programming"
    var version = 1.21  // Type inference
    
    // Short declaration
    language := "Go"
    year := 2009
    
    // Multiple declarations
    var x, y, z int = 1, 2, 3
    a, b := "Hello", "World"
    
    fmt.Printf("Name: %s, Version: %.2f\n", name, version)
    fmt.Printf("Language: %s, Year: %d\n", language, year)
    fmt.Printf("Numbers: %d, %d, %d\n", x, y, z)
    fmt.Printf("Words: %s %s\n", a, b)
}
```

### Example 2: Functions with Different Patterns
```go
// Simple function
func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

// Multiple return values
func divmod(a, b int) (int, int) {
    return a / b, a % b
}

// Named returns
func circleStats(radius float64) (area, circumference float64) {
    const pi = 3.14159
    area = pi * radius * radius
    circumference = 2 * pi * radius
    return
}

// Variadic function
func average(numbers ...float64) float64 {
    if len(numbers) == 0 {
        return 0
    }
    
    sum := 0.0
    for _, num := range numbers {
        sum += num
    }
    return sum / float64(len(numbers))
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-02/exercises/`:

1. **Variable Practice**: Work with different variable declaration methods
2. **Function Challenges**: Implement various function patterns
3. **Scope Understanding**: Practice with variable scope rules
4. **Error Handling**: Functions that return values and errors
5. **Variadic Functions**: Functions with variable arguments

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-02
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Go provides multiple ways to declare variables, each with its use case
- Short declaration (`:=`) can only be used inside functions
- Functions can return multiple values, making error handling elegant
- Named return values can improve code readability
- Variadic functions provide flexibility for variable arguments
- Variable scope determines where variables can be accessed

## 🎯 Real-World Applications

### How This is Used in GO-PRO Backend

The GO-PRO backend demonstrates professional function and variable patterns:

**Service Layer Functions:**
```go
// Multiple return values for error handling
func (s *lessonService) CreateLesson(ctx context.Context, req *domain.CreateLessonRequest) (*domain.Lesson, error) {
    // Validate request
    if err := s.validator.Validate(req); err != nil {
        return nil, errors.NewValidationError("invalid lesson data", err)
    }

    // Create lesson entity
    lesson := &domain.Lesson{
        ID:          generateID("lesson"),
        CourseID:    req.CourseID,
        Title:       strings.TrimSpace(req.Title),
        Description: strings.TrimSpace(req.Description),
        Content:     req.Content,
    }

    return lesson, s.repo.Create(ctx, lesson)
}
```

**Constructor Functions:**
```go
// Constructor pattern with dependency injection
func NewLessonService(repo repository.LessonRepository, config *Config) LessonService {
    return &lessonService{
        repo:      repo,
        logger:    config.Logger,
        validator: config.Validator,
        cache:     config.Cache,
    }
}
```

**Variadic Options Pattern:**
```go
// Flexible configuration with variadic functions
type ServerOption func(*Server)

func WithPort(port int) ServerOption {
    return func(s *Server) {
        s.port = port
    }
}

func NewServer(opts ...ServerOption) *Server {
    s := &Server{port: 8080} // default
    for _, opt := range opts {
        opt(s)
    }
    return s
}
```

## 🔒 Security Considerations

**Input Validation in Functions:**
```go
func validateUserInput(input string) error {
    // Always validate function parameters
    if len(input) == 0 {
        return errors.New("input cannot be empty")
    }
    if len(input) > 1000 {
        return errors.New("input too long")
    }
    return nil
}
```

**Avoid Global Mutable State:**
```go
// ❌ Bad: Global mutable variable
var userCount int

// ✅ Good: Pass state explicitly
func incrementUserCount(count int) int {
    return count + 1
}
```

**Secure Function Design:**
- Always validate inputs at function boundaries
- Return errors instead of panicking
- Use context for timeout and cancellation
- Avoid exposing internal state through return values

## ⚡ Performance Tips

**Function Inlining:**
```go
// Small functions may be inlined by the compiler
func add(a, b int) int {
    return a + b  // Likely to be inlined
}
```

**Avoid Unnecessary Allocations:**
```go
// ❌ Creates new slice on each call
func processData(data []int) []int {
    result := make([]int, len(data))
    // ...
    return result
}

// ✅ Reuse existing slice
func processDataInPlace(data []int) {
    // Modify data in place
}
```

**Named Returns for Clarity:**
```go
// Named returns can improve readability and performance
func parseConfig(path string) (cfg *Config, err error) {
    cfg = &Config{} // Allocated once
    // ... populate cfg
    return // No additional allocation
}
```

## 📊 Observability Insights

**Function Instrumentation:**
```go
import (
    "context"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
)

func (s *service) ProcessData(ctx context.Context, data string) error {
    // Create span for tracing
    ctx, span := otel.Tracer("service").Start(ctx, "ProcessData")
    defer span.End()

    // Add attributes for observability
    span.SetAttributes(
        attribute.String("data.length", fmt.Sprintf("%d", len(data))),
    )

    // Your logic here
    return nil
}
```

**Structured Logging in Functions:**
```go
import "log/slog"

func createUser(ctx context.Context, name string) error {
    slog.InfoContext(ctx, "creating user",
        "function", "createUser",
        "user_name", name,
    )

    // ... implementation

    slog.InfoContext(ctx, "user created successfully",
        "function", "createUser",
        "user_name", name,
    )
    return nil
}
```

## 🧪 Advanced Testing

**Table-Driven Tests for Functions:**
```go
func TestDivide(t *testing.T) {
    tests := []struct {
        name      string
        a, b      float64
        want      float64
        wantError bool
    }{
        {"normal division", 10, 2, 5, false},
        {"division by zero", 10, 0, 0, true},
        {"negative numbers", -10, 2, -5, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := divide(tt.a, tt.b)
            if (err != nil) != tt.wantError {
                t.Errorf("divide() error = %v, wantError %v", err, tt.wantError)
                return
            }
            if got != tt.want {
                t.Errorf("divide() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

**Benchmarking Functions:**
```go
func BenchmarkSum(b *testing.B) {
    numbers := []int{1, 2, 3, 4, 5}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        sum(numbers...)
    }
}
```

## 📖 Additional Resources

- [Go Tour - Variables](https://tour.golang.org/basics/8)
- [Go Tour - Functions](https://tour.golang.org/basics/4)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Go Code Review Comments - Function Names](https://github.com/golang/go/wiki/CodeReviewComments#function-names)
- [GO-PRO Service Layer Examples](../../backend/internal/service/)

## 🎓 Key Takeaways Summary

✅ **Multiple Declarations**: Use the right declaration style for the context
✅ **Error Handling**: Always return errors, never panic in libraries
✅ **Named Returns**: Use for clarity, but avoid in complex functions
✅ **Context First**: Always pass context.Context as first parameter
✅ **Dependency Injection**: Use constructor functions for clean architecture

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 3: Control Structures and Loops](../lesson-03/README.md)**

---

**Keep practicing!** 🚀

*Remember: Functions are the building blocks of Go programs. Master them well!*
