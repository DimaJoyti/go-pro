// Package main provides an interactive runner for all Go basic examples
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Example represents a runnable example
type Example struct {
	ID          int
	Name        string
	Description string
	Path        string
}

var examples = []Example{
	{1, "Hello World", "Basic Hello World program", "examples/01_hello"},
	{2, "Variables", "Variables, constants, and data types", "examples/02_variables"},
	{3, "Functions", "Function declarations and usage", "examples/03_functions"},
	{4, "Pointers", "Pointer basics and usage", "examples/04_pointers"},
	{5, "Arrays & Slices", "Working with arrays and slices", "examples/05_arrays_slices"},
	{6, "Control Flow", "If, switch, and loops", "examples/06_control_flow"},
	{7, "Maps", "Map operations", "examples/07_maps"},
	{8, "Structs", "Structs and methods", "examples/08_structs"},
	{9, "Interfaces", "Interface definitions and implementations", "examples/09_interfaces"},
	{10, "Errors", "Error handling patterns", "examples/10_errors"},
	{11, "Concurrency", "Goroutines and channels", "examples/11_concurrency"},
	{12, "Advanced", "Advanced Go topics", "examples/12_advanced"},
}

var projects = []Example{
	{1, "Calculator", "Simple calculator project", "projects/calculator"},
	{2, "Todo List", "Command-line todo list", "projects/todo_list"},
	{3, "Order System", "Order management system", "projects/order_system"},
}

func main() {
	clearScreen()
	printBanner()
	
	for {
		printMenu()
		choice := getUserInput()
		
		if choice == "q" || choice == "quit" || choice == "exit" {
			fmt.Println("\n👋 Thanks for learning Go! Happy coding!")
			break
		}
		
		handleChoice(choice)
	}
}

func printBanner() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║                                                            ║")
	fmt.Println("║           🚀 Go Programming - Basic Examples 🚀            ║")
	fmt.Println("║                                                            ║")
	fmt.Println("║              Interactive Learning Environment              ║")
	fmt.Println("║                                                            ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printMenu() {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📚 MAIN MENU")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
	
	fmt.Println("📖 Examples:")
	for _, ex := range examples {
		fmt.Printf("  [%2d] %-20s - %s\n", ex.ID, ex.Name, ex.Description)
	}
	
	fmt.Println("\n🎯 Projects:")
	for _, proj := range projects {
		fmt.Printf("  [p%d] %-20s - %s\n", proj.ID, proj.Name, proj.Description)
	}
	
	fmt.Println("\n⚙️  Options:")
	fmt.Println("  [a]  Run all examples")
	fmt.Println("  [t]  Run tests")
	fmt.Println("  [h]  Show help")
	fmt.Println("  [q]  Quit")
	
	fmt.Println(strings.Repeat("=", 60))
	fmt.Print("\n👉 Enter your choice: ")
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(input))
}

func handleChoice(choice string) {
	clearScreen()
	
	switch choice {
	case "a":
		runAllExamples()
	case "t":
		runTests()
	case "h":
		showHelp()
	default:
		if strings.HasPrefix(choice, "p") {
			// Project
			idStr := strings.TrimPrefix(choice, "p")
			if id, err := strconv.Atoi(idStr); err == nil {
				runProject(id)
			} else {
				fmt.Println("❌ Invalid project number")
			}
		} else {
			// Example
			if id, err := strconv.Atoi(choice); err == nil {
				runExample(id)
			} else {
				fmt.Println("❌ Invalid choice. Please try again.")
			}
		}
	}
	
	fmt.Print("\n\nPress Enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
	clearScreen()
}

func runExample(id int) {
	for _, ex := range examples {
		if ex.ID == id {
			fmt.Printf("🚀 Running: %s\n", ex.Name)
			fmt.Println(strings.Repeat("=", 60))
			fmt.Println()
			
			runGoFile(ex.Path)
			return
		}
	}
	fmt.Println("❌ Example not found")
}

func runProject(id int) {
	for _, proj := range projects {
		if proj.ID == id {
			fmt.Printf("🎯 Running Project: %s\n", proj.Name)
			fmt.Println(strings.Repeat("=", 60))
			fmt.Println()
			
			runGoFile(proj.Path)
			return
		}
	}
	fmt.Println("❌ Project not found")
}

func runAllExamples() {
	fmt.Println("🚀 Running All Examples")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
	
	for _, ex := range examples {
		fmt.Printf("\n📖 %d. %s\n", ex.ID, ex.Name)
		fmt.Println(strings.Repeat("-", 60))
		runGoFile(ex.Path)
		fmt.Println()
	}
}

func runTests() {
	fmt.Println("🧪 Running Tests")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
	
	cmd := exec.Command("go", "test", "./...")
	cmd.Dir = getBasicDir()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Tests failed: %v\n", err)
	} else {
		fmt.Println("✅ All tests passed!")
	}
}

func showHelp() {
	fmt.Println("📚 Help - How to Use This Interactive Runner")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
	fmt.Println("This interactive runner helps you explore Go programming examples.")
	fmt.Println()
	fmt.Println("Navigation:")
	fmt.Println("  • Enter a number (1-12) to run an example")
	fmt.Println("  • Enter 'p' followed by a number (p1, p2, p3) for projects")
	fmt.Println("  • Enter 'a' to run all examples sequentially")
	fmt.Println("  • Enter 't' to run all tests")
	fmt.Println("  • Enter 'q' to quit")
	fmt.Println()
	fmt.Println("Tips:")
	fmt.Println("  • Read the code in each example to understand the concepts")
	fmt.Println("  • Modify examples and run them again to experiment")
	fmt.Println("  • Check the README.md in each example folder for details")
	fmt.Println()
}

func runGoFile(relativePath string) {
	basicDir := getBasicDir()
	fullPath := filepath.Join(basicDir, relativePath)
	
	// Check if main.go exists
	mainFile := filepath.Join(fullPath, "main.go")
	if _, err := os.Stat(mainFile); os.IsNotExist(err) {
		fmt.Printf("⚠️  Example not yet implemented: %s\n", relativePath)
		return
	}
	
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = fullPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Error running example: %v\n", err)
	}
}

func getBasicDir() string {
	// Get the directory where this program is located
	_, filename, _, _ := runtime.Caller(0)
	// Go up from cmd/runner to basic directory
	return filepath.Join(filepath.Dir(filename), "..", "..")
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

