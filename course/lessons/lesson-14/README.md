# 📘 Lesson 14: Database Integration

Welcome to database programming with Go! This lesson covers connecting to databases, executing queries, handling transactions, and implementing data access patterns.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Connect to SQL databases using database/sql package
- Execute queries, inserts, updates, and deletes safely
- Handle database transactions and connection pooling
- Implement repository and data access patterns
- Work with popular database drivers (PostgreSQL, MySQL, SQLite)
- Handle database migrations and schema management
- Implement connection pooling and performance optimization
- Use ORMs and query builders effectively

## 📚 Theory

### Database Connection

Go's `database/sql` package provides a generic interface for SQL databases:

```go
import (
    "database/sql"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func connectDB() (*sql.DB, error) {
    dsn := "postgres://user:password@localhost/dbname?sslmode=disable"
    
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    
    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    // Test connection
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    return db, nil
}
```

### Query Execution

Execute different types of database operations:

```go
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

// Query single row
func getUserByID(db *sql.DB, id int) (*User, error) {
    query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
    
    var user User
    err := db.QueryRow(query, id).Scan(
        &user.ID,
        &user.Name,
        &user.Email,
        &user.CreatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user not found")
        }
        return nil, err
    }
    
    return &user, nil
}

// Query multiple rows
func getAllUsers(db *sql.DB) ([]User, error) {
    query := `SELECT id, name, email, created_at FROM users ORDER BY created_at DESC`
    
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    
    if err := rows.Err(); err != nil {
        return nil, err
    }
    
    return users, nil
}

// Insert data
func createUser(db *sql.DB, user *User) error {
    query := `
        INSERT INTO users (name, email, created_at) 
        VALUES ($1, $2, $3) 
        RETURNING id`
    
    err := db.QueryRow(query, user.Name, user.Email, time.Now()).Scan(&user.ID)
    return err
}

// Update data
func updateUser(db *sql.DB, user *User) error {
    query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
    
    result, err := db.Exec(query, user.Name, user.Email, user.ID)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}

// Delete data
func deleteUser(db *sql.DB, id int) error {
    query := `DELETE FROM users WHERE id = $1`
    
    result, err := db.Exec(query, id)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("user not found")
    }
    
    return nil
}
```

### Transaction Management

Handle database transactions for data consistency:

```go
func transferMoney(db *sql.DB, fromID, toID int, amount decimal.Decimal) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback() // Rollback if not committed
    
    // Check sender balance
    var balance decimal.Decimal
    err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1", fromID).Scan(&balance)
    if err != nil {
        return err
    }
    
    if balance.LessThan(amount) {
        return fmt.Errorf("insufficient funds")
    }
    
    // Debit sender
    _, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, fromID)
    if err != nil {
        return err
    }
    
    // Credit receiver
    _, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, toID)
    if err != nil {
        return err
    }
    
    // Record transaction
    _, err = tx.Exec(`
        INSERT INTO transactions (from_account, to_account, amount, created_at) 
        VALUES ($1, $2, $3, $4)`,
        fromID, toID, amount, time.Now())
    if err != nil {
        return err
    }
    
    return tx.Commit()
}
```

### Repository Pattern

Implement the repository pattern for data access:

```go
type UserRepository interface {
    GetByID(id int) (*User, error)
    GetByEmail(email string) (*User, error)
    GetAll() ([]User, error)
    Create(user *User) error
    Update(user *User) error
    Delete(id int) error
}

type PostgreSQLUserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
    return &PostgreSQLUserRepository{db: db}
}

func (r *PostgreSQLUserRepository) GetByID(id int) (*User, error) {
    query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
    
    var user User
    err := r.db.QueryRow(query, id).Scan(
        &user.ID, &user.Name, &user.Email, &user.CreatedAt,
    )
    
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, ErrUserNotFound
        }
        return nil, err
    }
    
    return &user, nil
}

func (r *PostgreSQLUserRepository) Create(user *User) error {
    query := `
        INSERT INTO users (name, email, created_at) 
        VALUES ($1, $2, $3) 
        RETURNING id, created_at`
    
    err := r.db.QueryRow(query, user.Name, user.Email, time.Now()).Scan(
        &user.ID, &user.CreatedAt,
    )
    
    return err
}
```

### Database Migrations

Manage database schema changes:

```go
type Migration struct {
    Version int
    Name    string
    Up      string
    Down    string
}

var migrations = []Migration{
    {
        Version: 1,
        Name:    "create_users_table",
        Up: `
            CREATE TABLE users (
                id SERIAL PRIMARY KEY,
                name VARCHAR(255) NOT NULL,
                email VARCHAR(255) UNIQUE NOT NULL,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            );
        `,
        Down: `DROP TABLE users;`,
    },
    {
        Version: 2,
        Name:    "add_users_updated_at",
        Up: `
            ALTER TABLE users 
            ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
        `,
        Down: `ALTER TABLE users DROP COLUMN updated_at;`,
    },
}

func runMigrations(db *sql.DB) error {
    // Create migrations table
    _, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS migrations (
            version INTEGER PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
    if err != nil {
        return err
    }
    
    // Get current version
    var currentVersion int
    err = db.QueryRow("SELECT COALESCE(MAX(version), 0) FROM migrations").Scan(&currentVersion)
    if err != nil {
        return err
    }
    
    // Apply pending migrations
    for _, migration := range migrations {
        if migration.Version <= currentVersion {
            continue
        }
        
        log.Printf("Applying migration %d: %s", migration.Version, migration.Name)
        
        tx, err := db.Begin()
        if err != nil {
            return err
        }
        
        _, err = tx.Exec(migration.Up)
        if err != nil {
            tx.Rollback()
            return fmt.Errorf("migration %d failed: %v", migration.Version, err)
        }
        
        _, err = tx.Exec("INSERT INTO migrations (version, name) VALUES ($1, $2)",
            migration.Version, migration.Name)
        if err != nil {
            tx.Rollback()
            return err
        }
        
        if err := tx.Commit(); err != nil {
            return err
        }
    }
    
    return nil
}
```

### Connection Pooling

Optimize database connections:

```go
func configureDB(db *sql.DB) {
    // Maximum number of open connections
    db.SetMaxOpenConns(25)
    
    // Maximum number of idle connections
    db.SetMaxIdleConns(5)
    
    // Maximum lifetime of a connection
    db.SetConnMaxLifetime(5 * time.Minute)
    
    // Maximum idle time for a connection
    db.SetConnMaxIdleTime(1 * time.Minute)
}

func healthCheck(db *sql.DB) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    return db.PingContext(ctx)
}
```

## 💻 Hands-On Examples

Navigate to `../../code/lesson-14/` to see and run these examples.

### Example 1: Complete User Service
```go
type UserService struct {
    repo UserRepository
}

func (s *UserService) CreateUser(name, email string) (*User, error) {
    // Validate input
    if name == "" || email == "" {
        return nil, fmt.Errorf("name and email are required")
    }
    
    // Check if user exists
    existing, _ := s.repo.GetByEmail(email)
    if existing != nil {
        return nil, fmt.Errorf("user with email %s already exists", email)
    }
    
    // Create user
    user := &User{
        Name:  name,
        Email: email,
    }
    
    if err := s.repo.Create(user); err != nil {
        return nil, err
    }
    
    return user, nil
}
```

### Example 2: Database Testing
```go
func setupTestDB(t *testing.T) *sql.DB {
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        t.Fatalf("Failed to open test database: %v", err)
    }
    
    // Create test schema
    schema := `
        CREATE TABLE users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
    `
    
    if _, err := db.Exec(schema); err != nil {
        t.Fatalf("Failed to create test schema: %v", err)
    }
    
    return db
}

func TestUserRepository_Create(t *testing.T) {
    db := setupTestDB(t)
    defer db.Close()
    
    repo := NewUserRepository(db)
    
    user := &User{
        Name:  "Test User",
        Email: "test@example.com",
    }
    
    err := repo.Create(user)
    if err != nil {
        t.Errorf("Create() error = %v", err)
    }
    
    if user.ID == 0 {
        t.Error("Expected user ID to be set")
    }
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-14/exercises/`:

1. **Database Connection**: Connect to different database types
2. **CRUD Operations**: Implement complete CRUD functionality
3. **Transaction Management**: Handle complex transactions
4. **Repository Pattern**: Build a repository layer
5. **Migration System**: Implement database migrations
6. **Performance Optimization**: Optimize queries and connections

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-14
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Always use parameterized queries to prevent SQL injection
- Handle database connections and transactions properly
- Implement repository pattern for clean data access
- Use connection pooling for performance
- Test database code with in-memory databases
- Handle database errors gracefully

## 📖 Additional Resources

- [database/sql Package](https://pkg.go.dev/database/sql)
- [SQL Drivers](https://github.com/golang/go/wiki/SQLDrivers)
- [Database Best Practices](https://go.dev/doc/database/index)
- [GORM ORM](https://gorm.io/)

## ➡️ Next Steps

Once you've completed all exercises and tests pass, move on to:
**[Lesson 15: Microservices Architecture](../lesson-15/README.md)**

---

**Master data persistence!** 💾

*Remember: Good database design and proper error handling are crucial for reliable applications!*
