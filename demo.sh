#!/bin/bash

# GO-PRO Learning Suite Demo Script
# This script demonstrates the complete learning platform

echo "🚀 GO-PRO: Complete Go Programming Learning Suite Demo"
echo "======================================================"
echo

# Check prerequisites
echo "📋 Checking Prerequisites..."
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21+ from https://go.dev/dl/"
    exit 1
fi

if ! command -v node &> /dev/null; then
    echo "❌ Node.js is not installed. Please install Node.js 18+ from https://nodejs.org/"
    exit 1
fi

echo "✅ Go version: $(go version)"
echo "✅ Node.js version: $(node --version)"
echo

# Demo 1: Course Content
echo "📚 Demo 1: Course Content Structure"
echo "-----------------------------------"
echo "Course overview:"
head -10 course/README.md
echo
echo "Available lessons:"
ls -la course/lessons/
echo
echo "Lesson 1 content preview:"
head -20 course/lessons/lesson-01/README.md
echo

# Demo 2: Interactive Exercises
echo "💻 Demo 2: Interactive Exercises"
echo "--------------------------------"
echo "Running student exercises (with TODOs):"
cd course/code/lesson-01
go run main.go
echo
echo "Running reference solutions:"
go run main.go solutions
echo

# Demo 3: Automated Testing
echo "🧪 Demo 3: Automated Testing System"
echo "-----------------------------------"
echo "Running tests (should fail for unimplemented exercises):"
go test ./exercises/... -v | head -20
echo "... (tests continue)"
echo
echo "This is expected! Students need to implement the exercises."
echo

# Demo 4: Backend API
echo "🔧 Demo 4: Backend Learning Platform API"
echo "----------------------------------------"
cd ../../../backend
echo "Installing backend dependencies..."
go mod tidy > /dev/null 2>&1

echo "Starting backend API in background..."
go run main.go &
BACKEND_PID=$!
sleep 3

echo "Testing API endpoints:"
echo
echo "Health check:"
curl -s http://localhost:8080/api/v1/health | jq '.'
echo
echo "Available courses:"
curl -s http://localhost:8080/api/v1/courses | jq '.data[0] | {id, title, description}'
echo
echo "Course lessons:"
curl -s http://localhost:8080/api/v1/courses/1/lessons | jq '.data[] | {id, title, description}'
echo

# Demo 5: Exercise Submission
echo "📝 Demo 5: Exercise Submission System"
echo "-------------------------------------"
echo "Submitting a sample solution:"
curl -s -X POST http://localhost:8080/api/v1/exercises/1/submit \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "demo_user",
    "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n  name := \"Demo User\"\n  age := 25\n  fmt.Printf(\"Name: %s, Age: %d\", name, age)\n}",
    "language": "go"
  }' | jq '.'
echo

# Demo 6: Progress Tracking
echo "📊 Demo 6: Progress Tracking"
echo "----------------------------"
echo "Checking user progress:"
curl -s http://localhost:8080/api/v1/progress/demo_user | jq '.'
echo

# Cleanup
echo "🧹 Cleaning up..."
kill $BACKEND_PID 2>/dev/null
wait $BACKEND_PID 2>/dev/null
echo

# Demo 7: Frontend (if available)
echo "🌐 Demo 7: Frontend Dashboard"
echo "-----------------------------"
cd ../frontend
if [ -f "package.json" ]; then
    echo "Frontend is available! To start the dashboard:"
    echo "  cd frontend"
    echo "  npm install"
    echo "  npm run dev"
    echo "  # Visit http://localhost:3000"
else
    echo "Frontend setup is ready for development."
fi
echo

# Summary
echo "✨ Demo Complete!"
echo "================"
echo
echo "🎯 What you've seen:"
echo "  ✅ Complete course structure with 15 lessons"
echo "  ✅ Interactive exercises with automated testing"
echo "  ✅ Go-based REST API for learning platform"
echo "  ✅ Progress tracking and exercise submission"
echo "  ✅ Real-time feedback and scoring system"
echo
echo "🚀 Next Steps:"
echo "  1. Start learning: cd course && cat README.md"
echo "  2. Try exercises: cd course/code/lesson-01 && go run main.go"
echo "  3. Run tests: go test ./exercises/..."
echo "  4. Start API: cd backend && go run main.go"
echo "  5. Build frontend: cd frontend && npm run dev"
echo
echo "📖 Full documentation: README.md"
echo "🌐 API docs: http://localhost:8080 (when backend is running)"
echo
echo "Happy learning! 🎉"
