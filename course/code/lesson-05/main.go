package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Lesson 5: Pointers and Memory Management ===\n")

	// Basic pointer operations
	fmt.Println("1. Basic Pointer Operations:")
	basicPointerOperations()
	fmt.Println()

	// Function parameters with pointers
	fmt.Println("2. Function Parameters:")
	demonstrateFunctionParameters()
	fmt.Println()

	// Struct pointers and methods
	fmt.Println("3. Struct Pointers and Methods:")
	demonstrateStructPointers()
	fmt.Println()

	// Memory allocation
	fmt.Println("4. Memory Allocation:")
	demonstrateMemoryAllocation()
	fmt.Println()

	// Performance comparison
	fmt.Println("5. Performance Comparison:")
	demonstratePerformance()
	fmt.Println()

	// Real-world example
	fmt.Println("6. Real-World Example - Linked List:")
	demonstrateLinkedList()
}

func basicPointerOperations() {
	// Basic variable and pointer
	x := 42
	xPtr := &x

	fmt.Printf("x = %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)
	fmt.Printf("xPtr = %p\n", xPtr)
	fmt.Printf("Value at xPtr: %d\n", *xPtr)

	// Modify through pointer
	*xPtr = 100
	fmt.Printf("After *xPtr = 100: x = %d\n", x)

	// Nil pointer
	var nilPtr *int
	fmt.Printf("nilPtr = %v\n", nilPtr)
	fmt.Printf("nilPtr == nil: %t\n", nilPtr == nil)

	// Safe dereferencing
	if nilPtr != nil {
		fmt.Printf("Value: %d\n", *nilPtr)
	} else {
		fmt.Println("Cannot dereference nil pointer")
	}
}

func demonstrateFunctionParameters() {
	// Pass by value
	a := 10
	fmt.Printf("Before modifyValue: a = %d\n", a)
	modifyValue(a)
	fmt.Printf("After modifyValue: a = %d\n", a)

	// Pass by pointer
	fmt.Printf("Before modifyPointer: a = %d\n", a)
	modifyPointer(&a)
	fmt.Printf("After modifyPointer: a = %d\n", a)

	// Swap function
	x, y := 5, 15
	fmt.Printf("Before swap: x=%d, y=%d\n", x, y)
	swap(&x, &y)
	fmt.Printf("After swap: x=%d, y=%d\n", x, y)
}

func modifyValue(x int) {
	x = 999 // Only modifies the copy
	fmt.Printf("Inside modifyValue: x = %d\n", x)
}

func modifyPointer(x *int) {
	*x = 999 // Modifies the original
	fmt.Printf("Inside modifyPointer: *x = %d\n", *x)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func demonstrateStructPointers() {
	// Create struct with pointer
	rect := &Rectangle{Width: 10, Height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())

	// Modify through pointer receiver
	rect.Scale(2)
	fmt.Printf("After scaling by 2: %+v\n", rect)
	fmt.Printf("New area: %.2f\n", rect.Area())

	// Compare value vs pointer receiver
	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %+v\n", person)

	// This won't modify the original (value receiver)
	person.TryToAge()
	fmt.Printf("After TryToAge: %+v\n", person)

	// This will modify the original (pointer receiver)
	person.HaveBirthday()
	fmt.Printf("After HaveBirthday: %+v\n", person)
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

type Person struct {
	Name string
	Age  int
}

// Value receiver - receives a copy
func (p Person) TryToAge() {
	p.Age++ // Only modifies the copy
	fmt.Printf("Inside TryToAge: %+v\n", p)
}

// Pointer receiver - receives a pointer
func (p *Person) HaveBirthday() {
	p.Age++ // Modifies the original
	fmt.Printf("Inside HaveBirthday: %+v\n", p)
}

func demonstrateMemoryAllocation() {
	// Using new()
	intPtr := new(int)
	*intPtr = 42
	fmt.Printf("new(int): %p -> %d\n", intPtr, *intPtr)

	// Using address operator
	value := 100
	valuePtr := &value
	fmt.Printf("&value: %p -> %d\n", valuePtr, *valuePtr)

	// Using make() for slices
	slice := make([]int, 5, 10)
	fmt.Printf("make([]int, 5, 10): len=%d, cap=%d\n", len(slice), cap(slice))

	// Using make() for maps
	m := make(map[string]int)
	m["key"] = 123
	fmt.Printf("make(map[string]int): %v\n", m)
}

func demonstratePerformance() {
	// Create a large struct
	type LargeStruct struct {
		Data [1000]int
		Name string
	}

	large := LargeStruct{Name: "Test"}
	for i := 0; i < 1000; i++ {
		large.Data[i] = i
	}

	// Measure value passing
	start := time.Now()
	for i := 0; i < 10000; i++ {
		processByValue(large)
	}
	valueTime := time.Since(start)

	// Measure pointer passing
	start = time.Now()
	for i := 0; i < 10000; i++ {
		processByPointer(&large)
	}
	pointerTime := time.Since(start)

	fmt.Printf("Processing by value: %v\n", valueTime)
	fmt.Printf("Processing by pointer: %v\n", pointerTime)
	fmt.Printf("Pointer is %.2fx faster\n", float64(valueTime)/float64(pointerTime))
}

func processByValue(ls LargeStruct) {
	// Simulate some work
	_ = ls.Data[0] + ls.Data[999]
}

func processByPointer(ls *LargeStruct) {
	// Simulate some work
	_ = ls.Data[0] + ls.Data[999]
}

func demonstrateLinkedList() {
	// Create a simple linked list
	list := &LinkedList{}

	// Add some nodes
	list.Add(1)
	list.Add(2)
	list.Add(3)

	fmt.Print("Linked list: ")
	list.Print()

	fmt.Printf("Length: %d\n", list.Length())
	fmt.Printf("Contains 2: %t\n", list.Contains(2))
	fmt.Printf("Contains 5: %t\n", list.Contains(5))
}

// Simple linked list implementation
type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Add(data int) {
	newNode := &Node{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}

func (ll *LinkedList) Length() int {
	count := 0
	current := ll.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func (ll *LinkedList) Contains(data int) bool {
	current := ll.Head
	for current != nil {
		if current.Data == data {
			return true
		}
		current = current.Next
	}
	return false
}
