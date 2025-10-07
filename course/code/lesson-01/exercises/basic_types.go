package exercises

import "fmt"

// Exercise 1: Basic Types Practice
// Complete the following functions to work with Go's basic types

// PersonInfo represents information about a person
type PersonInfo struct {
	Name   string
	Age    int
	Height float64
	IsStudent bool
}

// TODO: Complete this function
// CreatePersonInfo creates and returns a PersonInfo struct with the given values
// Parameters:
//   - name: person's name (string)
//   - age: person's age (int)
//   - height: person's height in meters (float64)
//   - isStudent: whether the person is a student (bool)
// Returns: PersonInfo struct with the provided values
func CreatePersonInfo(name string, age int, height float64, isStudent bool) PersonInfo {
	// TODO: Create and return a PersonInfo struct with the given parameters
	// Replace this return statement with your implementation
	return PersonInfo{}
}

// TODO: Complete this function
// FormatPersonInfo formats a PersonInfo struct into a readable string
// Parameter: info - PersonInfo struct to format
// Returns: formatted string with person's information
// Expected format: "Name: John, Age: 25, Height: 1.75m, Student: true"
func FormatPersonInfo(info PersonInfo) string {
	// TODO: Use fmt.Sprintf to format the person's information
	// Replace this return statement with your implementation
	return ""
}

// TODO: Complete this function
// CalculateBMI calculates the Body Mass Index given weight and height
// Parameters:
//   - weight: weight in kilograms (float64)
//   - height: height in meters (float64)
// Returns: BMI value (float64)
// Formula: BMI = weight / (height * height)
func CalculateBMI(weight, height float64) float64 {
	// TODO: Calculate and return the BMI
	// Replace this return statement with your implementation
	return 0.0
}

// TODO: Complete this function
// ClassifyBMI classifies BMI into categories
// Parameter: bmi - BMI value (float64)
// Returns: BMI category (string)
// Categories:
//   - BMI < 18.5: "Underweight"
//   - 18.5 <= BMI < 25: "Normal weight"
//   - 25 <= BMI < 30: "Overweight"
//   - BMI >= 30: "Obese"
func ClassifyBMI(bmi float64) string {
	// TODO: Classify the BMI and return the appropriate category
	// Replace this return statement with your implementation
	return ""
}

// TODO: Complete this function
// ConvertTemperature converts temperature between Celsius and Fahrenheit
// Parameters:
//   - temp: temperature value (float64)
//   - fromUnit: source unit, either "C" or "F" (string)
// Returns: converted temperature (float64)
// Formulas:
//   - Celsius to Fahrenheit: F = C * 9/5 + 32
//   - Fahrenheit to Celsius: C = (F - 32) * 5/9
func ConvertTemperature(temp float64, fromUnit string) float64 {
	// TODO: Convert temperature based on the fromUnit
	// If fromUnit is "C", convert from Celsius to Fahrenheit
	// If fromUnit is "F", convert from Fahrenheit to Celsius
	// Replace this return statement with your implementation
	return 0.0
}

// TODO: Complete this function
// IsValidAge checks if an age is valid (between 0 and 150)
// Parameter: age - age to validate (int)
// Returns: true if age is valid, false otherwise
func IsValidAge(age int) bool {
	// TODO: Check if age is between 0 and 150 (inclusive)
	// Replace this return statement with your implementation
	return false
}

// TODO: Complete this function
// GetAgeCategory returns the age category for a given age
// Parameter: age - age to categorize (int)
// Returns: age category (string)
// Categories:
//   - 0-12: "Child"
//   - 13-19: "Teenager"
//   - 20-64: "Adult"
//   - 65+: "Senior"
func GetAgeCategory(age int) string {
	// TODO: Determine and return the age category
	// Replace this return statement with your implementation
	return ""
}

// TODO: Complete this function
// CalculateCircleArea calculates the area of a circle given its radius
// Parameter: radius - radius of the circle (float64)
// Returns: area of the circle (float64)
// Formula: Area = π * radius²
// Use 3.14159 for π
func CalculateCircleArea(radius float64) float64 {
	// TODO: Calculate and return the circle area
	// Replace this return statement with your implementation
	return 0.0
}

// TODO: Complete this function
// IsEven checks if a number is even
// Parameter: number - number to check (int)
// Returns: true if number is even, false if odd
func IsEven(number int) bool {
	// TODO: Check if the number is even using the modulo operator (%)
	// Replace this return statement with your implementation
	return false
}

// TODO: Complete this function
// MaxOfThree returns the maximum of three integers
// Parameters: a, b, c - three integers to compare
// Returns: the largest of the three integers
func MaxOfThree(a, b, c int) int {
	// TODO: Find and return the maximum of the three numbers
	// Replace this return statement with your implementation
	return 0
}
