# 📘 Lesson 8: Error Handling Patterns

Welcome to Lesson 8! Error handling is fundamental to writing robust Go applications. This lesson covers Go's error handling philosophy, patterns, and best practices.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Understand Go's error handling philosophy
- Create and handle custom errors
- Use error wrapping and unwrapping
- Implement error handling patterns
- Design error-resilient APIs
- Apply panic and recover appropriately
- Build comprehensive error handling strategies

## 📚 Theory

### Go's Error Philosophy

Go treats errors as values, not exceptions:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

result, err := divide(10, 0)
if err != nil {
    log.Fatal(err)
}
```

### Custom Errors

Create meaningful error types:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation failed for %s: %s", e.Field, e.Message)
}
```

### Error Wrapping

Add context to errors:

```go
func processFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file %s: %w", filename, err)
    }
    defer file.Close()
    
    // Process file...
    return nil
}
```

### Panic and Recover

Use for unrecoverable errors:

```go
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic occurred: %v", r)
        }
    }()
    
    if b == 0 {
        panic("division by zero")
    }
    
    return a / b, nil
}
```

## 💻 Hands-On Examples

### Example 1: Custom Error Types
```go
type NetworkError struct {
    Op   string
    Addr string
    Err  error
}

func (e *NetworkError) Error() string {
    return fmt.Sprintf("network error during %s to %s: %v", e.Op, e.Addr, e.Err)
}

func (e *NetworkError) Unwrap() error {
    return e.Err
}
```

### Example 2: Error Handling Patterns
```go
func processData(data []byte) error {
    if len(data) == 0 {
        return errors.New("empty data")
    }
    
    if err := validateData(data); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    
    if err := saveData(data); err != nil {
        return fmt.Errorf("save failed: %w", err)
    }
    
    return nil
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-08/exercises/`:

1. **Basic Error Handling**: Create and handle simple errors
2. **Custom Error Types**: Implement meaningful error types
3. **Error Wrapping**: Add context to errors
4. **Error Patterns**: Implement common error handling patterns
5. **Panic and Recover**: Use panic/recover appropriately
6. **Error Testing**: Test error conditions thoroughly

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-08
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Errors are values in Go, not exceptions
- Always check and handle errors explicitly
- Create meaningful custom error types
- Use error wrapping to add context
- Panic only for unrecoverable errors
- Design APIs with clear error contracts
- Test error conditions thoroughly

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 9: Goroutines and Channels](../lesson-09/README.md)**

---

**Handle errors like a pro!** 🚀
