package main

import "fmt"

func main() {
	fmt.Println("=== Variables and Constants ===\n")

	// Variable declaration with var keyword
	var name string = "John"
	var age int = 30
	var isActive bool = true

	fmt.Println("1. Explicit Declaration:")
	fmt.Printf("   Name: %s, Age: %d, Active: %t\n\n", name, age, isActive)

	// Short variable declaration (type inference)
	city := "New York"
	temperature := 72.5

	fmt.Println("2. Short Declaration (type inference):")
	fmt.Printf("   City: %s, Temperature: %.1f°F\n\n", city, temperature)

	// Multiple variable declaration
	var (
		firstName = "Alice"
		lastName  = "Smith"
		score     = 95
	)

	fmt.Println("3. Multiple Declaration:")
	fmt.Printf("   %s %s scored %d\n\n", firstName, lastName, score)

	// Constants
	const Pi = 3.14159
	const AppName = "Go Basics"

	fmt.Println("4. Constants:")
	fmt.Printf("   Pi: %.5f\n", Pi)
	fmt.Printf("   App: %s\n\n", AppName)

	// Zero values
	var defaultInt int
	var defaultString string
	var defaultBool bool

	fmt.Println("5. Zero Values:")
	fmt.Printf("   int: %d, string: '%s', bool: %t\n\n", defaultInt, defaultString, defaultBool)

	// Type conversion
	var x int = 42
	var y float64 = float64(x)

	fmt.Println("6. Type Conversion:")
	fmt.Printf("   int: %d -> float64: %.2f\n", x, y)
}
