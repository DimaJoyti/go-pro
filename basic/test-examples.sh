#!/bin/bash

# Test all examples in the basic/examples directory

echo "======================================"
echo "Testing All Go Basic Examples"
echo "======================================"
echo ""

PASS=0
FAIL=0

for dir in basic/examples/0{1..9}_* basic/examples/{10,11,12}_*; do
    if [ -d "$dir" ]; then
        example_name=$(basename "$dir")
        echo "Testing: $example_name"
        
        if cd "$dir" && go run main.go > /dev/null 2>&1; then
            echo "✓ PASS"
            ((PASS++))
        else
            echo "✗ FAIL"
            ((FAIL++))
        fi
        
        cd - > /dev/null
        echo ""
    fi
done

echo "======================================"
echo "Summary:"
echo "  Passed: $PASS"
echo "  Failed: $FAIL"
echo "======================================"

if [ $FAIL -eq 0 ]; then
    echo "All examples passed! 🎉"
    exit 0
else
    echo "Some examples failed."
    exit 1
fi

