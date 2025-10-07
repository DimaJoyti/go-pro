package exercises

import (
	"strings"
	"testing"
)

func TestDeclareVariables(t *testing.T) {
	name, age, salary, isEmployed := DeclareVariables()
	
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"name", name, "John Doe"},
		{"age", age, 30},
		{"salary", salary, 75000.50},
		{"isEmployed", isEmployed, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("DeclareVariables() %s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestMultipleDeclarations(t *testing.T) {
	x, y, z, a, b := MultipleDeclarations()
	
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"x", x, 10},
		{"y", y, 20},
		{"z", z, 30},
		{"a", a, "Hello"},
		{"b", b, "World"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("MultipleDeclarations() %s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestBlockDeclaration(t *testing.T) {
	projectName, version, isStable := BlockDeclaration()
	
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"projectName", projectName, "GO-PRO"},
		{"version", version, 2.1},
		{"isStable", isStable, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("BlockDeclaration() %s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestTestVariableScope(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string // Parts that should be in the result
	}{
		{
			name:  "basic scope test",
			input: "test input",
			expected: []string{
				"Input: test input",
				"Package: I'm at package level",
				"Function: I'm in function",
				"Block:",
			},
		},
		{
			name:  "different input",
			input: "another test",
			expected: []string{
				"Input: another test",
				"Package: I'm at package level",
				"Function: I'm in function",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TestVariableScope(tt.input)
			for _, expected := range tt.expected {
				if !strings.Contains(result, expected) {
					t.Errorf("TestVariableScope() result should contain %q, got %q", expected, result)
				}
			}
		})
	}
}

func TestSwapVariables(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		wantA, wantB int
	}{
		{"positive numbers", 5, 10, 10, 5},
		{"negative numbers", -3, 7, 7, -3},
		{"zero values", 0, 42, 42, 0},
		{"same values", 15, 15, 15, 15},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapVariables(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("SwapVariables(%d, %d) = (%d, %d), want (%d, %d)", 
					tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestZeroValues(t *testing.T) {
	zeroInt, zeroFloat, zeroString, zeroBool := ZeroValues()
	
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"zeroInt", zeroInt, 0},
		{"zeroFloat", zeroFloat, 0.0},
		{"zeroString", zeroString, ""},
		{"zeroBool", zeroBool, false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("ZeroValues() %s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestConstantUsage(t *testing.T) {
	pi, maxRetries, appName := ConstantUsage()
	
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"pi", pi, 3.14159},
		{"maxRetries", maxRetries, 5},
		{"appName", appName, "Learning Go"},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("ConstantUsage() %s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestVariableReassignment(t *testing.T) {
	tests := []struct {
		name     string
		initial  int
		expected int
	}{
		{"positive number", 10, 35},  // (10 + 10) * 2 - 5 = 35
		{"zero", 0, 15},              // (0 + 10) * 2 - 5 = 15
		{"negative number", -5, 5},   // (-5 + 10) * 2 - 5 = 5
		{"large number", 100, 215},   // (100 + 10) * 2 - 5 = 215
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VariableReassignment(tt.initial)
			if got != tt.expected {
				t.Errorf("VariableReassignment(%d) = %d, want %d", tt.initial, got, tt.expected)
			}
		})
	}
}

func TestTypeInference(t *testing.T) {
	inferredInt, inferredFloat, inferredString, inferredBool := TypeInference()
	
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"inferredInt", inferredInt, 42},
		{"inferredFloat", inferredFloat, 3.14},
		{"inferredString", inferredString, "Go"},
		{"inferredBool", inferredBool, true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("TypeInference() %s = %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}

func TestShadowingExample(t *testing.T) {
	tests := []struct {
		name     string
		outer    int
		expected []string // Parts that should be in the result
	}{
		{
			name:  "basic shadowing",
			outer: 10,
			expected: []string{
				"Outer: 10",
				"Inner:",
			},
		},
		{
			name:  "negative outer",
			outer: -5,
			expected: []string{
				"Outer: -5",
				"Inner:",
			},
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ShadowingExample(tt.outer)
			for _, expected := range tt.expected {
				if !strings.Contains(result, expected) {
					t.Errorf("ShadowingExample() result should contain %q, got %q", expected, result)
				}
			}
		})
	}
}
