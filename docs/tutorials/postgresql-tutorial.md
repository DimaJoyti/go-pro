# 🐘 PostgreSQL with Go - Complete Tutorial

Master PostgreSQL database integration in Go applications with production-ready patterns.

## 🎯 Learning Objectives

By the end of this tutorial, you will be able to:
- Connect to PostgreSQL from Go applications
- Execute queries and handle results
- Implement the repository pattern
- Use prepared statements and transactions
- Handle NULL values and complex types
- Implement connection pooling
- Perform database migrations
- Write testable database code

---

## 📚 Table of Contents

1. [Setup and Connection](#setup-and-connection)
2. [Basic CRUD Operations](#basic-crud-operations)
3. [Advanced Queries](#advanced-queries)
4. [Transactions](#transactions)
5. [Connection Pooling](#connection-pooling)
6. [Repository Pattern](#repository-pattern)
7. [Migrations](#migrations)
8. [Testing](#testing)
9. [Best Practices](#best-practices)

---

## 1. Setup and Connection

### Install PostgreSQL Driver

```bash
go get github.com/lib/pq
# Or use pgx for better performance
go get github.com/jackc/pgx/v5
```

### Basic Connection

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    
    _ "github.com/lib/pq"
)

type Config struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func NewConnection(cfg Config) (*sql.DB, error) {
    // Build connection string
    connStr := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
    )
    
    // Open connection
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("opening database: %w", err)
    }
    
    // Verify connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("pinging database: %w", err)
    }
    
    return db, nil
}

func main() {
    cfg := Config{
        Host:     "localhost",
        Port:     5432,
        User:     "postgres",
        Password: "password",
        DBName:   "gopro",
        SSLMode:  "disable",
    }
    
    db, err := NewConnection(cfg)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    fmt.Println("Connected to PostgreSQL!")
}
```

### Using pgx (Recommended)

```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
    config, err := pgxpool.ParseConfig(connString)
    if err != nil {
        return nil, fmt.Errorf("parsing config: %w", err)
    }
    
    // Configure pool
    config.MaxConns = 25
    config.MinConns = 5
    
    pool, err := pgxpool.NewWithConfig(ctx, config)
    if err != nil {
        return nil, fmt.Errorf("creating pool: %w", err)
    }
    
    // Verify connection
    if err := pool.Ping(ctx); err != nil {
        return nil, fmt.Errorf("pinging database: %w", err)
    }
    
    return pool, nil
}

func main() {
    ctx := context.Background()
    
    connString := "postgres://postgres:password@localhost:5432/gopro?sslmode=disable"
    
    pool, err := NewPgxPool(ctx, connString)
    if err != nil {
        log.Fatal(err)
    }
    defer pool.Close()
    
    fmt.Println("Connected to PostgreSQL with pgx!")
}
```

---

## 2. Basic CRUD Operations

### Define Models

```go
package models

import (
    "time"
)

type User struct {
    ID        int64     `json:"id"`
    Email     string    `json:"email"`
    Name      string    `json:"name"`
    Active    bool      `json:"active"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Lesson struct {
    ID          int64     `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Duration    int       `json:"duration"` // minutes
    Difficulty  string    `json:"difficulty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### Create Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    duration INTEGER NOT NULL,
    difficulty VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_lessons_difficulty ON lessons(difficulty);
```

### Insert (Create)

```go
func CreateUser(ctx context.Context, db *sql.DB, user *User) error {
    query := `
        INSERT INTO users (email, name, active)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `
    
    err := db.QueryRowContext(
        ctx,
        query,
        user.Email,
        user.Name,
        user.Active,
    ).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
    
    if err != nil {
        return fmt.Errorf("creating user: %w", err)
    }
    
    return nil
}

// Batch insert
func CreateUsers(ctx context.Context, db *sql.DB, users []User) error {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return fmt.Errorf("beginning transaction: %w", err)
    }
    defer tx.Rollback()
    
    stmt, err := tx.PrepareContext(ctx, `
        INSERT INTO users (email, name, active)
        VALUES ($1, $2, $3)
        RETURNING id, created_at, updated_at
    `)
    if err != nil {
        return fmt.Errorf("preparing statement: %w", err)
    }
    defer stmt.Close()
    
    for i := range users {
        err := stmt.QueryRowContext(
            ctx,
            users[i].Email,
            users[i].Name,
            users[i].Active,
        ).Scan(&users[i].ID, &users[i].CreatedAt, &users[i].UpdatedAt)
        
        if err != nil {
            return fmt.Errorf("inserting user %d: %w", i, err)
        }
    }
    
    if err := tx.Commit(); err != nil {
        return fmt.Errorf("committing transaction: %w", err)
    }
    
    return nil
}
```

### Read (Query)

```go
// Get single user
func GetUserByID(ctx context.Context, db *sql.DB, id int64) (*User, error) {
    query := `
        SELECT id, email, name, active, created_at, updated_at
        FROM users
        WHERE id = $1
    `
    
    var user User
    err := db.QueryRowContext(ctx, query, id).Scan(
        &user.ID,
        &user.Email,
        &user.Name,
        &user.Active,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("user not found")
    }
    if err != nil {
        return nil, fmt.Errorf("querying user: %w", err)
    }
    
    return &user, nil
}

// Get multiple users
func GetAllUsers(ctx context.Context, db *sql.DB) ([]User, error) {
    query := `
        SELECT id, email, name, active, created_at, updated_at
        FROM users
        ORDER BY created_at DESC
    `
    
    rows, err := db.QueryContext(ctx, query)
    if err != nil {
        return nil, fmt.Errorf("querying users: %w", err)
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(
            &user.ID,
            &user.Email,
            &user.Name,
            &user.Active,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("scanning user: %w", err)
        }
        users = append(users, user)
    }
    
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("iterating rows: %w", err)
    }
    
    return users, nil
}

// Pagination
func GetUsersPaginated(ctx context.Context, db *sql.DB, limit, offset int) ([]User, error) {
    query := `
        SELECT id, email, name, active, created_at, updated_at
        FROM users
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `
    
    rows, err := db.QueryContext(ctx, query, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("querying users: %w", err)
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(
            &user.ID, &user.Email, &user.Name,
            &user.Active, &user.CreatedAt, &user.UpdatedAt,
        ); err != nil {
            return nil, fmt.Errorf("scanning user: %w", err)
        }
        users = append(users, user)
    }
    
    return users, rows.Err()
}
```

### Update

```go
func UpdateUser(ctx context.Context, db *sql.DB, user *User) error {
    query := `
        UPDATE users
        SET email = $1, name = $2, active = $3, updated_at = CURRENT_TIMESTAMP
        WHERE id = $4
        RETURNING updated_at
    `
    
    err := db.QueryRowContext(
        ctx,
        query,
        user.Email,
        user.Name,
        user.Active,
        user.ID,
    ).Scan(&user.UpdatedAt)
    
    if err == sql.ErrNoRows {
        return fmt.Errorf("user not found")
    }
    if err != nil {
        return fmt.Errorf("updating user: %w", err)
    }
    
    return nil
}

// Partial update
func UpdateUserEmail(ctx context.Context, db *sql.DB, id int64, email string) error {
    query := `
        UPDATE users
        SET email = $1, updated_at = CURRENT_TIMESTAMP
        WHERE id = $2
    `
    
    result, err := db.ExecContext(ctx, query, email, id)
    if err != nil {
        return fmt.Errorf("updating email: %w", err)
    }
    
    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("getting rows affected: %w", err)
    }
    
    if rows == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}
```

### Delete

```go
func DeleteUser(ctx context.Context, db *sql.DB, id int64) error {
    query := `DELETE FROM users WHERE id = $1`
    
    result, err := db.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("deleting user: %w", err)
    }
    
    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("getting rows affected: %w", err)
    }
    
    if rows == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}

// Soft delete
func SoftDeleteUser(ctx context.Context, db *sql.DB, id int64) error {
    query := `
        UPDATE users
        SET active = false, updated_at = CURRENT_TIMESTAMP
        WHERE id = $1
    `
    
    result, err := db.ExecContext(ctx, query, id)
    if err != nil {
        return fmt.Errorf("soft deleting user: %w", err)
    }
    
    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("getting rows affected: %w", err)
    }
    
    if rows == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}
```

---


