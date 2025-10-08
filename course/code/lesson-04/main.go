package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("=== Lesson 4: Arrays, Slices, and Maps ===\n")

	// Arrays demonstration
	fmt.Println("1. Arrays:")
	demonstrateArrays()
	fmt.Println()

	// Slices demonstration
	fmt.Println("2. Slices:")
	demonstrateSlices()
	fmt.Println()

	// Maps demonstration
	fmt.Println("3. Maps:")
	demonstrateMaps()
	fmt.Println()

	// Advanced operations
	fmt.Println("4. Advanced Operations:")
	demonstrateAdvancedOperations()
	fmt.Println()

	// Real-world example
	fmt.Println("5. Real-World Example - Student Management:")
	studentManagementExample()
}

func demonstrateArrays() {
	// Array declaration and initialization
	var numbers [5]int // Zero-valued array
	fmt.Printf("Zero-valued array: %v\n", numbers)

	// Array with initial values
	fruits := [3]string{"apple", "banana", "orange"}
	fmt.Printf("Fruits array: %v\n", fruits)

	// Array with inferred size
	cities := [...]string{"New York", "London", "Tokyo", "Paris"}
	fmt.Printf("Cities array: %v (length: %d)\n", cities, len(cities))

	// Accessing and modifying array elements
	numbers[0] = 10
	numbers[1] = 20
	fmt.Printf("Modified numbers: %v\n", numbers)

	// Arrays are value types (copied when assigned)
	numbersCopy := numbers
	numbersCopy[0] = 999
	fmt.Printf("Original: %v, Copy: %v\n", numbers, numbersCopy)
}

func demonstrateSlices() {
	// Slice declaration and initialization
	var numbers []int // nil slice
	fmt.Printf("Nil slice: %v (len: %d, cap: %d)\n", numbers, len(numbers), cap(numbers))

	// Creating slices with make
	scores := make([]int, 3, 5) // length 3, capacity 5
	fmt.Printf("Made slice: %v (len: %d, cap: %d)\n", scores, len(scores), cap(scores))

	// Slice literal
	fruits := []string{"apple", "banana", "orange"}
	fmt.Printf("Fruit slice: %v\n", fruits)

	// Appending to slices
	fruits = append(fruits, "grape", "kiwi")
	fmt.Printf("After append: %v (len: %d, cap: %d)\n", fruits, len(fruits), cap(fruits))

	// Slicing operations
	fmt.Printf("Subslice [1:3]: %v\n", fruits[1:3])
	fmt.Printf("Prefix [:2]: %v\n", fruits[:2])
	fmt.Printf("Suffix [2:]: %v\n", fruits[2:])

	// Copying slices
	fruitsCopy := make([]string, len(fruits))
	copy(fruitsCopy, fruits)
	fmt.Printf("Copied slice: %v\n", fruitsCopy)

	// Slice of slices (2D slice)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Printf("Matrix: %v\n", matrix)
}

func demonstrateMaps() {
	// Map declaration and initialization
	var ages map[string]int // nil map
	fmt.Printf("Nil map: %v\n", ages)

	// Creating map with make
	ages = make(map[string]int)
	ages["Alice"] = 30
	ages["Bob"] = 25
	fmt.Printf("Ages map: %v\n", ages)

	// Map literal
	scores := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
	}
	fmt.Printf("Scores map: %v\n", scores)

	// Accessing map values
	aliceScore := scores["Alice"]
	fmt.Printf("Alice's score: %d\n", aliceScore)

	// Checking if key exists
	if score, exists := scores["David"]; exists {
		fmt.Printf("David's score: %d\n", score)
	} else {
		fmt.Println("David not found in scores")
	}

	// Adding and updating
	scores["David"] = 88
	scores["Alice"] = 97 // Update existing
	fmt.Printf("Updated scores: %v\n", scores)

	// Deleting from map
	delete(scores, "Bob")
	fmt.Printf("After deleting Bob: %v\n", scores)

	// Iterating over map
	fmt.Println("All scores:")
	for name, score := range scores {
		fmt.Printf("  %s: %d\n", name, score)
	}
}

func demonstrateAdvancedOperations() {
	// Working with slice capacity
	slice := make([]int, 0, 10)
	fmt.Printf("Initial slice: len=%d, cap=%d\n", len(slice), cap(slice))

	for i := 0; i < 15; i++ {
		slice = append(slice, i)
		fmt.Printf("After append %d: len=%d, cap=%d\n", i, len(slice), cap(slice))
	}

	// Sorting slices
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Before sort: %v\n", numbers)
	sort.Ints(numbers)
	fmt.Printf("After sort: %v\n", numbers)

	// Sorting strings
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("Before sort: %v\n", words)
	sort.Strings(words)
	fmt.Printf("After sort: %v\n", words)

	// Map with slice values
	groups := map[string][]string{
		"fruits":     {"apple", "banana", "orange"},
		"vegetables": {"carrot", "broccoli", "spinach"},
		"grains":     {"rice", "wheat", "oats"},
	}

	fmt.Println("Food groups:")
	for category, items := range groups {
		fmt.Printf("  %s: %v\n", category, items)
	}
}

func studentManagementExample() {
	// Student management system using maps and slices
	type Student struct {
		Name   string
		Grades []int
	}

	students := map[string]*Student{
		"S001": {Name: "Alice Johnson", Grades: []int{95, 87, 92}},
		"S002": {Name: "Bob Smith", Grades: []int{78, 85, 90}},
		"S003": {Name: "Carol Davis", Grades: []int{92, 94, 89}},
	}

	// Calculate and display averages
	fmt.Println("Student Report:")
	for id, student := range students {
		total := 0
		for _, grade := range student.Grades {
			total += grade
		}
		average := float64(total) / float64(len(student.Grades))
		fmt.Printf("  %s - %s: Average %.1f\n", id, student.Name, average)
	}

	// Add new grade to a student
	students["S001"].Grades = append(students["S001"].Grades, 98)
	fmt.Printf("Updated grades for %s: %v\n", 
		students["S001"].Name, students["S001"].Grades)

	// Find top performer
	var topStudent string
	var topAverage float64

	for _, student := range students {
		total := 0
		for _, grade := range student.Grades {
			total += grade
		}
		average := float64(total) / float64(len(student.Grades))
		
		if average > topAverage {
			topAverage = average
			topStudent = student.Name
		}
	}

	fmt.Printf("Top performer: %s with average %.1f\n", topStudent, topAverage)
}
