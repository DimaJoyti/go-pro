# 📘 Lesson 5: Pointers and Memory Management

Welcome to Lesson 5! Understanding pointers is crucial for effective Go programming. This lesson covers pointer concepts, memory allocation, and best practices for memory management in Go.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Understand what pointers are and why they're useful
- Declare and use pointer variables
- Understand the difference between value and pointer receivers
- Work with pointer arithmetic (where applicable)
- Manage memory efficiently with proper allocation strategies
- Avoid common pointer-related pitfalls and memory leaks
- Use pointers for efficient function parameter passing

## 📚 Theory

### What are Pointers?

A pointer is a variable that stores the memory address of another variable. Instead of holding a value directly, it "points to" where the value is stored in memory.

```go
var x int = 42
var p *int = &x  // p is a pointer to int, holds address of x

fmt.Println(x)   // 42 (value of x)
fmt.Println(&x)  // 0xc000014098 (address of x)
fmt.Println(p)   // 0xc000014098 (value of p, which is address of x)
fmt.Println(*p)  // 42 (value at address stored in p)
```

### Pointer Operations

**Address Operator (`&`)**: Gets the memory address of a variable
**Dereference Operator (`*`)**: Gets the value at a memory address

```go
var name string = "Go"
var namePtr *string = &name

fmt.Println(name)     // "Go"
fmt.Println(&name)    // 0xc000010230
fmt.Println(namePtr)  // 0xc000010230
fmt.Println(*namePtr) // "Go"

*namePtr = "Golang"   // Modify value through pointer
fmt.Println(name)     // "Golang"
```

### Zero Value of Pointers

The zero value of a pointer is `nil`:

```go
var p *int
fmt.Println(p == nil) // true

// Always check for nil before dereferencing
if p != nil {
    fmt.Println(*p)
}
```

### Pointers vs Values

**Pass by Value** (default in Go):
```go
func modifyValue(x int) {
    x = 100  // Only modifies the copy
}

func main() {
    num := 42
    modifyValue(num)
    fmt.Println(num) // Still 42
}
```

**Pass by Pointer**:
```go
func modifyPointer(x *int) {
    *x = 100  // Modifies the original value
}

func main() {
    num := 42
    modifyPointer(&num)
    fmt.Println(num) // Now 100
}
```

### Memory Allocation

Go provides two ways to allocate memory:

**Using `new()`**:
```go
p := new(int)    // Allocates memory for int, returns *int
*p = 42
fmt.Println(*p)  // 42
```

**Using `make()`** (for slices, maps, channels):
```go
slice := make([]int, 5)    // Creates slice with length 5
m := make(map[string]int)  // Creates empty map
ch := make(chan int)       // Creates channel
```

### Pointer Receivers vs Value Receivers

```go
type Person struct {
    Name string
    Age  int
}

// Value receiver - receives a copy
func (p Person) GetName() string {
    return p.Name
}

// Pointer receiver - receives a pointer to the original
func (p *Person) SetAge(age int) {
    p.Age = age  // Modifies the original
}

func (p *Person) HaveBirthday() {
    p.Age++      // Modifies the original
}
```

## 💻 Hands-On Examples

Navigate to `../../code/lesson-05/` to see and run these examples.

### Example 1: Basic Pointer Operations
```go
func basicPointerOperations() {
    // Declare variables
    x := 42
    y := "Hello"
    
    // Create pointers
    xPtr := &x
    yPtr := &y
    
    fmt.Printf("x = %d, address = %p\n", x, &x)
    fmt.Printf("xPtr = %p, value = %d\n", xPtr, *xPtr)
    
    // Modify through pointer
    *xPtr = 100
    fmt.Printf("After modification: x = %d\n", x)
}
```

### Example 2: Function Parameters
```go
func swap(a, b *int) {
    *a, *b = *b, *a
}

func demonstrateSwap() {
    x, y := 10, 20
    fmt.Printf("Before swap: x=%d, y=%d\n", x, y)
    
    swap(&x, &y)
    fmt.Printf("After swap: x=%d, y=%d\n", x, y)
}
```

### Example 3: Struct Pointers
```go
type Rectangle struct {
    Width, Height float64
}

func (r *Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

func demonstrateStructPointers() {
    rect := &Rectangle{Width: 10, Height: 5}
    fmt.Printf("Area: %.2f\n", rect.Area())
    
    rect.Scale(2)
    fmt.Printf("After scaling: %+v\n", rect)
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-05/exercises/`:

1. **Pointer Basics**: Practice with pointer declaration and dereferencing
2. **Function Parameters**: Implement functions that modify values through pointers
3. **Struct Methods**: Create methods with both value and pointer receivers
4. **Memory Management**: Practice efficient memory allocation patterns
5. **Linked List**: Implement a simple linked list using pointers
6. **Performance Comparison**: Compare performance of value vs pointer passing

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-05
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Pointers store memory addresses, not values directly
- Use `&` to get address, `*` to dereference
- Always check for `nil` before dereferencing pointers
- Pointer receivers allow methods to modify the receiver
- Use pointers for large structs to avoid copying overhead
- Go's garbage collector manages memory automatically
- Prefer value receivers unless you need to modify the receiver

## ⚠️ Common Pitfalls

1. **Dereferencing nil pointers** - Always check for nil
2. **Returning pointers to local variables** - Can cause issues
3. **Unnecessary pointer usage** - Don't use pointers everywhere
4. **Confusing pointer syntax** - Practice makes perfect

## 📖 Additional Resources

- [Go Tour - Pointers](https://tour.golang.org/moretypes/1)
- [Effective Go - Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go Memory Model](https://go.dev/ref/mem)
- [Go Blog - Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 6: Structs and Methods](../lesson-06/README.md)**

---

**Master memory management!** 🚀
