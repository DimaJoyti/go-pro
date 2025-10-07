# 📘 Lesson 17: Security Best Practices

Welcome to Lesson 17! Security is paramount in modern applications. This lesson covers security best practices, common vulnerabilities, and building secure Go applications.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Implement authentication and authorization
- Secure HTTP communications with TLS
- Prevent common security vulnerabilities
- Handle sensitive data securely
- Implement rate limiting and DDoS protection
- Apply cryptographic best practices
- Build security-first applications

## 📚 Theory

### Authentication & Authorization

**JWT Implementation:**
```go
func generateJWT(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
        "iat":     time.Now().Unix(),
    })
    
    return token.SignedString([]byte(secretKey))
}

func validateJWT(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return []byte(secretKey), nil
    })
}
```

### Input Validation

**SQL Injection Prevention:**
```go
// Bad: Vulnerable to SQL injection
func getUserBad(db *sql.DB, userID string) (*User, error) {
    query := fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", userID)
    // This is vulnerable!
}

// Good: Use prepared statements
func getUser(db *sql.DB, userID string) (*User, error) {
    query := "SELECT id, name, email FROM users WHERE id = ?"
    row := db.QueryRow(query, userID)
    
    var user User
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    return &user, err
}
```

### Cryptography

**Password Hashing:**
```go
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

func checkPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

**Encryption:**
```go
func encrypt(plaintext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
    return ciphertext, nil
}
```

## 💻 Hands-On Examples

### Example 1: Secure HTTP Server
```go
func secureServer() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", secureHandler)
    
    server := &http.Server{
        Addr:         ":8443",
        Handler:      securityMiddleware(mux),
        TLSConfig:    &tls.Config{MinVersion: tls.VersionTLS12},
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    
    log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
}

func securityMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Security headers
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("Strict-Transport-Security", "max-age=31536000")
        
        next.ServeHTTP(w, r)
    })
}
```

### Example 2: Rate Limiting
```go
type RateLimiter struct {
    visitors map[string]*visitor
    mu       sync.RWMutex
    rate     rate.Limit
    burst    int
}

type visitor struct {
    limiter  *rate.Limiter
    lastSeen time.Time
}

func (rl *RateLimiter) Allow(ip string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()
    
    v, exists := rl.visitors[ip]
    if !exists {
        limiter := rate.NewLimiter(rl.rate, rl.burst)
        rl.visitors[ip] = &visitor{limiter, time.Now()}
        return limiter.Allow()
    }
    
    v.lastSeen = time.Now()
    return v.limiter.Allow()
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-17/exercises/`:

1. **Authentication System**: Build JWT-based authentication
2. **Input Validation**: Implement comprehensive input validation
3. **Secure Communication**: Set up TLS and secure headers
4. **Rate Limiting**: Implement rate limiting middleware
5. **Cryptography**: Use encryption and hashing properly
6. **Security Audit**: Audit and secure an existing application

## ✅ Validation

Run the tests to validate your security implementations:

```bash
cd ../../code/lesson-17
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Never trust user input - validate everything
- Use prepared statements to prevent SQL injection
- Implement proper authentication and authorization
- Use HTTPS everywhere with strong TLS configuration
- Hash passwords with bcrypt or similar
- Implement rate limiting to prevent abuse
- Keep dependencies updated for security patches
- Follow the principle of least privilege

## 🛡️ Security Checklist

- [ ] Input validation and sanitization
- [ ] SQL injection prevention
- [ ] XSS protection
- [ ] CSRF protection
- [ ] Secure headers implementation
- [ ] TLS/HTTPS configuration
- [ ] Authentication and authorization
- [ ] Rate limiting and DDoS protection
- [ ] Secure password handling
- [ ] Dependency security scanning

## ➡️ Next Steps

Once you've completed all exercises and security tests pass, move on to:
**[Lesson 18: Deployment and DevOps](../lesson-18/README.md)**

---

**Security first!** 🛡️
