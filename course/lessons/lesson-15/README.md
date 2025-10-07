# 📘 Lesson 15: Microservices Architecture

Welcome to Lesson 15! This advanced lesson covers microservices architecture patterns, service communication, and building distributed systems with Go.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Design microservices architectures
- Implement service-to-service communication
- Handle distributed system challenges
- Apply microservices patterns and best practices
- Build resilient and scalable services
- Implement service discovery and load balancing
- Monitor and observe distributed systems

## 📚 Theory

### Microservices Principles

Key characteristics of microservices:

- **Single Responsibility**: Each service has one business capability
- **Decentralized**: Independent deployment and scaling
- **Technology Agnostic**: Services can use different technologies
- **Fault Tolerant**: Failures in one service don't cascade
- **Observable**: Comprehensive logging and monitoring

### Service Communication

**Synchronous Communication:**
- HTTP/REST APIs
- gRPC for high-performance communication
- GraphQL for flexible data fetching

**Asynchronous Communication:**
- Message queues (RabbitMQ, Apache Kafka)
- Event-driven architecture
- Pub/Sub patterns

### Common Patterns

**API Gateway Pattern:**
```go
type APIGateway struct {
    userService    UserService
    orderService   OrderService
    paymentService PaymentService
}

func (gw *APIGateway) HandleRequest(w http.ResponseWriter, r *http.Request) {
    switch r.URL.Path {
    case "/users":
        gw.userService.Handle(w, r)
    case "/orders":
        gw.orderService.Handle(w, r)
    case "/payments":
        gw.paymentService.Handle(w, r)
    }
}
```

**Circuit Breaker Pattern:**
```go
type CircuitBreaker struct {
    maxFailures int
    failures    int
    state       State
    timeout     time.Duration
}

func (cb *CircuitBreaker) Call(fn func() error) error {
    if cb.state == Open {
        return errors.New("circuit breaker is open")
    }
    
    err := fn()
    if err != nil {
        cb.failures++
        if cb.failures >= cb.maxFailures {
            cb.state = Open
        }
        return err
    }
    
    cb.failures = 0
    return nil
}
```

## 💻 Hands-On Examples

### Example 1: Service Structure
```go
type UserService struct {
    db     Database
    logger Logger
    config Config
}

func NewUserService(db Database, logger Logger, config Config) *UserService {
    return &UserService{
        db:     db,
        logger: logger,
        config: config,
    }
}

func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := s.db.GetUser(ctx, id)
    if err != nil {
        s.logger.Error("failed to get user", "id", id, "error", err)
        return nil, err
    }
    return user, nil
}
```

### Example 2: Service Discovery
```go
type ServiceRegistry interface {
    Register(service Service) error
    Discover(serviceName string) ([]Service, error)
    Deregister(serviceID string) error
}

type Service struct {
    ID      string
    Name    string
    Address string
    Port    int
    Health  string
}
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-15/exercises/`:

1. **Service Design**: Design microservices for a business domain
2. **API Gateway**: Implement an API gateway pattern
3. **Service Communication**: Build inter-service communication
4. **Circuit Breaker**: Implement fault tolerance patterns
5. **Service Discovery**: Create service registry and discovery
6. **Distributed Tracing**: Add observability to services

## ✅ Validation

Run the tests to validate your understanding:

```bash
cd ../../code/lesson-15
go test -v ./exercises/...
```

## 🔍 Key Takeaways

- Microservices enable independent scaling and deployment
- Service communication requires careful design
- Fault tolerance is critical in distributed systems
- Observability is essential for debugging and monitoring
- API gateways provide centralized entry points
- Circuit breakers prevent cascade failures
- Service discovery enables dynamic service location

## 🏗️ Architecture Patterns

- **Strangler Fig**: Gradually replace monolith
- **Database per Service**: Each service owns its data
- **Saga Pattern**: Manage distributed transactions
- **CQRS**: Separate read and write models
- **Event Sourcing**: Store events instead of state

## 📊 Monitoring and Observability

- **Metrics**: Performance and business metrics
- **Logging**: Structured logging with correlation IDs
- **Tracing**: Distributed request tracing
- **Health Checks**: Service health monitoring
- **Alerting**: Proactive issue detection

## ➡️ Next Steps

Congratulations! You've completed the core Go curriculum. Consider exploring:
- **Advanced Projects**: Build complex distributed systems
- **Performance Optimization**: Profile and optimize Go applications
- **Security**: Implement security best practices
- **DevOps**: Deploy and operate Go services in production

---

**Build distributed systems!** 🚀
