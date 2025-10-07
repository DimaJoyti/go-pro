# API Gateway Service

The API Gateway is the single entry point for all client requests in the GO-PRO microservices architecture.

## Features

- **Request Routing**: Routes requests to appropriate backend services
- **Authentication**: JWT-based authentication and authorization
- **Rate Limiting**: Protects backend services from abuse
- **Circuit Breaking**: Prevents cascading failures
- **Request/Response Transformation**: Adds headers, modifies requests
- **Health Checks**: Monitors backend service health
- **Distributed Tracing**: Propagates trace context
- **CORS**: Handles cross-origin requests
- **Logging**: Structured logging with request IDs

## Architecture

```
Client → API Gateway → Backend Services
                    ├── User Service (8081)
                    ├── Course Service (8082)
                    └── Progress Service (8083)
```

## Routing Rules

| Path Pattern | Backend Service | Authentication |
|-------------|----------------|----------------|
| `/api/v1/auth/*` | User Service | Optional |
| `/api/v1/users/*` | User Service | Required |
| `/api/v1/courses` | Course Service | Optional |
| `/api/v1/courses/*` | Course Service | Required |
| `/api/v1/lessons/*` | Course Service | Required |
| `/api/v1/progress/*` | Progress Service | Required |
| `/api/v1/achievements/*` | Progress Service | Required |

## Configuration

Configuration is loaded from environment variables:

### Server Configuration
- `PORT` - Server port (default: 8080)
- `READ_TIMEOUT` - Read timeout (default: 15s)
- `WRITE_TIMEOUT` - Write timeout (default: 15s)
- `SHUTDOWN_TIMEOUT` - Graceful shutdown timeout (default: 30s)
- `MAX_HEADER_BYTES` - Max header size (default: 1MB)

### Backend Services
- `USER_SERVICE_URL` - User service URL (default: http://localhost:8081)
- `COURSE_SERVICE_URL` - Course service URL (default: http://localhost:8082)
- `PROGRESS_SERVICE_URL` - Progress service URL (default: http://localhost:8083)

### Authentication
- `JWT_SECRET` - JWT signing secret (required)
- `TOKEN_DURATION` - Token validity duration (default: 24h)

### Rate Limiting
- `RATE_LIMIT_RPM` - Requests per minute (default: 60)
- `RATE_LIMIT_BURST` - Burst size (default: 10)

### Redis
- `REDIS_URL` - Redis URL (default: localhost:6379)
- `REDIS_PASSWORD` - Redis password (optional)
- `REDIS_DB` - Redis database (default: 0)

### Logging
- `LOG_LEVEL` - Log level (default: info)
- `LOG_FORMAT` - Log format (default: json)

## Running Locally

### Prerequisites
- Go 1.22 or later
- Backend services running

### Development Mode
```bash
# Set environment variables
export USER_SERVICE_URL=http://localhost:8081
export COURSE_SERVICE_URL=http://localhost:8082
export PROGRESS_SERVICE_URL=http://localhost:8083
export JWT_SECRET=your-secret-key

# Run the service
go run cmd/main.go
```

### Using Docker
```bash
# Build image
docker build -t gopro-api-gateway .

# Run container
docker run -p 8080:8080 \
  -e USER_SERVICE_URL=http://user-service:8081 \
  -e COURSE_SERVICE_URL=http://course-service:8082 \
  -e PROGRESS_SERVICE_URL=http://progress-service:8083 \
  -e JWT_SECRET=your-secret-key \
  gopro-api-gateway
```

### Using Docker Compose
```bash
# From services directory
docker-compose up api-gateway
```

## API Endpoints

### Health Check
```bash
GET /health
```

Response:
```json
{
  "status": "healthy",
  "service": "api-gateway",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

### Readiness Check
```bash
GET /ready
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
  "timestamp": "2024-01-01T00:00:00Z"
}
```

### Proxied Endpoints

All `/api/v1/*` endpoints are proxied to backend services.

Example:
```bash
# Login (proxied to User Service)
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}

# Get courses (proxied to Course Service)
GET /api/v1/courses
Authorization: Bearer <token>

# Get user progress (proxied to Progress Service)
GET /api/v1/progress/user/123
Authorization: Bearer <token>
```

## Authentication

The API Gateway validates JWT tokens and adds user information to request headers:

```
Authorization: Bearer <jwt-token>
```

After validation, the following headers are added to backend requests:
- `X-User-ID`: User ID from token
- `X-User-Role`: User role from token

## Middleware Chain

Requests pass through the following middleware:

1. **Request ID**: Adds unique request ID
2. **Logger**: Logs request details
3. **Recovery**: Recovers from panics
4. **CORS**: Handles cross-origin requests
5. **Timeout**: Enforces request timeout
6. **Service Info**: Adds service metadata headers
7. **Authentication**: Validates JWT tokens (for protected routes)
8. **Rate Limiting**: Enforces rate limits (optional)

## Error Handling

The API Gateway returns standard HTTP error codes:

- `400 Bad Request` - Invalid request
- `401 Unauthorized` - Missing or invalid authentication
- `403 Forbidden` - Insufficient permissions
- `404 Not Found` - Service or endpoint not found
- `429 Too Many Requests` - Rate limit exceeded
- `500 Internal Server Error` - Server error
- `502 Bad Gateway` - Backend service error
- `503 Service Unavailable` - Service not ready
- `504 Gateway Timeout` - Request timeout

## Monitoring

### Metrics
- Request count by endpoint
- Request duration
- Error rate
- Backend service health

### Logging
All requests are logged with:
- Request ID
- Method and path
- Status code
- Duration
- User ID (if authenticated)

### Tracing
Distributed tracing is enabled with OpenTelemetry:
- Trace ID propagation
- Span creation for each request
- Backend service correlation

## Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test -run TestAuthMiddleware ./internal/auth/
```

## Development

### Project Structure
```
api-gateway/
├── cmd/
│   └── main.go              # Entry point
├── internal/
│   ├── config/              # Configuration
│   ├── handler/             # HTTP handlers
│   ├── proxy/               # Reverse proxy
│   └── auth/                # Authentication
├── Dockerfile               # Container image
├── go.mod                   # Dependencies
└── README.md                # This file
```

### Adding New Routes

1. Update routing rules in `internal/proxy/proxy.go`
2. Add authentication requirements in `internal/handler/handler.go`
3. Update documentation

### Adding New Middleware

1. Create middleware function in `internal/middleware/`
2. Add to middleware chain in `internal/handler/handler.go`
3. Update configuration if needed

## Security

- JWT tokens are validated on every request
- Secrets are loaded from environment variables
- HTTPS is recommended for production
- Rate limiting prevents abuse
- CORS is configurable
- Request timeouts prevent resource exhaustion

## Performance

- Connection pooling for backend services
- Circuit breaker prevents cascading failures
- Request timeout prevents hanging requests
- Efficient routing with minimal overhead
- Stateless design for horizontal scaling

## Troubleshooting

### Gateway returns 502 Bad Gateway
- Check if backend services are running
- Verify service URLs in configuration
- Check network connectivity

### Gateway returns 401 Unauthorized
- Verify JWT secret matches User Service
- Check token expiration
- Ensure Authorization header is present

### Gateway returns 504 Gateway Timeout
- Increase timeout configuration
- Check backend service performance
- Verify network latency

## Production Deployment

### Environment Variables
Set all required environment variables:
```bash
export PORT=8080
export USER_SERVICE_URL=http://user-service:8081
export COURSE_SERVICE_URL=http://course-service:8082
export PROGRESS_SERVICE_URL=http://progress-service:8083
export JWT_SECRET=<strong-secret-key>
export REDIS_URL=redis:6379
export LOG_LEVEL=info
```

### Health Checks
Configure Kubernetes liveness and readiness probes:
```yaml
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 10
  periodSeconds: 30

readinessProbe:
  httpGet:
    path: /ready
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
```

### Scaling
The API Gateway is stateless and can be scaled horizontally:
```bash
kubectl scale deployment api-gateway --replicas=3
```

