# 📘 Lesson 1: Go Syntax and Basic Types

Welcome to your first Go lesson! In this lesson, you'll learn the fundamental syntax of Go and work with basic data types.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Set up a Go development environment
- Understand Go's basic syntax and program structure
- Work with primitive data types (int, float, string, bool)
- Declare and use constants
- Perform type conversions
- Use the `iota` identifier for enumerated constants

## 📚 Theory

### Go Program Structure

Every Go program starts with a package declaration, followed by imports, and then the program code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### Basic Types

Go has several built-in basic types:

#### **Numeric Types**
- **Integers**: `int`, `int8`, `int16`, `int32`, `int64`
- **Unsigned integers**: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- **Floating point**: `float32`, `float64`
- **Complex numbers**: `complex64`, `complex128`

#### **Other Types**
- **Boolean**: `bool` (true or false)
- **String**: `string` (UTF-8 encoded)
- **Byte**: `byte` (alias for uint8)
- **Rune**: `rune` (alias for int32, represents Unicode code points)

### Variable Declarations

```go
// Explicit type declaration
var name string = "Go"
var age int = 13

// Type inference
var language = "Go"
var year = 2009

// Short variable declaration (inside functions only)
version := "1.21"
```

### Constants

Constants are declared with the `const` keyword:

```go
const Pi = 3.14159
const Language = "Go"

// Multiple constants
const (
    StatusOK = 200
    StatusNotFound = 404
    StatusError = 500
)
```

### The `iota` Identifier

`iota` is used to create enumerated constants:

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
    Thursday         // 4
    Friday           // 5
    Saturday         // 6
)
```

### Type Conversions

Go requires explicit type conversions:

```go
var i int = 42
var f float64 = float64(i)  // Convert int to float64
var s string = fmt.Sprintf("%d", i)  // Convert int to string
```

## 💻 Hands-On Examples

Let's look at some practical examples. Navigate to `../../code/lesson-01/` to see and run these examples.

### Example 1: Basic Types Demo
```go
package main

import "fmt"

func main() {
    // Integer types
    var age int = 25
    var population int64 = 7800000000
    
    // Floating point types
    var temperature float32 = 36.5
    var pi float64 = 3.14159265359
    
    // Boolean type
    var isActive bool = true
    
    // String type
    var name string = "Go Programming"
    
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Population: %d\n", population)
    fmt.Printf("Temperature: %.1f°C\n", temperature)
    fmt.Printf("Pi: %.5f\n", pi)
    fmt.Printf("Active: %t\n", isActive)
    fmt.Printf("Language: %s\n", name)
}
```

### Example 2: Constants and iota
```go
package main

import "fmt"

const (
    // HTTP status codes using iota
    StatusContinue = 100 + iota
    StatusSwitchingProtocols
    StatusProcessing
)

const (
    // File permissions
    ReadPermission = 1 << iota  // 1 (binary: 001)
    WritePermission             // 2 (binary: 010)
    ExecutePermission           // 4 (binary: 100)
)

func main() {
    fmt.Printf("HTTP Status Codes:\n")
    fmt.Printf("Continue: %d\n", StatusContinue)
    fmt.Printf("Switching Protocols: %d\n", StatusSwitchingProtocols)
    fmt.Printf("Processing: %d\n", StatusProcessing)
    
    fmt.Printf("\nFile Permissions:\n")
    fmt.Printf("Read: %d\n", ReadPermission)
    fmt.Printf("Write: %d\n", WritePermission)
    fmt.Printf("Execute: %d\n", ExecutePermission)
    fmt.Printf("Read+Write: %d\n", ReadPermission|WritePermission)
}
```

## 🧪 Exercises

Now it's time to practice! Complete the exercises in `../../code/lesson-01/exercises/`:

1. **Basic Types Practice**: Work with different data types
2. **Constants Challenge**: Create meaningful constants using iota
3. **Type Conversion**: Practice converting between types safely
4. **Real-World Scenario**: Build a simple program using all concepts

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-01
go test -v ./exercises/...
```

All tests should pass before moving to the next lesson.

## 🔍 Key Takeaways

- Go is statically typed with type inference capabilities
- Constants are compile-time values that cannot be changed
- `iota` provides an elegant way to create enumerated constants
- Type conversions must be explicit in Go
- Go's type system helps catch errors at compile time

## 🎯 Real-World Applications

### How This is Used in GO-PRO Backend

The GO-PRO backend extensively uses these basic types and constants. Here are real examples:

**Status Constants with iota:**
```go
// From backend error handling
type ErrorCode int

const (
    ErrCodeValidation ErrorCode = iota + 1000
    ErrCodeNotFound
    ErrCodeUnauthorized
    ErrCodeConflict
    ErrCodeInternal
)
```

**Type Safety in Domain Models:**
```go
// Lesson model uses strict typing
type Lesson struct {
    ID          string    `json:"id"`
    CourseID    string    `json:"course_id"`
    Title       string    `json:"title"`
    Order       int       `json:"order"`
    CreatedAt   time.Time `json:"created_at"`
}
```

## 🔒 Security Considerations

**Input Validation:**
- Always validate string inputs for length and content
- Use type constraints to prevent invalid data
- Sanitize user inputs before processing

**Type Safety:**
- Go's strong typing prevents many common vulnerabilities
- Explicit type conversions prevent unexpected behavior
- Constants prevent accidental modification of critical values

## ⚡ Performance Tips

**Choose the Right Type:**
- Use `int` for general integers (platform-optimized)
- Use specific sizes (`int32`, `int64`) only when needed
- Prefer `string` over `[]byte` for immutable text
- Use `const` for compile-time optimization

**Memory Efficiency:**
```go
// Efficient: compile-time constant
const MaxUsers = 1000

// Less efficient: runtime variable
var maxUsers = 1000
```

## 📊 Observability Insights

When working with basic types in production:

**Logging Best Practices:**
```go
import "log/slog"

// Structured logging with types
slog.Info("user created",
    "user_id", userID,        // string
    "age", age,               // int
    "active", isActive,       // bool
    "created_at", time.Now(), // time.Time
)
```

**Metrics and Tracing:**
- Use typed constants for metric names
- Ensure consistent type usage across services
- Type safety helps prevent metric cardinality issues

## 🧪 Advanced Testing

**Testing Type Conversions:**
```go
func TestTypeConversion(t *testing.T) {
    tests := []struct {
        name    string
        input   int
        want    float64
        wantErr bool
    }{
        {"positive", 42, 42.0, false},
        {"zero", 0, 0.0, false},
        {"negative", -10, -10.0, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := float64(tt.input)
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

## 📖 Additional Resources

- [Go Tour - Basics](https://tour.golang.org/basics/1)
- [Go Specification - Types](https://go.dev/ref/spec#Types)
- [Effective Go - Constants](https://go.dev/doc/effective_go#constants)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [GO-PRO Backend Examples](../../backend/internal/domain/models.go)

## 🎓 Key Takeaways Summary

✅ **Type Safety**: Go's strong typing catches errors at compile time
✅ **Constants**: Use `const` for immutable values and `iota` for enumerations
✅ **Explicit Conversions**: All type conversions must be explicit
✅ **Performance**: Choose appropriate types for your use case
✅ **Security**: Type safety is your first line of defense

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 2: Variables, Constants, and Functions](../lesson-02/README.md)**

---

**Happy coding!** 🚀

*Remember: The best way to learn Go is by writing Go code. Don't just read - practice!*
