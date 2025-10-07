# GO-PRO Microservices Deployment Guide

This guide covers deploying the GO-PRO microservices architecture in various environments.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Local Development](#local-development)
3. [Docker Deployment](#docker-deployment)
4. [Kubernetes Deployment](#kubernetes-deployment)
5. [Production Deployment](#production-deployment)
6. [Monitoring & Observability](#monitoring--observability)
7. [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Tools
- Docker 24.0+
- Docker Compose 2.20+
- Go 1.22+
- kubectl 1.28+ (for Kubernetes)
- Helm 3.12+ (for Kubernetes)

### Optional Tools
- k9s (Kubernetes CLI)
- Lens (Kubernetes IDE)
- Postman (API testing)

## Local Development

### 1. Clone Repository
```bash
git clone https://github.com/DimaJoyti/go-pro.git
cd go-pro/services
```

### 2. Set Environment Variables
```bash
# Copy example environment file
cp .env.example .env

# Edit environment variables
vim .env
```

### 3. Start Infrastructure
```bash
# Start PostgreSQL, Redis, Kafka
docker-compose up -d postgres redis kafka zookeeper

# Wait for services to be ready
make health
```

### 4. Run Database Migrations
```bash
make db-migrate
```

### 5. Start Services

**Option A: Using Docker Compose (Recommended)**
```bash
# Start all services
make up

# View logs
make logs

# Check health
make health
```

**Option B: Running Services Individually**
```bash
# Terminal 1: API Gateway
cd api-gateway
go run cmd/main.go

# Terminal 2: User Service
cd user-service
go run cmd/main.go

# Terminal 3: Course Service
cd course-service
go run cmd/main.go

# Terminal 4: Progress Service
cd progress-service
go run cmd/main.go
```

### 6. Verify Deployment
```bash
# Check API Gateway
curl http://localhost:8080/health

# Check User Service
curl http://localhost:8081/health

# Check Course Service
curl http://localhost:8082/health

# Check Progress Service
curl http://localhost:8083/health
```

### 7. Access Monitoring Tools
- **Jaeger UI**: http://localhost:16686
- **Prometheus**: http://localhost:9090
- **Grafana**: http://localhost:3000 (admin/admin)

## Docker Deployment

### Build Images
```bash
# Build all services
make build

# Build specific service
make build-api-gateway
make build-user
make build-course
make build-progress
```

### Run with Docker Compose
```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Clean up volumes
docker-compose down -v
```

### Environment Configuration
Create a `.env` file:
```env
# Server
PORT=8080

# Database
DB_HOST=postgres
DB_PORT=5432
DB_USER=gopro
DB_PASSWORD=gopro_password
DB_NAME=gopro

# Redis
REDIS_URL=redis:6379

# Kafka
KAFKA_BROKERS=kafka:9092

# JWT
JWT_SECRET=your-secret-key-change-in-production

# Services
USER_SERVICE_URL=http://user-service:8081
COURSE_SERVICE_URL=http://course-service:8082
PROGRESS_SERVICE_URL=http://progress-service:8083
```

## Kubernetes Deployment

### 1. Create Namespace
```bash
kubectl create namespace gopro
kubectl config set-context --current --namespace=gopro
```

### 2. Create Secrets
```bash
# Database credentials
kubectl create secret generic db-credentials \
  --from-literal=username=gopro \
  --from-literal=password=gopro_password

# JWT secret
kubectl create secret generic jwt-secret \
  --from-literal=secret=your-secret-key

# Redis password (if needed)
kubectl create secret generic redis-credentials \
  --from-literal=password=redis_password
```

### 3. Deploy Infrastructure

**PostgreSQL**
```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install postgres bitnami/postgresql \
  --set auth.username=gopro \
  --set auth.password=gopro_password \
  --set auth.database=gopro \
  --set primary.persistence.size=10Gi
```

**Redis**
```bash
helm install redis bitnami/redis \
  --set auth.enabled=false \
  --set master.persistence.size=5Gi
```

**Kafka**
```bash
helm install kafka bitnami/kafka \
  --set persistence.size=10Gi \
  --set zookeeper.persistence.size=5Gi
```

### 4. Deploy Services

**API Gateway**
```yaml
# api-gateway-deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: api-gateway
spec:
  replicas: 3
  selector:
    matchLabels:
      app: api-gateway
  template:
    metadata:
      labels:
        app: api-gateway
    spec:
      containers:
      - name: api-gateway
        image: gopro-api-gateway:latest
        ports:
        - containerPort: 8080
        env:
        - name: PORT
          value: "8080"
        - name: USER_SERVICE_URL
          value: "http://user-service:8081"
        - name: COURSE_SERVICE_URL
          value: "http://course-service:8082"
        - name: PROGRESS_SERVICE_URL
          value: "http://progress-service:8083"
        - name: JWT_SECRET
          valueFrom:
            secretKeyRef:
              name: jwt-secret
              key: secret
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
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "256Mi"
            cpu: "200m"
---
apiVersion: v1
kind: Service
metadata:
  name: api-gateway
spec:
  selector:
    app: api-gateway
  ports:
  - port: 8080
    targetPort: 8080
  type: LoadBalancer
```

Apply deployment:
```bash
kubectl apply -f api-gateway-deployment.yaml
```

### 5. Deploy Monitoring

**Prometheus**
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm install prometheus prometheus-community/prometheus \
  --set server.persistentVolume.size=10Gi
```

**Grafana**
```bash
helm install grafana grafana/grafana \
  --set persistence.enabled=true \
  --set persistence.size=5Gi \
  --set adminPassword=admin
```

**Jaeger**
```bash
helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm install jaeger jaegertracing/jaeger
```

### 6. Configure Ingress
```yaml
# ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: gopro-ingress
  annotations:
    kubernetes.io/ingress.class: nginx
    cert-manager.io/cluster-issuer: letsencrypt-prod
spec:
  tls:
  - hosts:
    - api.gopro.com
    secretName: gopro-tls
  rules:
  - host: api.gopro.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: api-gateway
            port:
              number: 8080
```

Apply ingress:
```bash
kubectl apply -f ingress.yaml
```

## Production Deployment

### 1. Pre-Deployment Checklist

- [ ] All environment variables configured
- [ ] Secrets properly managed
- [ ] Database migrations tested
- [ ] Load testing completed
- [ ] Security scanning passed
- [ ] Monitoring configured
- [ ] Backup strategy in place
- [ ] Disaster recovery plan documented
- [ ] SSL certificates configured
- [ ] DNS records updated

### 2. Environment Configuration

**Production Environment Variables**
```env
# Server
PORT=8080
LOG_LEVEL=info
LOG_FORMAT=json

# Database (use managed service)
DB_HOST=prod-postgres.example.com
DB_PORT=5432
DB_USER=gopro_prod
DB_PASSWORD=<strong-password>
DB_NAME=gopro_prod
DB_SSL_MODE=require

# Redis (use managed service)
REDIS_URL=prod-redis.example.com:6379
REDIS_PASSWORD=<strong-password>
REDIS_TLS=true

# Kafka (use managed service)
KAFKA_BROKERS=prod-kafka-1:9092,prod-kafka-2:9092,prod-kafka-3:9092
KAFKA_TLS=true

# JWT
JWT_SECRET=<strong-secret-key>
TOKEN_DURATION=24h

# Rate Limiting
RATE_LIMIT_RPM=100
RATE_LIMIT_BURST=20

# Timeouts
READ_TIMEOUT=30s
WRITE_TIMEOUT=30s
SHUTDOWN_TIMEOUT=60s
```

### 3. Database Setup

**Run Migrations**
```bash
# Connect to production database
psql -h prod-postgres.example.com -U gopro_prod -d gopro_prod

# Run migration script
\i init-db.sql

# Verify tables
\dt users.*
\dt courses.*
\dt progress.*
```

### 4. Deployment Strategy

**Blue-Green Deployment**
```bash
# Deploy new version (green)
kubectl apply -f deployments/green/

# Test green deployment
curl https://green.api.gopro.com/health

# Switch traffic to green
kubectl patch service api-gateway -p '{"spec":{"selector":{"version":"green"}}}'

# Monitor for issues
kubectl logs -f deployment/api-gateway-green

# Rollback if needed
kubectl patch service api-gateway -p '{"spec":{"selector":{"version":"blue"}}}'
```

**Canary Deployment**
```bash
# Deploy canary with 10% traffic
kubectl apply -f deployments/canary/

# Monitor metrics
# If successful, gradually increase traffic
# If issues, rollback immediately
```

### 5. Scaling Configuration

**Horizontal Pod Autoscaler**
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: api-gateway-hpa
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: api-gateway
  minReplicas: 3
  maxReplicas: 10
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 70
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 80
```

## Monitoring & Observability

### Metrics to Monitor

**Service Metrics**
- Request rate (requests/second)
- Error rate (errors/second)
- Response time (p50, p95, p99)
- Availability (uptime percentage)

**Infrastructure Metrics**
- CPU usage
- Memory usage
- Disk I/O
- Network I/O

**Business Metrics**
- User registrations
- Course enrollments
- Lesson completions
- Active users

### Alerting Rules

**Critical Alerts**
- Service down (availability < 99%)
- High error rate (> 5%)
- Database connection failures
- Kafka consumer lag > 1000

**Warning Alerts**
- High response time (p95 > 1s)
- High CPU usage (> 80%)
- High memory usage (> 85%)
- Disk space low (< 20%)

## Troubleshooting

### Common Issues

**Service Not Starting**
```bash
# Check logs
kubectl logs deployment/api-gateway

# Check events
kubectl describe pod api-gateway-xxx

# Check configuration
kubectl get configmap
kubectl get secret
```

**Database Connection Issues**
```bash
# Test database connectivity
kubectl run -it --rm debug --image=postgres:15 --restart=Never -- \
  psql -h postgres -U gopro -d gopro

# Check database logs
kubectl logs deployment/postgres
```

**High Latency**
```bash
# Check service metrics
kubectl top pods

# Check database performance
# Check Redis performance
# Check network latency
```

### Health Check Commands

```bash
# Check all services
make health

# Check specific service
curl http://api-gateway:8080/health
curl http://user-service:8081/health
curl http://course-service:8082/health
curl http://progress-service:8083/health

# Check database
psql -h postgres -U gopro -c "SELECT 1"

# Check Redis
redis-cli ping

# Check Kafka
kafka-topics.sh --list --bootstrap-server kafka:9092
```

## Backup & Recovery

### Database Backup
```bash
# Automated daily backups
pg_dump -h postgres -U gopro gopro > backup-$(date +%Y%m%d).sql

# Restore from backup
psql -h postgres -U gopro gopro < backup-20240101.sql
```

### Disaster Recovery
1. Restore database from latest backup
2. Deploy services from last known good version
3. Verify all services are healthy
4. Restore Redis cache (optional)
5. Resume Kafka consumers
6. Monitor for issues

## Security Checklist

- [ ] All secrets stored in secret management system
- [ ] TLS enabled for all services
- [ ] Database connections encrypted
- [ ] JWT tokens properly validated
- [ ] Rate limiting configured
- [ ] CORS properly configured
- [ ] Security headers added
- [ ] Regular security scans
- [ ] Dependency updates automated
- [ ] Access logs enabled

