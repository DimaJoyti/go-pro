# 💻 Lesson 1: Code Examples and Exercises

This directory contains the practical code examples and exercises for Lesson 1: Go Syntax and Basic Types.

## 📁 Directory Structure

```
lesson-01/
├── main.go              # Runnable examples demonstrating all concepts
├── go.mod              # Go module file
├── exercises/          # Practice exercises for students
│   ├── basic_types.go     # Basic types exercises
│   ├── basic_types_test.go # Tests for basic types
│   ├── constants.go       # Constants and iota exercises
│   └── constants_test.go  # Tests for constants
└── solutions/          # Reference solutions
    ├── basic_types_solution.go
    └── constants_solution.go
```

## 🚀 Getting Started

### 1. Run the Examples
```bash
go run main.go
```

This will demonstrate all the concepts covered in Lesson 1:
- Basic data types (int, float, string, bool)
- Constants and their usage
- Type conversions
- iota for enumerated constants
- Bit operations for permissions

### 2. Complete the Exercises

Navigate to the `exercises/` directory and complete the TODO items in:
- `basic_types.go` - Practice with Go's basic data types
- `constants.go` - Work with constants and iota

### 3. Run the Tests

Test your solutions with:
```bash
# Test basic types exercises
go test -v ./exercises -run TestBasicTypes

# Test constants exercises  
go test -v ./exercises -run TestConstants

# Run all tests
go test -v ./exercises/...
```

### 4. Check Solutions

If you get stuck, reference solutions are available in the `solutions/` directory.

## 🎯 Exercise Overview

### Basic Types Exercises
- **PersonInfo struct**: Practice with struct creation and formatting
- **BMI Calculator**: Work with floating-point calculations
- **Temperature Converter**: Implement Celsius/Fahrenheit conversion
- **Age Validator**: Practice with boolean logic and ranges
- **Circle Calculator**: Apply mathematical formulas
- **Number Operations**: Work with integer operations

### Constants Exercises
- **Mathematical Constants**: Define Pi, E, and Golden Ratio
- **HTTP Status Codes**: Use iota for sequential constants
- **Log Levels**: Create enumerated constants starting from 0
- **File Permissions**: Use bit operations with iota
- **Weekdays**: Create constants starting from 1
- **Permission Functions**: Implement bitwise operations

## ✅ Success Criteria

All tests should pass before moving to Lesson 2:

```bash
go test -v ./exercises/...
```

Expected output:
```
=== RUN TestCreatePersonInfo
--- PASS: TestCreatePersonInfo
=== RUN TestFormatPersonInfo  
--- PASS: TestFormatPersonInfo
... (all tests passing)
PASS
```

## 🔍 Key Learning Points

After completing these exercises, you should understand:

1. **Go's Type System**: How Go handles different data types
2. **Constants**: When and how to use constants vs variables
3. **iota**: Creating enumerated constants elegantly
4. **Type Conversions**: Explicit conversion requirements in Go
5. **Bit Operations**: Using bitwise operators for flags
6. **Testing**: How Go's testing framework works

## 📚 Additional Practice

Try these additional challenges:
1. Create your own constants using iota for a different domain (colors, sizes, etc.)
2. Implement additional mathematical functions using the constants
3. Create a simple calculator using the type conversion functions
4. Experiment with different number bases and bit operations

## ➡️ Next Steps

Once all tests pass, proceed to:
**[Lesson 2: Variables, Constants, and Functions](../../lessons/lesson-02/README.md)**

---

**Happy coding!** 🚀
