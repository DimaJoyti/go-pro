package main

import (
	"fmt"
	"math"
)

// Basic interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function that accepts interface
func printShapeInfo(s Shape) {
	fmt.Printf("   Area: %.2f\n", s.Area())
	fmt.Printf("   Perimeter: %.2f\n", s.Perimeter())
}

// Stringer interface (from fmt package)
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

// Multiple interfaces
type Writer interface {
	Write(data string) error
}

type Reader interface {
	Read() (string, error)
}

type ReadWriter interface {
	Reader
	Writer
}

// Implementation
type File struct {
	content string
}

func (f *File) Write(data string) error {
	f.content = data
	return nil
}

func (f *File) Read() (string, error) {
	return f.content, nil
}

// Empty interface
func printAnything(v interface{}) {
	fmt.Printf("   Value: %v, Type: %T\n", v, v)
}

// Type assertion
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("   Integer: %d\n", v)
	case string:
		fmt.Printf("   String: %s\n", v)
	case bool:
		fmt.Printf("   Boolean: %t\n", v)
	default:
		fmt.Printf("   Unknown type: %T\n", v)
	}
}

func main() {
	fmt.Println("=== Interfaces ===\n")

	// Basic interface usage
	fmt.Println("1. Basic Interface:")
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	fmt.Println("   Rectangle:")
	printShapeInfo(rect)
	fmt.Println("   Circle:")
	printShapeInfo(circle)
	fmt.Println()

	// Interface slice
	fmt.Println("2. Interface Slice:")
	shapes := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		Rectangle{Width: 8, Height: 2},
	}

	for i, shape := range shapes {
		fmt.Printf("   Shape %d: Area = %.2f\n", i+1, shape.Area())
	}
	fmt.Println()

	// Stringer interface
	fmt.Println("3. Stringer Interface:")
	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("   %s\n\n", person) // Automatically calls String() method

	// Multiple interfaces
	fmt.Println("4. Multiple Interfaces (ReadWriter):")
	file := &File{}
	file.Write("Hello, Go!")
	content, _ := file.Read()
	fmt.Printf("   File content: %s\n\n", content)

	// Empty interface
	fmt.Println("5. Empty Interface:")
	printAnything(42)
	printAnything("Hello")
	printAnything(true)
	printAnything(3.14)
	fmt.Println()

	// Type assertion
	fmt.Println("6. Type Assertion:")
	var i interface{} = "hello"
	s, ok := i.(string)
	if ok {
		fmt.Printf("   String value: %s\n", s)
	}

	n, ok := i.(int)
	if !ok {
		fmt.Printf("   Not an integer (value: %d, ok: %t)\n\n", n, ok)
	}

	// Type switch
	fmt.Println("7. Type Switch:")
	describe(42)
	describe("hello")
	describe(true)
	describe(3.14)
	fmt.Println()

	// Interface nil check
	fmt.Println("8. Interface Nil Check:")
	var shape Shape
	if shape == nil {
		fmt.Println("   Shape interface is nil")
	}
}
