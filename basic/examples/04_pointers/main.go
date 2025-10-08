package main

import "fmt"

// Function that modifies value through pointer
func increment(x *int) {
	*x = *x + 1
}

// Function without pointer (pass by value)
func incrementValue(x int) {
	x = x + 1
}

// Swap values using pointers
func swap(a, b *int) {
	*a, *b = *b, *a
}

type Person struct {
	Name string
	Age  int
}

// Method with pointer receiver
func (p *Person) birthday() {
	p.Age++
}

// Method with value receiver
func (p Person) greet() {
	fmt.Printf("   Hello, I'm %s\n", p.Name)
}

func main() {
	fmt.Println("=== Pointers ===\n")

	// Basic pointer usage
	fmt.Println("1. Basic Pointers:")
	x := 42
	p := &x // p is a pointer to x
	fmt.Printf("   Value of x: %d\n", x)
	fmt.Printf("   Address of x: %p\n", p)
	fmt.Printf("   Value at pointer: %d\n\n", *p)

	// Modifying through pointer
	fmt.Println("2. Modifying Through Pointer:")
	*p = 100
	fmt.Printf("   After *p = 100, x = %d\n\n", x)

	// Function with pointer
	fmt.Println("3. Function with Pointer:")
	num := 5
	fmt.Printf("   Before increment: %d\n", num)
	increment(&num)
	fmt.Printf("   After increment: %d\n\n", num)

	// Function without pointer
	fmt.Println("4. Function without Pointer (pass by value):")
	num2 := 5
	fmt.Printf("   Before incrementValue: %d\n", num2)
	incrementValue(num2)
	fmt.Printf("   After incrementValue: %d (unchanged)\n\n", num2)

	// Swap example
	fmt.Println("5. Swap Using Pointers:")
	a, b := 10, 20
	fmt.Printf("   Before swap: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("   After swap: a=%d, b=%d\n\n", a, b)

	// Struct with pointer receiver
	fmt.Println("6. Struct with Pointer Receiver:")
	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("   Before birthday: %s is %d years old\n", person.Name, person.Age)
	person.birthday()
	fmt.Printf("   After birthday: %s is %d years old\n\n", person.Name, person.Age)

	// Nil pointer
	fmt.Println("7. Nil Pointer:")
	var ptr *int
	fmt.Printf("   Nil pointer: %v\n", ptr)
	if ptr == nil {
		fmt.Println("   Pointer is nil")
	}
}
