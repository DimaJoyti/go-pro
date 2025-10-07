# 📘 Lesson 7: Interfaces and Polymorphism

Welcome to Lesson 7! Interfaces are one of Go's most powerful features, enabling clean abstractions and polymorphic behavior. This lesson covers interface design, implementation, and best practices.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Define and implement interfaces in Go
- Understand implicit interface satisfaction
- Use empty interfaces and type assertions
- Apply interface composition and embedding
- Design clean APIs using interfaces
- Implement common interface patterns
- Use interfaces for testing and mocking

## 📚 Theory

### Interface Basics

Interfaces define method signatures without implementation:

```go
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// Interface composition
type ReadWriter interface {
    Reader
    Writer
}
```

### Implicit Implementation

Go uses implicit interface satisfaction - no explicit "implements" keyword:

```go
type File struct {
    name string
}

func (f File) Write(data []byte) (int, error) {
    // Implementation
    return len(data), nil
}

// File automatically satisfies Writer interface
var w Writer = File{name: "test.txt"}
```

### Empty Interface

The empty interface `interface{}` can hold any type:

```go
func PrintAnything(v interface{}) {
    fmt.Println(v)
}

PrintAnything(42)
PrintAnything("hello")
PrintAnything([]int{1, 2, 3})
```

### Type Assertions

Extract concrete types from interfaces:

```go
var i interface{} = "hello"

// Type assertion
s := i.(string)
fmt.Println(s) // "hello"

// Safe type assertion
s, ok := i.(string)
if ok {
    fmt.Println("String:", s)
}
```

## 💻 Hands-On Examples

### Example 1: Shape Interface
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
```

### Example 2: Interface Composition
```go
type Stringer interface {
    String() string
}

type Validator interface {
    Validate() error
}

type Model interface {
    Stringer
    Validator
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-07/exercises/`:

1. **Basic Interfaces**: Define and implement simple interfaces
2. **Interface Composition**: Combine multiple interfaces
3. **Type Assertions**: Work with empty interfaces and type checking
4. **Polymorphism**: Use interfaces for different implementations
5. **Testing with Interfaces**: Create mockable interfaces
6. **Real-World Application**: Build a plugin system using interfaces

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-07
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Interfaces define behavior, not data
- Implementation is implicit in Go
- Small interfaces are better than large ones
- Use interfaces for abstraction and testing
- Empty interface accepts any type
- Type assertions extract concrete types
- Interface composition enables flexible design

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 8: Error Handling Patterns](../lesson-08/README.md)**

---

**Master Go interfaces!** 🚀
