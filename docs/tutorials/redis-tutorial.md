# 🔴 Redis with Go - Complete Tutorial

Master Redis for caching, session management, and real-time features in Go applications.

## 🎯 Learning Objectives

By the end of this tutorial, you will be able to:
- Connect to Redis from Go applications
- Implement caching strategies
- Use Redis data structures (strings, hashes, lists, sets, sorted sets)
- Implement pub/sub messaging
- Handle sessions and rate limiting
- Use Redis for distributed locks
- Implement leaderboards and real-time features
- Monitor Redis performance

---

## 📚 Table of Contents

1. [Setup and Connection](#setup-and-connection)
2. [Basic Operations](#basic-operations)
3. [Data Structures](#data-structures)
4. [Caching Patterns](#caching-patterns)
5. [Pub/Sub Messaging](#pubsub-messaging)
6. [Advanced Features](#advanced-features)
7. [Best Practices](#best-practices)

---

## 1. Setup and Connection

### Install Redis Client

```bash
# Using go-redis (recommended)
go get github.com/redis/go-redis/v9

# Or redigo
go get github.com/gomodule/redigo/redis
```

### Docker Compose for Local Redis

```yaml
# docker-compose.yml
version: '3.8'

services:
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
    command: redis-server --appendonly yes
    volumes:
      - redis-data:/data

  redis-commander:
    image: rediscommander/redis-commander:latest
    environment:
      - REDIS_HOSTS=local:redis:6379
    ports:
      - "8081:8081"
    depends_on:
      - redis

volumes:
  redis-data:
```

### Basic Connection

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    "github.com/redis/go-redis/v9"
)

type RedisClient struct {
    client *redis.Client
}

func NewRedisClient(addr, password string, db int) (*RedisClient, error) {
    client := redis.NewClient(&redis.Options{
        Addr:         addr,
        Password:     password,
        DB:           db,
        DialTimeout:  5 * time.Second,
        ReadTimeout:  3 * time.Second,
        WriteTimeout: 3 * time.Second,
        PoolSize:     10,
        MinIdleConns: 5,
    })
    
    ctx := context.Background()
    if err := client.Ping(ctx).Err(); err != nil {
        return nil, fmt.Errorf("pinging redis: %w", err)
    }
    
    return &RedisClient{client: client}, nil
}

func (rc *RedisClient) Close() error {
    return rc.client.Close()
}

func main() {
    client, err := NewRedisClient("localhost:6379", "", 0)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    
    fmt.Println("Connected to Redis!")
}
```

---

## 2. Basic Operations

### String Operations

```go
func (rc *RedisClient) Set(ctx context.Context, key, value string, expiration time.Duration) error {
    return rc.client.Set(ctx, key, value, expiration).Err()
}

func (rc *RedisClient) Get(ctx context.Context, key string) (string, error) {
    val, err := rc.client.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", fmt.Errorf("key does not exist")
    }
    return val, err
}

func (rc *RedisClient) Delete(ctx context.Context, keys ...string) error {
    return rc.client.Del(ctx, keys...).Err()
}

func (rc *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
    n, err := rc.client.Exists(ctx, key).Result()
    return n > 0, err
}

// Example usage
func main() {
    ctx := context.Background()
    client, _ := NewRedisClient("localhost:6379", "", 0)
    defer client.Close()
    
    // Set with 1 hour expiration
    client.Set(ctx, "user:123:session", "session-token-xyz", 1*time.Hour)
    
    // Get value
    val, _ := client.Get(ctx, "user:123:session")
    fmt.Println("Session:", val)
    
    // Check existence
    exists, _ := client.Exists(ctx, "user:123:session")
    fmt.Println("Exists:", exists)
}
```

### Increment/Decrement

```go
func (rc *RedisClient) Increment(ctx context.Context, key string) (int64, error) {
    return rc.client.Incr(ctx, key).Result()
}

func (rc *RedisClient) IncrementBy(ctx context.Context, key string, value int64) (int64, error) {
    return rc.client.IncrBy(ctx, key, value).Result()
}

func (rc *RedisClient) Decrement(ctx context.Context, key string) (int64, error) {
    return rc.client.Decr(ctx, key).Result()
}

// Example: Page view counter
func IncrementPageViews(ctx context.Context, rc *RedisClient, pageID string) (int64, error) {
    key := fmt.Sprintf("page:%s:views", pageID)
    return rc.Increment(ctx, key)
}
```

---

## 3. Data Structures

### Hashes (User Profiles)

```go
type UserProfile struct {
    ID       string
    Name     string
    Email    string
    LastSeen time.Time
}

func (rc *RedisClient) SaveUserProfile(ctx context.Context, user *UserProfile) error {
    key := fmt.Sprintf("user:%s:profile", user.ID)
    
    return rc.client.HSet(ctx, key,
        "name", user.Name,
        "email", user.Email,
        "last_seen", user.LastSeen.Unix(),
    ).Err()
}

func (rc *RedisClient) GetUserProfile(ctx context.Context, userID string) (*UserProfile, error) {
    key := fmt.Sprintf("user:%s:profile", userID)
    
    data, err := rc.client.HGetAll(ctx, key).Result()
    if err != nil {
        return nil, err
    }
    
    if len(data) == 0 {
        return nil, fmt.Errorf("user not found")
    }
    
    lastSeen, _ := strconv.ParseInt(data["last_seen"], 10, 64)
    
    return &UserProfile{
        ID:       userID,
        Name:     data["name"],
        Email:    data["email"],
        LastSeen: time.Unix(lastSeen, 0),
    }, nil
}

func (rc *RedisClient) UpdateUserField(ctx context.Context, userID, field, value string) error {
    key := fmt.Sprintf("user:%s:profile", userID)
    return rc.client.HSet(ctx, key, field, value).Err()
}
```

### Lists (Activity Feed)

```go
func (rc *RedisClient) AddActivity(ctx context.Context, userID, activity string) error {
    key := fmt.Sprintf("user:%s:activities", userID)
    
    // Add to beginning of list
    if err := rc.client.LPush(ctx, key, activity).Err(); err != nil {
        return err
    }
    
    // Keep only last 100 activities
    return rc.client.LTrim(ctx, key, 0, 99).Err()
}

func (rc *RedisClient) GetActivities(ctx context.Context, userID string, limit int) ([]string, error) {
    key := fmt.Sprintf("user:%s:activities", userID)
    return rc.client.LRange(ctx, key, 0, int64(limit-1)).Result()
}

// Queue implementation
func (rc *RedisClient) EnqueueJob(ctx context.Context, queue, job string) error {
    return rc.client.RPush(ctx, queue, job).Err()
}

func (rc *RedisClient) DequeueJob(ctx context.Context, queue string) (string, error) {
    return rc.client.LPop(ctx, queue).Result()
}
```

### Sets (Tags, Followers)

```go
func (rc *RedisClient) AddTag(ctx context.Context, lessonID, tag string) error {
    key := fmt.Sprintf("lesson:%s:tags", lessonID)
    return rc.client.SAdd(ctx, key, tag).Err()
}

func (rc *RedisClient) GetTags(ctx context.Context, lessonID string) ([]string, error) {
    key := fmt.Sprintf("lesson:%s:tags", lessonID)
    return rc.client.SMembers(ctx, key).Result()
}

func (rc *RedisClient) HasTag(ctx context.Context, lessonID, tag string) (bool, error) {
    key := fmt.Sprintf("lesson:%s:tags", lessonID)
    return rc.client.SIsMember(ctx, key, tag).Result()
}

// Followers
func (rc *RedisClient) Follow(ctx context.Context, userID, targetID string) error {
    followingKey := fmt.Sprintf("user:%s:following", userID)
    followersKey := fmt.Sprintf("user:%s:followers", targetID)
    
    pipe := rc.client.Pipeline()
    pipe.SAdd(ctx, followingKey, targetID)
    pipe.SAdd(ctx, followersKey, userID)
    
    _, err := pipe.Exec(ctx)
    return err
}

func (rc *RedisClient) GetFollowers(ctx context.Context, userID string) ([]string, error) {
    key := fmt.Sprintf("user:%s:followers", userID)
    return rc.client.SMembers(ctx, key).Result()
}
```

### Sorted Sets (Leaderboard)

```go
func (rc *RedisClient) UpdateScore(ctx context.Context, userID string, score float64) error {
    return rc.client.ZAdd(ctx, "leaderboard", redis.Z{
        Score:  score,
        Member: userID,
    }).Err()
}

func (rc *RedisClient) GetTopUsers(ctx context.Context, limit int) ([]string, error) {
    return rc.client.ZRevRange(ctx, "leaderboard", 0, int64(limit-1)).Result()
}

func (rc *RedisClient) GetUserRank(ctx context.Context, userID string) (int64, error) {
    rank, err := rc.client.ZRevRank(ctx, "leaderboard", userID).Result()
    if err != nil {
        return 0, err
    }
    return rank + 1, nil // 1-based ranking
}

func (rc *RedisClient) GetUserScore(ctx context.Context, userID string) (float64, error) {
    return rc.client.ZScore(ctx, "leaderboard", userID).Result()
}

// Leaderboard with scores
func (rc *RedisClient) GetTopUsersWithScores(ctx context.Context, limit int) ([]redis.Z, error) {
    return rc.client.ZRevRangeWithScores(ctx, "leaderboard", 0, int64(limit-1)).Result()
}
```

---

## 4. Caching Patterns

### Cache-Aside Pattern

```go
type LessonCache struct {
    redis *RedisClient
    db    *sql.DB
}

func (lc *LessonCache) GetLesson(ctx context.Context, id int64) (*Lesson, error) {
    cacheKey := fmt.Sprintf("lesson:%d", id)
    
    // Try cache first
    cached, err := lc.redis.Get(ctx, cacheKey)
    if err == nil {
        var lesson Lesson
        if err := json.Unmarshal([]byte(cached), &lesson); err == nil {
            return &lesson, nil
        }
    }
    
    // Cache miss - fetch from database
    lesson, err := lc.fetchFromDB(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Store in cache
    data, _ := json.Marshal(lesson)
    lc.redis.Set(ctx, cacheKey, string(data), 1*time.Hour)
    
    return lesson, nil
}

func (lc *LessonCache) InvalidateLesson(ctx context.Context, id int64) error {
    cacheKey := fmt.Sprintf("lesson:%d", id)
    return lc.redis.Delete(ctx, cacheKey)
}
```

### Write-Through Cache

```go
func (lc *LessonCache) UpdateLesson(ctx context.Context, lesson *Lesson) error {
    // Update database
    if err := lc.updateDB(ctx, lesson); err != nil {
        return err
    }
    
    // Update cache
    cacheKey := fmt.Sprintf("lesson:%d", lesson.ID)
    data, _ := json.Marshal(lesson)
    return lc.redis.Set(ctx, cacheKey, string(data), 1*time.Hour)
}
```

### Cache with TTL Refresh

```go
func (lc *LessonCache) GetWithRefresh(ctx context.Context, id int64) (*Lesson, error) {
    cacheKey := fmt.Sprintf("lesson:%d", id)
    
    // Get with TTL check
    ttl, err := lc.redis.client.TTL(ctx, cacheKey).Result()
    if err != nil || ttl < 5*time.Minute {
        // Refresh cache
        lesson, err := lc.fetchFromDB(ctx, id)
        if err != nil {
            return nil, err
        }
        
        data, _ := json.Marshal(lesson)
        lc.redis.Set(ctx, cacheKey, string(data), 1*time.Hour)
        return lesson, nil
    }
    
    // Use cached value
    cached, _ := lc.redis.Get(ctx, cacheKey)
    var lesson Lesson
    json.Unmarshal([]byte(cached), &lesson)
    return &lesson, nil
}
```

---

## 5. Pub/Sub Messaging

### Publisher

```go
func (rc *RedisClient) Publish(ctx context.Context, channel, message string) error {
    return rc.client.Publish(ctx, channel, message).Err()
}

// Example: Notify lesson completion
func NotifyLessonCompleted(ctx context.Context, rc *RedisClient, userID string, lessonID int64) error {
    event := map[string]interface{}{
        "user_id":   userID,
        "lesson_id": lessonID,
        "timestamp": time.Now().Unix(),
    }
    
    data, _ := json.Marshal(event)
    return rc.Publish(ctx, "lesson-completed", string(data))
}
```

### Subscriber

```go
func (rc *RedisClient) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
    return rc.client.Subscribe(ctx, channels...)
}

func ListenForEvents(ctx context.Context, rc *RedisClient) {
    pubsub := rc.Subscribe(ctx, "lesson-completed", "user-registered")
    defer pubsub.Close()
    
    ch := pubsub.Channel()
    
    for msg := range ch {
        log.Printf("Received message from %s: %s", msg.Channel, msg.Payload)
        
        switch msg.Channel {
        case "lesson-completed":
            handleLessonCompleted(msg.Payload)
        case "user-registered":
            handleUserRegistered(msg.Payload)
        }
    }
}

func handleLessonCompleted(payload string) {
    var event map[string]interface{}
    json.Unmarshal([]byte(payload), &event)
    log.Printf("User %v completed lesson %v", event["user_id"], event["lesson_id"])
}
```

---


