package exercises

import (
	"math"
	"testing"
)

func TestCreatePersonInfo(t *testing.T) {
	tests := []struct {
		name      string
		inputName string
		age       int
		height    float64
		isStudent bool
		want      PersonInfo
	}{
		{
			name:      "Adult student",
			inputName: "Alice",
			age:       22,
			height:    1.65,
			isStudent: true,
			want:      PersonInfo{Name: "Alice", Age: 22, Height: 1.65, IsStudent: true},
		},
		{
			name:      "Working professional",
			inputName: "Bob",
			age:       35,
			height:    1.80,
			isStudent: false,
			want:      PersonInfo{Name: "Bob", Age: 35, Height: 1.80, IsStudent: false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreatePersonInfo(tt.inputName, tt.age, tt.height, tt.isStudent)
			if got != tt.want {
				t.Errorf("CreatePersonInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatPersonInfo(t *testing.T) {
	tests := []struct {
		name string
		info PersonInfo
		want string
	}{
		{
			name: "Student info",
			info: PersonInfo{Name: "Alice", Age: 22, Height: 1.65, IsStudent: true},
			want: "Name: Alice, Age: 22, Height: 1.65m, Student: true",
		},
		{
			name: "Professional info",
			info: PersonInfo{Name: "Bob", Age: 35, Height: 1.80, IsStudent: false},
			want: "Name: Bob, Age: 35, Height: 1.80m, Student: false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatPersonInfo(tt.info)
			if got != tt.want {
				t.Errorf("FormatPersonInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateBMI(t *testing.T) {
	tests := []struct {
		name   string
		weight float64
		height float64
		want   float64
	}{
		{
			name:   "Normal BMI",
			weight: 70.0,
			height: 1.75,
			want:   22.86,
		},
		{
			name:   "Underweight BMI",
			weight: 50.0,
			height: 1.70,
			want:   17.30,
		},
		{
			name:   "Overweight BMI",
			weight: 85.0,
			height: 1.65,
			want:   31.22,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateBMI(tt.weight, tt.height)
			if math.Abs(got-tt.want) > 0.01 {
				t.Errorf("CalculateBMI() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestClassifyBMI(t *testing.T) {
	tests := []struct {
		name string
		bmi  float64
		want string
	}{
		{
			name: "Underweight",
			bmi:  17.5,
			want: "Underweight",
		},
		{
			name: "Normal weight",
			bmi:  22.0,
			want: "Normal weight",
		},
		{
			name: "Overweight",
			bmi:  27.5,
			want: "Overweight",
		},
		{
			name: "Obese",
			bmi:  32.0,
			want: "Obese",
		},
		{
			name: "Boundary normal",
			bmi:  18.5,
			want: "Normal weight",
		},
		{
			name: "Boundary overweight",
			bmi:  25.0,
			want: "Overweight",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClassifyBMI(tt.bmi)
			if got != tt.want {
				t.Errorf("ClassifyBMI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		name     string
		temp     float64
		fromUnit string
		want     float64
	}{
		{
			name:     "Celsius to Fahrenheit - freezing",
			temp:     0.0,
			fromUnit: "C",
			want:     32.0,
		},
		{
			name:     "Celsius to Fahrenheit - boiling",
			temp:     100.0,
			fromUnit: "C",
			want:     212.0,
		},
		{
			name:     "Fahrenheit to Celsius - freezing",
			temp:     32.0,
			fromUnit: "F",
			want:     0.0,
		},
		{
			name:     "Fahrenheit to Celsius - boiling",
			temp:     212.0,
			fromUnit: "F",
			want:     100.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertTemperature(tt.temp, tt.fromUnit)
			if math.Abs(got-tt.want) > 0.01 {
				t.Errorf("ConvertTemperature() = %.2f, want %.2f", got, tt.want)
			}
		})
	}
}

func TestIsValidAge(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{
			name: "Valid age - child",
			age:  10,
			want: true,
		},
		{
			name: "Valid age - adult",
			age:  30,
			want: true,
		},
		{
			name: "Valid age - senior",
			age:  80,
			want: true,
		},
		{
			name: "Invalid age - negative",
			age:  -5,
			want: false,
		},
		{
			name: "Invalid age - too old",
			age:  200,
			want: false,
		},
		{
			name: "Boundary - minimum valid",
			age:  0,
			want: true,
		},
		{
			name: "Boundary - maximum valid",
			age:  150,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidAge(tt.age)
			if got != tt.want {
				t.Errorf("IsValidAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAgeCategory(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want string
	}{
		{
			name: "Child",
			age:  8,
			want: "Child",
		},
		{
			name: "Teenager",
			age:  16,
			want: "Teenager",
		},
		{
			name: "Adult",
			age:  35,
			want: "Adult",
		},
		{
			name: "Senior",
			age:  70,
			want: "Senior",
		},
		{
			name: "Boundary - child to teenager",
			age:  13,
			want: "Teenager",
		},
		{
			name: "Boundary - teenager to adult",
			age:  20,
			want: "Adult",
		},
		{
			name: "Boundary - adult to senior",
			age:  65,
			want: "Senior",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAgeCategory(tt.age)
			if got != tt.want {
				t.Errorf("GetAgeCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateCircleArea(t *testing.T) {
	tests := []struct {
		name   string
		radius float64
		want   float64
	}{
		{
			name:   "Small circle",
			radius: 1.0,
			want:   3.14159,
		},
		{
			name:   "Medium circle",
			radius: 5.0,
			want:   78.53975,
		},
		{
			name:   "Large circle",
			radius: 10.0,
			want:   314.159,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateCircleArea(tt.radius)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("CalculateCircleArea() = %.5f, want %.5f", got, tt.want)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name   string
		number int
		want   bool
	}{
		{
			name:   "Even positive",
			number: 4,
			want:   true,
		},
		{
			name:   "Odd positive",
			number: 7,
			want:   false,
		},
		{
			name:   "Even negative",
			number: -6,
			want:   true,
		},
		{
			name:   "Odd negative",
			number: -3,
			want:   false,
		},
		{
			name:   "Zero",
			number: 0,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEven(tt.number)
			if got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxOfThree(t *testing.T) {
	tests := []struct {
		name    string
		a, b, c int
		want    int
	}{
		{
			name: "First is max",
			a:    10, b: 5, c: 3,
			want: 10,
		},
		{
			name: "Second is max",
			a:    3, b: 15, c: 7,
			want: 15,
		},
		{
			name: "Third is max",
			a:    4, b: 8, c: 20,
			want: 20,
		},
		{
			name: "All equal",
			a:    5, b: 5, c: 5,
			want: 5,
		},
		{
			name: "Negative numbers",
			a:    -10, b: -5, c: -15,
			want: -5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxOfThree(tt.a, tt.b, tt.c)
			if got != tt.want {
				t.Errorf("MaxOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
