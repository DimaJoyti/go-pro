package main

import "fmt"

// Package-level variables (global scope)
var globalMessage = "I'm accessible throughout the package"
var globalCounter int = 0

// Constants for demonstration
const (
	AppName    = "GO-PRO Lesson 2"
	MaxRetries = 3
	Pi         = 3.14159
)

func main() {
	fmt.Println("=== GO-PRO Lesson 2: Variables, Constants, and Functions ===\n")

	// Demonstrate variable declarations
	demonstrateVariableDeclarations()
	fmt.Println()

	// Demonstrate variable scope
	demonstrateVariableScope()
	fmt.Println()

	// Demonstrate basic functions
	demonstrateBasicFunctions()
	fmt.Println()

	// Demonstrate multiple return values
	demonstrateMultipleReturns()
	fmt.Println()

	// Demonstrate named returns
	demonstrateNamedReturns()
	fmt.Println()

	// Demonstrate variadic functions
	demonstrateVariadicFunctions()
	fmt.Println()

	// Demonstrate function variables
	demonstrateFunctionVariables()
}

// demonstrateVariableDeclarations shows different ways to declare variables
func demonstrateVariableDeclarations() {
	fmt.Println("📝 Variable Declarations:")

	// Method 1: Explicit declaration with var
	var name string = "Go Programming"
	var version float64 = 1.21
	var isActive bool = true

	// Method 2: Type inference with var
	var language = "Go"
	var year = 2009

	// Method 3: Short variable declaration (only inside functions)
	framework := "Standard Library"
	popularity := 95

	// Method 4: Multiple variable declarations
	var x, y, z int = 10, 20, 30
	a, b := "Hello", "World"

	// Method 5: Block declaration
	var (
		projectName = "GO-PRO"
		difficulty  = "Beginner to Advanced"
		duration    = "12 weeks"
	)

	// Print all variables
	fmt.Printf("Explicit declaration: %s v%.2f (Active: %t)\n", name, version, isActive)
	fmt.Printf("Type inference: %s from %d\n", language, year)
	fmt.Printf("Short declaration: %s (Popularity: %d%%)\n", framework, popularity)
	fmt.Printf("Multiple variables: %d, %d, %d\n", x, y, z)
	fmt.Printf("Multiple short: %s %s\n", a, b)
	fmt.Printf("Block declaration: %s - %s (%s)\n", projectName, difficulty, duration)
}

// demonstrateVariableScope shows different variable scopes
func demonstrateVariableScope() {
	fmt.Println("🔍 Variable Scope:")

	// Function-level variable
	localVar := "I'm local to this function"

	// Access global variable
	globalCounter++
	fmt.Printf("Global message: %s\n", globalMessage)
	fmt.Printf("Global counter: %d\n", globalCounter)
	fmt.Printf("Local variable: %s\n", localVar)

	// Block scope
	if true {
		blockVar := "I'm only in this block"
		fmt.Printf("Block variable: %s\n", blockVar)
		
		// Can access outer scopes
		fmt.Printf("Accessing local from block: %s\n", localVar)
	}
	// blockVar is not accessible here

	// Loop scope
	for i := 0; i < 3; i++ {
		loopVar := fmt.Sprintf("Loop iteration %d", i)
		fmt.Printf("Loop variable: %s\n", loopVar)
	}
	// i and loopVar are not accessible here
}

// demonstrateBasicFunctions shows basic function patterns
func demonstrateBasicFunctions() {
	fmt.Println("⚙️ Basic Functions:")

	// Call simple functions
	greeting := greet("Go Developer")
	fmt.Printf("Greeting: %s\n", greeting)

	sum := add(15, 25)
	fmt.Printf("Sum: %d\n", sum)

	product := multiply(3, 4, 5)
	fmt.Printf("Product: %d\n", product)

	// Function with same type parameters
	result := calculate(10, 5, 2)
	fmt.Printf("Calculation result: %d\n", result)
}

// demonstrateMultipleReturns shows functions with multiple return values
func demonstrateMultipleReturns() {
	fmt.Println("↩️ Multiple Return Values:")

	// Function returning multiple values
	quotient, remainder := divmod(17, 5)
	fmt.Printf("17 ÷ 5 = %d remainder %d\n", quotient, remainder)

	// Function returning value and error
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Safe division result: %.2f\n", result)
	}

	// Test error case
	_, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	// Ignoring return values with blank identifier
	_, remainder2 := divmod(23, 7)
	fmt.Printf("Only remainder: %d\n", remainder2)
}

// demonstrateNamedReturns shows named return values
func demonstrateNamedReturns() {
	fmt.Println("🏷️ Named Return Values:")

	area, perimeter := rectangleStats(5.0, 3.0)
	fmt.Printf("Rectangle (5×3): Area=%.2f, Perimeter=%.2f\n", area, perimeter)

	circumference, area2 := circleStats(4.0)
	fmt.Printf("Circle (r=4): Circumference=%.2f, Area=%.2f\n", circumference, area2)
}

// demonstrateVariadicFunctions shows functions with variable arguments
func demonstrateVariadicFunctions() {
	fmt.Println("📊 Variadic Functions:")

	// Functions with variable number of arguments
	sum1 := sum(1, 2, 3)
	fmt.Printf("Sum of 1,2,3: %d\n", sum1)

	sum2 := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Sum of 1-10: %d\n", sum2)

	avg1 := average(85.5, 92.0, 78.5, 96.0)
	fmt.Printf("Average grade: %.2f\n", avg1)

	// Passing slice to variadic function
	numbers := []int{5, 10, 15, 20, 25}
	sum3 := sum(numbers...) // Spread operator
	fmt.Printf("Sum of slice: %d\n", sum3)

	// String joining
	message := joinStrings(" ", "Go", "is", "awesome", "for", "learning!")
	fmt.Printf("Joined message: %s\n", message)
}

// demonstrateFunctionVariables shows functions as variables
func demonstrateFunctionVariables() {
	fmt.Println("🔧 Function Variables:")

	// Function as variable
	var operation func(int, int) int
	operation = add
	result1 := operation(10, 5)
	fmt.Printf("Using add function: %d\n", result1)

	// Anonymous function
	multiply := func(a, b int) int {
		return a * b
	}
	result2 := multiply(7, 8)
	fmt.Printf("Anonymous multiply: %d\n", result2)

	// Function returning function
	adder := makeAdder(10)
	result3 := adder(5)
	fmt.Printf("Closure result: %d\n", result3)

	// Using function as parameter
	result4 := applyOperation(20, 4, subtract)
	fmt.Printf("Applied subtraction: %d\n", result4)
}

// Basic function examples
func greet(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to Go programming.", name)
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b, c int) int {
	return a * b * c
}

// Function with same type parameters (shorthand)
func calculate(x, y, z int) int {
	return (x + y) * z
}

// Multiple return values
func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Named return values
func rectangleStats(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

func circleStats(radius float64) (circumference, area float64) {
	circumference = 2 * Pi * radius
	area = Pi * radius * radius
	return
}

// Variadic functions
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func average(scores ...float64) float64 {
	if len(scores) == 0 {
		return 0
	}
	
	total := 0.0
	for _, score := range scores {
		total += score
	}
	return total / float64(len(scores))
}

func joinStrings(separator string, strings ...string) string {
	if len(strings) == 0 {
		return ""
	}
	
	result := strings[0]
	for i := 1; i < len(strings); i++ {
		result += separator + strings[i]
	}
	return result
}

// Function variables and closures
func subtract(a, b int) int {
	return a - b
}

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}
