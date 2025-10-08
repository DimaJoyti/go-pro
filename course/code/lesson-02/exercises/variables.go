package exercises

import "fmt"

// Suppress unused import warning - fmt is used in exercise solutions
var _ = fmt.Sprint

// Exercise 1: Variable Declarations and Scope
// Complete the following functions to practice with Go variables

// TODO: Complete this function
// DeclareVariables demonstrates different variable declaration methods
// It should declare and return variables using different approaches
// Returns: name (string), age (int), salary (float64), isEmployed (bool)
func DeclareVariables() (string, int, float64, bool) {
	// TODO: Declare variables using different methods:
	// 1. Use explicit var declaration for name = "John Doe"
	// 2. Use type inference for age = 30
	// 3. Use short declaration for salary = 75000.50
	// 4. Use explicit var for isEmployed = true

	// Replace these return values with your declarations
	return "", 0, 0.0, false
}

// TODO: Complete this function
// MultipleDeclarations practices declaring multiple variables at once
// Returns: x, y, z (all int), a, b (both string)
func MultipleDeclarations() (int, int, int, string, string) {
	// TODO: Declare multiple variables:
	// 1. Declare x, y, z as int with values 10, 20, 30 in one line
	// 2. Declare a, b as string with values "Hello", "World" using short declaration

	// Replace these return values with your declarations
	return 0, 0, 0, "", ""
}

// TODO: Complete this function
// BlockDeclaration uses block declaration syntax
// Returns: projectName (string), version (float64), isStable (bool)
func BlockDeclaration() (string, float64, bool) {
	// TODO: Use block declaration syntax (var (...)) to declare:
	// - projectName = "GO-PRO"
	// - version = 2.1
	// - isStable = true

	// Replace these return values with your block declaration
	return "", 0.0, false
}

// Package-level variable for scope testing
var packageVar = "I'm at package level"

// TODO: Complete this function
// TestVariableScope demonstrates variable scope rules
// Parameter: input (string)
// Returns: result (string) containing information about variable access
func TestVariableScope(input string) string {
	// TODO:
	// 1. Create a function-level variable: functionVar = "I'm in function"
	// 2. Create a result string that includes:
	//    - The input parameter
	//    - The packageVar
	//    - The functionVar
	//    - A block-scoped variable created inside an if statement
	// 3. Format: "Input: [input], Package: [packageVar], Function: [functionVar], Block: [blockVar]"

	// Replace this return statement with your implementation
	return ""
}

// TODO: Complete this function
// SwapVariables swaps the values of two variables
// Parameters: a, b (both int)
// Returns: swapped values (b, a)
func SwapVariables(a, b int) (int, int) {
	// TODO: Swap the values of a and b and return them
	// Hint: You can use multiple assignment

	// Replace this return statement with your implementation
	return 0, 0
}

// TODO: Complete this function
// ZeroValues returns the zero values of different types
// Returns: zeroInt (int), zeroFloat (float64), zeroString (string), zeroBool (bool)
func ZeroValues() (int, float64, string, bool) {
	// TODO: Declare variables without initialization to get zero values
	// Return the zero values of int, float64, string, and bool

	// Replace these return values with your zero value declarations
	return 1, 1.0, "not zero", true
}

// TODO: Complete this function
// ConstantUsage demonstrates working with constants
// Returns: pi (float64), maxRetries (int), appName (string)
func ConstantUsage() (float64, int, string) {
	// TODO: Declare constants:
	// - pi = 3.14159
	// - maxRetries = 5
	// - appName = "Learning Go"
	// Return these constants

	// Replace these return values with your constant declarations
	return 0.0, 0, ""
}

// TODO: Complete this function
// VariableReassignment demonstrates variable reassignment
// Parameter: initial (int)
// Returns: final value after operations (int)
func VariableReassignment(initial int) int {
	// TODO:
	// 1. Create a variable 'value' and assign it the initial parameter
	// 2. Add 10 to value
	// 3. Multiply value by 2
	// 4. Subtract 5 from value
	// 5. Return the final value

	// Replace this return statement with your implementation
	return 0
}

// TODO: Complete this function
// TypeInference demonstrates Go's type inference
// Returns: inferredInt, inferredFloat, inferredString, inferredBool with their inferred types
func TypeInference() (int, float64, string, bool) {
	// TODO: Use type inference (var x = value) to declare:
	// - A variable that Go infers as int (value: 42)
	// - A variable that Go infers as float64 (value: 3.14)
	// - A variable that Go infers as string (value: "Go")
	// - A variable that Go infers as bool (value: true)

	// Replace these return values with your type inference declarations
	return 0, 0.0, "", false
}

// TODO: Complete this function
// ShadowingExample demonstrates variable shadowing
// Parameter: outer (int)
// Returns: result describing the shadowing behavior (string)
func ShadowingExample(outer int) string {
	// TODO:
	// 1. Create a variable 'value' with the outer parameter
	// 2. In an if block, create another 'value' variable that shadows the outer one
	// 3. Return a string describing both values
	// Format: "Outer: [outer_value], Inner: [inner_value]"

	// Replace this return statement with your implementation
	return ""
}
