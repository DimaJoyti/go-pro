# 📘 Lesson 4: Arrays, Slices, and Maps

Welcome to Lesson 4! Now that you understand control structures, let's explore Go's collection types - arrays, slices, and maps. These are fundamental data structures you'll use constantly in Go programming.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Understand the difference between arrays and slices
- Create and manipulate arrays with fixed sizes
- Work with slices as dynamic arrays
- Use slice operations like append, copy, and slicing
- Create and manipulate maps for key-value storage
- Iterate over collections using range loops
- Apply best practices for memory-efficient collection usage

## 📚 Theory

### Arrays

Arrays in Go have a fixed size that's part of their type. They're useful when you know the exact number of elements you need.

```go
// Array declaration and initialization
var numbers [5]int                    // Zero-valued array
var fruits = [3]string{"apple", "banana", "orange"}
cities := [...]string{"New York", "London", "Tokyo"} // Size inferred
```

**Key Points:**
- Size is part of the type: `[5]int` and `[10]int` are different types
- Arrays are value types (copied when assigned)
- Size must be known at compile time

### Slices

Slices are dynamic arrays built on top of arrays. They're more flexible and commonly used.

```go
// Slice declaration and initialization
var numbers []int                     // nil slice
var fruits = []string{"apple", "banana", "orange"}
cities := make([]string, 3, 5)       // length 3, capacity 5
```

**Slice Operations:**
```go
// Slicing operations
slice := []int{1, 2, 3, 4, 5}
subSlice := slice[1:4]    // [2, 3, 4]
prefix := slice[:3]       // [1, 2, 3]
suffix := slice[2:]       // [3, 4, 5]

// Appending elements
slice = append(slice, 6, 7, 8)
slice = append(slice, anotherSlice...)
```

### Maps

Maps are Go's built-in hash table implementation for key-value pairs.

```go
// Map declaration and initialization
var ages map[string]int               // nil map
ages = make(map[string]int)           // empty map
scores := map[string]int{             // map literal
    "Alice": 95,
    "Bob":   87,
    "Carol": 92,
}
```

**Map Operations:**
```go
// Adding and accessing elements
ages["David"] = 25
age := ages["Alice"]

// Checking if key exists
age, exists := ages["Eve"]
if !exists {
    fmt.Println("Eve not found")
}

// Deleting elements
delete(ages, "Bob")
```

## 💻 Hands-On Examples

Navigate to `../../code/lesson-04/` to see and run these examples.

### Example 1: Array vs Slice Comparison
```go
func demonstrateArraysVsSlices() {
    // Arrays - fixed size
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Printf("Array: %v, Length: %d\n", arr, len(arr))
    
    // Slices - dynamic
    var slice []int = []int{1, 2, 3}
    slice = append(slice, 4, 5)
    fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", 
               slice, len(slice), cap(slice))
}
```

### Example 2: Slice Manipulation
```go
func sliceOperations() {
    numbers := []int{1, 2, 3, 4, 5}
    
    // Slicing
    fmt.Println("Original:", numbers)
    fmt.Println("Middle:", numbers[1:4])
    fmt.Println("First 3:", numbers[:3])
    fmt.Println("Last 2:", numbers[3:])
    
    // Appending
    numbers = append(numbers, 6, 7)
    fmt.Println("After append:", numbers)
    
    // Copying
    copied := make([]int, len(numbers))
    copy(copied, numbers)
    fmt.Println("Copied:", copied)
}
```

### Example 3: Map Operations
```go
func mapOperations() {
    // Student grades
    grades := map[string]int{
        "Alice":   95,
        "Bob":     87,
        "Charlie": 92,
    }
    
    // Adding new grade
    grades["Diana"] = 98
    
    // Checking and accessing
    if grade, exists := grades["Alice"]; exists {
        fmt.Printf("Alice's grade: %d\n", grade)
    }
    
    // Iterating over map
    for name, grade := range grades {
        fmt.Printf("%s: %d\n", name, grade)
    }
    
    // Deleting
    delete(grades, "Bob")
    fmt.Println("After deletion:", grades)
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-04/exercises/`:

1. **Array Practice**: Work with fixed-size arrays and understand their limitations
2. **Slice Manipulation**: Master slice operations, appending, and slicing
3. **Map Operations**: Implement CRUD operations with maps
4. **Collection Algorithms**: Implement common algorithms like search and sort
5. **Memory Efficiency**: Practice efficient slice and map usage
6. **Real-World Scenario**: Build a simple inventory management system

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-04
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Arrays have fixed size and are value types
- Slices are dynamic and more commonly used than arrays
- Use `make()` to create slices and maps with specific capacity
- Maps provide O(1) average-case lookup time
- Always check if a map key exists when needed
- Use `range` loops for clean iteration over collections
- Be mindful of slice capacity to avoid unnecessary allocations

## 📖 Additional Resources

- [Go Tour - Arrays and Slices](https://tour.golang.org/moretypes/6)
- [Go Blog - Arrays, Slices, and Strings](https://go.dev/blog/slices-intro)
- [Effective Go - Arrays and Slices](https://go.dev/doc/effective_go#arrays)
- [Go Specification - Array and Slice Types](https://go.dev/ref/spec#Array_types)

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 5: Pointers and Memory Management](../lesson-05/README.md)**

---

**Master Go's collections!** 🚀
