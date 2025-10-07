package solutions

import "fmt"

// PersonInfo represents information about a person
type PersonInfo struct {
	Name      string
	Age       int
	Height    float64
	IsStudent bool
}

// CreatePersonInfo creates and returns a PersonInfo struct with the given values
func CreatePersonInfo(name string, age int, height float64, isStudent bool) PersonInfo {
	return PersonInfo{
		Name:      name,
		Age:       age,
		Height:    height,
		IsStudent: isStudent,
	}
}

// FormatPersonInfo formats a PersonInfo struct into a readable string
func FormatPersonInfo(info PersonInfo) string {
	return fmt.Sprintf("Name: %s, Age: %d, Height: %.2fm, Student: %t",
		info.Name, info.Age, info.Height, info.IsStudent)
}

// CalculateBMI calculates the Body Mass Index given weight and height
func CalculateBMI(weight, height float64) float64 {
	return weight / (height * height)
}

// ClassifyBMI classifies BMI into categories
func ClassifyBMI(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi < 25:
		return "Normal weight"
	case bmi < 30:
		return "Overweight"
	default:
		return "Obese"
	}
}

// ConvertTemperature converts temperature between Celsius and Fahrenheit
func ConvertTemperature(temp float64, fromUnit string) float64 {
	if fromUnit == "C" {
		// Celsius to Fahrenheit: F = C * 9/5 + 32
		return temp*9/5 + 32
	} else if fromUnit == "F" {
		// Fahrenheit to Celsius: C = (F - 32) * 5/9
		return (temp - 32) * 5 / 9
	}
	return temp // Return original if unit is not recognized
}

// IsValidAge checks if an age is valid (between 0 and 150)
func IsValidAge(age int) bool {
	return age >= 0 && age <= 150
}

// GetAgeCategory returns the age category for a given age
func GetAgeCategory(age int) string {
	switch {
	case age <= 12:
		return "Child"
	case age <= 19:
		return "Teenager"
	case age <= 64:
		return "Adult"
	default:
		return "Senior"
	}
}

// CalculateCircleArea calculates the area of a circle given its radius
func CalculateCircleArea(radius float64) float64 {
	const pi = 3.14159
	return pi * radius * radius
}

// IsEven checks if a number is even
func IsEven(number int) bool {
	return number%2 == 0
}

// MaxOfThree returns the maximum of three integers
func MaxOfThree(a, b, c int) int {
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	return max
}
