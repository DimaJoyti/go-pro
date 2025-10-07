# Development Setup Guide

This guide will help you set up the development environment for the GO-PRO Learning Platform Backend.

## Prerequisites

- **Go**: Version 1.21 or higher
- **Docker**: Version 20.10 or higher
- **Docker Compose**: Version 2.0 or higher
- **Make**: For running automation commands
- **Git**: For version control
- **Python**: Version 3.8+ (for pre-commit hooks)

## Quick Start

1. **Clone and setup the repository**:
   ```bash
   git clone <your-repo-url>
   cd go-pro
   make setup
   ```

2. **Start the development environment**:
   ```bash
   make dev-docker
   ```

3. **Access the services**:
   - Backend API: <http://localhost:8080>
   - Database Admin: <http://localhost:8081>
   - Redis Commander: <http://localhost:8082>
   - Prometheus: <http://localhost:9090>
   - Grafana: <http://localhost:3000>

## Development Workflow

### Local Development (without Docker)

1. **Install dependencies**:
   ```bash
   make deps
   ```

2. **Start the backend with hot reload**:
   ```bash
   make dev
   ```

The API will be available at <http://localhost:8080> with automatic reloading on code changes.

### Docker Development Environment

For a complete development setup with all services:

```bash
# Start all services
make docker-dev

# View logs
make logs

# Stop all services
make docker-stop
```

### Code Quality Checks

Run all quality checks before committing:

```bash
# Format code
make fmt

# Run linter
make lint

# Run security scan
make security

# Run tests
make test

# Run all quality checks
make quality
```

### Pre-commit Hooks

Pre-commit hooks are automatically installed with `make setup`. They run:

- Go formatting and imports
- Linting with golangci-lint
- Security scanning
- Tests
- Secret detection

To run hooks manually:

```bash
make pre-commit
```

## Testing

### Unit Tests

```bash
# Run unit tests
make test

# Run tests with coverage
make test-coverage

# Run integration tests
make test-integration
```

### Load Testing

```bash
# Start the backend
make dev

# In another terminal, run load tests
make test-load
```

## Building and Deployment

### Local Build

```bash
# Build the application
make build

# Build for multiple platforms
make build-all
```

### Docker Build

```bash
# Build Docker image
make docker-build

# Build multi-platform image
make docker-build-multi
```

### Production Environment

```bash
# Start production stack
make docker-prod

# Check application health
make health
```

## Database Management

### Development Database

The development environment includes:
- PostgreSQL database
- Adminer for database management
- Redis for caching

Access Adminer at <http://localhost:8081> with:
- Server: `postgres`
- Username: `gopro_user`
- Password: `gopro_password`
- Database: `gopro_dev`

### Migrations

```bash
# Run migrations (when implemented)
make db-migrate

# Seed database with test data
make db-seed
```

## Monitoring and Observability

### Metrics

- **Prometheus**: <http://localhost:9090>
- **Grafana**: <http://localhost:3000> (admin/admin)

### Logs

```bash
# View application logs
make logs

# View production logs
make logs-prod
```

## Common Commands

| Command | Description |
|---------|-------------|
| `make help` | Show all available commands |
| `make dev` | Start development server with hot reload |
| `make test` | Run unit tests |
| `make lint` | Run code linter |
| `make build` | Build the application |
| `make docker-dev` | Start full development environment |
| `make clean` | Clean build artifacts |
| `make health` | Check application health |

## IDE Setup

### VS Code

Recommended extensions:
- Go (official)
- Docker
- YAML
- Markdown All in One
- GitLens
- Thunder Client (API testing)

### Settings

Create `.vscode/settings.json`:

```json
{
  "go.formatTool": "goimports",
  "go.lintTool": "golangci-lint",
  "go.vetOnSave": "package",
  "go.testFlags": ["-v", "-race"],
  "go.coverOnSave": true,
  "files.eol": "\n",
  "files.insertFinalNewline": true,
  "files.trimTrailingWhitespace": true
}
```

## Troubleshooting

### Common Issues

1. **Port conflicts**: Ensure ports 8080, 5432, 6379, 9090, 3000 are available
2. **Docker issues**: Try `docker-compose down -v` to reset volumes
3. **Go modules**: Run `go mod tidy` if dependency issues occur
4. **Pre-commit hooks**: Run `pre-commit install --install-hooks` if hooks fail

### Reset Development Environment

```bash
# Stop all containers and remove volumes
make docker-stop
docker-compose -f docker-compose.dev.yml down -v

# Clean Go cache
make clean

# Restart
make docker-dev
```

### Debug Mode

To run the application in debug mode:

```bash
# Set debug environment variables
export GO_ENV=development
export LOG_LEVEL=debug

# Start with debugging
make dev
```

## API Documentation

The API documentation is available at <http://localhost:8080> when running the backend.

### API Testing

Use the included load testing script or tools like curl, Postman, or Thunder Client:

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Get courses
curl http://localhost:8080/api/v1/courses

# Get specific course
curl http://localhost:8080/api/v1/courses/go-pro
```

## Contributing

1. Create a feature branch from `develop`
2. Make your changes
3. Run quality checks: `make quality`
4. Commit your changes (pre-commit hooks will run)
5. Push and create a pull request

## Environment Variables

### Development

| Variable | Default | Description |
|----------|---------|-------------|
| `GO_ENV` | `development` | Environment mode |
| `PORT` | `8080` | Server port |
| `LOG_LEVEL` | `debug` | Logging level |

### Production

| Variable | Required | Description |
|----------|----------|-------------|
| `DATABASE_URL` | Yes | PostgreSQL connection string |
| `REDIS_URL` | Yes | Redis connection string |
| `GO_ENV` | Yes | Set to `production` |
| `PORT` | No | Server port (default: 8080) |
| `LOG_LEVEL` | No | Logging level (default: info) |

## Support

For questions and support:
1. Check this documentation
2. Review existing issues in the repository
3. Create a new issue with detailed information
4. Contact the development team