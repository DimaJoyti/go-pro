# 📘 Lesson 18: Deployment and DevOps

Welcome to Lesson 18! This lesson covers deploying Go applications, containerization, CI/CD pipelines, and DevOps best practices.

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
- Build and deploy Go applications
- Create Docker containers for Go apps
- Set up CI/CD pipelines
- Deploy to cloud platforms
- Implement monitoring and logging
- Manage configuration and secrets
- Apply DevOps best practices

## 📚 Theory

### Building Go Applications

**Cross-Platform Builds:**
```bash
# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o app-linux main.go
GOOS=windows GOARCH=amd64 go build -o app-windows.exe main.go
GOOS=darwin GOARCH=amd64 go build -o app-macos main.go

# Optimized production build
go build -ldflags="-s -w" -o app main.go
```

### Docker Containerization

**Multi-stage Dockerfile:**
```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
```

### Configuration Management

**Environment-based Config:**
```go
type Config struct {
    Port        string `env:"PORT" envDefault:"8080"`
    DatabaseURL string `env:"DATABASE_URL,required"`
    JWTSecret   string `env:"JWT_SECRET,required"`
    LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
}

func LoadConfig() (*Config, error) {
    cfg := &Config{}
    if err := env.Parse(cfg); err != nil {
        return nil, err
    }
    return cfg, nil
}
```

### Health Checks

**Health Check Endpoint:**
```go
type HealthChecker struct {
    db    *sql.DB
    redis *redis.Client
}

func (h *HealthChecker) Check(ctx context.Context) map[string]string {
    status := make(map[string]string)
    
    // Database health
    if err := h.db.PingContext(ctx); err != nil {
        status["database"] = "unhealthy: " + err.Error()
    } else {
        status["database"] = "healthy"
    }
    
    // Redis health
    if err := h.redis.Ping(ctx).Err(); err != nil {
        status["redis"] = "unhealthy: " + err.Error()
    } else {
        status["redis"] = "healthy"
    }
    
    return status
}
```

## 💻 Hands-On Examples

### Example 1: CI/CD Pipeline (GitHub Actions)
```yaml
name: CI/CD Pipeline

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.21
    
    - name: Run tests
      run: |
        go test -v ./...
        go test -race ./...
    
    - name: Run security scan
      run: |
        go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
        gosec ./...

  build-and-deploy:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Build Docker image
      run: |
        docker build -t myapp:${{ github.sha }} .
        docker tag myapp:${{ github.sha }} myapp:latest
    
    - name: Deploy to production
      run: |
        # Deploy commands here
        echo "Deploying to production..."
```

### Example 2: Kubernetes Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: go-app
  template:
    metadata:
      labels:
        app: go-app
    spec:
      containers:
      - name: go-app
        image: myapp:latest
        ports:
        - containerPort: 8080
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: app-secrets
              key: database-url
        livenessProbe:
          httpGet:
            path: /health
            port: 8080
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
---
apiVersion: v1
kind: Service
metadata:
  name: go-app-service
spec:
  selector:
    app: go-app
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
  type: LoadBalancer
```

## 🧪 Exercises

Complete the exercises in `../../code/lesson-18/exercises/`:

1. **Docker Containerization**: Create optimized Docker images
2. **CI/CD Pipeline**: Set up automated testing and deployment
3. **Cloud Deployment**: Deploy to AWS/GCP/Azure
4. **Monitoring Setup**: Implement application monitoring
5. **Configuration Management**: Handle secrets and config
6. **Production Readiness**: Prepare app for production

## ✅ Validation

Test your deployment setup:

```bash
cd ../../code/lesson-18
docker build -t lesson18-app .
docker run -p 8080:8080 lesson18-app
```

## 🔍 Key Takeaways

- Use multi-stage Docker builds for smaller images
- Implement comprehensive health checks
- Automate testing and deployment with CI/CD
- Manage configuration through environment variables
- Monitor applications in production
- Use container orchestration for scalability
- Implement proper logging and observability

## 🚀 Deployment Strategies

- **Blue-Green Deployment**: Zero-downtime deployments
- **Rolling Updates**: Gradual replacement of instances
- **Canary Releases**: Test with small traffic percentage
- **Feature Flags**: Control feature rollouts
- **Circuit Breakers**: Handle service failures gracefully

## 📊 Monitoring Stack

- **Metrics**: Prometheus + Grafana
- **Logging**: ELK Stack or Fluentd
- **Tracing**: Jaeger or Zipkin
- **Alerting**: PagerDuty or Slack integration
- **Health Checks**: Kubernetes probes

## ➡️ Next Steps

Once you've successfully deployed your application, move on to:
**[Lesson 19: Advanced Design Patterns](../lesson-19/README.md)**

---

**Deploy with confidence!** 🚀
