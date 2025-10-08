# 📖 GO-PRO TUTORIALS

Step-by-step tutorials for getting started with Go-Pro projects.

---

## 🚀 Quick Start Tutorial

### Tutorial 1: Your First Go Project (15 minutes)

**Goal**: Build and run the URL Shortener Service

#### Step 1: Navigate to the Project
```bash
cd basic/projects/url-shortener
```

#### Step 2: Explore the Structure
```bash
# View the README
cat README.md

# Check the project structure
tree -L 2
```

#### Step 3: Run Tests
```bash
# Run all tests
make test

# You should see:
# PASS
# ok  	github.com/DimaJoyti/go-pro/basic/projects/url-shortener/tests	0.004s
```

#### Step 4: Build and Run
```bash
# Build the application
make build

# Run the server
make run

# Server starts at http://localhost:8080
```

#### Step 5: Test the API
```bash
# In another terminal, shorten a URL
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com/DimaJoyti/go-pro"}'

# Response:
# {
#   "short_code": "abc123",
#   "short_url": "http://localhost:8080/abc123",
#   "original_url": "https://github.com/DimaJoyti/go-pro"
# }

# Visit the short URL
curl -L http://localhost:8080/abc123

# Get analytics
curl http://localhost:8080/api/analytics/abc123
```

**🎉 Congratulations!** You've successfully run your first Go-Pro project!

---

## 🌤️ Tutorial 2: Building a CLI Application (20 minutes)

**Goal**: Use the Weather CLI to fetch weather data

#### Step 1: Get an API Key
1. Visit [OpenWeatherMap](https://openweathermap.org/api)
2. Sign up for a free account
3. Get your API key

#### Step 2: Setup the Project
```bash
cd basic/projects/weather-cli

# Set your API key
export WEATHER_API_KEY="your-api-key-here"
```

#### Step 3: Build the CLI
```bash
make build

# Binary created at: bin/weather
```

#### Step 4: Get Current Weather
```bash
# Get weather for a city
./bin/weather current --city "London"

# Output:
# ☀️  Weather in London
# ════════════════════════════════════════
# Temperature: 15°C
# Feels Like:  13°C
# Conditions:  Clear sky
# Humidity:    65%
# Wind:        12 km/h NW
```

#### Step 5: Get Forecast
```bash
# Get 5-day forecast
./bin/weather forecast --city "Tokyo"

# Get detailed forecast
./bin/weather forecast --city "Paris" --detailed

# Get JSON output
./bin/weather current --city "New York" --format json
```

#### Step 6: Explore Caching
```bash
# First request (hits API)
time ./bin/weather current --city "Berlin"

# Second request (uses cache - much faster!)
time ./bin/weather current --city "Berlin"
```

**🎉 You've built a production-ready CLI tool!**

---

## 🔐 Tutorial 3: File Encryption (15 minutes)

**Goal**: Encrypt and decrypt files securely

#### Step 1: Setup
```bash
cd basic/projects/file-encryptor
make build
```

#### Step 2: Create a Test File
```bash
echo "This is a secret message!" > secret.txt
cat secret.txt
```

#### Step 3: Encrypt the File
```bash
./bin/encrypt encrypt --input secret.txt

# You'll be prompted for a password
# Enter password: ********
# Confirm password: ********

# Output:
# 🔐 File Encryption Tool
# ════════════════════════════════════════
# Encrypting: secret.txt
# Size: 27 B
# [████████████████████] 100%
# ✓ Encryption complete!
#   Output: secret.txt.enc
```

#### Step 4: Decrypt the File
```bash
./bin/encrypt decrypt --input secret.txt.enc

# Enter password: ********

# Output:
# 🔓 File Decryption Tool
# ════════════════════════════════════════
# Decrypting: secret.txt.enc
# [████████████████████] 100%
# ✓ Decryption complete!
#   Output: secret.txt.dec
```

#### Step 5: Verify
```bash
# Compare original and decrypted
diff secret.txt secret.txt.dec

# No output means files are identical!
```

#### Step 6: Run the Demo
```bash
# Automated demo
make demo
```

**🎉 You've mastered file encryption in Go!**

---

## 📝 Tutorial 4: Building a Blog API (30 minutes)

**Goal**: Create a blog with authentication

#### Step 1: Setup Database
```bash
cd basic/projects/blog-engine

# Create PostgreSQL database
make db-setup

# Run migrations
make db-migrate
```

#### Step 2: Start the Server
```bash
# Set environment variables
export DATABASE_URL="postgres://localhost/blogdb?sslmode=disable"
export JWT_SECRET="your-secret-key"

# Run the server
make run

# Server starts at http://localhost:8080
```

#### Step 3: Register a User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "john",
    "email": "john@example.com",
    "password": "password123",
    "full_name": "John Doe"
  }'

# Response includes user data (without password)
```

#### Step 4: Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'

# Response:
# {
#   "token": "eyJhbGciOiJIUzI1NiIs...",
#   "refresh_token": "eyJhbGciOiJIUzI1NiIs...",
#   "user": {...}
# }

# Save the token for next steps
export TOKEN="your-token-here"
```

#### Step 5: Create a Post
```bash
curl -X POST http://localhost:8080/api/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "My First Blog Post",
    "content": "# Hello World\n\nThis is my first post!",
    "status": "published"
  }'
```

#### Step 6: Get Posts
```bash
# Get all posts (public)
curl http://localhost:8080/api/posts

# Get specific post
curl http://localhost:8080/api/posts/1

# Get post by slug
curl http://localhost:8080/api/posts/slug/my-first-blog-post
```

**🎉 You've built a complete blog API with authentication!**

---

## 🎯 Next Steps

### Beginner Path
1. ✅ Complete Tutorial 1 (URL Shortener)
2. ✅ Complete Tutorial 2 (Weather CLI)
3. ✅ Complete Tutorial 3 (File Encryptor)
4. 📋 Modify projects to add your own features
5. 📋 Combine concepts from multiple projects

### Intermediate Path
1. ✅ Complete Tutorial 4 (Blog Engine)
2. 📋 Explore Job Queue System
3. 📋 Build Rate Limiter middleware
4. 📋 Create Log Aggregation pipeline

### Advanced Path
1. 📋 Study Service Mesh architecture
2. 📋 Implement Time Series Database
3. 📋 Build Container Orchestrator
4. 📋 Deploy to production

---

## 💡 Tips for Success

### Learning Tips
- **Start Small**: Begin with beginner projects
- **Read Code**: Study the implementation details
- **Run Tests**: Understand how tests validate behavior
- **Modify**: Change features and see what happens
- **Document**: Write notes about what you learn

### Development Tips
- **Use Make**: All projects have Makefiles for common tasks
- **Read READMEs**: Each project has comprehensive documentation
- **Check Tests**: Tests show expected behavior
- **Use Docker**: Many projects include Docker support
- **Ask Questions**: Review code comments for explanations

### Debugging Tips
- **Check Logs**: Look at server output for errors
- **Use curl**: Test APIs from command line
- **Read Errors**: Go error messages are descriptive
- **Use Debugger**: Delve is great for Go debugging
- **Test Incrementally**: Test small changes frequently

---

## 📚 Additional Resources

### Official Documentation
- [Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)

### Project-Specific
- Each project has detailed README
- Architecture documentation in `docs/` folders
- API documentation where applicable
- Example usage in test files

### Community
- [Go Forum](https://forum.golangbridge.org/)
- [Go Slack](https://gophers.slack.com/)
- [r/golang](https://reddit.com/r/golang)

---

**Happy Learning! 🚀**

