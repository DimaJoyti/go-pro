// Package client provides HTTP client utilities for service-to-service communication
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPClient is a wrapper around http.Client with additional features
type HTTPClient struct {
	client  *http.Client
	baseURL string
	headers map[string]string
}

// NewHTTPClient creates a new HTTP client
func NewHTTPClient(baseURL string, timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		baseURL: baseURL,
		headers: make(map[string]string),
	}
}

// SetHeader sets a default header for all requests
func (c *HTTPClient) SetHeader(key, value string) {
	c.headers[key] = value
}

// Get performs a GET request
func (c *HTTPClient) Get(ctx context.Context, path string) (*Response, error) {
	return c.do(ctx, http.MethodGet, path, nil)
}

// Post performs a POST request
func (c *HTTPClient) Post(ctx context.Context, path string, body interface{}) (*Response, error) {
	return c.do(ctx, http.MethodPost, path, body)
}

// Put performs a PUT request
func (c *HTTPClient) Put(ctx context.Context, path string, body interface{}) (*Response, error) {
	return c.do(ctx, http.MethodPut, path, body)
}

// Delete performs a DELETE request
func (c *HTTPClient) Delete(ctx context.Context, path string) (*Response, error) {
	return c.do(ctx, http.MethodDelete, path, nil)
}

// do performs the actual HTTP request
func (c *HTTPClient) do(ctx context.Context, method, path string, body interface{}) (*Response, error) {
	url := c.baseURL + path

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set default headers
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	// Set content type for POST/PUT requests
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Perform the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return &Response{Response: resp}, nil
}

// Response wraps http.Response with additional methods
type Response struct {
	*http.Response
}

// JSON decodes the response body as JSON
func (r *Response) JSON(v interface{}) error {
	defer r.Body.Close()

	if r.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", r.StatusCode)
	}

	return json.NewDecoder(r.Body).Decode(v)
}

// Bytes returns the response body as bytes
func (r *Response) Bytes() ([]byte, error) {
	defer r.Body.Close()
	return io.ReadAll(r.Body)
}

// String returns the response body as string
func (r *Response) String() (string, error) {
	b, err := r.Bytes()
	return string(b), err
}

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	maxFailures  int
	resetTimeout time.Duration
	failures     int
	lastFailTime time.Time
	state        CircuitState
}

// CircuitState represents the state of the circuit breaker
type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(maxFailures int, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures:  maxFailures,
		resetTimeout: resetTimeout,
		state:        StateClosed,
	}
}

// Call executes a function with circuit breaker protection
func (cb *CircuitBreaker) Call(fn func() error) error {
	if cb.state == StateOpen {
		if time.Since(cb.lastFailTime) > cb.resetTimeout {
			cb.state = StateHalfOpen
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}

	err := fn()
	if err != nil {
		cb.recordFailure()
		return err
	}

	cb.recordSuccess()
	return nil
}

func (cb *CircuitBreaker) recordFailure() {
	cb.failures++
	cb.lastFailTime = time.Now()

	if cb.failures >= cb.maxFailures {
		cb.state = StateOpen
	}
}

func (cb *CircuitBreaker) recordSuccess() {
	cb.failures = 0
	cb.state = StateClosed
}

// RetryConfig configures retry behavior
type RetryConfig struct {
	MaxRetries     int
	InitialBackoff time.Duration
	MaxBackoff     time.Duration
	Multiplier     float64
}

// DefaultRetryConfig returns default retry configuration
func DefaultRetryConfig() *RetryConfig {
	return &RetryConfig{
		MaxRetries:     3,
		InitialBackoff: 100 * time.Millisecond,
		MaxBackoff:     5 * time.Second,
		Multiplier:     2.0,
	}
}

// RetryableHTTPClient wraps HTTPClient with retry logic
type RetryableHTTPClient struct {
	client *HTTPClient
	config *RetryConfig
}

// NewRetryableHTTPClient creates a new retryable HTTP client
func NewRetryableHTTPClient(baseURL string, timeout time.Duration, config *RetryConfig) *RetryableHTTPClient {
	if config == nil {
		config = DefaultRetryConfig()
	}

	return &RetryableHTTPClient{
		client: NewHTTPClient(baseURL, timeout),
		config: config,
	}
}

// Get performs a GET request with retry logic
func (rc *RetryableHTTPClient) Get(ctx context.Context, path string) (*Response, error) {
	return rc.doWithRetry(ctx, func() (*Response, error) {
		return rc.client.Get(ctx, path)
	})
}

// Post performs a POST request with retry logic
func (rc *RetryableHTTPClient) Post(ctx context.Context, path string, body interface{}) (*Response, error) {
	return rc.doWithRetry(ctx, func() (*Response, error) {
		return rc.client.Post(ctx, path, body)
	})
}

// doWithRetry executes a function with retry logic
func (rc *RetryableHTTPClient) doWithRetry(ctx context.Context, fn func() (*Response, error)) (*Response, error) {
	var lastErr error
	backoff := rc.config.InitialBackoff

	for attempt := 0; attempt <= rc.config.MaxRetries; attempt++ {
		resp, err := fn()
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}

		lastErr = err
		if attempt < rc.config.MaxRetries {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(backoff):
				backoff = time.Duration(float64(backoff) * rc.config.Multiplier)
				if backoff > rc.config.MaxBackoff {
					backoff = rc.config.MaxBackoff
				}
			}
		}
	}

	return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}

