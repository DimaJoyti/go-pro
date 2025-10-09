# Completion Guide - Fun Directory Restructure

This guide provides step-by-step instructions to complete the remaining phases of the restructure.

## Quick Commands Reference

```bash
# Navigate to project
cd /home/dima/Desktop/FUN/go-pro/basic/examples/fun

# Run tests
go test ./... -v

# Run specific demo
go run cmd/examples/datastructures/stack_demo.go

# Build all
go build ./...

# Format code
go fmt ./...

# Tidy dependencies
go mod tidy
```

---

## Phase 4: Concurrency Patterns

### Step 1: Create Basic Goroutines Package

```bash
# Create file
touch pkg/concurrency/basics.go
```

**Content Template:**
```go
package concurrency

import (
    "fmt"
    "sync"
    "time"
)

// RunConcurrent demonstrates basic goroutine usage
func RunConcurrent(tasks []func()) {
    var wg sync.WaitGroup
    for _, task := range tasks {
        wg.Add(1)
        go func(t func()) {
            defer wg.Done()
            t()
        }(task)
    }
    wg.Wait()
}

// More functions...
```

### Step 2: Migrate Producer-Consumer

Source: `producer_consumer.go`
Target: `pkg/concurrency/producer_consumer.go`

**Key fixes:**
- Replace `"-".repeat(60)` with `strings.Repeat("-", 60)`
- Use `utils.PrintSubHeader()` instead

### Step 3: Migrate Rate Limiter

Source: `rate_limiter.go`
Target: `pkg/concurrency/ratelimiter.go`

**Key fixes:**
- Remove `func (s string) repeat(count int)` method
- Use `strings.Repeat()` from stdlib

### Step 4: Create Demos

```bash
touch cmd/examples/concurrency/goroutines_demo.go
touch cmd/examples/concurrency/channels_demo.go
touch cmd/examples/concurrency/producer_consumer_demo.go
```

---

## Phase 5: Cache & Advanced

### Step 1: Upgrade Cache with Generics

Source: `cache-demo/cache.go`
Target: `pkg/cache/cache.go`

**Changes needed:**
```go
// Old
type CacheItem struct {
    Value      interface{}
    Expiration time.Time
}

// New
type CacheItem[V any] struct {
    Value      V
    Expiration time.Time
}

type Cache[K comparable, V any] struct {
    items map[K]*CacheItem[V]
    mu    sync.RWMutex
    ttl   time.Duration
}
```

### Step 2: Handle Weak Pointers

**Option A:** Remove `weekPointers.go` (recommended)
**Option B:** Upgrade to Go 1.24 in go.mod

```bash
# If removing
rm weekPointers.go

# If upgrading
# Edit go.mod: change go 1.23 to go 1.24
```

### Step 3: Create LRU Cache

```bash
touch pkg/cache/lru.go
```

---

## Phase 6: Basics & Fundamentals

### Migration Pattern

For each file, follow this pattern:

**Example: variables.go**

```bash
# 1. Copy to new location
cp variables.go cmd/examples/basics/variables.go

# 2. Edit the file
# - Keep package main
# - Keep main() function
# - Remove conflicts with other files
# - Add imports for utils

# 3. Test
go run cmd/examples/basics/variables.go
```

### Files to Migrate:

1. `variables.go` → `cmd/examples/basics/variables.go`
   - No changes needed, already standalone

2. `function.go` → `cmd/examples/basics/functions.go`
   - Fix: Rename `sum` function to `sumNumbers` to avoid conflict

3. `pointer.go` → `cmd/examples/basics/pointers.go`
   - No changes needed

4. `struct.go` → `cmd/examples/basics/structs.go`
   - No changes needed

5. `interface.go` → `cmd/examples/basics/interfaces.go`
   - No changes needed

6. `loop.go` → `cmd/examples/basics/loops.go`
   - No changes needed

7. `iota.go` → `cmd/examples/basics/iota.go`
   - No changes needed

8. `select.go` → `cmd/examples/basics/select.go`
   - No changes needed

---

## Phase 7: Specialized Examples

### JSON Parser Fix

Source: `json_parser.go`

**Problem:** Incomplete User struct

**Fix:**
```go
type User struct {
    ID       int      `json:"id"`
    Name     string   `json:"name"`
    Email    string   `json:"email"`
    Age      int      `json:"age"`
    Active   bool     `json:"active"`
    Tags     []string `json:"tags"`
    Metadata Metadata `json:"metadata"`
}

type Metadata struct {
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    Role      string `json:"role"`
}
```

### Orders Merge

Combine `order.go` and `orders.go`:

```bash
# Create new file
touch cmd/examples/advanced/orders_demo.go

# Merge both files, rename main functions
# order.go: mainOrder() → demoSimpleOrders()
# orders.go: mainOrders() → demoOrderManagement()
# Create new main() that calls both
```

### Word Counter

Source: `word_counter.go`
Target: `cmd/examples/advanced/wordcount_demo.go`

**Fixes:**
- Already mostly good
- Just move and test

---

## Phase 8: Cleanup & Documentation

### Step 1: Remove Old Files

```bash
# List files to remove
ls -la basic/examples/fun/*.go

# Remove them (be careful!)
rm basic/examples/fun/LIFO.go
rm basic/examples/fun/binary_search.go
rm basic/examples/fun/context_timeout.go
# ... etc for all migrated files

# Keep only:
# - go.mod
# - go.sum  
# - README.md
# - IMPLEMENTATION_STATUS.md
# - COMPLETION_GUIDE.md
```

### Step 2: Verify All Imports

```bash
# Check for broken imports
go build ./...

# If errors, fix imports in affected files
```

### Step 3: Run All Tests

```bash
# Run tests
go test ./... -v

# Run with coverage
go test ./... -cover

# Run benchmarks
go test ./... -bench=.
```

### Step 4: Create Missing Tests

Create `test/algorithms_test.go`:

```go
package test

import (
    "testing"
    "github.com/DimaJoyti/go-pro/basic/examples/fun/pkg/algorithms"
)

func TestBinarySearch(t *testing.T) {
    arr := []int{1, 3, 5, 7, 9}
    
    tests := []struct {
        target   int
        expected int
    }{
        {5, 2},
        {1, 0},
        {9, 4},
        {4, -1},
    }
    
    for _, tt := range tests {
        result := algorithms.BinarySearch(arr, tt.target)
        if result != tt.expected {
            t.Errorf("BinarySearch(%d) = %d; want %d", tt.target, result, tt.expected)
        }
    }
}

// More tests...
```

---

## Common Issues & Solutions

### Issue 1: Multiple main() Functions

**Error:** `main redeclared in this block`

**Solution:** Each `.go` file in `cmd/examples/` should be in its own subdirectory OR have unique package names.

### Issue 2: string.repeat() Method

**Error:** `cannot define new methods on non-local type string`

**Solution:** Replace with `strings.Repeat(s, count)`

### Issue 3: Import Errors

**Error:** `package X is not in GOROOT`

**Solution:**
```bash
go mod tidy
go get <missing-package>
```

### Issue 4: Function Conflicts

**Error:** `sum redeclared in this block`

**Solution:** Rename one of the functions or move to different packages

---

## Testing Checklist

- [ ] All data structure tests pass
- [ ] All algorithm tests pass
- [ ] All concurrency tests pass
- [ ] All cache tests pass
- [ ] All benchmarks run
- [ ] No compilation errors
- [ ] All demos run successfully
- [ ] Code coverage > 80%

---

## Final Verification

```bash
# 1. Clean build
go clean ./...
go build ./...

# 2. Run all tests
go test ./... -v -cover

# 3. Run all demos
for demo in cmd/examples/*/*.go; do
    echo "Running $demo"
    go run "$demo"
done

# 4. Check formatting
go fmt ./...

# 5. Run linter (if installed)
golangci-lint run ./...
```

---

## Success Criteria

✅ All phases complete
✅ All tests passing
✅ All demos running
✅ No old files in root
✅ Documentation complete
✅ Code coverage > 80%
✅ No linter warnings
✅ README updated
✅ Examples work as documented

---

## Need Help?

1. Check `IMPLEMENTATION_STATUS.md` for current progress
2. Review completed phases for patterns
3. Look at existing tests for examples
4. Check `pkg/utils/helpers.go` for utility functions

---

**Good luck completing the restructure! 🚀**

