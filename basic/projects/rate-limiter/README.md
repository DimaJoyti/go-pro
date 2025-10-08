# 🚦 API Rate Limiter Middleware

Advanced rate limiting middleware with multiple algorithms, distributed support, and configurable rules.

## 🎯 Features

- ✅ Token bucket algorithm
- ✅ Sliding window implementation
- ✅ Fixed window counter
- ✅ Distributed rate limiting (Redis)
- ✅ Per-user and per-IP limits
- ✅ Configurable rules engine
- ✅ Metrics and monitoring
- ✅ Custom headers (X-RateLimit-*)

## 🏗️ Algorithms

### Token Bucket
- Smooth rate limiting
- Burst handling
- Refill rate configuration

### Sliding Window
- Accurate rate limiting
- No burst issues
- Memory efficient

### Fixed Window
- Simple implementation
- Fast performance
- Predictable behavior

## 📖 Usage

```go
// Basic usage
limiter := ratelimiter.New(ratelimiter.Config{
    Algorithm: ratelimiter.TokenBucket,
    Rate:      100,  // requests
    Per:       time.Minute,
    Burst:     10,
})

http.Handle("/api/", limiter.Middleware(apiHandler))
```

## 🎓 Learning Objectives

- Rate limiting algorithms
- Middleware patterns
- Distributed systems
- Redis integration
- HTTP headers

---

**Status**: Planned | **Difficulty**: Intermediate | **Time**: 6-8 hours

