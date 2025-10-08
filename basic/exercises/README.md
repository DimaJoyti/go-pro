# Go Practice Exercises

This directory contains hands-on coding exercises organized by difficulty level. Each exercise includes both a challenge version and a solution.

## 📂 Directory Structure

```
exercises/
├── 01_basics/           # Beginner-level exercises
├── 02_intermediate/     # Intermediate-level exercises
└── 03_advanced/         # Advanced-level exercises
```

## 🎯 Exercise List

### 01_basics - Beginner Level

Perfect for those just starting with Go.

#### 1. FizzBuzz
- **File**: `fizzbuzz.go` / `fizzbuzz_solution.go`
- **Concepts**: Control flow, conditionals, loops
- **Challenge**: Print numbers 1-100, replacing multiples of 3 with "Fizz", multiples of 5 with "Buzz", and multiples of both with "FizzBuzz"
- **Skills**: 
  - If-else statements
  - Modulo operator
  - String formatting

#### 2. Reverse String
- **File**: `reverse_string.go` / `reverse_string_solution.go`
- **Concepts**: Strings, runes, Unicode
- **Challenge**: Reverse a string while properly handling Unicode characters
- **Skills**:
  - String to rune conversion
  - Slice manipulation
  - Unicode awareness

### 02_intermediate - Intermediate Level

For those comfortable with Go basics.

#### 1. URL Shortener
- **File**: `url_shortener.go` / `url_shortener_solution.go`
- **Concepts**: Maps, structs, random generation
- **Challenge**: Build a URL shortening service
- **Skills**:
  - Map operations
  - Random string generation
  - Collision handling
  - Service design patterns

**Features to implement**:
- Generate short codes for long URLs
- Store URL mappings
- Retrieve original URLs
- Handle duplicate URLs
- Prevent collisions

### 03_advanced - Advanced Level

For experienced Go developers.

#### 1. Concurrent Web Crawler
- **File**: `web_crawler.go` / `web_crawler_solution.go`
- **Concepts**: Goroutines, channels, synchronization
- **Challenge**: Build a concurrent web crawler
- **Skills**:
  - Goroutine management
  - Channel communication
  - WaitGroups
  - Mutex for thread safety
  - Depth-limited traversal

**Features to implement**:
- Concurrent URL fetching
- Avoid visiting URLs twice
- Limit crawling depth
- Thread-safe visited tracking
- Graceful shutdown

## 🚀 How to Use

### 1. Try the Challenge First

```bash
cd 01_basics
go run fizzbuzz.go
```

### 2. Check Your Solution

After attempting the exercise, compare with the solution:

```bash
go run fizzbuzz_solution.go
```

### 3. Experiment and Modify

- Modify the code to add new features
- Try different approaches
- Optimize for performance
- Add error handling

## 📝 Exercise Guidelines

### Before Starting
1. Read the exercise description carefully
2. Understand the requirements
3. Think about the approach
4. Consider edge cases

### While Coding
1. Write clean, readable code
2. Add comments for complex logic
3. Handle errors appropriately
4. Test with different inputs

### After Completing
1. Compare with the solution
2. Understand different approaches
3. Optimize if needed
4. Add additional features

## 🎓 Learning Tips

### For Beginners (01_basics)
- Focus on understanding basic syntax
- Don't worry about optimization yet
- Practice writing clean code
- Test with various inputs

### For Intermediate (02_intermediate)
- Think about code organization
- Consider edge cases
- Practice using Go's standard library
- Focus on idiomatic Go

### For Advanced (03_advanced)
- Master concurrency patterns
- Understand synchronization primitives
- Practice debugging concurrent code
- Consider performance implications

## 🔍 Common Patterns to Learn

### From FizzBuzz
- Conditional logic
- Loop patterns
- String formatting

### From Reverse String
- Working with Unicode
- Slice operations
- In-place algorithms

### From URL Shortener
- Map-based storage
- Random generation
- Collision detection
- Service patterns

### From Web Crawler
- Worker pools
- Channel patterns
- Synchronization
- Concurrent data structures

## 📊 Difficulty Progression

```
01_basics          → Foundation concepts
    ↓
02_intermediate    → Practical applications
    ↓
03_advanced        → Complex patterns & concurrency
```

## 🎯 Next Steps

After completing these exercises:

1. **Explore More Examples**: Check `../examples/` for more code samples
2. **Build Projects**: Apply concepts in `../projects/`
3. **Read Documentation**: Study Go's official documentation
4. **Contribute**: Add your own exercises or improvements

## 📚 Additional Resources

- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Tour](https://tour.golang.org/)
- [Go Playground](https://play.golang.org/)

## 🤝 Contributing

Have an idea for a new exercise?

1. Create the challenge file
2. Create the solution file
3. Add documentation
4. Test thoroughly
5. Submit a pull request

## ✅ Checklist

Track your progress:

### Basics
- [ ] FizzBuzz
- [ ] Reverse String

### Intermediate
- [ ] URL Shortener

### Advanced
- [ ] Web Crawler

## 💡 Challenge Yourself

After completing the exercises:

1. **Optimize**: Make your solutions more efficient
2. **Extend**: Add new features
3. **Test**: Write unit tests
4. **Benchmark**: Measure performance
5. **Refactor**: Improve code quality

## 🏆 Mastery Goals

- ✅ Complete all basic exercises
- ✅ Complete all intermediate exercises
- ✅ Complete all advanced exercises
- ✅ Write tests for your solutions
- ✅ Optimize for performance
- ✅ Create your own exercises

Happy Coding! 🚀

