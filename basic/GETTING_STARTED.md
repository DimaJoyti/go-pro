# Getting Started with Go Basics

Welcome to the upgraded Go Basics learning environment! This guide will help you get started quickly.

## 🚀 Quick Start (30 seconds)

```bash
cd basic
go run cmd/runner/main.go
```

This launches an interactive menu where you can explore all examples!

## 📚 What's Available

### ✅ Implemented Examples (Ready to Run)

1. **Hello World** (`examples/01_hello/`)
   - Basic program structure
   - Print functions
   - String formatting

2. **Variables** (`examples/02_variables/`)
   - Variable declarations
   - Data types
   - Constants and iota
   - Type conversion
   - Scope

3. **Functions** (`examples/03_functions/`)
   - Basic functions
   - Multiple return values
   - Named returns
   - Variadic functions
   - Closures
   - Higher-order functions
   - Defer

4. **Concurrency** (`examples/11_concurrency/`)
   - Goroutines
   - WaitGroups
   - Channels
   - Buffered channels
   - Select statement
   - Worker pools

### ✅ Implemented Projects

1. **Order System** (`projects/order_system/`)
   - Complete order management system
   - Status transitions
   - Error handling
   - Struct methods
   - Real-world example

### 🔨 Coming Soon

- Pointers (04)
- Arrays & Slices (05)
- Control Flow (06)
- Maps (07)
- Structs (08)
- Interfaces (09)
- Errors (10)
- Advanced (12)
- Calculator Project
- Todo List Project

## 📖 Learning Paths

### Path 1: Complete Beginner

```bash
# Start here if you're new to Go
cd basic

# 1. Hello World
go run examples/01_hello/main.go

# 2. Variables
go run examples/02_variables/main.go

# 3. Functions
go run examples/03_functions/main.go

# 4. Try the interactive runner
go run cmd/runner/main.go
```

### Path 2: Intermediate Developer

```bash
# If you know basics, jump to advanced topics
cd basic

# Concurrency
go run examples/11_concurrency/main.go

# Real project
go run projects/order_system/main.go

# Interactive exploration
go run cmd/runner/main.go
```

### Path 3: Hands-On Learner

```bash
# Use the interactive runner
cd basic
go run cmd/runner/main.go

# Then modify examples and re-run them
# Edit examples/01_hello/main.go
# Run again to see your changes
```

## 🎯 Interactive Runner Features

The interactive runner (`cmd/runner/main.go`) provides:

- **📖 Browse Examples** - See all 12 categories
- **🎯 Run Projects** - Execute mini-projects
- **🧪 Run Tests** - Test your code
- **📚 Help System** - Built-in documentation
- **🎨 Clean UI** - Beautiful terminal interface

### Menu Options

```
[1-12]  Run specific example
[p1-p3] Run specific project
[a]     Run all examples
[t]     Run tests
[h]     Show help
[q]     Quit
```

## 📂 Directory Structure

```
basic/
├── cmd/
│   └── runner/              # Interactive runner ✅
│       └── main.go
├── examples/
│   ├── 01_hello/           # Hello World ✅
│   ├── 02_variables/       # Variables ✅
│   ├── 03_functions/       # Functions ✅
│   ├── 04_pointers/        # Coming soon
│   ├── 05_arrays_slices/   # Coming soon
│   ├── 06_control_flow/    # Coming soon
│   ├── 07_maps/            # Coming soon
│   ├── 08_structs/         # Coming soon
│   ├── 09_interfaces/      # Coming soon
│   ├── 10_errors/          # Coming soon
│   ├── 11_concurrency/     # Concurrency ✅
│   └── 12_advanced/        # Coming soon
├── projects/
│   ├── calculator/         # Coming soon
│   ├── todo_list/          # Coming soon
│   └── order_system/       # Order System ✅
├── exercises/              # Practice exercises (coming soon)
├── tests/                  # Test examples (coming soon)
└── _legacy/                # Old files (preserved for reference)
```

## 🧪 Testing

### Test Everything

```bash
cd basic
./test-upgraded.sh
```

### Test Specific Example

```bash
cd examples/01_hello
go run main.go
```

### Build All

```bash
cd basic
go build ./...
```

## 💡 Tips for Learning

1. **Read the Code** - Each example is heavily commented
2. **Run Examples** - See the output
3. **Modify Code** - Change values and re-run
4. **Read READMEs** - Each example has a README.md
5. **Use the Runner** - Interactive exploration is fun!

## 🔧 Development

### Prerequisites

- Go 1.21 or higher
- Terminal/Command Prompt

### Setup

```bash
cd basic
go mod download
```

### Format Code

```bash
go fmt ./...
```

### Run Linter (if installed)

```bash
golangci-lint run
```

## 📝 Example Workflow

```bash
# 1. Start with interactive runner
cd basic
go run cmd/runner/main.go

# 2. Select an example (e.g., press "1" for Hello World)

# 3. After running, read the code
cat examples/01_hello/main.go

# 4. Read the documentation
cat examples/01_hello/README.md

# 5. Modify the example
# Edit examples/01_hello/main.go in your editor

# 6. Run again to see changes
go run examples/01_hello/main.go

# 7. Move to next example
go run cmd/runner/main.go
```

## 🎓 Learning Objectives

By working through these examples, you'll learn:

- ✅ Go syntax and program structure
- ✅ Variables, types, and constants
- ✅ Functions and closures
- ✅ Concurrency with goroutines
- ✅ Error handling patterns
- ✅ Real-world project structure
- 🔨 Pointers and memory (coming soon)
- 🔨 Data structures (coming soon)
- 🔨 Interfaces and polymorphism (coming soon)

## 🆘 Getting Help

### In the Interactive Runner

Press `h` for help

### Documentation

- Each example has a `README.md`
- Check `UPGRADE_SUMMARY.md` for details
- See `_legacy/` for old examples

### Common Issues

**Q: Example not found?**
A: Some examples are not yet implemented. Check the test output.

**Q: Build errors?**
A: Run `go mod tidy` in the basic directory.

**Q: Can't find old files?**
A: They're in `_legacy/` directory.

## 🎉 Next Steps

1. ✅ Run the interactive runner
2. ✅ Complete Hello World example
3. ✅ Work through Variables
4. ✅ Study Functions
5. ✅ Explore Concurrency
6. ✅ Build the Order System project
7. 🔨 Wait for more examples (or contribute!)

## 🤝 Contributing

Want to add more examples? The structure is ready:

1. Create `examples/XX_topic/main.go`
2. Add `examples/XX_topic/README.md`
3. Update `cmd/runner/main.go` to include it
4. Test with `./test-upgraded.sh`

Happy Learning! 🚀

