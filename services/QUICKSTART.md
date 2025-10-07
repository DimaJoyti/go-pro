# GO-PRO Microservices Quick Start Guide

Get the GO-PRO microservices up and running in 5 minutes!

## Prerequisites

- Docker 24.0+
- Docker Compose 2.20+
- Go 1.22+ (for local development)
- Make (optional, for convenience commands)

## Quick Start (Docker Compose)

### 1. Clone and Navigate
```bash
git clone https://github.com/DimaJoyti/go-pro.git
cd go-pro/services
```

### 2. Start All Services
```bash
# Using Make (recommended)
make up

# Or using Docker Compose directly
docker-compose up -d
```

This will start:
- PostgreSQL (port 5432)
- Redis (port 6379)
- Kafka + Zookeeper (ports 9092, 9093, 2181)
- API Gateway (port 8080)
- User Service (port 8081) - *when implemented*
- Course Service (port 8082) - *when implemented*
- Progress Service (port 8083) - *when implemented*
- Jaeger (port 16686)
- Prometheus (port 9090)
- Grafana (port 3000)

### 3. Check Service Health
```bash
# Using Make
make health

# Or manually
curl http://localhost:8080/health
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
```

### 4. View Logs
```bash
# All services
make logs

# Specific service
make logs-api-gateway
make logs-user
make logs-course
make logs-progress

# Or using Docker Compose
docker-compose logs -f api-gateway
```

### 5. Access Monitoring Tools

**Jaeger (Distributed Tracing)**
```bash
# Open in browser
make jaeger
# Or visit: http://localhost:16686
```

**Prometheus (Metrics)**
```bash
# Open in browser
make prometheus
# Or visit: http://localhost:9090
```

**Grafana (Dashboards)**
```bash
# Open in browser
make grafana
# Or visit: http://localhost:3000
# Default credentials: admin/admin
```

### 6. Test the API

**Health Check**
```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "healthy",
  "service": "api-gateway",
  "timestamp": "2024-01-07T00:00:00Z"
}
```

**Readiness Check**
```bash
curl http://localhost:8080/ready
```

Response:
```json
{
  "status": "ready",
  "services": {
    "user-service": "healthy",
    "course-service": "healthy",
    "progress-service": "healthy"
  },
  "timestamp": "2024-01-07T00:00:00Z"
}
```

## Local Development (Without Docker)

### 1. Start Infrastructure
```bash
# Start only PostgreSQL, Redis, Kafka
docker-compose up -d postgres redis kafka zookeeper

# Wait for services to be ready
sleep 10
```

### 2. Run Database Migrations
```bash
docker-compose exec postgres psql -U gopro -d gopro -f /docker-entrypoint-initdb.d/init-db.sql
```

### 3. Set Environment Variables
```bash
export PORT=8080
export USER_SERVICE_URL=http://localhost:8081
export COURSE_SERVICE_URL=http://localhost:8082
export PROGRESS_SERVICE_URL=http://localhost:8083
export JWT_SECRET=your-secret-key
export REDIS_URL=localhost:6379
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=gopro
export DB_PASSWORD=gopro_password
export DB_NAME=gopro
export KAFKA_BROKERS=localhost:9093
```

### 4. Run Services

**Terminal 1: API Gateway**
```bash
cd api-gateway
go mod download
go run cmd/main.go
```

**Terminal 2: User Service** *(when implemented)*
```bash
cd user-service
go mod download
go run cmd/main.go
```

**Terminal 3: Course Service** *(when implemented)*
```bash
cd course-service
go mod download
go run cmd/main.go
```

**Terminal 4: Progress Service** *(when implemented)*
```bash
cd progress-service
go mod download
go run cmd/main.go
```

## Common Commands

### Service Management
```bash
make up          # Start all services
make down        # Stop all services
make restart     # Restart all services
make ps          # Show running containers
make stats       # Show container stats
```

### Logs
```bash
make logs                # All services
make logs-api-gateway    # API Gateway only
make logs-user           # User Service only
make logs-course         # Course Service only
make logs-progress       # Progress Service only
```

### Database
```bash
make db-migrate   # Run migrations
make db-shell     # Open PostgreSQL shell
make db-reset     # Reset database (WARNING: deletes all data)
```

### Testing
```bash
make test                # Run all tests
make test-api-gateway    # Test API Gateway
make test-user           # Test User Service
make test-course         # Test Course Service
make test-progress       # Test Progress Service
make test-integration    # Integration tests
```

### Code Quality
```bash
make lint        # Run linters
make fmt         # Format code
make vet         # Run go vet
make deps        # Download dependencies
make tidy        # Tidy dependencies
```

### Monitoring
```bash
make health      # Check service health
make jaeger      # Open Jaeger UI
make prometheus  # Open Prometheus UI
make grafana     # Open Grafana UI
```

### Cleanup
```bash
make clean       # Clean up containers and volumes
make clean-all   # Clean up everything including images
```

## Troubleshooting

### Services Not Starting

**Check Docker**
```bash
docker --version
docker-compose --version
docker ps
```

**Check Logs**
```bash
docker-compose logs postgres
docker-compose logs redis
docker-compose logs kafka
```

**Restart Services**
```bash
make down
make up
```

### Database Connection Issues

**Test Database Connection**
```bash
docker-compose exec postgres psql -U gopro -d gopro -c "SELECT 1"
```

**Check Database Logs**
```bash
docker-compose logs postgres
```

**Reset Database**
```bash
make db-reset
```

### Port Already in Use

**Find Process Using Port**
```bash
# Linux/Mac
lsof -i :8080
lsof -i :5432

# Windows
netstat -ano | findstr :8080
```

**Kill Process**
```bash
# Linux/Mac
kill -9 <PID>

# Windows
taskkill /PID <PID> /F
```

### API Gateway Returns 502

**Check Backend Services**
```bash
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
```

**Check Service URLs**
```bash
docker-compose exec api-gateway env | grep SERVICE_URL
```

### High Memory Usage

**Check Container Stats**
```bash
docker stats
```

**Limit Container Resources**
Edit `docker-compose.yml`:
```yaml
services:
  api-gateway:
    deploy:
      resources:
        limits:
          memory: 512M
          cpus: '0.5'
```

## Next Steps

### 1. Explore the API
- Read the API documentation
- Test endpoints with Postman
- Review example requests

### 2. Monitor Services
- Check Jaeger for traces
- View Prometheus metrics
- Create Grafana dashboards

### 3. Develop Features
- Implement User Service
- Implement Course Service
- Implement Progress Service
- Add new endpoints

### 4. Run Tests
- Write unit tests
- Write integration tests
- Run load tests
- Check test coverage

### 5. Deploy to Production
- Review deployment guide
- Configure production environment
- Set up CI/CD pipeline
- Deploy to Kubernetes

## Useful Links

- **Architecture**: [README.md](README.md)
- **Deployment**: [DEPLOYMENT.md](DEPLOYMENT.md)
- **Implementation Status**: [IMPLEMENTATION_STATUS.md](IMPLEMENTATION_STATUS.md)
- **API Gateway**: [api-gateway/README.md](api-gateway/README.md)

## Getting Help

### Documentation
- Check the README files in each service directory
- Review the architecture documentation
- Read the deployment guide

### Logs
- Check service logs for errors
- Review database logs
- Check Kafka logs

### Community
- Open an issue on GitHub
- Ask in discussions
- Check existing issues

## Summary

You should now have:
- ✅ All services running
- ✅ Database initialized
- ✅ Monitoring tools accessible
- ✅ API Gateway responding to requests

**Next**: Start implementing the backend services or explore the existing API Gateway implementation!

Happy coding! 🚀

