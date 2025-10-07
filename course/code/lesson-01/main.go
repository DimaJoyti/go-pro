package main

import "fmt"

// Constants demonstration
const (
	// Basic constants
	AppName    = "GO-PRO Learning Platform"
	AppVersion = "1.0.0"
	MaxUsers   = 1000
)

// Using iota for enumerated constants
const (
	// Days of the week
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	// HTTP status codes
	StatusOK       = 200
	StatusNotFound = 404
	StatusError    = 500
)

const (
	// File permissions using bit operations with iota
	ReadPermission = 1 << iota // 1 (binary: 001)
	WritePermission            // 2 (binary: 010)
	ExecutePermission          // 4 (binary: 100)
)

func main() {
	fmt.Println("=== GO-PRO Lesson 1: Go Syntax and Basic Types ===\n")

	// Demonstrate basic types
	demonstrateBasicTypes()
	fmt.Println()

	// Demonstrate constants
	demonstrateConstants()
	fmt.Println()

	// Demonstrate type conversions
	demonstrateTypeConversions()
	fmt.Println()

	// Demonstrate iota usage
	demonstrateIota()
}

// demonstrateBasicTypes shows Go's basic data types
func demonstrateBasicTypes() {
	fmt.Println("📊 Basic Types Demonstration:")

	// Integer types
	var age int = 25
	var population int64 = 7800000000
	var smallNumber int8 = 127

	// Unsigned integer types
	var positiveNumber uint = 42
	var largePositiveNumber uint64 = 18446744073709551615

	// Floating point types
	var temperature float32 = 36.5
	var pi float64 = 3.14159265359

	// Boolean type
	var isLearning bool = true
	var isCompleted bool = false

	// String type
	var language string = "Go"
	var greeting = "Hello, World!" // Type inference

	// Character types
	var letter byte = 'A'           // byte is alias for uint8
	var unicodeChar rune = '🚀'     // rune is alias for int32

	// Print all values with their types
	fmt.Printf("Integer (int): %d\n", age)
	fmt.Printf("Large Integer (int64): %d\n", population)
	fmt.Printf("Small Integer (int8): %d\n", smallNumber)
	fmt.Printf("Unsigned Integer (uint): %d\n", positiveNumber)
	fmt.Printf("Large Unsigned (uint64): %d\n", largePositiveNumber)
	fmt.Printf("Float32: %.2f\n", temperature)
	fmt.Printf("Float64: %.10f\n", pi)
	fmt.Printf("Boolean (learning): %t\n", isLearning)
	fmt.Printf("Boolean (completed): %t\n", isCompleted)
	fmt.Printf("String: %s\n", language)
	fmt.Printf("String with inference: %s\n", greeting)
	fmt.Printf("Byte (as number): %d\n", letter)
	fmt.Printf("Byte (as character): %c\n", letter)
	fmt.Printf("Rune (as number): %d\n", unicodeChar)
	fmt.Printf("Rune (as character): %c\n", unicodeChar)
}

// demonstrateConstants shows constant usage
func demonstrateConstants() {
	fmt.Println("🔒 Constants Demonstration:")

	fmt.Printf("App Name: %s\n", AppName)
	fmt.Printf("App Version: %s\n", AppVersion)
	fmt.Printf("Max Users: %d\n", MaxUsers)

	fmt.Printf("HTTP Status OK: %d\n", StatusOK)
	fmt.Printf("HTTP Status Not Found: %d\n", StatusNotFound)
	fmt.Printf("HTTP Status Error: %d\n", StatusError)
}

// demonstrateTypeConversions shows explicit type conversions
func demonstrateTypeConversions() {
	fmt.Println("🔄 Type Conversions Demonstration:")

	// Numeric conversions
	var intValue int = 42
	var floatValue float64 = float64(intValue)
	var float32Value float32 = float32(floatValue)

	fmt.Printf("Original int: %d\n", intValue)
	fmt.Printf("Converted to float64: %.2f\n", floatValue)
	fmt.Printf("Converted to float32: %.2f\n", float32Value)

	// String conversions (using fmt.Sprintf)
	var number int = 123
	var numberAsString string = fmt.Sprintf("%d", number)
	fmt.Printf("Number as int: %d\n", number)
	fmt.Printf("Number as string: %s\n", numberAsString)

	// Boolean to string
	var isActive bool = true
	var boolAsString string = fmt.Sprintf("%t", isActive)
	fmt.Printf("Boolean: %t\n", isActive)
	fmt.Printf("Boolean as string: %s\n", boolAsString)
}

// demonstrateIota shows iota usage patterns
func demonstrateIota() {
	fmt.Println("🔢 Iota Demonstration:")

	// Days of the week
	fmt.Println("Days of the week:")
	fmt.Printf("Sunday: %d\n", Sunday)
	fmt.Printf("Monday: %d\n", Monday)
	fmt.Printf("Tuesday: %d\n", Tuesday)
	fmt.Printf("Wednesday: %d\n", Wednesday)
	fmt.Printf("Thursday: %d\n", Thursday)
	fmt.Printf("Friday: %d\n", Friday)
	fmt.Printf("Saturday: %d\n", Saturday)

	// File permissions with bit operations
	fmt.Println("\nFile permissions (bit flags):")
	fmt.Printf("Read: %d (binary: %08b)\n", ReadPermission, ReadPermission)
	fmt.Printf("Write: %d (binary: %08b)\n", WritePermission, WritePermission)
	fmt.Printf("Execute: %d (binary: %08b)\n", ExecutePermission, ExecutePermission)

	// Combined permissions
	readWrite := ReadPermission | WritePermission
	fullPermissions := ReadPermission | WritePermission | ExecutePermission

	fmt.Printf("Read+Write: %d (binary: %08b)\n", readWrite, readWrite)
	fmt.Printf("Full permissions: %d (binary: %08b)\n", fullPermissions, fullPermissions)

	// Check permissions
	fmt.Printf("\nPermission checks:")
	fmt.Printf("Has read permission: %t\n", fullPermissions&ReadPermission != 0)
	fmt.Printf("Has write permission: %t\n", fullPermissions&WritePermission != 0)
	fmt.Printf("Has execute permission: %t\n", fullPermissions&ExecutePermission != 0)
}
