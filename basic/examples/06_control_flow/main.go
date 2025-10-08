package main

import "fmt"

func main() {
	fmt.Println("=== Control Flow ===\n")

	// If statement
	fmt.Println("1. If Statement:")
	x := 10
	if x > 5 {
		fmt.Printf("   %d is greater than 5\n\n", x)
	}

	// If-else
	fmt.Println("2. If-Else:")
	age := 18
	if age >= 18 {
		fmt.Println("   You are an adult\n")
	} else {
		fmt.Println("   You are a minor\n")
	}

	// If-else if-else
	fmt.Println("3. If-Else If-Else:")
	score := 85
	if score >= 90 {
		fmt.Println("   Grade: A")
	} else if score >= 80 {
		fmt.Println("   Grade: B")
	} else if score >= 70 {
		fmt.Println("   Grade: C")
	} else {
		fmt.Println("   Grade: F")
	}
	fmt.Println()

	// If with short statement
	fmt.Println("4. If with Short Statement:")
	if num := 42; num%2 == 0 {
		fmt.Printf("   %d is even\n\n", num)
	}

	// Switch statement
	fmt.Println("5. Switch Statement:")
	day := 3
	switch day {
	case 1:
		fmt.Println("   Monday")
	case 2:
		fmt.Println("   Tuesday")
	case 3:
		fmt.Println("   Wednesday")
	case 4:
		fmt.Println("   Thursday")
	case 5:
		fmt.Println("   Friday")
	default:
		fmt.Println("   Weekend")
	}
	fmt.Println()

	// Switch with multiple cases
	fmt.Println("6. Switch with Multiple Cases:")
	letter := "a"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Printf("   '%s' is a vowel\n\n", letter)
	default:
		fmt.Printf("   '%s' is a consonant\n\n", letter)
	}

	// Switch without condition (like if-else chain)
	fmt.Println("7. Switch without Condition:")
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("   Freezing")
	case temperature < 15:
		fmt.Println("   Cold")
	case temperature < 25:
		fmt.Println("   Mild")
	default:
		fmt.Println("   Warm")
	}
	fmt.Println()

	// For loop - basic
	fmt.Println("8. For Loop (Basic):")
	for i := 0; i < 5; i++ {
		fmt.Printf("   %d ", i)
	}
	fmt.Println("\n")

	// For loop - while style
	fmt.Println("9. For Loop (While Style):")
	count := 0
	for count < 3 {
		fmt.Printf("   Count: %d\n", count)
		count++
	}
	fmt.Println()

	// For loop - infinite with break
	fmt.Println("10. For Loop with Break:")
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Printf("   Iteration: %d\n", i)
		i++
	}
	fmt.Println()

	// For range - slice
	fmt.Println("11. For Range (Slice):")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("   [%d] %s\n", index, fruit)
	}
	fmt.Println()

	// For range - map
	fmt.Println("12. For Range (Map):")
	ages := map[string]int{"Alice": 25, "Bob": 30}
	for name, age := range ages {
		fmt.Printf("   %s is %d years old\n", name, age)
	}
	fmt.Println()

	// Continue
	fmt.Println("13. Continue:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Printf("   %d ", i)
	}
	fmt.Println()
}
