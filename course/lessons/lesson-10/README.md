# 📘 Lesson 10: Packages and Modules

Welcome to Lesson 10! Understanding Go's package system and modules is essential for building maintainable applications. This lesson covers package organization, modules, and dependency management.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Organize code into packages effectively
- Understand package visibility and naming conventions
- Create and manage Go modules
- Handle dependencies with go.mod and go.sum
- Use internal packages for encapsulation
- Apply package design best practices
- Publish and version modules

## 📚 Theory

### Package Basics

Every Go file belongs to a package:

```go
package main // Executable package

package utils // Library package

// Import packages
import (
    "fmt"
    "strings"
    "github.com/user/repo/pkg/utils"
)
```

### Package Visibility

Capitalized names are exported (public):

```go
package calculator

// Exported function
func Add(a, b int) int {
    return a + b
}

// Unexported function
func multiply(a, b int) int {
    return a * b
}

// Exported type
type Calculator struct {
    Name string // Exported field
    version int // Unexported field
}
```

### Go Modules

Modules are collections of packages:

```bash
# Initialize module
go mod init github.com/user/myproject

# Add dependency
go get github.com/gorilla/mux

# Update dependencies
go mod tidy

# Vendor dependencies
go mod vendor
```

### Module Structure

```
myproject/
├── go.mod
├── go.sum
├── main.go
├── internal/
│   └── auth/
│       └── auth.go
├── pkg/
│   └── utils/
│       └── utils.go
└── cmd/
    └── server/
        └── main.go
```

## 💻 Hands-On Examples

### Example 1: Package Organization
```go
// pkg/calculator/calculator.go
package calculator

type Calculator struct {
    name string
}

func New(name string) *Calculator {
    return &Calculator{name: name}
}

func (c *Calculator) Add(a, b float64) float64 {
    return a + b
}

func (c *Calculator) Multiply(a, b float64) float64 {
    return a * b
}
```

### Example 2: Internal Packages
```go
// internal/config/config.go
package config

type Config struct {
    Port     int
    Database string
}

func Load() *Config {
    return &Config{
        Port:     8080,
        Database: "localhost:5432",
    }
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-10/exercises/`:

1. **Package Creation**: Create well-organized packages
2. **Module Management**: Work with go.mod and dependencies
3. **Internal Packages**: Use internal packages for encapsulation
4. **Package Documentation**: Write effective package documentation
5. **Dependency Management**: Handle external dependencies
6. **Module Publishing**: Prepare modules for publication

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-10
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Packages organize related functionality
- Exported names start with capital letters
- Modules manage dependencies and versioning
- Internal packages provide encapsulation
- Use semantic versioning for modules
- Keep packages focused and cohesive
- Document public APIs thoroughly

## 📖 Best Practices

- One package per directory
- Use descriptive package names
- Avoid circular dependencies
- Keep internal implementation details private
- Use go.mod for dependency management
- Regular dependency updates and security checks

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 12: Testing and Benchmarking](../lesson-12/README.md)**

---

**Organize code like a pro!** 🚀
