package main

import "fmt"

// Basic function with parameters and return value
func add(x, y int) int {
	return x + y
}

// Function with multiple return values
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return x / y, nil
}

// Function with named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

// Variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Function type declaration
type Operation func(int, int) int

// Higher-order function that returns a function
func getOperation(op string) Operation {
	switch op {
	case "+":
		return func(x, y int) int { return x + y }
	case "-":
		return func(x, y int) int { return x - y }
	case "*":
		return func(x, y int) int { return x * y }
	default:
		return func(x, y int) int { return 0 }
	}
}

// Closure example
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println("=== Functions ===\n")

	// Basic function
	fmt.Println("1. Basic Function:")
	fmt.Printf("   add(5, 3) = %d\n\n", add(5, 3))

	// Multiple return values
	fmt.Println("2. Multiple Return Values:")
	if result, err := divide(10, 2); err == nil {
		fmt.Printf("   divide(10, 2) = %.2f\n", result)
	}
	if _, err := divide(10, 0); err != nil {
		fmt.Printf("   divide(10, 0) = Error: %v\n\n", err)
	}

	// Named return values
	fmt.Println("3. Named Return Values:")
	x, y := split(17)
	fmt.Printf("   split(17) = %d, %d\n\n", x, y)

	// Variadic function
	fmt.Println("4. Variadic Function:")
	fmt.Printf("   sum(1, 2, 3, 4, 5) = %d\n\n", sum(1, 2, 3, 4, 5))

	// Higher-order function
	fmt.Println("5. Higher-Order Function:")
	plus := getOperation("+")
	minus := getOperation("-")
	multiply := getOperation("*")
	fmt.Printf("   plus(10, 5) = %d\n", plus(10, 5))
	fmt.Printf("   minus(10, 5) = %d\n", minus(10, 5))
	fmt.Printf("   multiply(10, 5) = %d\n\n", multiply(10, 5))

	// Closure
	fmt.Println("6. Closure:")
	nextCount := counter()
	fmt.Printf("   counter() = %d\n", nextCount())
	fmt.Printf("   counter() = %d\n", nextCount())
	fmt.Printf("   counter() = %d\n", nextCount())
}
