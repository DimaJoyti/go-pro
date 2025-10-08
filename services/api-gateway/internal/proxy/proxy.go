// Package proxy provides HTTP reverse proxy functionality
package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ServiceProxy handles proxying requests to backend services
type ServiceProxy struct {
	client  *http.Client
	timeout time.Duration
}

// NewServiceProxy creates a new service proxy
func NewServiceProxy(timeout time.Duration) *ServiceProxy {
	return &ServiceProxy{
		client: &http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 10,
				IdleConnTimeout:     90 * time.Second,
			},
		},
		timeout: timeout,
	}
}

// ProxyRequest proxies a request to a backend service
func (p *ServiceProxy) ProxyRequest(w http.ResponseWriter, r *http.Request, targetURL string) error {
	// Parse target URL
	target, err := url.Parse(targetURL)
	if err != nil {
		return fmt.Errorf("invalid target URL: %w", err)
	}

	// Create proxy request
	proxyReq, err := p.createProxyRequest(r, target)
	if err != nil {
		return fmt.Errorf("failed to create proxy request: %w", err)
	}

	// Execute request
	resp, err := p.client.Do(proxyReq)
	if err != nil {
		return fmt.Errorf("proxy request failed: %w", err)
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set status code
	w.WriteHeader(resp.StatusCode)

	// Copy response body
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to copy response body: %w", err)
	}

	return nil
}

// createProxyRequest creates a new HTTP request for proxying
func (p *ServiceProxy) createProxyRequest(r *http.Request, target *url.URL) (*http.Request, error) {
	// Create new URL
	proxyURL := &url.URL{
		Scheme:   target.Scheme,
		Host:     target.Host,
		Path:     r.URL.Path,
		RawQuery: r.URL.RawQuery,
	}

	// Create new request
	proxyReq, err := http.NewRequestWithContext(r.Context(), r.Method, proxyURL.String(), r.Body)
	if err != nil {
		return nil, err
	}

	// Copy headers
	for key, values := range r.Header {
		// Skip hop-by-hop headers
		if isHopByHopHeader(key) {
			continue
		}
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	// Add X-Forwarded headers
	if clientIP := getClientIP(r); clientIP != "" {
		proxyReq.Header.Set("X-Forwarded-For", clientIP)
	}
	proxyReq.Header.Set("X-Forwarded-Proto", getScheme(r))
	proxyReq.Header.Set("X-Forwarded-Host", r.Host)

	return proxyReq, nil
}

// Router handles routing requests to appropriate backend services
type Router struct {
	proxy              *ServiceProxy
	userServiceURL     string
	courseServiceURL   string
	progressServiceURL string
}

// NewRouter creates a new router
func NewRouter(userServiceURL, courseServiceURL, progressServiceURL string, timeout time.Duration) *Router {
	return &Router{
		proxy:              NewServiceProxy(timeout),
		userServiceURL:     userServiceURL,
		courseServiceURL:   courseServiceURL,
		progressServiceURL: progressServiceURL,
	}
}

// Route routes a request to the appropriate backend service
func (rt *Router) Route(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	var targetURL string

	switch {
	case strings.HasPrefix(path, "/api/v1/users") || strings.HasPrefix(path, "/api/v1/auth"):
		targetURL = rt.userServiceURL
	case strings.HasPrefix(path, "/api/v1/courses") || strings.HasPrefix(path, "/api/v1/lessons"):
		targetURL = rt.courseServiceURL
	case strings.HasPrefix(path, "/api/v1/progress") || strings.HasPrefix(path, "/api/v1/achievements"):
		targetURL = rt.progressServiceURL
	default:
		http.Error(w, "Service not found", http.StatusNotFound)
		return
	}

	// Proxy the request
	if err := rt.proxy.ProxyRequest(w, r, targetURL); err != nil {
		http.Error(w, fmt.Sprintf("Proxy error: %v", err), http.StatusBadGateway)
		return
	}
}

// CircuitBreaker implements circuit breaker pattern for service calls
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

// Helper functions

func isHopByHopHeader(header string) bool {
	hopByHopHeaders := []string{
		"Connection",
		"Keep-Alive",
		"Proxy-Authenticate",
		"Proxy-Authorization",
		"Te",
		"Trailers",
		"Transfer-Encoding",
		"Upgrade",
	}

	for _, h := range hopByHopHeaders {
		if strings.EqualFold(header, h) {
			return true
		}
	}
	return false
}

func getClientIP(r *http.Request) string {
	// Check X-Forwarded-For header
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		return strings.TrimSpace(ips[0])
	}

	// Check X-Real-IP header
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}

	// Use remote address
	ip := r.RemoteAddr
	if colon := strings.LastIndex(ip, ":"); colon != -1 {
		ip = ip[:colon]
	}
	return ip
}

func getScheme(r *http.Request) string {
	if r.TLS != nil {
		return "https"
	}
	if scheme := r.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	return "http"
}
