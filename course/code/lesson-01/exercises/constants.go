package exercises

// Exercise 2: Constants and iota Practice
// Complete the constant declarations and functions below

// TODO: Declare constants for mathematical values
// Create constants for:
// - Pi (3.14159)
// - E (2.71828)
// - GoldenRatio (1.61803)
const (
// TODO: Add your mathematical constants here
)

// TODO: Create an enumeration for HTTP status codes using iota
// Create constants for:
// - StatusOK = 200
// - StatusCreated = 201
// - StatusAccepted = 202
// - StatusNoContent = 204
// Hint: Start with StatusOK = 200 + iota, then adjust others
const (
// TODO: Add your HTTP status constants here using iota
)

// TODO: Create an enumeration for log levels using iota
// Create constants for:
// - LogDebug = 0
// - LogInfo = 1
// - LogWarning = 2
// - LogError = 3
// - LogFatal = 4
const (
// TODO: Add your log level constants here using iota
)

// TODO: Create bit flag constants for file permissions using iota
// Create constants for:
// - PermissionRead = 1 (binary: 001)
// - PermissionWrite = 2 (binary: 010)
// - PermissionExecute = 4 (binary: 100)
// Hint: Use 1 << iota for bit shifting
const (
// TODO: Add your permission constants here using bit operations with iota
)

// TODO: Create an enumeration for days of the week starting from 1
// Create constants for:
// - Monday = 1
// - Tuesday = 2
// - ... and so on until Sunday = 7
// Hint: Use iota + 1
const (
// TODO: Add your weekday constants here
)

// TODO: Complete this function
// GetHTTPStatusMessage returns a message for the given HTTP status code
// Parameter: statusCode - HTTP status code (int)
// Returns: status message (string)
// Handle these cases:
// - 200: "OK"
// - 201: "Created"
// - 202: "Accepted"
// - 204: "No Content"
// - default: "Unknown Status"
func GetHTTPStatusMessage(statusCode int) string {
	// TODO: Use a switch statement to return appropriate messages
	// Replace this return statement with your implementation
	return ""
}

// TODO: Complete this function
// GetLogLevelName returns the name of the log level
// Parameter: level - log level constant (int)
// Returns: log level name (string)
// Handle these cases:
// - LogDebug: "DEBUG"
// - LogInfo: "INFO"
// - LogWarning: "WARNING"
// - LogError: "ERROR"
// - LogFatal: "FATAL"
// - default: "UNKNOWN"
func GetLogLevelName(level int) string {
	// TODO: Use a switch statement to return appropriate log level names
	// Replace this return statement with your implementation
	return ""
}

// TODO: Complete this function
// HasPermission checks if the given permissions include the specified permission
// Parameters:
//   - permissions - combined permissions (int)
//   - permission - permission to check (int)
//
// Returns: true if permission is included, false otherwise
// Hint: Use bitwise AND (&) operator
func HasPermission(permissions, permission int) bool {
	// TODO: Check if the permission is set using bitwise operations
	// Replace this return statement with your implementation
	return false
}

// TODO: Complete this function
// AddPermission adds a permission to the existing permissions
// Parameters:
//   - permissions - current permissions (int)
//   - permission - permission to add (int)
//
// Returns: updated permissions (int)
// Hint: Use bitwise OR (|) operator
func AddPermission(permissions, permission int) int {
	// TODO: Add the permission using bitwise operations
	// Replace this return statement with your implementation
	return 0
}

// TODO: Complete this function
// RemovePermission removes a permission from the existing permissions
// Parameters:
//   - permissions - current permissions (int)
//   - permission - permission to remove (int)
//
// Returns: updated permissions (int)
// Hint: Use bitwise AND with NOT (&^) operator
func RemovePermission(permissions, permission int) int {
	// TODO: Remove the permission using bitwise operations
	// Replace this return statement with your implementation
	return 0
}

// TODO: Complete this function
// GetWeekdayName returns the name of the weekday
// Parameter: day - weekday constant (int)
// Returns: weekday name (string)
// Handle these cases:
// - Monday: "Monday"
// - Tuesday: "Tuesday"
// - ... and so on
// - default: "Invalid Day"
func GetWeekdayName(day int) string {
	// TODO: Use a switch statement to return appropriate weekday names
	// Replace this return statement with your implementation
	return ""
}

// TODO: Complete this function
// IsWeekend checks if the given day is a weekend day
// Parameter: day - weekday constant (int)
// Returns: true if it's Saturday or Sunday, false otherwise
func IsWeekend(day int) bool {
	// TODO: Check if the day is Saturday or Sunday
	// Replace this return statement with your implementation
	return false
}

// TODO: Complete this function
// CalculateCircleProperties calculates area and circumference of a circle
// Parameter: radius - radius of the circle (float64)
// Returns: area and circumference (both float64)
// Formulas:
// - Area = π * radius²
// - Circumference = 2 * π * radius
// Use the Pi constant you declared above
func CalculateCircleProperties(radius float64) (float64, float64) {
	// TODO: Calculate area and circumference using your Pi constant
	// Replace these return values with your implementation
	return 0.0, 0.0
}

// TODO: Complete this function
// FormatPermissions returns a string representation of permissions
// Parameter: permissions - combined permissions (int)
// Returns: formatted permissions string (e.g., "rwx", "r--", "rw-")
// Format: "r" for read, "w" for write, "x" for execute, "-" for missing permission
func FormatPermissions(permissions int) string {
	// TODO: Build a string showing which permissions are set
	// Replace this return statement with your implementation
	return ""
}
