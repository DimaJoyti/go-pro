package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator provides basic arithmetic operations
type Calculator struct {
	history []string
}

// NewCalculator creates a new calculator instance
func NewCalculator() *Calculator {
	return &Calculator{
		history: make([]string, 0),
	}
}

// Add performs addition
func (c *Calculator) Add(a, b float64) float64 {
	result := a + b
	c.addToHistory(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

// Subtract performs subtraction
func (c *Calculator) Subtract(a, b float64) float64 {
	result := a - b
	c.addToHistory(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

// Multiply performs multiplication
func (c *Calculator) Multiply(a, b float64) float64 {
	result := a * b
	c.addToHistory(fmt.Sprintf("%.2f × %.2f = %.2f", a, b, result))
	return result
}

// Divide performs division
func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	c.addToHistory(fmt.Sprintf("%.2f ÷ %.2f = %.2f", a, b, result))
	return result, nil
}

// Power calculates a raised to the power of b
func (c *Calculator) Power(a, b float64) float64 {
	result := 1.0
	for i := 0; i < int(b); i++ {
		result *= a
	}
	c.addToHistory(fmt.Sprintf("%.2f ^ %.0f = %.2f", a, b, result))
	return result
}

// addToHistory adds an operation to the history
func (c *Calculator) addToHistory(operation string) {
	c.history = append(c.history, operation)
}

// ShowHistory displays the calculation history
func (c *Calculator) ShowHistory() {
	if len(c.history) == 0 {
		fmt.Println("No history available")
		return
	}

	fmt.Println("\n📜 Calculation History:")
	fmt.Println(strings.Repeat("=", 50))
	for i, op := range c.history {
		fmt.Printf("%2d. %s\n", i+1, op)
	}
	fmt.Println(strings.Repeat("=", 50))
}

// ClearHistory clears the calculation history
func (c *Calculator) ClearHistory() {
	c.history = make([]string, 0)
	fmt.Println("✓ History cleared")
}

func printBanner() {
	fmt.Println("\n╔════════════════════════════════════════════════╗")
	fmt.Println("║                                                ║")
	fmt.Println("║           🧮  Go Calculator  🧮                ║")
	fmt.Println("║                                                ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
}

func printMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📋 Operations:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("  [1] Add (+)")
	fmt.Println("  [2] Subtract (-)")
	fmt.Println("  [3] Multiply (×)")
	fmt.Println("  [4] Divide (÷)")
	fmt.Println("  [5] Power (^)")
	fmt.Println("  [h] Show History")
	fmt.Println("  [c] Clear History")
	fmt.Println("  [q] Quit")
	fmt.Println(strings.Repeat("=", 50))
}

func getInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getNumber(reader *bufio.Reader, prompt string) (float64, error) {
	input := getInput(reader, prompt)
	return strconv.ParseFloat(input, 64)
}

func performOperation(calc *Calculator, reader *bufio.Reader, operation string) {
	var result float64
	var err error

	// Get first number
	num1, err := getNumber(reader, "Enter first number: ")
	if err != nil {
		fmt.Println("❌ Invalid number")
		return
	}

	// Get second number
	num2, err := getNumber(reader, "Enter second number: ")
	if err != nil {
		fmt.Println("❌ Invalid number")
		return
	}

	// Perform operation
	switch operation {
	case "1":
		result = calc.Add(num1, num2)
		fmt.Printf("\n✓ Result: %.2f\n", result)
	case "2":
		result = calc.Subtract(num1, num2)
		fmt.Printf("\n✓ Result: %.2f\n", result)
	case "3":
		result = calc.Multiply(num1, num2)
		fmt.Printf("\n✓ Result: %.2f\n", result)
	case "4":
		result, err = calc.Divide(num1, num2)
		if err != nil {
			fmt.Printf("\n❌ Error: %v\n", err)
		} else {
			fmt.Printf("\n✓ Result: %.2f\n", result)
		}
	case "5":
		result = calc.Power(num1, num2)
		fmt.Printf("\n✓ Result: %.2f\n", result)
	}
}

func main() {
	calc := NewCalculator()
	reader := bufio.NewReader(os.Stdin)

	printBanner()

	for {
		printMenu()
		choice := getInput(reader, "\n👉 Enter your choice: ")

		switch choice {
		case "1", "2", "3", "4", "5":
			performOperation(calc, reader, choice)
		case "h", "H":
			calc.ShowHistory()
		case "c", "C":
			calc.ClearHistory()
		case "q", "Q", "quit", "exit":
			fmt.Println("\n👋 Thanks for using Go Calculator! Goodbye!")
			return
		default:
			fmt.Println("\n❌ Invalid choice. Please try again.")
		}
	}
}
