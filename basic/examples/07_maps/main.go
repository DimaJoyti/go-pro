package main

import "fmt"

func main() {
	fmt.Println("=== Maps ===\n")

	// Create map using make
	fmt.Println("1. Create Map with make:")
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	ages["Charlie"] = 35
	fmt.Printf("   Ages: %v\n\n", ages)

	// Map literal
	fmt.Println("2. Map Literal:")
	scores := map[string]int{
		"Math":    95,
		"English": 88,
		"Science": 92,
	}
	fmt.Printf("   Scores: %v\n\n", scores)

	// Access map values
	fmt.Println("3. Access Map Values:")
	fmt.Printf("   Math score: %d\n", scores["Math"])
	fmt.Printf("   English score: %d\n\n", scores["English"])

	// Check if key exists
	fmt.Println("4. Check if Key Exists:")
	if score, exists := scores["History"]; exists {
		fmt.Printf("   History score: %d\n", score)
	} else {
		fmt.Println("   History score not found")
	}
	fmt.Println()

	// Add/Update values
	fmt.Println("5. Add/Update Values:")
	scores["History"] = 90
	scores["Math"] = 98 // Update existing
	fmt.Printf("   Updated scores: %v\n\n", scores)

	// Delete from map
	fmt.Println("6. Delete from Map:")
	delete(scores, "English")
	fmt.Printf("   After deleting English: %v\n\n", scores)

	// Iterate over map
	fmt.Println("7. Iterate Over Map:")
	for subject, score := range scores {
		fmt.Printf("   %s: %d\n", subject, score)
	}
	fmt.Println()

	// Map length
	fmt.Println("8. Map Length:")
	fmt.Printf("   Number of subjects: %d\n\n", len(scores))

	// Map with struct values
	fmt.Println("9. Map with Struct Values:")
	type Person struct {
		Name string
		Age  int
	}

	people := map[string]Person{
		"emp1": {Name: "Alice", Age: 25},
		"emp2": {Name: "Bob", Age: 30},
	}

	for id, person := range people {
		fmt.Printf("   %s: %s (%d years old)\n", id, person.Name, person.Age)
	}
	fmt.Println()

	// Nested maps
	fmt.Println("10. Nested Maps:")
	grades := map[string]map[string]int{
		"Alice": {
			"Math":    95,
			"English": 88,
		},
		"Bob": {
			"Math":    85,
			"English": 92,
		},
	}

	for student, subjects := range grades {
		fmt.Printf("   %s's grades:\n", student)
		for subject, grade := range subjects {
			fmt.Printf("     %s: %d\n", subject, grade)
		}
	}
}
