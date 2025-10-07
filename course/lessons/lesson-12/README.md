# 📘 Lesson 12: Testing and Benchmarking

Welcome to Lesson 12! Testing is crucial for reliable software. This lesson covers Go's testing framework, best practices, and performance benchmarking.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Write comprehensive unit tests
- Use table-driven tests effectively
- Create benchmarks for performance testing
- Apply test-driven development (TDD)
- Mock dependencies for isolated testing
- Measure test coverage
- Write integration and end-to-end tests

## 📚 Theory

### Basic Testing

Go's built-in testing framework:

```go
// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

### Table-Driven Tests

Test multiple scenarios efficiently:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -2, -3},
        {"zero", 0, 5, 5},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}
```

### Benchmarking

Measure performance:

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = "hello" + "world"
    }
}
```

### Test Helpers

Create reusable test utilities:

```go
func assertEqual(t *testing.T, got, want interface{}) {
    t.Helper()
    if got != want {
        t.Errorf("got %v, want %v", got, want)
    }
}
```

## 💻 Hands-On Examples

### Example 1: Comprehensive Testing
```go
func TestCalculator(t *testing.T) {
    calc := NewCalculator()
    
    t.Run("Addition", func(t *testing.T) {
        result := calc.Add(5, 3)
        assertEqual(t, result, 8)
    })
    
    t.Run("Division by zero", func(t *testing.T) {
        _, err := calc.Divide(5, 0)
        if err == nil {
            t.Error("expected error for division by zero")
        }
    })
}
```

### Example 2: Mocking
```go
type MockDatabase struct {
    users map[int]User
}

func (m *MockDatabase) GetUser(id int) (User, error) {
    user, exists := m.users[id]
    if !exists {
        return User{}, errors.New("user not found")
    }
    return user, nil
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-12/exercises/`:

1. **Unit Testing**: Write comprehensive unit tests
2. **Table-Driven Tests**: Implement efficient test scenarios
3. **Benchmarking**: Create performance benchmarks
4. **Test Coverage**: Achieve high test coverage
5. **Mocking**: Test with mock dependencies
6. **Integration Testing**: Write integration tests

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-12
go test -v ./exercises/...
go test -bench=. ./exercises/...
go test -cover ./exercises/...
```

## 🔍 Key Takeaways

- Testing is built into Go's toolchain
- Table-driven tests handle multiple scenarios
- Benchmarks measure performance objectively
- Test helpers reduce code duplication
- Mocking enables isolated unit testing
- High test coverage indicates code quality
- Integration tests verify system behavior

## 📊 Testing Commands

```bash
go test                    # Run tests
go test -v                 # Verbose output
go test -run TestAdd       # Run specific test
go test -bench=.           # Run benchmarks
go test -cover             # Show coverage
go test -race              # Race condition detection
```

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 15: Microservices Architecture](../lesson-15/README.md)**

---

**Test everything!** 🚀
