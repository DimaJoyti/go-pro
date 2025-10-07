# 📘 Lesson 19: Advanced Design Patterns

Welcome to Lesson 19! This advanced lesson covers sophisticated design patterns, architectural patterns, and advanced Go programming techniques.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Implement advanced design patterns in Go
- Apply functional programming concepts
- Use generics effectively (Go 1.18+)
- Build plugin architectures
- Implement event-driven systems
- Apply domain-driven design principles
- Create extensible and maintainable systems

## 📚 Theory

### Functional Programming Patterns

**Higher-Order Functions:**
```go
type Predicate[T any] func(T) bool
type Mapper[T, U any] func(T) U

func Filter[T any](slice []T, predicate Predicate[T]) []T {
    var result []T
    for _, item := range slice {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}

func Map[T, U any](slice []T, mapper Mapper[T, U]) []U {
    result := make([]U, len(slice))
    for i, item := range slice {
        result[i] = mapper(item)
    }
    return result
}
```

### Builder Pattern with Generics

**Generic Builder:**
```go
type Builder[T any] struct {
    build func() T
    steps []func(*T)
}

func NewBuilder[T any](constructor func() T) *Builder[T] {
    return &Builder[T]{build: constructor}
}

func (b *Builder[T]) With(step func(*T)) *Builder[T] {
    b.steps = append(b.steps, step)
    return b
}

func (b *Builder[T]) Build() T {
    obj := b.build()
    for _, step := range b.steps {
        step(&obj)
    }
    return obj
}
```

### Plugin Architecture

**Plugin Interface:**
```go
type Plugin interface {
    Name() string
    Version() string
    Execute(ctx context.Context, input interface{}) (interface{}, error)
}

type PluginManager struct {
    plugins map[string]Plugin
    mu      sync.RWMutex
}

func (pm *PluginManager) Register(plugin Plugin) {
    pm.mu.Lock()
    defer pm.mu.Unlock()
    pm.plugins[plugin.Name()] = plugin
}

func (pm *PluginManager) Execute(name string, ctx context.Context, input interface{}) (interface{}, error) {
    pm.mu.RLock()
    plugin, exists := pm.plugins[name]
    pm.mu.RUnlock()
    
    if !exists {
        return nil, fmt.Errorf("plugin %s not found", name)
    }
    
    return plugin.Execute(ctx, input)
}
```

### Event-Driven Architecture

**Event Bus:**
```go
type Event interface {
    Type() string
    Timestamp() time.Time
}

type EventHandler func(Event) error

type EventBus struct {
    handlers map[string][]EventHandler
    mu       sync.RWMutex
}

func (eb *EventBus) Subscribe(eventType string, handler EventHandler) {
    eb.mu.Lock()
    defer eb.mu.Unlock()
    eb.handlers[eventType] = append(eb.handlers[eventType], handler)
}

func (eb *EventBus) Publish(event Event) error {
    eb.mu.RLock()
    handlers := eb.handlers[event.Type()]
    eb.mu.RUnlock()
    
    for _, handler := range handlers {
        if err := handler(event); err != nil {
            return err
        }
    }
    return nil
}
```

## 💻 Hands-On Examples

### Example 1: Command Pattern
```go
type Command interface {
    Execute() error
    Undo() error
}

type CommandInvoker struct {
    history []Command
    current int
}

func (ci *CommandInvoker) Execute(cmd Command) error {
    if err := cmd.Execute(); err != nil {
        return err
    }
    
    // Truncate history if we're not at the end
    ci.history = ci.history[:ci.current]
    ci.history = append(ci.history, cmd)
    ci.current++
    
    return nil
}

func (ci *CommandInvoker) Undo() error {
    if ci.current == 0 {
        return errors.New("nothing to undo")
    }
    
    ci.current--
    return ci.history[ci.current].Undo()
}
```

### Example 2: Strategy Pattern with Generics
```go
type Strategy[T any] interface {
    Execute(data T) (T, error)
}

type Context[T any] struct {
    strategy Strategy[T]
}

func (c *Context[T]) SetStrategy(strategy Strategy[T]) {
    c.strategy = strategy
}

func (c *Context[T]) ExecuteStrategy(data T) (T, error) {
    return c.strategy.Execute(data)
}

// Concrete strategies
type SortStrategy struct{}

func (s SortStrategy) Execute(data []int) ([]int, error) {
    sorted := make([]int, len(data))
    copy(sorted, data)
    sort.Ints(sorted)
    return sorted, nil
}
```

### Example 3: Observer Pattern
```go
type Observer[T any] interface {
    Update(data T)
}

type Subject[T any] struct {
    observers []Observer[T]
    data      T
}

func (s *Subject[T]) Attach(observer Observer[T]) {
    s.observers = append(s.observers, observer)
}

func (s *Subject[T]) Detach(observer Observer[T]) {
    for i, obs := range s.observers {
        if obs == observer {
            s.observers = append(s.observers[:i], s.observers[i+1:]...)
            break
        }
    }
}

func (s *Subject[T]) Notify() {
    for _, observer := range s.observers {
        observer.Update(s.data)
    }
}

func (s *Subject[T]) SetData(data T) {
    s.data = data
    s.Notify()
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-19/exercises/`:

1. **Functional Programming**: Implement functional utilities with generics
2. **Builder Pattern**: Create flexible builders for complex objects
3. **Plugin System**: Build a extensible plugin architecture
4. **Event System**: Implement event-driven communication
5. **Command Pattern**: Create undoable operations
6. **Domain-Driven Design**: Apply DDD principles to a business domain

## ✅ Validation

Run the tests to validate your pattern implementations:

```bash
cd ../../code/lesson-19
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Design patterns solve common architectural problems
- Generics enable type-safe, reusable patterns
- Functional programming concepts enhance code expressiveness
- Plugin architectures provide extensibility
- Event-driven systems enable loose coupling
- Command pattern enables undo/redo functionality
- Observer pattern facilitates reactive programming

## 🏗️ Architectural Patterns

- **Hexagonal Architecture**: Ports and adapters
- **Clean Architecture**: Dependency inversion
- **CQRS**: Command Query Responsibility Segregation
- **Event Sourcing**: Store events, not state
- **Saga Pattern**: Manage distributed transactions
- **Strangler Fig**: Gradually replace legacy systems

## 🎯 When to Use Patterns

- **Creational**: Object creation complexity
- **Structural**: Object composition and relationships
- **Behavioral**: Communication between objects
- **Concurrency**: Thread-safe operations
- **Architectural**: System-level organization

## ➡️ Next Steps

Once you've mastered these advanced patterns, move on to:
**[Lesson 20: Building Production Systems](../lesson-20/README.md)**

---

**Design with patterns!** 🎨
