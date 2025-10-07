# Testing Guide

This document provides comprehensive information about testing in the GO-PRO backend application.

## Table of Contents

- [Overview](#overview)
- [Test Types](#test-types)
- [Running Tests](#running-tests)
- [Writing Tests](#writing-tests)
- [Test Utilities](#test-utilities)
- [Mocking](#mocking)
- [Coverage](#coverage)
- [Best Practices](#best-practices)

## Overview

The GO-PRO backend uses a comprehensive testing strategy that includes:

- **Unit Tests**: Test individual functions and methods in isolation
- **Integration Tests**: Test interactions between components
- **Table-Driven Tests**: Test multiple scenarios with different inputs
- **Benchmark Tests**: Measure performance of critical code paths
- **Mock Objects**: Simulate dependencies for isolated testing

## Test Types

### Unit Tests

Unit tests focus on testing individual functions or methods in isolation. They use mocks to simulate dependencies.

**Location**: `*_test.go` files alongside the code being tested

**Example**:
```go
func TestCourseService_Create(t *testing.T) {
    // Arrange
    mockRepo := testutil.NewMockCourseRepository()
    service := NewCourseService(mockRepo, ...)
    
    // Act
    err := service.Create(ctx, course)
    
    // Assert
    require.NoError(t, err)
}
```

### Integration Tests

Integration tests verify that different components work together correctly. They use real database connections and external services.

**Location**: Files with `// +build integration` tag

**Example**:
```go
// +build integration

func TestCourseRepository_Create(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }
    
    db := testutil.NewTestDB(t)
    defer db.Cleanup()
    // ... test implementation
}
```

### Table-Driven Tests

Table-driven tests allow testing multiple scenarios with different inputs in a structured way.

**Example**:
```go
func TestValidation(t *testing.T) {
    tests := []struct {
        name        string
        input       string
        wantErr     bool
        errContains string
    }{
        {
            name:    "valid input",
            input:   "valid@example.com",
            wantErr: false,
        },
        {
            name:        "invalid input",
            input:       "invalid",
            wantErr:     true,
            errContains: "invalid email",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := Validate(tt.input)
            if tt.wantErr {
                require.Error(t, err)
                assert.Contains(t, err.Error(), tt.errContains)
            } else {
                require.NoError(t, err)
            }
        })
    }
}
```

### Benchmark Tests

Benchmark tests measure the performance of code.

**Example**:
```go
func BenchmarkCourseService_Create(b *testing.B) {
    service := setupService()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = service.Create(ctx, course)
    }
}
```

## Running Tests

### All Tests

```bash
make test
```

### Unit Tests Only

```bash
make test-unit
# or
go test -short ./...
```

### Integration Tests Only

```bash
make test-integration
# or
go test -tags=integration ./...
```

### With Coverage

```bash
make test-coverage
```

### With Race Detection

```bash
make test-race
```

### Benchmark Tests

```bash
make test-bench
```

### Specific Package

```bash
go test ./internal/service/...
```

### Specific Test

```bash
go test -run TestCourseService_Create ./internal/service/
```

### Verbose Output

```bash
make test-verbose
# or
go test -v ./...
```

## Writing Tests

### Test Structure

Follow the **Arrange-Act-Assert** pattern:

```go
func TestFunction(t *testing.T) {
    // Arrange: Set up test data and dependencies
    mockRepo := testutil.NewMockCourseRepository()
    service := NewService(mockRepo)
    
    // Act: Execute the function being tested
    result, err := service.DoSomething(ctx, input)
    
    // Assert: Verify the results
    require.NoError(t, err)
    assert.Equal(t, expected, result)
}
```

### Using Test Utilities

The `testutil` package provides helpful utilities:

```go
// Create test database
db := testutil.NewTestDB(t)
defer db.Cleanup()

// Create test logger
logger := testutil.NewTestLogger(t)

// Create test data
course := testutil.CreateTestCourse("id", "title")
user := testutil.CreateTestUser("id", "username", "email")

// Assertions
testutil.AssertNoError(t, err)
testutil.AssertEqual(t, expected, actual)
testutil.RequireNoError(t, err)
```

### Using Mocks

Mock implementations are available in `testutil/mocks.go`:

```go
// Mock cache
mockCache := testutil.NewMockCacheManager()
mockCache.Set(ctx, "key", "value", 0)

// Mock repository
mockRepo := testutil.NewMockCourseRepository()
mockRepo.AddCourse(course)

// Mock messaging
mockMessaging := testutil.NewMockMessagingService()

// Verify calls
assert.Equal(t, 1, mockRepo.GetCallCount("Create"))
```

## Test Utilities

### TestDB

Provides a test database connection with cleanup:

```go
db := testutil.NewTestDB(t)
defer db.Cleanup()

// Truncate tables for clean state
db.TruncateTables(ctx, "courses", "lessons")
```

### Test Data Creators

Helper functions to create test data:

- `CreateTestCourse(id, title)` - Create a test course
- `CreateTestLesson(id, courseID, title, order)` - Create a test lesson
- `CreateTestExercise(id, lessonID, title)` - Create a test exercise
- `CreateTestProgress(id, userID, lessonID, status)` - Create progress record
- `CreateTestUser(id, username, email)` - Create a test user

### Assertions

Wrapper functions for common assertions:

- `AssertNoError(t, err)` - Assert no error occurred
- `AssertError(t, err)` - Assert an error occurred
- `AssertEqual(t, expected, actual)` - Assert equality
- `AssertNotNil(t, object)` - Assert not nil
- `RequireNoError(t, err)` - Require no error (fails immediately)

## Mocking

### Available Mocks

1. **MockCacheManager**: Mock implementation of cache.CacheManager
2. **MockCourseRepository**: Mock implementation of repository.CourseRepository
3. **MockMessagingService**: Mock implementation of messaging service

### Creating Custom Mocks

```go
type MockService struct {
    mu    sync.RWMutex
    calls map[string]int
}

func (m *MockService) Method(ctx context.Context) error {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.calls["Method"]++
    return nil
}

func (m *MockService) GetCallCount(method string) int {
    m.mu.RLock()
    defer m.mu.RUnlock()
    return m.calls[method]
}
```

## Coverage

### Generate Coverage Report

```bash
make test-coverage
```

This generates:
- `coverage/coverage.out` - Coverage data
- `coverage/coverage.html` - HTML report

### View Coverage Report

```bash
make test-coverage-view
```

### Coverage Goals

- **Overall**: Aim for 80%+ coverage
- **Critical Paths**: 90%+ coverage for authentication, authorization, payment
- **Business Logic**: 85%+ coverage for service layer
- **Handlers**: 75%+ coverage for HTTP handlers

## Best Practices

### 1. Test Naming

Use descriptive test names that explain what is being tested:

```go
func TestCourseService_Create_WithValidData_ShouldSucceed(t *testing.T)
func TestCourseService_Create_WithInvalidData_ShouldReturnError(t *testing.T)
```

### 2. Table-Driven Tests

Use table-driven tests for multiple scenarios:

```go
tests := []struct {
    name    string
    input   Input
    want    Output
    wantErr bool
}{
    // test cases
}
```

### 3. Parallel Tests

Run independent tests in parallel:

```go
func TestSomething(t *testing.T) {
    t.Parallel()
    // test implementation
}
```

### 4. Test Isolation

Each test should be independent and not rely on other tests:

```go
func TestFunction(t *testing.T) {
    // Clean state
    db.TruncateTables(ctx, "table")
    
    // Test implementation
}
```

### 5. Use Subtests

Group related tests using subtests:

```go
func TestCourseService(t *testing.T) {
    t.Run("Create", func(t *testing.T) {
        // test create
    })
    
    t.Run("Update", func(t *testing.T) {
        // test update
    })
}
```

### 6. Mock Verification

Verify that mocks were called as expected:

```go
assert.Equal(t, 1, mockRepo.GetCallCount("Create"))
assert.Equal(t, 1, mockCache.GetCallCount("Set"))
```

### 7. Error Testing

Always test error cases:

```go
if tt.wantErr {
    require.Error(t, err)
    assert.Contains(t, err.Error(), tt.errContains)
} else {
    require.NoError(t, err)
}
```

### 8. Cleanup

Always clean up resources:

```go
db := testutil.NewTestDB(t)
defer db.Cleanup()
```

### 9. Context Usage

Always use context in tests:

```go
ctx := context.Background()
// or with timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```

### 10. Skip Long Tests

Skip long-running tests in short mode:

```go
if testing.Short() {
    t.Skip("Skipping integration test")
}
```

## CI/CD Integration

Tests are automatically run in CI/CD pipelines:

```bash
# Run all CI checks
make ci

# Run only tests
make ci-test

# Run only linter
make ci-lint
```

## Troubleshooting

### Tests Failing Locally

1. Ensure test database is running:
   ```bash
   docker-compose -f docker-compose.test.yml up -d
   ```

2. Check environment variables:
   ```bash
   source .env.test
   ```

3. Clean and rebuild:
   ```bash
   make clean
   make build
   ```

### Race Conditions

Run with race detector:
```bash
make test-race
```

### Slow Tests

Identify slow tests:
```bash
go test -v ./... | grep -E "PASS|FAIL" | sort -k3 -n
```

## Resources

- [Go Testing Package](https://pkg.go.dev/testing)
- [Testify Documentation](https://github.com/stretchr/testify)
- [Table-Driven Tests](https://github.com/golang/go/wiki/TableDrivenTests)
- [Go Test Comments](https://github.com/golang/go/wiki/TestComments)
