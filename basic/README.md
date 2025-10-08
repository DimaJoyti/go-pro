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
│   ├── 01. Hello Go/       # Hello World
│   ├── 02. Variables/      # Variables & Constants
│   ├── 03. Functions/      # Functions
│   ├── 04. Pointers/       # Pointers
│   ├── 05. Arrays and Slices/  # Arrays & Slices
│   ├── 06. Control Flow/   # Control Flow
│   ├── 07. Maps/           # Maps
│   ├── 08. Structs/        # Structs
│   ├── 09. Interfaces/     # Interfaces
│   ├── 10. Errors/         # Error Handling
│   ├── 11. Goroutines/     # Goroutines & Channels
│   ├── 12. File IO/        # File I/O Operations (NEW!)
│   ├── 13. Testing/        # Testing & Benchmarks (NEW!)
│   ├── prime_numbers.go    # Concurrent prime finder (NEW!)
│   ├── word_counter.go     # Word frequency counter (NEW!)
│   ├── json_parser.go      # JSON handling (NEW!)
│   ├── rate_limiter.go     # Rate limiting (NEW!)
│   ├── cache.go            # In-memory cache (NEW!)
│   ├── queue.go            # FIFO queue (NEW!)
│   ├── linked_list.go      # Linked list (NEW!)
│   ├── binary_search.go    # Binary search (NEW!)
│   ├── producer_consumer.go # Producer-consumer pattern (NEW!)
│   ├── context_timeout.go  # Context usage (NEW!)
│   ├── merge_sort.go       # Concurrent merge sort (NEW!)
│   └── ... (other examples)
├── projects/               # Mini projects
│   ├── calculator/         # Calculator app
│   ├── todo_list/          # Todo list manager
│   └── order_system/       # Order management
├── exercises/              # Practice exercises
│   ├── 01_basics/          # FizzBuzz, Reverse String (NEW!)
│   ├── 02_intermediate/    # URL Shortener (NEW!)
│   └── 03_advanced/        # Web Crawler (NEW!)
└── tests/                  # Test examples
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

### Advanced (Examples 9-13)
9. **Interfaces** - Polymorphism and abstraction
10. **Errors** - Error handling patterns
11. **Concurrency** - Goroutines, channels, patterns
12. **File I/O** - File operations and directory management (NEW!)
13. **Testing** - Unit tests, table-driven tests, benchmarks (NEW!)

## 🆕 New Exercises (24 Total!)

### Standalone Examples (11 files)
- **prime_numbers.go** - Concurrent prime number finder
- **word_counter.go** - Word frequency analysis
- **json_parser.go** - JSON encoding/decoding
- **rate_limiter.go** - Rate limiting implementation
- **cache.go** - In-memory cache with expiration
- **queue.go** - FIFO queue data structure
- **linked_list.go** - Linked list implementation
- **binary_search.go** - Binary search algorithms
- **producer_consumer.go** - Concurrency patterns
- **context_timeout.go** - Context package usage
- **merge_sort.go** - Concurrent sorting

### Structured Directories
- **12. File IO/** - 6 file operation exercises
- **13. Testing/** - 3 testing examples with benchmarks

### Practice Exercises
- **01_basics/** - FizzBuzz, Reverse String
- **02_intermediate/** - URL Shortener
- **03_advanced/** - Concurrent Web Crawler

📖 **See [NEW_EXERCISES.md](examples/NEW_EXERCISES.md) for detailed documentation**

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

