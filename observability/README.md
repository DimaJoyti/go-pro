# GO-PRO Observability with OpenTelemetry

This directory contains the comprehensive observability implementation for the GO-PRO learning platform using OpenTelemetry for distributed tracing, metrics collection, and structured logging.

## Architecture Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                    Application Services                          │
│  (Backend, Frontend, Microservices)                             │
└────────────────────────┬────────────────────────────────────────┘
                         │ OpenTelemetry SDK
                         │ (Traces, Metrics, Logs)
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│              OpenTelemetry Collector                             │
│  - Receive (OTLP, Jaeger, Zipkin, Prometheus)                  │
│  - Process (Batch, Filter, Transform)                           │
│  - Export (Jaeger, Prometheus, Loki, Cloud providers)          │
└────────────────────────┬────────────────────────────────────────┘
                         │
         ┌───────────────┼───────────────┬──────────────┐
         │               │               │              │
         ▼               ▼               ▼              ▼
┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐
│   Jaeger    │  │ Prometheus  │  │    Loki     │  │   Cloud     │
│  (Traces)   │  │  (Metrics)  │  │   (Logs)    │  │ Providers   │
└──────┬──────┘  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘
       │                │                │                │
       └────────────────┴────────────────┴────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Grafana                                     │
│  - Distributed Tracing (Jaeger)                                │
│  - Metrics Dashboards (Prometheus)                              │
│  - Log Aggregation (Loki)                                       │
│  - Unified Observability                                        │
└─────────────────────────────────────────────────────────────────┘
```

## Features

### Distributed Tracing
- ✅ End-to-end request tracing across services
- ✅ Automatic context propagation
- ✅ Span attributes and events
- ✅ Error tracking and stack traces
- ✅ Performance bottleneck identification
- ✅ Service dependency mapping

### Metrics Collection
- ✅ Request rate, duration, and error rate (RED metrics)
- ✅ Resource utilization (CPU, memory, disk)
- ✅ Custom business metrics
- ✅ Histogram and summary metrics
- ✅ Multi-dimensional metrics with labels
- ✅ Automatic metric aggregation

### Structured Logging
- ✅ JSON-formatted logs
- ✅ Trace context correlation
- ✅ Log levels (debug, info, warn, error)
- ✅ Structured fields for filtering
- ✅ Centralized log aggregation
- ✅ Log sampling for high-volume services

## Components

### 1. OpenTelemetry SDK Integration

**Go Backend**:
```go
// otel/tracing.go - Tracing setup
// otel/metrics.go - Metrics setup
// otel/logging.go - Logging setup
```

**Next.js Frontend**:
```typescript
// otel/tracing.ts - Browser tracing
// otel/metrics.ts - Web vitals metrics
```

### 2. OpenTelemetry Collector

**Configuration**: `configs/otel-collector-config.yaml`
- Receivers: OTLP, Jaeger, Zipkin, Prometheus
- Processors: Batch, Filter, Transform, Sampling
- Exporters: Jaeger, Prometheus, Loki, CloudWatch, Cloud Trace

### 3. Observability Stack

**Jaeger**: Distributed tracing backend
**Prometheus**: Metrics storage and querying
**Loki**: Log aggregation
**Grafana**: Unified visualization

### 4. Dashboards

**Pre-built Grafana Dashboards**:
- Service Overview
- Request Performance
- Error Tracking
- Resource Utilization
- Business Metrics
- SLO/SLA Monitoring

## Quick Start

### 1. Deploy OpenTelemetry Collector

```bash
# Using Kubernetes
kubectl apply -f k8s/observability/

# Using Docker Compose
docker-compose -f docker-compose.observability.yml up -d
```

### 2. Configure Application

**Backend (Go)**:
```go
import "go-pro/observability/otel"

func main() {
    // Initialize OpenTelemetry
    shutdown := otel.InitTracing("backend-service")
    defer shutdown()
    
    shutdown = otel.InitMetrics("backend-service")
    defer shutdown()
    
    // Your application code
}
```

**Frontend (Next.js)**:
```typescript
import { initTracing } from '@/observability/otel/tracing';

// In _app.tsx
initTracing('frontend-service');
```

### 3. Access Dashboards

```bash
# Jaeger UI
http://localhost:16686

# Prometheus UI
http://localhost:9090

# Grafana
http://localhost:3000
# Default credentials: admin/admin
```

## Instrumentation Guide

### HTTP Server Instrumentation

**Go (net/http)**:
```go
import (
    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
    handler := otelhttp.NewHandler(
        http.HandlerFunc(handleRequest),
        "http-server",
    )
    http.ListenAndServe(":8080", handler)
}
```

### HTTP Client Instrumentation

**Go**:
```go
client := &http.Client{
    Transport: otelhttp.NewTransport(http.DefaultTransport),
}
```

### Database Instrumentation

**PostgreSQL (Go)**:
```go
import (
    "github.com/uptrace/opentelemetry-go-extra/otelgorm"
)

db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err := db.Use(otelgorm.NewPlugin()); err != nil {
    panic(err)
}
```

### Redis Instrumentation

**Go**:
```go
import (
    "github.com/go-redis/redis/extra/redisotel/v8"
)

rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})
rdb.AddHook(redisotel.NewTracingHook())
```

### Custom Spans

**Go**:
```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/attribute"
)

func processOrder(ctx context.Context, orderID string) error {
    tracer := otel.Tracer("order-service")
    ctx, span := tracer.Start(ctx, "processOrder")
    defer span.End()
    
    span.SetAttributes(
        attribute.String("order.id", orderID),
        attribute.String("order.status", "processing"),
    )
    
    // Business logic
    if err := validateOrder(ctx, orderID); err != nil {
        span.RecordError(err)
        span.SetStatus(codes.Error, err.Error())
        return err
    }
    
    return nil
}
```

### Custom Metrics

**Go**:
```go
import (
    "go.opentelemetry.io/otel/metric"
)

var (
    meter = otel.Meter("order-service")
    orderCounter, _ = meter.Int64Counter(
        "orders.processed",
        metric.WithDescription("Number of orders processed"),
    )
)

func processOrder(ctx context.Context, orderID string) {
    orderCounter.Add(ctx, 1,
        metric.WithAttributes(
            attribute.String("status", "success"),
        ),
    )
}
```

### Structured Logging

**Go**:
```go
import (
    "go.uber.org/zap"
    "go.opentelemetry.io/otel/trace"
)

func handleRequest(ctx context.Context) {
    logger := zap.L()
    
    // Add trace context to logs
    span := trace.SpanFromContext(ctx)
    logger = logger.With(
        zap.String("trace_id", span.SpanContext().TraceID().String()),
        zap.String("span_id", span.SpanContext().SpanID().String()),
    )
    
    logger.Info("Processing request",
        zap.String("user_id", "123"),
        zap.String("action", "create_order"),
    )
}
```

## Metrics Reference

### RED Metrics (Requests, Errors, Duration)

**Request Rate**:
```promql
rate(http_requests_total[5m])
```

**Error Rate**:
```promql
rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m])
```

**Duration (p99)**:
```promql
histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m]))
```

### USE Metrics (Utilization, Saturation, Errors)

**CPU Utilization**:
```promql
rate(process_cpu_seconds_total[5m])
```

**Memory Utilization**:
```promql
process_resident_memory_bytes / node_memory_MemTotal_bytes
```

**Disk I/O**:
```promql
rate(node_disk_io_time_seconds_total[5m])
```

## Alerting

### Example Alerts

**High Error Rate**:
```yaml
- alert: HighErrorRate
  expr: rate(http_requests_total{status=~"5.."}[5m]) / rate(http_requests_total[5m]) > 0.05
  for: 5m
  labels:
    severity: critical
  annotations:
    summary: "High error rate detected"
    description: "Error rate is {{ $value | humanizePercentage }}"
```

**High Latency**:
```yaml
- alert: HighLatency
  expr: histogram_quantile(0.99, rate(http_request_duration_seconds_bucket[5m])) > 1
  for: 10m
  labels:
    severity: warning
  annotations:
    summary: "High latency detected"
    description: "P99 latency is {{ $value }}s"
```

## Best Practices

### Tracing
✅ Use semantic conventions for span names
✅ Add meaningful attributes to spans
✅ Record errors with stack traces
✅ Use sampling for high-volume services
✅ Propagate context across service boundaries
✅ Keep span names consistent

### Metrics
✅ Use descriptive metric names
✅ Add relevant labels (but avoid high cardinality)
✅ Use histograms for latency metrics
✅ Use counters for event counts
✅ Use gauges for current values
✅ Aggregate metrics at collection time

### Logging
✅ Use structured logging (JSON)
✅ Include trace context in logs
✅ Use appropriate log levels
✅ Avoid logging sensitive data
✅ Use log sampling for high-volume logs
✅ Include relevant context fields

## Performance Considerations

### Sampling

**Head-based Sampling**:
```yaml
# Sample 10% of traces
processors:
  probabilistic_sampler:
    sampling_percentage: 10
```

**Tail-based Sampling**:
```yaml
# Sample all errors and slow requests
processors:
  tail_sampling:
    policies:
      - name: error-traces
        type: status_code
        status_code: {status_codes: [ERROR]}
      - name: slow-traces
        type: latency
        latency: {threshold_ms: 1000}
```

### Batching

```yaml
processors:
  batch:
    timeout: 10s
    send_batch_size: 1024
```

## Troubleshooting

### Common Issues

**1. Missing Traces**:
- Check collector is running
- Verify OTLP endpoint configuration
- Check network connectivity
- Review sampling configuration

**2. High Memory Usage**:
- Reduce batch size
- Increase export interval
- Enable sampling
- Limit span attributes

**3. Missing Metrics**:
- Verify Prometheus scrape configuration
- Check metric names and labels
- Review metric export interval
- Check collector logs

## Additional Resources

- [OpenTelemetry Documentation](https://opentelemetry.io/docs/)
- [OpenTelemetry Go SDK](https://github.com/open-telemetry/opentelemetry-go)
- [OpenTelemetry Collector](https://github.com/open-telemetry/opentelemetry-collector)
- [Jaeger Documentation](https://www.jaegertracing.io/docs/)
- [Prometheus Documentation](https://prometheus.io/docs/)

