# Go Basics - Interactive Learning Environment

A comprehensive, well-organized collection of Go programming examples, exercises, and projects for learning Go from basics to advanced concepts.

## 🚀 Quick Start

### Interactive Runner (Recommended)

Run the interactive menu to explore all examples:

```bash
cd basic
go run cmd/runner/main.go
```

This launches an interactive menu where you can:
- Browse and run all examples
- Execute projects
- Run tests
- Get help and documentation

### Run Individual Examples

```bash
# Navigate to any example
cd examples/01_hello
go run main.go

# Or run from basic directory
cd basic
go run examples/01_hello/main.go
```

## 📁 New Structure

```
basic/
├── cmd/
│   └── runner/              # Interactive example runner
├── examples/                # Organized learning examples
│   ├── 01_hello/           # Hello World
│   ├── 02_variables/       # Variables & Constants
│   ├── 03_functions/       # Functions
│   ├── 04_pointers/        # Pointers
│   ├── 05_arrays_slices/   # Arrays & Slices
│   ├── 06_control_flow/    # Control Flow
│   ├── 07_maps/            # Maps
│   ├── 08_structs/         # Structs
│   ├── 09_interfaces/      # Interfaces
│   ├── 10_errors/          # Error Handling
│   ├── 11_concurrency/     # Goroutines & Channels
│   └── 12_advanced/        # Advanced Topics
├── projects/               # Mini projects
│   ├── calculator/         # Calculator app
│   ├── todo_list/          # Todo list manager
│   └── order_system/       # Order management
├── exercises/              # Practice exercises
│   ├── 01_basics/
│   ├── 02_intermediate/
│   └── 03_advanced/
└── tests/                  # Test examples

Legacy directories (being migrated):
├── 01. Hello Go/           # Original lesson format
├── 02. Variables/
└── ... (other numbered lessons)
```

## 📚 Learning Path

### Beginner (Examples 1-4)
1. **Hello World** - Basic program structure
2. **Variables** - Data types, declarations, constants
3. **Functions** - Function basics, closures, defer
4. **Pointers** - Memory addresses, pointer operations

### Intermediate (Examples 5-8)
5. **Arrays & Slices** - Collections and operations
6. **Control Flow** - If, switch, loops
7. **Maps** - Key-value data structures
8. **Structs** - Custom types and methods

### Advanced (Examples 9-12)
9. **Interfaces** - Polymorphism and abstraction
10. **Errors** - Error handling patterns
11. **Concurrency** - Goroutines, channels, patterns
12. **Advanced** - Advanced Go features

## 🎯 Projects

Real-world mini-projects to practice:

1. **Calculator** - Command-line calculator
2. **Todo List** - Task management system
3. **Order System** - E-commerce order management

Run projects:
```bash
go run projects/order_system/main.go
```

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Run tests for specific example:
```bash
cd examples/03_functions
go test -v
```

## 📖 Documentation

Each example includes:
- `main.go` - Runnable code with comments
- `README.md` - Learning objectives and explanations
- `*_test.go` - Test examples (where applicable)

## 🛠️ Development

### Prerequisites
- Go 1.21 or higher

### Setup
```bash
cd basic
go mod download
```

### Build
```bash
go build ./...
```

### Format Code
```bash
go fmt ./...
```

