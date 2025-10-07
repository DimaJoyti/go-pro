# 📘 Lesson 6: Structs and Methods

Welcome to Lesson 6! Now that you understand pointers, let's explore Go's approach to object-oriented programming through structs and methods. This lesson covers struct definition, method implementation, and design patterns.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Define and use custom struct types
- Create struct literals and initialize structs
- Implement methods with both value and pointer receivers
- Understand method sets and receiver types
- Use struct embedding for composition
- Apply struct tags for metadata
- Design clean APIs using structs and methods

## 📚 Theory

### Defining Structs

Structs are Go's way of creating custom types that group related data together:

```go
type Person struct {
    Name    string
    Age     int
    Email   string
    Active  bool
}
```

### Creating and Initializing Structs

```go
// Zero value initialization
var p1 Person

// Struct literal with field names
p2 := Person{
    Name:   "Alice",
    Age:    30,
    Email:  "alice@example.com",
    Active: true,
}

// Struct literal without field names (not recommended)
p3 := Person{"Bob", 25, "bob@example.com", true}

// Pointer to struct
p4 := &Person{Name: "Charlie", Age: 35}
```

### Methods

Methods are functions with a special receiver argument:

```go
// Value receiver - receives a copy
func (p Person) GetFullInfo() string {
    return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// Pointer receiver - receives a pointer to the original
func (p *Person) HaveBirthday() {
    p.Age++
}

func (p *Person) UpdateEmail(newEmail string) {
    p.Email = newEmail
}
```

### Value vs Pointer Receivers

**Use value receivers when:**
- The method doesn't modify the receiver
- The receiver is small (basic types, small structs)
- You want to work with copies

**Use pointer receivers when:**
- The method modifies the receiver
- The receiver is large (to avoid copying)
- You want consistency (if some methods use pointer receivers)

### Struct Embedding

Go supports composition through struct embedding:

```go
type Address struct {
    Street  string
    City    string
    Country string
}

type Employee struct {
    Person          // Embedded struct
    Address         // Embedded struct
    EmployeeID string
    Department string
}

// Can access embedded fields directly
emp := Employee{}
emp.Name = "Alice"        // From Person
emp.Street = "123 Main"   // From Address
emp.EmployeeID = "E001"
```

### Struct Tags

Struct tags provide metadata for fields:

```go
type User struct {
    ID       int    `json:"id" db:"user_id"`
    Name     string `json:"name" db:"full_name"`
    Email    string `json:"email" db:"email_address"`
    Password string `json:"-" db:"password_hash"`
}
```

## 💻 Hands-On Examples

Navigate to `../../code/lesson-06/` to see and run these examples.

### Example 1: Basic Struct and Methods
```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

### Example 2: Struct Embedding
```go
type Animal struct {
    Name    string
    Species string
}

func (a Animal) Speak() string {
    return fmt.Sprintf("%s makes a sound", a.Name)
}

type Dog struct {
    Animal
    Breed string
}

func (d Dog) Speak() string {
    return fmt.Sprintf("%s barks", d.Name)
}

func (d Dog) Fetch() string {
    return fmt.Sprintf("%s fetches the ball", d.Name)
}
```

### Example 3: Method Sets and Interfaces
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-06/exercises/`:

1. **Basic Structs**: Define and work with simple struct types
2. **Methods Implementation**: Create methods with appropriate receivers
3. **Struct Embedding**: Use composition to build complex types
4. **JSON Serialization**: Work with struct tags for JSON marshaling
5. **Design Patterns**: Implement common patterns using structs
6. **Real-World Application**: Build a complete system using structs and methods

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-06
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Structs group related data into custom types
- Methods provide behavior for struct types
- Use pointer receivers for modification or large structs
- Struct embedding enables composition over inheritance
- Struct tags provide metadata for serialization and validation
- Method sets determine interface satisfaction
- Design clean APIs by grouping related functionality

## 📖 Additional Resources

- [Go Tour - Structs](https://tour.golang.org/moretypes/2)
- [Go Tour - Methods](https://tour.golang.org/methods/1)
- [Effective Go - Methods](https://go.dev/doc/effective_go#methods)
- [Go Blog - Methods and Interfaces](https://go.dev/blog/methods-interfaces)

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 7: Interfaces and Polymorphism](../lesson-07/README.md)**

---

**Build with structs and methods!** 🚀
