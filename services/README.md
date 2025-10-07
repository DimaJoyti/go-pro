# GO-PRO Microservices Architecture

This directory contains the microservices implementation of the GO-PRO learning platform.

## Architecture Overview

The application is decomposed into the following microservices:

```
┌─────────────────────────────────────────────────────────────────┐
│                         Client Layer                             │
│                    (Web, Mobile, CLI)                            │
└────────────────────────┬────────────────────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                      API Gateway                                 │
│  - Request Routing                                               │
│  - Authentication/Authorization                                  │
│  - Rate Limiting                                                 │
│  - Request/Response Transformation                               │
│  - Circuit Breaking                                              │
└────────────────────────┬────────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┬──────────────┐
         │               │               │              │
         ▼               ▼               ▼              ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│    User     │  │   Course    │  │  Progress   │  │ Notification│
│   Service   │  │   Service   │  │   Service   │  │   Service   │
│             │  │             │  │             │  │             │
│ - Auth      │  │ - Courses   │  │ - Tracking  │  │ - Email     │
│ - Users     │  │ - Lessons   │  │ - Analytics │  │ - Push      │
│ - Profiles  │  │ - Exercises │  │ - Reports   │  │ - SMS       │
└──────┬──────┘  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘
       │                │                │                │
       └────────────────┴────────────────┴────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Event Bus (Kafka)                           │
│  - User Events                                                   │
│  - Course Events                                                 │
│  - Progress Events                                               │
│  - Notification Events                                           │
└─────────────────────────────────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┬──────────────┐
         │               │               │              │
         ▼               ▼               ▼              ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│  PostgreSQL │  │    Redis    │  │     S3      │  │  Analytics  │
│  (Primary)  │  │   (Cache)   │  │  (Storage)  │  │   (Kafka)   │
└─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘
```

## Services

### 1. API Gateway
**Port**: 8080  
**Responsibilities**:
- Single entry point for all client requests
- Request routing to appropriate microservices
- Authentication and authorization
- Rate limiting and throttling
- Request/response transformation
- Circuit breaking and fallback
- API versioning
- Metrics and logging aggregation

**Technology Stack**:
- Go with net/http
- JWT authentication
- Redis for rate limiting
- OpenTelemetry for tracing

### 2. User Service
**Port**: 8081  
**Responsibilities**:
- User registration and authentication
- User profile management
- Password management
- Role and permission management
- Session management
- OAuth integration

**Database**: PostgreSQL (users schema)  
**Cache**: Redis (sessions, user data)  
**Events Published**:
- UserCreated
- UserUpdated
- UserDeleted
- UserLoggedIn
- UserLoggedOut

### 3. Course Service
**Port**: 8082  
**Responsibilities**:
- Course CRUD operations
- Lesson management
- Exercise management
- Course categorization
- Content versioning
- Search and filtering

**Database**: PostgreSQL (courses schema)  
**Cache**: Redis (course catalog, popular courses)  
**Events Published**:
- CourseCreated
- CourseUpdated
- CourseDeleted
- LessonCreated
- ExerciseCreated

### 4. Progress Service
**Port**: 8083  
**Responsibilities**:
- Track user progress
- Calculate completion percentages
- Generate learning analytics
- Provide recommendations
- Track achievements and badges
- Generate reports

**Database**: PostgreSQL (progress schema)  
**Cache**: Redis (progress data, leaderboards)  
**Events Published**:
- ProgressUpdated
- LessonCompleted
- CourseCompleted
- AchievementUnlocked

### 5. Notification Service
**Port**: 8084  
**Responsibilities**:
- Send email notifications
- Send push notifications
- Send SMS notifications
- Manage notification preferences
- Template management
- Delivery tracking

**Database**: PostgreSQL (notifications schema)  
**Queue**: Kafka (notification queue)  
**Events Consumed**:
- UserCreated
- CourseCompleted
- AchievementUnlocked

## Communication Patterns

### Synchronous Communication (HTTP/gRPC)
- API Gateway → Services: HTTP REST
- Service → Service: gRPC (for low-latency requirements)

### Asynchronous Communication (Event-Driven)
- Services → Kafka: Event publishing
- Services ← Kafka: Event consumption
- Use for:
  - User activity tracking
  - Notifications
  - Analytics
  - Audit logging

## Data Management

### Database Per Service Pattern
Each service has its own database schema:
- **users**: User service data
- **courses**: Course service data
- **progress**: Progress service data
- **notifications**: Notification service data

### Shared Data Challenges
- **Eventual Consistency**: Services sync via events
- **Distributed Transactions**: Saga pattern for multi-service operations
- **Data Duplication**: Acceptable for read performance

## Service Discovery

### Development
- Static configuration in environment variables
- Docker Compose service names

### Production
- Kubernetes Service Discovery
- Consul/Etcd for service registry
- DNS-based discovery

## Resilience Patterns

### Circuit Breaker
- Prevent cascading failures
- Fail fast when service is down
- Automatic recovery detection

### Retry with Exponential Backoff
- Retry failed requests
- Exponential backoff to prevent overwhelming
- Maximum retry attempts

### Timeout
- Request timeouts for all service calls
- Prevent indefinite waiting
- Configurable per service

### Bulkhead
- Isolate resources per service
- Prevent resource exhaustion
- Thread pool isolation

## Security

### Authentication
- JWT tokens issued by User Service
- Token validation in API Gateway
- Service-to-service authentication with mTLS

### Authorization
- Role-Based Access Control (RBAC)
- Permission checks in each service
- Centralized policy management

### Data Encryption
- TLS for all service communication
- Encryption at rest for sensitive data
- Secrets management with Vault/Secrets Manager

## Monitoring & Observability

### Distributed Tracing
- OpenTelemetry for trace collection
- Jaeger for trace visualization
- Trace context propagation across services

### Metrics
- Prometheus for metrics collection
- Grafana for visualization
- Service-level metrics (latency, throughput, errors)

### Logging
- Structured logging (JSON format)
- Centralized log aggregation (ELK/Loki)
- Correlation IDs for request tracking

### Health Checks
- Liveness probes (is service running?)
- Readiness probes (is service ready to accept traffic?)
- Dependency health checks

## Deployment

### Docker
Each service has its own Dockerfile:
```
services/
├── api-gateway/Dockerfile
├── user-service/Dockerfile
├── course-service/Dockerfile
├── progress-service/Dockerfile
└── notification-service/Dockerfile
```

### Kubernetes
Each service has its own K8s manifests:
- Deployment
- Service
- ConfigMap
- Secret
- HPA (Horizontal Pod Autoscaler)
- PDB (Pod Disruption Budget)

### CI/CD
- Independent deployment pipelines
- Automated testing per service
- Canary deployments
- Blue-green deployments

## Development Workflow

### Local Development
```bash
# Start all services with Docker Compose
docker-compose up

# Start individual service
cd services/user-service
go run cmd/main.go

# Run tests
make test

# Build service
make build
```

### Testing
- Unit tests per service
- Integration tests with test containers
- Contract tests between services
- End-to-end tests for critical flows

## Migration Strategy

### Phase 1: Strangler Pattern
1. Keep existing monolith running
2. Route new features to microservices
3. Gradually migrate existing features

### Phase 2: Service Extraction
1. Extract User Service
2. Extract Course Service
3. Extract Progress Service
4. Extract Notification Service

### Phase 3: Decommission Monolith
1. Migrate all traffic to microservices
2. Remove monolith dependencies
3. Decommission monolith

## Best Practices

1. **Single Responsibility**: Each service owns one business capability
2. **Loose Coupling**: Services communicate via well-defined APIs
3. **High Cohesion**: Related functionality grouped together
4. **Autonomous**: Services can be deployed independently
5. **Resilient**: Handle failures gracefully
6. **Observable**: Comprehensive monitoring and logging
7. **Scalable**: Scale services independently based on load

## Directory Structure

```
services/
├── README.md                    # This file
├── docker-compose.yml           # Local development setup
├── shared/                      # Shared libraries
│   ├── events/                  # Event definitions
│   ├── proto/                   # gRPC definitions
│   └── utils/                   # Common utilities
├── api-gateway/                 # API Gateway service
│   ├── cmd/                     # Entry point
│   ├── internal/                # Internal packages
│   ├── Dockerfile               # Container image
│   └── go.mod                   # Dependencies
├── user-service/                # User service
├── course-service/              # Course service
├── progress-service/            # Progress service
└── notification-service/        # Notification service
```

## Next Steps

1. Implement shared libraries (events, proto, utils)
2. Implement API Gateway
3. Implement User Service
4. Implement Course Service
5. Implement Progress Service
6. Implement Notification Service
7. Set up service mesh (Istio)
8. Implement distributed tracing
9. Set up monitoring and alerting
10. Performance testing and optimization

