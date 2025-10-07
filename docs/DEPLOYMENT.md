# Deployment Guide

This guide covers deployment options and strategies for the GO-PRO Learning Platform Backend.

## Deployment Options

### 1. Docker Compose (Recommended for Development/Small Production)

#### Development Environment
```bash
# Start development environment
make docker-dev

# Services available:
# - Backend API: http://localhost:8080
# - Database Admin: http://localhost:8081
# - Redis Commander: http://localhost:8082
# - Prometheus: http://localhost:9090
# - Grafana: http://localhost:3000
```

#### Production Environment
```bash
# Set environment variables
export POSTGRES_PASSWORD=your-secure-password

# Start production environment
make docker-prod
```

### 2. Kubernetes (Recommended for Production)

#### Prerequisites
- Kubernetes cluster (v1.24+)
- kubectl configured
- NGINX Ingress Controller
- cert-manager (for TLS certificates)

#### Deployment Steps

1. **Create secrets**:
```bash
# Create namespace
kubectl apply -f deploy/k8s/namespace.yaml

# Create secrets
kubectl create secret generic go-pro-secrets \
  --from-literal=database-url="postgres://user:pass@host:5432/db" \
  --from-literal=redis-url="redis://host:6379" \
  --from-literal=jwt-secret="your-jwt-secret" \
  -n go-pro

# Create image pull secret (if using private registry)
kubectl create secret docker-registry regcred \
  --docker-server=ghcr.io \
  --docker-username=your-username \
  --docker-password=your-token \
  -n go-pro
```

2. **Deploy the application**:
```bash
# Deploy application
kubectl apply -f deploy/k8s/deployment.yaml
kubectl apply -f deploy/k8s/service.yaml
kubectl apply -f deploy/k8s/ingress.yaml

# Check deployment status
kubectl get pods -n go-pro
kubectl get services -n go-pro
kubectl get ingress -n go-pro
```

3. **Verify deployment**:
```bash
# Check application health
curl -k https://api.go-pro.example.com/api/v1/health

# View logs
kubectl logs -f deployment/go-pro-backend -n go-pro
```

### 3. Cloud Platforms

#### AWS ECS
```bash
# Build and push image
make docker-build
docker tag go-pro-backend:latest your-account.dkr.ecr.region.amazonaws.com/go-pro:latest
docker push your-account.dkr.ecr.region.amazonaws.com/go-pro:latest

# Deploy using ECS CLI or AWS Console
```

#### Google Cloud Run
```bash
# Build and push to GCR
make docker-build
docker tag go-pro-backend:latest gcr.io/your-project/go-pro:latest
docker push gcr.io/your-project/go-pro:latest

# Deploy to Cloud Run
gcloud run deploy go-pro-backend \
  --image gcr.io/your-project/go-pro:latest \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated
```

#### Azure Container Instances
```bash
# Build and push to ACR
make docker-build
az acr build --registry your-registry --image go-pro:latest .

# Deploy to ACI
az container create \
  --resource-group your-rg \
  --name go-pro-backend \
  --image your-registry.azurecr.io/go-pro:latest \
  --dns-name-label go-pro-api \
  --ports 8080
```

## Environment Configuration

### Required Environment Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `GO_ENV` | Environment mode | `production` |
| `PORT` | Server port | `8080` |
| `LOG_LEVEL` | Logging level | `info` |
| `DATABASE_URL` | PostgreSQL connection string | `postgres://user:pass@host:5432/db` |
| `REDIS_URL` | Redis connection string | `redis://host:6379` |
| `JWT_SECRET` | JWT signing secret | `your-secure-secret` |

### Optional Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `READ_TIMEOUT` | `15s` | HTTP read timeout |
| `WRITE_TIMEOUT` | `15s` | HTTP write timeout |
| `IDLE_TIMEOUT` | `60s` | HTTP idle timeout |
| `MAX_HEADER_BYTES` | `1MB` | Maximum header size |

## Health Checks and Monitoring

### Health Check Endpoint
```bash
GET /api/v1/health
```

Response:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T00:00:00Z",
  "version": "1.0.0"
}
```

### Metrics Endpoint
```bash
GET /metrics
```

Prometheus format metrics for:
- HTTP request duration
- HTTP request count
- Active connections
- Memory usage
- Go runtime metrics

### Logging

Application logs are structured JSON with fields:
- `timestamp`: ISO 8601 timestamp
- `level`: Log level (debug, info, warn, error)
- `message`: Log message
- `method`: HTTP method (for HTTP logs)
- `path`: Request path (for HTTP logs)
- `status`: HTTP status code (for HTTP logs)
- `duration`: Request duration (for HTTP logs)

## Security Considerations

### Container Security
- Runs as non-root user (UID 1001)
- Read-only root filesystem
- No privileged capabilities
- Minimal base image (Alpine Linux)

### Network Security
- TLS termination at ingress/load balancer
- Rate limiting configured
- CORS properly configured
- No direct database access from outside

### Secrets Management
- Database credentials in Kubernetes secrets
- JWT secrets in environment variables
- No secrets in container images
- Regular secret rotation recommended

## Scaling and Performance

### Horizontal Scaling
```bash
# Scale Kubernetes deployment
kubectl scale deployment go-pro-backend --replicas=5 -n go-pro

# Docker Compose scaling
docker-compose -f docker-compose.prod.yml up -d --scale go-pro-backend=3
```

### Performance Tuning

#### Application Level
- Connection pooling for database
- Redis caching for frequently accessed data
- Gzip compression for responses
- Request timeouts and rate limiting

#### Infrastructure Level
- Load balancer health checks
- Database read replicas
- Redis clustering
- CDN for static assets

### Resource Requirements

#### Minimum (Development)
- CPU: 0.25 cores
- Memory: 256MB
- Storage: 1GB

#### Recommended (Production)
- CPU: 0.5 cores per instance
- Memory: 512MB per instance
- Storage: 10GB (logs and temp files)

#### High Traffic (Production)
- CPU: 1 core per instance
- Memory: 1GB per instance
- Storage: 20GB (logs and temp files)

## Backup and Disaster Recovery

### Database Backup
```bash
# PostgreSQL backup
pg_dump -h host -U user -d gopro_prod > backup.sql

# Automated backup with cron
0 2 * * * pg_dump -h host -U user -d gopro_prod | gzip > /backups/gopro_$(date +\%Y\%m\%d).sql.gz
```

### Application State
- Stateless application design
- Configuration in environment variables
- User uploaded files in object storage (if applicable)

### Recovery Procedures
1. Restore database from backup
2. Deploy latest application version
3. Verify health checks
4. Update DNS if necessary

## Monitoring and Alerting

### Prometheus Metrics
- `http_requests_total`: Total HTTP requests
- `http_request_duration_seconds`: Request duration histogram
- `go_memstats_alloc_bytes`: Memory allocation
- `go_goroutines`: Number of goroutines

### Grafana Dashboards
- Application overview
- HTTP request metrics
- System resource usage
- Error rates and response times

### Alerting Rules
```yaml
# Example Prometheus alerting rules
groups:
- name: go-pro-backend
  rules:
  - alert: HighErrorRate
    expr: rate(http_requests_total{status=~"5.."}[5m]) > 0.1
    for: 2m
    labels:
      severity: warning
    annotations:
      summary: High error rate detected

  - alert: HighResponseTime
    expr: histogram_quantile(0.95, http_request_duration_seconds_bucket) > 1
    for: 5m
    labels:
      severity: warning
    annotations:
      summary: High response time detected
```

## Troubleshooting

### Common Issues

#### Application Won't Start
```bash
# Check logs
kubectl logs deployment/go-pro-backend -n go-pro

# Check configuration
kubectl describe pod <pod-name> -n go-pro

# Check secrets
kubectl get secrets -n go-pro
```

#### High Memory Usage
```bash
# Monitor memory usage
kubectl top pods -n go-pro

# Check for memory leaks in logs
kubectl logs deployment/go-pro-backend -n go-pro | grep -i "memory\|leak"
```

#### Database Connection Issues
```bash
# Test database connectivity
kubectl exec -it deployment/go-pro-backend -n go-pro -- /bin/sh
# Inside container:
# wget -qO- http://localhost:8080/api/v1/health
```

#### Performance Issues
```bash
# Check resource usage
kubectl top pods -n go-pro

# Monitor metrics
curl http://localhost:9090/metrics

# Check slow queries in database
# Monitor Grafana dashboards
```

### Rolling Back Deployments

#### Kubernetes
```bash
# View rollout history
kubectl rollout history deployment/go-pro-backend -n go-pro

# Rollback to previous version
kubectl rollout undo deployment/go-pro-backend -n go-pro

# Rollback to specific revision
kubectl rollout undo deployment/go-pro-backend --to-revision=2 -n go-pro
```

#### Docker Compose
```bash
# Update image tag in compose file
# Then restart services
docker-compose -f docker-compose.prod.yml up -d
```

## Maintenance

### Regular Tasks
- Update base images monthly
- Rotate secrets quarterly
- Update dependencies monthly
- Review and update scaling policies
- Monitor and analyze performance metrics
- Test backup and recovery procedures

### Updates and Patches
1. Test updates in staging environment
2. Schedule maintenance window
3. Deploy with rolling update strategy
4. Monitor health checks and metrics
5. Rollback if issues detected

### Database Maintenance
- Regular VACUUM and ANALYZE
- Monitor query performance
- Update statistics
- Check for long-running queries
- Review and optimize indexes

## Support and Maintenance

For deployment issues:
1. Check application logs
2. Verify configuration
3. Test connectivity
4. Review monitoring dashboards
5. Consult this documentation
6. Contact the development team