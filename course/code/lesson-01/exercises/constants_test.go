package exercises

import (
	"math"
	"testing"
)

func TestMathematicalConstants(t *testing.T) {
	// Test that mathematical constants are defined correctly
	tests := []struct {
		name      string
		constant  float64
		expected  float64
		tolerance float64
	}{
		{"Pi", Pi, 3.14159, 0.00001},
		{"E", E, 2.71828, 0.00001},
		{"GoldenRatio", GoldenRatio, 1.61803, 0.00001},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if math.Abs(tt.constant-tt.expected) > tt.tolerance {
				t.Errorf("%s = %f, want %f", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestHTTPStatusConstants(t *testing.T) {
	// Test that HTTP status constants are defined correctly
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"StatusOK", StatusOK, 200},
		{"StatusCreated", StatusCreated, 201},
		{"StatusAccepted", StatusAccepted, 202},
		{"StatusNoContent", StatusNoContent, 204},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestLogLevelConstants(t *testing.T) {
	// Test that log level constants are defined correctly
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"LogDebug", LogDebug, 0},
		{"LogInfo", LogInfo, 1},
		{"LogWarning", LogWarning, 2},
		{"LogError", LogError, 3},
		{"LogFatal", LogFatal, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestPermissionConstants(t *testing.T) {
	// Test that permission constants are defined correctly
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"PermissionRead", PermissionRead, 1},
		{"PermissionWrite", PermissionWrite, 2},
		{"PermissionExecute", PermissionExecute, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestWeekdayConstants(t *testing.T) {
	// Test that weekday constants are defined correctly
	tests := []struct {
		name     string
		constant int
		expected int
	}{
		{"Monday", Monday, 1},
		{"Tuesday", Tuesday, 2},
		{"Wednesday", Wednesday, 3},
		{"Thursday", Thursday, 4},
		{"Friday", Friday, 5},
		{"Saturday", Saturday, 6},
		{"Sunday", Sunday, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.constant != tt.expected {
				t.Errorf("%s = %d, want %d", tt.name, tt.constant, tt.expected)
			}
		})
	}
}

func TestGetHTTPStatusMessage(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		want       string
	}{
		{"OK", 200, "OK"},
		{"Created", 201, "Created"},
		{"Accepted", 202, "Accepted"},
		{"No Content", 204, "No Content"},
		{"Unknown", 404, "Unknown Status"},
		{"Unknown negative", -1, "Unknown Status"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetHTTPStatusMessage(tt.statusCode)
			if got != tt.want {
				t.Errorf("GetHTTPStatusMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLogLevelName(t *testing.T) {
	tests := []struct {
		name  string
		level int
		want  string
	}{
		{"Debug", LogDebug, "DEBUG"},
		{"Info", LogInfo, "INFO"},
		{"Warning", LogWarning, "WARNING"},
		{"Error", LogError, "ERROR"},
		{"Fatal", LogFatal, "FATAL"},
		{"Unknown", 99, "UNKNOWN"},
		{"Unknown negative", -1, "UNKNOWN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLogLevelName(tt.level)
			if got != tt.want {
				t.Errorf("GetLogLevelName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasPermission(t *testing.T) {
	tests := []struct {
		name        string
		permissions int
		permission  int
		want        bool
	}{
		{"Has read", PermissionRead, PermissionRead, true},
		{"Has write", PermissionWrite, PermissionWrite, true},
		{"Has execute", PermissionExecute, PermissionExecute, true},
		{"Has read+write, check read", PermissionRead | PermissionWrite, PermissionRead, true},
		{"Has read+write, check write", PermissionRead | PermissionWrite, PermissionWrite, true},
		{"Has read+write, check execute", PermissionRead | PermissionWrite, PermissionExecute, false},
		{"Has all permissions", PermissionRead | PermissionWrite | PermissionExecute, PermissionRead, true},
		{"No permissions", 0, PermissionRead, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasPermission(tt.permissions, tt.permission)
			if got != tt.want {
				t.Errorf("HasPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddPermission(t *testing.T) {
	tests := []struct {
		name        string
		permissions int
		permission  int
		want        int
	}{
		{"Add read to empty", 0, PermissionRead, PermissionRead},
		{"Add write to read", PermissionRead, PermissionWrite, PermissionRead | PermissionWrite},
		{"Add execute to read+write", PermissionRead | PermissionWrite, PermissionExecute, PermissionRead | PermissionWrite | PermissionExecute},
		{"Add existing permission", PermissionRead, PermissionRead, PermissionRead},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddPermission(tt.permissions, tt.permission)
			if got != tt.want {
				t.Errorf("AddPermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemovePermission(t *testing.T) {
	tests := []struct {
		name        string
		permissions int
		permission  int
		want        int
	}{
		{"Remove read from read", PermissionRead, PermissionRead, 0},
		{"Remove write from read+write", PermissionRead | PermissionWrite, PermissionWrite, PermissionRead},
		{"Remove execute from all", PermissionRead | PermissionWrite | PermissionExecute, PermissionExecute, PermissionRead | PermissionWrite},
		{"Remove non-existing permission", PermissionRead, PermissionWrite, PermissionRead},
		{"Remove from empty", 0, PermissionRead, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemovePermission(tt.permissions, tt.permission)
			if got != tt.want {
				t.Errorf("RemovePermission() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWeekdayName(t *testing.T) {
	tests := []struct {
		name string
		day  int
		want string
	}{
		{"Monday", Monday, "Monday"},
		{"Tuesday", Tuesday, "Tuesday"},
		{"Wednesday", Wednesday, "Wednesday"},
		{"Thursday", Thursday, "Thursday"},
		{"Friday", Friday, "Friday"},
		{"Saturday", Saturday, "Saturday"},
		{"Sunday", Sunday, "Sunday"},
		{"Invalid day", 0, "Invalid Day"},
		{"Invalid day negative", -1, "Invalid Day"},
		{"Invalid day large", 10, "Invalid Day"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetWeekdayName(tt.day)
			if got != tt.want {
				t.Errorf("GetWeekdayName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		name string
		day  int
		want bool
	}{
		{"Monday", Monday, false},
		{"Tuesday", Tuesday, false},
		{"Wednesday", Wednesday, false},
		{"Thursday", Thursday, false},
		{"Friday", Friday, false},
		{"Saturday", Saturday, true},
		{"Sunday", Sunday, true},
		{"Invalid day", 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsWeekend(tt.day)
			if got != tt.want {
				t.Errorf("IsWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateCircleProperties(t *testing.T) {
	tests := []struct {
		name              string
		radius            float64
		wantArea          float64
		wantCircumference float64
	}{
		{
			name:              "Unit circle",
			radius:            1.0,
			wantArea:          Pi,
			wantCircumference: 2 * Pi,
		},
		{
			name:              "Small circle",
			radius:            2.0,
			wantArea:          4 * Pi,
			wantCircumference: 4 * Pi,
		},
		{
			name:              "Large circle",
			radius:            5.0,
			wantArea:          25 * Pi,
			wantCircumference: 10 * Pi,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotArea, gotCircumference := CalculateCircleProperties(tt.radius)
			if math.Abs(gotArea-tt.wantArea) > 0.001 {
				t.Errorf("CalculateCircleProperties() area = %v, want %v", gotArea, tt.wantArea)
			}
			if math.Abs(gotCircumference-tt.wantCircumference) > 0.001 {
				t.Errorf("CalculateCircleProperties() circumference = %v, want %v", gotCircumference, tt.wantCircumference)
			}
		})
	}
}

func TestFormatPermissions(t *testing.T) {
	tests := []struct {
		name        string
		permissions int
		want        string
	}{
		{"No permissions", 0, "---"},
		{"Read only", PermissionRead, "r--"},
		{"Write only", PermissionWrite, "-w-"},
		{"Execute only", PermissionExecute, "--x"},
		{"Read and write", PermissionRead | PermissionWrite, "rw-"},
		{"Read and execute", PermissionRead | PermissionExecute, "r-x"},
		{"Write and execute", PermissionWrite | PermissionExecute, "-wx"},
		{"All permissions", PermissionRead | PermissionWrite | PermissionExecute, "rwx"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatPermissions(tt.permissions)
			if got != tt.want {
				t.Errorf("FormatPermissions() = %v, want %v", got, tt.want)
			}
		})
	}
}
