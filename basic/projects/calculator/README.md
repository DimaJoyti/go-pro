# Calculator Project

A simple command-line calculator built with Go.

## Features

- ✅ Basic arithmetic operations (Add, Subtract, Multiply, Divide)
- ✅ Power operation
- ✅ Calculation history
- ✅ Interactive menu
- ✅ Error handling (division by zero)
- ✅ Clean, user-friendly interface

## Operations

1. **Addition** - Add two numbers
2. **Subtraction** - Subtract two numbers
3. **Multiplication** - Multiply two numbers
4. **Division** - Divide two numbers (with zero check)
5. **Power** - Raise a number to a power
6. **History** - View calculation history
7. **Clear** - Clear calculation history

## How to Run

```bash
cd basic/projects/calculator
go run main.go
```

## Usage Example

```
╔════════════════════════════════════════════════╗
║                                                ║
║           🧮  Go Calculator  🧮                ║
║                                                ║
╚════════════════════════════════════════════════╝

==================================================
📋 Operations:
==================================================
  [1] Add (+)
  [2] Subtract (-)
  [3] Multiply (×)
  [4] Divide (÷)
  [5] Power (^)
  [h] Show History
  [c] Clear History
  [q] Quit
==================================================

👉 Enter your choice: 1
Enter first number: 10
Enter second number: 5

✓ Result: 15.00
```

## Code Structure

- `Calculator` struct - Main calculator with history tracking
- `Add`, `Subtract`, `Multiply`, `Divide`, `Power` - Arithmetic operations
- `ShowHistory` - Display calculation history
- `ClearHistory` - Clear history
- Interactive menu system

## Learning Objectives

- Struct methods
- Error handling
- User input handling
- String formatting
- Slice operations
- Interactive CLI applications

