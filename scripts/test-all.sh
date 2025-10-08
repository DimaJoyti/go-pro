#!/bin/bash

# Comprehensive test script for the entire go-pro repository
# Tests all Go modules and reports results

set -o pipefail

echo "=========================================="
echo "Go-Pro Repository Build Test"
echo "=========================================="
echo ""

# Color codes
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Counters
TOTAL=0
PASSED=0
FAILED=0

# Function to test a module
test_module() {
    local dir=$1
    local name=$2
    
    TOTAL=$((TOTAL + 1))
    
    echo -n "Testing $name... "
    
    if cd "$dir" 2>/dev/null && go build ./... >/dev/null 2>&1; then
        echo -e "${GREEN}✓ PASS${NC}"
        PASSED=$((PASSED + 1))
        cd - >/dev/null 2>&1
        return 0
    else
        echo -e "${RED}✗ FAIL${NC}"
        FAILED=$((FAILED + 1))
        cd - >/dev/null 2>&1
        return 1
    fi
}

# Test basic directory
echo "=== Testing Basic Examples ==="
if [ -f "basic/test-basics.sh" ]; then
    if bash basic/test-basics.sh >/dev/null 2>&1; then
        echo -e "${GREEN}✓ Basic examples: PASS${NC}"
        PASSED=$((PASSED + 1))
    else
        echo -e "${RED}✗ Basic examples: FAIL${NC}"
        FAILED=$((FAILED + 1))
    fi
    TOTAL=$((TOTAL + 1))
fi
echo ""

# Test backend
echo "=== Testing Backend ==="
test_module "backend" "Backend API"
echo ""

# Test services
echo "=== Testing Services ==="
test_module "services/api-gateway" "API Gateway"
test_module "services/shared" "Shared Libraries"
echo ""

# Test course lessons
echo "=== Testing Course Lessons ==="
for lesson_dir in course/code/lesson-*; do
    if [ -d "$lesson_dir" ] && [ -f "$lesson_dir/go.mod" ]; then
        lesson_name=$(basename "$lesson_dir")
        test_module "$lesson_dir" "$lesson_name"
    fi
done
echo ""

# Summary
echo "=========================================="
echo "Test Summary"
echo "=========================================="
echo -e "Total modules tested: $TOTAL"
echo -e "${GREEN}Passed: $PASSED${NC}"
if [ $FAILED -gt 0 ]; then
    echo -e "${RED}Failed: $FAILED${NC}"
else
    echo "Failed: $FAILED"
fi
echo "=========================================="

# Exit with appropriate code
if [ $FAILED -gt 0 ]; then
    echo ""
    echo -e "${YELLOW}Some tests failed. Run individual module tests for details.${NC}"
    exit 1
else
    echo ""
    echo -e "${GREEN}All tests passed! ✓${NC}"
    exit 0
fi

