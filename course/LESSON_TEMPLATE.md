# 📘 Lesson Template: [Lesson Title]

This template shows the enhanced structure for all GO-PRO lessons. Use this as a guide when creating or updating lessons.

## Standard Lesson Structure

### 1. Header and Introduction
```markdown
# 📘 Lesson X: [Title]

Welcome to Lesson X! [Brief introduction to the topic and its importance]
```

### 2. Learning Objectives
```markdown
## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- [Specific, measurable objective 1]
- [Specific, measurable objective 2]
- [Specific, measurable objective 3]
- [Continue with 4-7 objectives]
```

### 3. Theory Section
```markdown
## 📚 Theory

### [Main Concept 1]

[Explanation with code examples]

```go
// Code example
func example() {
    // Implementation
}
```

### [Main Concept 2]

[Explanation with code examples]
```

### 4. Hands-On Examples
```markdown
## 💻 Hands-On Examples

Navigate to `../../code/lesson-XX/` to see and run these examples.

### Example 1: [Example Name]
```go
// Complete, runnable code example
package main

import "fmt"

func main() {
    // Example implementation
}
```

### Example 2: [Example Name]
[Continue with 2-4 examples]
```

### 5. Real-World Applications ✨ NEW
```markdown
## 🎯 Real-World Applications

### How This is Used in GO-PRO Backend

[Explain how this concept is used in the actual backend code]

**[Specific Pattern/Feature]:**
```go
// Real code from backend/internal/...
func (s *service) Method(ctx context.Context, req *Request) (*Response, error) {
    // Actual implementation pattern
}
```

**[Another Pattern]:**
```go
// Another real example
```

[Link to actual backend files]
```

### 6. Security Considerations ✨ NEW
```markdown
## 🔒 Security Considerations

**[Security Topic 1]:**
```go
// ❌ Bad: Insecure pattern
func insecureExample() {
    // What not to do
}

// ✅ Good: Secure pattern
func secureExample() {
    // Correct implementation
}
```

**[Security Topic 2]:**
- Security best practice 1
- Security best practice 2
- Common vulnerability to avoid

**Key Security Points:**
- Always validate inputs
- Use proper authentication
- Prevent common vulnerabilities
```

### 7. Performance Tips ✨ NEW
```markdown
## ⚡ Performance Tips

**[Performance Topic 1]:**
```go
// ❌ Less efficient
func slowVersion() {
    // Inefficient implementation
}

// ✅ More efficient
func fastVersion() {
    // Optimized implementation
}
```

**[Performance Topic 2]:**
- Optimization tip 1
- Optimization tip 2
- When to optimize vs when not to

**Benchmarking:**
```go
func BenchmarkOperation(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Operation to benchmark
    }
}
```
```

### 8. Observability Insights ✨ NEW
```markdown
## 📊 Observability Insights

**Tracing:**
```go
import "go.opentelemetry.io/otel"

func (s *service) Operation(ctx context.Context) error {
    ctx, span := otel.Tracer("service").Start(ctx, "Operation")
    defer span.End()
    
    // Add attributes
    span.SetAttributes(
        attribute.String("key", "value"),
    )
    
    return nil
}
```

**Metrics:**
```go
var (
    operationCounter metric.Int64Counter
    operationDuration metric.Float64Histogram
)

func recordMetrics(ctx context.Context, duration time.Duration) {
    operationCounter.Add(ctx, 1)
    operationDuration.Record(ctx, duration.Seconds())
}
```

**Structured Logging:**
```go
import "log/slog"

func logOperation(ctx context.Context, id string) {
    slog.InfoContext(ctx, "operation started",
        "operation_id", id,
        "timestamp", time.Now(),
    )
}
```
```

### 9. Advanced Testing ✨ NEW
```markdown
## 🧪 Advanced Testing

**Table-Driven Tests:**
```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name    string
        input   InputType
        want    OutputType
        wantErr bool
    }{
        {"case 1", input1, output1, false},
        {"case 2", input2, output2, false},
        {"error case", input3, nil, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Function(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
```

**Benchmarks:**
```go
func BenchmarkFunction(b *testing.B) {
    // Setup
    data := setupTestData()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Function(data)
    }
}
```
```

### 10. Exercises
```markdown
## 🧪 Exercises

Complete the exercises in `../../code/lesson-XX/exercises/`:

1. **[Exercise 1 Name]**: [Brief description]
2. **[Exercise 2 Name]**: [Brief description]
3. **[Exercise 3 Name]**: [Brief description]
4. **[Exercise 4 Name]**: [Brief description]
5. **[Exercise 5 Name]**: [Brief description]
6. **[Exercise 6 Name]**: [Brief description]
```

### 11. Validation
```markdown
## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-XX
go test -v ./exercises/...

# Run with coverage
go test -v -cover ./exercises/...

# Run benchmarks
go test -bench=. ./exercises/...
```

All tests should pass before moving to the next lesson.
```

### 12. Key Takeaways
```markdown
## 🔍 Key Takeaways

- [Key point 1 with brief explanation]
- [Key point 2 with brief explanation]
- [Key point 3 with brief explanation]
- [Key point 4 with brief explanation]
- [Key point 5 with brief explanation]
```

### 13. Additional Resources
```markdown
## 📖 Additional Resources

- [Official Go Documentation Link](https://go.dev/doc/)
- [Go Blog Article](https://go.dev/blog/)
- [Effective Go Section](https://go.dev/doc/effective_go)
- [GO-PRO Backend Examples](../../backend/internal/...)
- [Related Package Documentation](https://pkg.go.dev/...)
```

### 14. Key Takeaways Summary ✨ NEW
```markdown
## 🎓 Key Takeaways Summary

✅ **[Topic 1]**: Brief summary  
✅ **[Topic 2]**: Brief summary  
✅ **[Topic 3]**: Brief summary  
✅ **[Topic 4]**: Brief summary  
✅ **[Topic 5]**: Brief summary  
```

### 15. Next Steps
```markdown
## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson X+1: Next Lesson Title](../lesson-XX/README.md)**

---

**[Motivational closing]** 🚀

*[Helpful reminder or tip]*
```

## Section Guidelines

### Real-World Applications
- Show actual code from the GO-PRO backend
- Explain how the concept is used in production
- Link to specific files in the backend
- Demonstrate practical patterns

### Security Considerations
- Always include ❌ Bad and ✅ Good examples
- Cover common vulnerabilities
- Provide security best practices
- Link to security resources

### Performance Tips
- Show performance comparisons
- Include benchmark examples
- Explain when to optimize
- Provide profiling guidance

### Observability Insights
- Include OpenTelemetry examples
- Show structured logging patterns
- Demonstrate metrics collection
- Explain tracing strategies

### Advanced Testing
- Use table-driven test patterns
- Include benchmark examples
- Show test coverage strategies
- Demonstrate integration tests

## Code Example Standards

### All Code Examples Should:
- Be complete and runnable
- Include proper imports
- Follow Go conventions
- Include comments
- Handle errors properly
- Be production-ready

### Example Format:
```go
package main

import (
    "context"
    "fmt"
    "log"
)

// Function demonstrates [concept]
func ExampleFunction(ctx context.Context, input string) (string, error) {
    // Validate input
    if input == "" {
        return "", fmt.Errorf("input cannot be empty")
    }
    
    // Process
    result := processInput(input)
    
    // Return
    return result, nil
}

func main() {
    ctx := context.Background()
    result, err := ExampleFunction(ctx, "example")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
}
```

## Lesson Metadata

Each lesson should have:
- **Duration**: Estimated time in hours
- **Difficulty**: Beginner/Intermediate/Advanced/Expert
- **Prerequisites**: Previous lessons required
- **Exercises**: Number of practice problems
- **Projects**: Related projects (if any)

---

**Use this template to maintain consistency across all lessons!** 📚

