package otel

import (
	"context"
	"os"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggingConfig holds the configuration for structured logging
type LoggingConfig struct {
	Level       string
	Environment string
	ServiceName string
	OutputPaths []string
}

// DefaultLoggingConfig returns a default logging configuration
func DefaultLoggingConfig(serviceName string) *LoggingConfig {
	return &LoggingConfig{
		Level:       getEnv("LOG_LEVEL", "info"),
		Environment: getEnv("ENVIRONMENT", "development"),
		ServiceName: serviceName,
		OutputPaths: []string{"stdout"},
	}
}

// InitLogging initializes structured logging with trace context
func InitLogging(serviceName string) (*zap.Logger, error) {
	return InitLoggingWithConfig(DefaultLoggingConfig(serviceName))
}

// InitLoggingWithConfig initializes structured logging with custom configuration
func InitLoggingWithConfig(config *LoggingConfig) (*zap.Logger, error) {
	// Parse log level
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	// Create encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Create core
	var core zapcore.Core
	if config.Environment == "production" {
		// JSON encoder for production
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		)
	} else {
		// Console encoder for development
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		core = zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderConfig),
			zapcore.AddSync(os.Stdout),
			level,
		)
	}

	// Create logger with initial fields
	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(
			zap.String("service", config.ServiceName),
			zap.String("environment", config.Environment),
		),
	)

	// Set as global logger
	zap.ReplaceGlobals(logger)

	return logger, nil
}

// Logger returns a logger with trace context
func Logger(ctx context.Context) *zap.Logger {
	logger := zap.L()
	
	// Add trace context if available
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		logger = logger.With(
			zap.String("trace_id", span.SpanContext().TraceID().String()),
			zap.String("span_id", span.SpanContext().SpanID().String()),
			zap.Bool("trace_sampled", span.SpanContext().IsSampled()),
		)
	}
	
	return logger
}

// LoggerWithFields returns a logger with additional fields
func LoggerWithFields(ctx context.Context, fields ...zap.Field) *zap.Logger {
	return Logger(ctx).With(fields...)
}

// Debug logs a debug message with trace context
func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	Logger(ctx).Debug(msg, fields...)
}

// Info logs an info message with trace context
func Info(ctx context.Context, msg string, fields ...zap.Field) {
	Logger(ctx).Info(msg, fields...)
}

// Warn logs a warning message with trace context
func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	Logger(ctx).Warn(msg, fields...)
}

// Error logs an error message with trace context
func Error(ctx context.Context, msg string, fields ...zap.Field) {
	Logger(ctx).Error(msg, fields...)
}

// Fatal logs a fatal message with trace context and exits
func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	Logger(ctx).Fatal(msg, fields...)
}

// LoggingMiddleware is an HTTP middleware that logs requests
type LoggingMiddleware struct {
	logger *zap.Logger
}

// NewLoggingMiddleware creates a new logging middleware
func NewLoggingMiddleware(logger *zap.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

// Handler wraps an HTTP handler with logging
func (m *LoggingMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := r.Context()
		
		// Create response writer wrapper
		rw := &loggingResponseWriter{ResponseWriter: w, statusCode: 200}
		
		// Log request
		logger := Logger(ctx).With(
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.String("user_agent", r.UserAgent()),
		)
		
		logger.Info("Request started")
		
		// Call next handler
		next.ServeHTTP(rw, r.WithContext(ctx))
		
		// Log response
		duration := time.Since(start)
		logger = logger.With(
			zap.Int("status_code", rw.statusCode),
			zap.Duration("duration", duration),
			zap.Int64("response_size", rw.bytesWritten),
		)
		
		if rw.statusCode >= 500 {
			logger.Error("Request completed with error")
		} else if rw.statusCode >= 400 {
			logger.Warn("Request completed with client error")
		} else {
			logger.Info("Request completed")
		}
	})
}

// loggingResponseWriter wraps http.ResponseWriter to capture status code and bytes written
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode   int
	bytesWritten int64
}

func (rw *loggingResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *loggingResponseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytesWritten += int64(n)
	return n, err
}

// StructuredError represents a structured error with additional context
type StructuredError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
	Err     error                  `json:"-"`
}

// Error implements the error interface
func (e *StructuredError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

// Unwrap implements the errors.Unwrap interface
func (e *StructuredError) Unwrap() error {
	return e.Err
}

// NewStructuredError creates a new structured error
func NewStructuredError(code, message string, err error) *StructuredError {
	return &StructuredError{
		Code:    code,
		Message: message,
		Err:     err,
		Details: make(map[string]interface{}),
	}
}

// WithDetail adds a detail to the structured error
func (e *StructuredError) WithDetail(key string, value interface{}) *StructuredError {
	e.Details[key] = value
	return e
}

// LogError logs a structured error with trace context
func LogError(ctx context.Context, err error, msg string, fields ...zap.Field) {
	logger := Logger(ctx)
	
	// Add error to fields
	fields = append(fields, zap.Error(err))
	
	// If it's a structured error, add additional fields
	if structErr, ok := err.(*StructuredError); ok {
		fields = append(fields,
			zap.String("error_code", structErr.Code),
			zap.Any("error_details", structErr.Details),
		)
	}
	
	logger.Error(msg, fields...)
	
	// Also record error in span
	RecordError(ctx, err)
}

// Example usage:
//
// func main() {
//     // Initialize logging
//     logger, err := otel.InitLogging("my-service")
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer logger.Sync()
//
//     // Use logger with trace context
//     ctx := context.Background()
//     otel.Info(ctx, "Application started",
//         zap.String("version", "1.0.0"),
//     )
//
//     // Create HTTP handler with logging
//     mux := http.NewServeMux()
//     mux.HandleFunc("/api/users", handleUsers)
//
//     // Wrap with logging middleware
//     middleware := otel.NewLoggingMiddleware(logger)
//     handler := middleware.Handler(mux)
//
//     // Start server
//     http.ListenAndServe(":8080", handler)
// }
//
// func handleUsers(w http.ResponseWriter, r *http.Request) {
//     ctx := r.Context()
//
//     // Log with trace context
//     otel.Info(ctx, "Fetching users",
//         zap.String("user_id", "123"),
//     )
//
//     // Handle error
//     if err := fetchUsers(ctx); err != nil {
//         structErr := otel.NewStructuredError(
//             "USER_FETCH_ERROR",
//             "Failed to fetch users",
//             err,
//         ).WithDetail("user_id", "123")
//
//         otel.LogError(ctx, structErr, "Error fetching users")
//         http.Error(w, structErr.Message, http.StatusInternalServerError)
//         return
//     }
//
//     otel.Info(ctx, "Users fetched successfully")
// }

