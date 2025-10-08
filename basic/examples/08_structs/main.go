package main

import "fmt"

// Basic struct
type Person struct {
	Name string
	Age  int
}

// Struct with embedded struct
type Address struct {
	Street string
	City   string
	Zip    string
}

type Employee struct {
	Person  // Embedded struct
	Address // Embedded struct
	ID      int
	Salary  float64
}

// Method with value receiver
func (p Person) Greet() {
	fmt.Printf("   Hello, I'm %s\n", p.Name)
}

// Method with pointer receiver
func (p *Person) Birthday() {
	p.Age++
}

// Method that returns a value
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Struct with methods
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

func main() {
	fmt.Println("=== Structs ===\n")

	// Create struct
	fmt.Println("1. Create Struct:")
	person1 := Person{Name: "Alice", Age: 25}
	fmt.Printf("   %+v\n\n", person1)

	// Access struct fields
	fmt.Println("2. Access Struct Fields:")
	fmt.Printf("   Name: %s\n", person1.Name)
	fmt.Printf("   Age: %d\n\n", person1.Age)

	// Modify struct fields
	fmt.Println("3. Modify Struct Fields:")
	person1.Age = 26
	fmt.Printf("   Updated age: %d\n\n", person1.Age)

	// Struct with partial initialization
	fmt.Println("4. Partial Initialization:")
	person2 := Person{Name: "Bob"}
	fmt.Printf("   %+v (Age defaults to 0)\n\n", person2)

	// Pointer to struct
	fmt.Println("5. Pointer to Struct:")
	p := &person1
	p.Age = 27 // Go automatically dereferences
	fmt.Printf("   Modified through pointer: %+v\n\n", person1)

	// Anonymous struct
	fmt.Println("6. Anonymous Struct:")
	car := struct {
		Make  string
		Model string
		Year  int
	}{
		Make:  "Toyota",
		Model: "Camry",
		Year:  2023,
	}
	fmt.Printf("   %+v\n\n", car)

	// Embedded struct
	fmt.Println("7. Embedded Struct:")
	emp := Employee{
		Person: Person{Name: "Charlie", Age: 30},
		Address: Address{
			Street: "123 Main St",
			City:   "New York",
			Zip:    "10001",
		},
		ID:     1001,
		Salary: 75000,
	}
	fmt.Printf("   Employee: %s, Age: %d\n", emp.Name, emp.Age) // Can access embedded fields directly
	fmt.Printf("   Address: %s, %s %s\n", emp.Street, emp.City, emp.Zip)
	fmt.Printf("   ID: %d, Salary: $%.2f\n\n", emp.ID, emp.Salary)

	// Methods with value receiver
	fmt.Println("8. Methods with Value Receiver:")
	person1.Greet()
	fmt.Println()

	// Methods with pointer receiver
	fmt.Println("9. Methods with Pointer Receiver:")
	fmt.Printf("   Before birthday: %d years old\n", person1.Age)
	person1.Birthday()
	fmt.Printf("   After birthday: %d years old\n\n", person1.Age)

	// Methods that return values
	fmt.Println("10. Methods that Return Values:")
	if person1.IsAdult() {
		fmt.Printf("   %s is an adult\n\n", person1.Name)
	}

	// Struct with methods (Rectangle)
	fmt.Println("11. Struct with Multiple Methods:")
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Printf("   Rectangle: %+v\n", rect)
	fmt.Printf("   Area: %.2f\n", rect.Area())
	fmt.Printf("   Perimeter: %.2f\n", rect.Perimeter())
}
