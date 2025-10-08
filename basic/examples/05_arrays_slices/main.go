package main

import "fmt"

func main() {
	fmt.Println("=== Arrays and Slices ===\n")

	// Arrays - fixed size
	fmt.Println("1. Arrays (Fixed Size):")
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	fmt.Printf("   Array: %v\n", arr)
	fmt.Printf("   Length: %d\n\n", len(arr))

	// Array literal
	fmt.Println("2. Array Literal:")
	colors := [3]string{"red", "green", "blue"}
	fmt.Printf("   Colors: %v\n\n", colors)

	// Slices - dynamic size
	fmt.Println("3. Slices (Dynamic Size):")
	var slice []int
	fmt.Printf("   Empty slice: %v (len=%d, cap=%d)\n", slice, len(slice), cap(slice))

	slice = append(slice, 1, 2, 3)
	fmt.Printf("   After append: %v (len=%d, cap=%d)\n\n", slice, len(slice), cap(slice))

	// Slice literal
	fmt.Println("4. Slice Literal:")
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Printf("   Numbers: %v\n\n", numbers)

	// Slicing
	fmt.Println("5. Slicing Operations:")
	fmt.Printf("   numbers[1:4] = %v\n", numbers[1:4])
	fmt.Printf("   numbers[:3] = %v\n", numbers[:3])
	fmt.Printf("   numbers[2:] = %v\n\n", numbers[2:])

	// Make slice
	fmt.Println("6. Make Slice:")
	s := make([]int, 5) // length 5
	fmt.Printf("   make([]int, 5): %v (len=%d, cap=%d)\n", s, len(s), cap(s))

	s2 := make([]int, 3, 10) // length 3, capacity 10
	fmt.Printf("   make([]int, 3, 10): %v (len=%d, cap=%d)\n\n", s2, len(s2), cap(s2))

	// Append and capacity
	fmt.Println("7. Append and Capacity:")
	nums := make([]int, 0, 3)
	fmt.Printf("   Initial: %v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))

	for i := 1; i <= 5; i++ {
		nums = append(nums, i)
		fmt.Printf("   After append %d: %v (len=%d, cap=%d)\n", i, nums, len(nums), cap(nums))
	}
	fmt.Println()

	// Copy slice
	fmt.Println("8. Copy Slice:")
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Printf("   Source: %v\n", src)
	fmt.Printf("   Destination: %v\n\n", dst)

	// Iteration
	fmt.Println("9. Iteration:")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("   [%d] %s\n", i, fruit)
	}
	fmt.Println()

	// 2D slice
	fmt.Println("10. 2D Slice:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i, row := range matrix {
		fmt.Printf("   Row %d: %v\n", i, row)
	}
}
