# 📘 Lesson 3: Control Structures and Loops

Welcome to Lesson 3! Now that you understand variables and functions, let's explore Go's control structures and loops - the tools that control program flow.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Use if/else statements for conditional logic
- Implement switch statements for multi-way branching
- Write various types of for loops (traditional, range, infinite)
- Apply break and continue statements for loop control
- Understand and use defer statements for cleanup
- Combine control structures for complex program flow

## 📚 Theory

### If/Else Statements

Go's if statements are straightforward but powerful:

#### **Basic If Statement**
```go
if condition {
    // code to execute if condition is true
}
```

#### **If/Else**
```go
if condition {
    // code if true
} else {
    // code if false
}
```

#### **If/Else If/Else**
```go
if condition1 {
    // code if condition1 is true
} else if condition2 {
    // code if condition2 is true
} else {
    // code if all conditions are false
}
```

#### **If with Short Statement**
```go
if x := getValue(); x > 0 {
    // x is available in this scope
    fmt.Println("Positive:", x)
}
// x is not available here
```

### Switch Statements

Switch statements provide a clean way to handle multiple conditions:

#### **Basic Switch**
```go
switch value {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

#### **Switch with Expression**
```go
switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
    fmt.Println("Weekend!")
default:
    fmt.Println("Weekday")
}
```

#### **Switch without Expression (Type Switch)**
```go
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
default:
    grade = "F"
}
```

### For Loops

Go has only one loop construct: the for loop, but it's very flexible:

#### **Traditional For Loop**
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

#### **While-style Loop**
```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

#### **Infinite Loop**
```go
for {
    // infinite loop
    if condition {
        break
    }
}
```

#### **Range Loop**
```go
// Over slice
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Over map
scores := map[string]int{"Alice": 95, "Bob": 87}
for name, score := range scores {
    fmt.Printf("%s: %d\n", name, score)
}

// Over string (runes)
for i, char := range "Hello" {
    fmt.Printf("Position %d: %c\n", i, char)
}
```

### Loop Control

#### **Break Statement**
```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Exit the loop
    }
    fmt.Println(i)
}
```

#### **Continue Statement**
```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Println(i) // Only prints odd numbers
}
```

#### **Labeled Break/Continue**
```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer // Break out of both loops
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

### Defer Statement

The defer statement schedules a function call to be run after the function completes:

```go
func example() {
    defer fmt.Println("This runs last")
    fmt.Println("This runs first")
    defer fmt.Println("This runs second to last")
    fmt.Println("This runs second")
}
// Output:
// This runs first
// This runs second
// This runs second to last
// This runs last
```

#### **Common Defer Use Cases**
```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close() // Ensures file is closed when function returns
    
    // Read file operations...
    return nil
}
```

## 💻 Hands-On Examples

Navigate to `../../code/lesson-03/` to see and run these examples.

### Example 1: Conditional Logic
```go
func checkGrade(score int) string {
    if score >= 90 {
        return "A"
    } else if score >= 80 {
        return "B"
    } else if score >= 70 {
        return "C"
    } else if score >= 60 {
        return "D"
    } else {
        return "F"
    }
}
```

### Example 2: Switch Statements
```go
func getDayType(day string) string {
    switch day {
    case "Saturday", "Sunday":
        return "Weekend"
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        return "Weekday"
    default:
        return "Invalid day"
    }
}
```

### Example 3: Various Loop Patterns
```go
func demonstrateLoops() {
    // Traditional for loop
    for i := 0; i < 5; i++ {
        fmt.Printf("Count: %d\n", i)
    }
    
    // Range over slice
    fruits := []string{"apple", "banana", "orange"}
    for index, fruit := range fruits {
        fmt.Printf("%d: %s\n", index, fruit)
    }
    
    // Range over map
    ages := map[string]int{"Alice": 30, "Bob": 25}
    for name, age := range ages {
        fmt.Printf("%s is %d years old\n", name, age)
    }
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-03/exercises/`:

1. **Conditional Logic**: Implement various if/else patterns
2. **Switch Statements**: Practice with different switch variations
3. **Loop Patterns**: Master different types of for loops
4. **Loop Control**: Use break and continue effectively
5. **Defer Practice**: Understand defer statement behavior

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-03
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- If statements can include short variable declarations
- Switch statements don't fall through by default (no break needed)
- For loops are the only loop construct but very flexible
- Range loops provide clean iteration over collections
- Break and continue control loop execution
- Defer statements ensure cleanup code runs

## 📖 Additional Resources

- [Go Tour - Flow Control](https://tour.golang.org/flowcontrol/1)
- [Effective Go - Control Structures](https://go.dev/doc/effective_go#control-structures)
- [Go Specification - Statements](https://go.dev/ref/spec#Statements)

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 4: Arrays, Slices, and Maps](../lesson-04/README.md)**

---

**Keep practicing!** 🚀

*Remember: Control structures are the building blocks of program logic. Master them to write clear, efficient code!*
