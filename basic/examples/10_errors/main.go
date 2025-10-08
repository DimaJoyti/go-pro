package main

import (
	"errors"
	"fmt"
)

// Custom error type
type DivisionError struct {
	Dividend int
	Divisor  int
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
}

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Function with custom error
func divideInt(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{Dividend: a, Divisor: b}
	}
	return a / b, nil
}

// Function with formatted error
func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("invalid age: %d (must be non-negative)", age)
	}
	if age > 150 {
		return fmt.Errorf("invalid age: %d (must be less than 150)", age)
	}
	return nil
}

// Function with multiple error checks
func processUser(name string, age int) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if err := validateAge(age); err != nil {
		return err
	}
	return nil
}

// Panic and recover
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	result = a / b
	return result, nil
}

func main() {
	fmt.Println("=== Error Handling ===\n")

	// Basic error handling
	fmt.Println("1. Basic Error Handling:")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Result: %.2f\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("   Error: %v\n\n", err)
	}

	// Custom error
	fmt.Println("2. Custom Error:")
	_, err = divideInt(10, 0)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		// Type assertion to check error type
		if divErr, ok := err.(*DivisionError); ok {
			fmt.Printf("   Custom error details: dividend=%d, divisor=%d\n\n",
				divErr.Dividend, divErr.Divisor)
		}
	}

	// Formatted error
	fmt.Println("3. Formatted Error:")
	if err := validateAge(-5); err != nil {
		fmt.Printf("   Error: %v\n", err)
	}
	if err := validateAge(200); err != nil {
		fmt.Printf("   Error: %v\n", err)
	}
	if err := validateAge(25); err == nil {
		fmt.Println("   Age 25 is valid")
	}
	fmt.Println()

	// Multiple error checks
	fmt.Println("4. Multiple Error Checks:")
	if err := processUser("", 25); err != nil {
		fmt.Printf("   Error: %v\n", err)
	}
	if err := processUser("Alice", -5); err != nil {
		fmt.Printf("   Error: %v\n", err)
	}
	if err := processUser("Alice", 25); err == nil {
		fmt.Println("   User Alice (25) is valid")
	}
	fmt.Println()

	// Defer
	fmt.Println("5. Defer:")
	func() {
		defer fmt.Println("   This runs last (defer 3)")
		defer fmt.Println("   This runs second (defer 2)")
		defer fmt.Println("   This runs first (defer 1)")
		fmt.Println("   This runs immediately")
	}()
	fmt.Println()

	// Panic and recover
	fmt.Println("6. Panic and Recover:")
	result2, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Result: %d\n", result2)
	}

	result2, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	}
	fmt.Println()

	// Error wrapping (Go 1.13+)
	fmt.Println("7. Error Wrapping:")
	baseErr := errors.New("base error")
	wrappedErr := fmt.Errorf("wrapped: %w", baseErr)
	fmt.Printf("   Wrapped error: %v\n", wrappedErr)
	fmt.Printf("   Is base error: %t\n", errors.Is(wrappedErr, baseErr))
	fmt.Println()

	// Sentinel errors
	fmt.Println("8. Sentinel Errors:")
	var ErrNotFound = errors.New("not found")
	var ErrInvalidInput = errors.New("invalid input")

	err = ErrNotFound
	if errors.Is(err, ErrNotFound) {
		fmt.Println("   Error is ErrNotFound")
	}
	if !errors.Is(err, ErrInvalidInput) {
		fmt.Println("   Error is not ErrInvalidInput")
	}
}
