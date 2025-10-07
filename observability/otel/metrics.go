package otel

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MetricsConfig holds the configuration for OpenTelemetry metrics
type MetricsConfig struct {
	ServiceName    string
	ServiceVersion string
	Environment    string
	OTLPEndpoint   string
	ExportInterval time.Duration
}

// DefaultMetricsConfig returns a default metrics configuration
func DefaultMetricsConfig(serviceName string) *MetricsConfig {
	return &MetricsConfig{
		ServiceName:    serviceName,
		ServiceVersion: getEnv("SERVICE_VERSION", "1.0.0"),
		Environment:    getEnv("ENVIRONMENT", "development"),
		OTLPEndpoint:   getEnv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317"),
		ExportInterval: 60 * time.Second,
	}
}

// InitMetrics initializes OpenTelemetry metrics
func InitMetrics(serviceName string) (func(), error) {
	return InitMetricsWithConfig(DefaultMetricsConfig(serviceName))
}

// InitMetricsWithConfig initializes OpenTelemetry metrics with custom configuration
func InitMetricsWithConfig(config *MetricsConfig) (func(), error) {
	ctx := context.Background()

	// Create resource with service information
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(config.ServiceName),
			semconv.ServiceVersion(config.ServiceVersion),
			semconv.DeploymentEnvironment(config.Environment),
		),
		resource.WithHost(),
		resource.WithOS(),
		resource.WithProcess(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Create OTLP metric exporter
	metricExporter, err := createMetricExporter(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}

	// Create meter provider
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExporter,
				sdkmetric.WithInterval(config.ExportInterval),
			),
		),
	)

	// Set global meter provider
	otel.SetMeterProvider(mp)

	// Return shutdown function
	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mp.Shutdown(ctx); err != nil {
			fmt.Printf("Error shutting down meter provider: %v\n", err)
		}
	}, nil
}

// createMetricExporter creates an OTLP metric exporter
func createMetricExporter(ctx context.Context, config *MetricsConfig) (sdkmetric.Exporter, error) {
	// Create gRPC connection to OTLP collector
	conn, err := grpc.DialContext(ctx, config.OTLPEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection: %w", err)
	}

	// Create OTLP metric exporter
	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithGRPCConn(conn),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP exporter: %w", err)
	}

	return exporter, nil
}

// Meter returns a meter for the given name
func Meter(name string) metric.Meter {
	return otel.Meter(name)
}

// HTTPMetrics holds HTTP-related metrics
type HTTPMetrics struct {
	requestCounter    metric.Int64Counter
	requestDuration   metric.Float64Histogram
	requestSize       metric.Int64Histogram
	responseSize      metric.Int64Histogram
	activeRequests    metric.Int64UpDownCounter
}

// NewHTTPMetrics creates a new HTTPMetrics instance
func NewHTTPMetrics(serviceName string) (*HTTPMetrics, error) {
	meter := Meter(serviceName)

	requestCounter, err := meter.Int64Counter(
		"http.server.requests",
		metric.WithDescription("Total number of HTTP requests"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		return nil, err
	}

	requestDuration, err := meter.Float64Histogram(
		"http.server.duration",
		metric.WithDescription("HTTP request duration"),
		metric.WithUnit("ms"),
	)
	if err != nil {
		return nil, err
	}

	requestSize, err := meter.Int64Histogram(
		"http.server.request.size",
		metric.WithDescription("HTTP request size"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}

	responseSize, err := meter.Int64Histogram(
		"http.server.response.size",
		metric.WithDescription("HTTP response size"),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}

	activeRequests, err := meter.Int64UpDownCounter(
		"http.server.active_requests",
		metric.WithDescription("Number of active HTTP requests"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		return nil, err
	}

	return &HTTPMetrics{
		requestCounter:  requestCounter,
		requestDuration: requestDuration,
		requestSize:     requestSize,
		responseSize:    responseSize,
		activeRequests:  activeRequests,
	}, nil
}

// RecordRequest records HTTP request metrics
func (m *HTTPMetrics) RecordRequest(ctx context.Context, method, path string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	attrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.route", path),
		attribute.Int("http.status_code", statusCode),
	}

	m.requestCounter.Add(ctx, 1, metric.WithAttributes(attrs...))
	m.requestDuration.Record(ctx, float64(duration.Milliseconds()), metric.WithAttributes(attrs...))
	m.requestSize.Record(ctx, requestSize, metric.WithAttributes(attrs...))
	m.responseSize.Record(ctx, responseSize, metric.WithAttributes(attrs...))
}

// DatabaseMetrics holds database-related metrics
type DatabaseMetrics struct {
	queryCounter  metric.Int64Counter
	queryDuration metric.Float64Histogram
	connections   metric.Int64UpDownCounter
}

// NewDatabaseMetrics creates a new DatabaseMetrics instance
func NewDatabaseMetrics(serviceName string) (*DatabaseMetrics, error) {
	meter := Meter(serviceName)

	queryCounter, err := meter.Int64Counter(
		"db.queries",
		metric.WithDescription("Total number of database queries"),
		metric.WithUnit("{query}"),
	)
	if err != nil {
		return nil, err
	}

	queryDuration, err := meter.Float64Histogram(
		"db.query.duration",
		metric.WithDescription("Database query duration"),
		metric.WithUnit("ms"),
	)
	if err != nil {
		return nil, err
	}

	connections, err := meter.Int64UpDownCounter(
		"db.connections",
		metric.WithDescription("Number of active database connections"),
		metric.WithUnit("{connection}"),
	)
	if err != nil {
		return nil, err
	}

	return &DatabaseMetrics{
		queryCounter:  queryCounter,
		queryDuration: queryDuration,
		connections:   connections,
	}, nil
}

// RecordQuery records database query metrics
func (m *DatabaseMetrics) RecordQuery(ctx context.Context, operation, table string, duration time.Duration, err error) {
	status := "success"
	if err != nil {
		status = "error"
	}

	attrs := []attribute.KeyValue{
		attribute.String("db.operation", operation),
		attribute.String("db.table", table),
		attribute.String("status", status),
	}

	m.queryCounter.Add(ctx, 1, metric.WithAttributes(attrs...))
	m.queryDuration.Record(ctx, float64(duration.Milliseconds()), metric.WithAttributes(attrs...))
}

// BusinessMetrics holds business-related metrics
type BusinessMetrics struct {
	userRegistrations metric.Int64Counter
	orderCreated      metric.Int64Counter
	orderValue        metric.Float64Histogram
	courseEnrollments metric.Int64Counter
	lessonCompleted   metric.Int64Counter
}

// NewBusinessMetrics creates a new BusinessMetrics instance
func NewBusinessMetrics(serviceName string) (*BusinessMetrics, error) {
	meter := Meter(serviceName)

	userRegistrations, err := meter.Int64Counter(
		"business.user.registrations",
		metric.WithDescription("Total number of user registrations"),
		metric.WithUnit("{user}"),
	)
	if err != nil {
		return nil, err
	}

	orderCreated, err := meter.Int64Counter(
		"business.order.created",
		metric.WithDescription("Total number of orders created"),
		metric.WithUnit("{order}"),
	)
	if err != nil {
		return nil, err
	}

	orderValue, err := meter.Float64Histogram(
		"business.order.value",
		metric.WithDescription("Order value distribution"),
		metric.WithUnit("USD"),
	)
	if err != nil {
		return nil, err
	}

	courseEnrollments, err := meter.Int64Counter(
		"business.course.enrollments",
		metric.WithDescription("Total number of course enrollments"),
		metric.WithUnit("{enrollment}"),
	)
	if err != nil {
		return nil, err
	}

	lessonCompleted, err := meter.Int64Counter(
		"business.lesson.completed",
		metric.WithDescription("Total number of lessons completed"),
		metric.WithUnit("{lesson}"),
	)
	if err != nil {
		return nil, err
	}

	return &BusinessMetrics{
		userRegistrations: userRegistrations,
		orderCreated:      orderCreated,
		orderValue:        orderValue,
		courseEnrollments: courseEnrollments,
		lessonCompleted:   lessonCompleted,
	}, nil
}

// RecordUserRegistration records a user registration
func (m *BusinessMetrics) RecordUserRegistration(ctx context.Context, source string) {
	m.userRegistrations.Add(ctx, 1, metric.WithAttributes(
		attribute.String("source", source),
	))
}

// RecordOrder records an order
func (m *BusinessMetrics) RecordOrder(ctx context.Context, value float64, currency string) {
	m.orderCreated.Add(ctx, 1, metric.WithAttributes(
		attribute.String("currency", currency),
	))
	m.orderValue.Record(ctx, value, metric.WithAttributes(
		attribute.String("currency", currency),
	))
}

// RecordCourseEnrollment records a course enrollment
func (m *BusinessMetrics) RecordCourseEnrollment(ctx context.Context, courseID string) {
	m.courseEnrollments.Add(ctx, 1, metric.WithAttributes(
		attribute.String("course.id", courseID),
	))
}

// RecordLessonCompletion records a lesson completion
func (m *BusinessMetrics) RecordLessonCompletion(ctx context.Context, courseID, lessonID string) {
	m.lessonCompleted.Add(ctx, 1, metric.WithAttributes(
		attribute.String("course.id", courseID),
		attribute.String("lesson.id", lessonID),
	))
}

// Example usage:
//
// func main() {
//     // Initialize metrics
//     shutdown, err := otel.InitMetrics("my-service")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer shutdown()
//
//     // Create HTTP metrics
//     httpMetrics, err := otel.NewHTTPMetrics("my-service")
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     // Create business metrics
//     businessMetrics, err := otel.NewBusinessMetrics("my-service")
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     // Use metrics in handlers
//     http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
//         start := time.Now()
//         ctx := r.Context()
//
//         // Business logic
//         registerUser(ctx)
//
//         // Record metrics
//         duration := time.Since(start)
//         httpMetrics.RecordRequest(ctx, r.Method, r.URL.Path, 200, duration, 0, 0)
//         businessMetrics.RecordUserRegistration(ctx, "web")
//     })
// }

