package circuitbreaker

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrCircuitOpen     = errors.New("circuit breaker is open")
	ErrTooManyRequests = errors.New("too many requests")
)

// State represents the circuit breaker state
type State int

const (
	StateClosed State = iota
	StateHalfOpen
	StateOpen
)

// String returns the string representation of the state
func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateHalfOpen:
		return "half-open"
	case StateOpen:
		return "open"
	default:
		return "unknown"
	}
}

// Config holds circuit breaker configuration
type Config struct {
	MaxRequests      uint32        // Max requests allowed in half-open state
	Interval         time.Duration // Interval to clear counts
	Timeout          time.Duration // Timeout to switch from open to half-open
	FailureThreshold uint32        // Number of failures to open circuit
	SuccessThreshold uint32        // Number of successes to close circuit from half-open
}

// DefaultConfig returns default circuit breaker configuration
func DefaultConfig() *Config {
	return &Config{
		MaxRequests:      1,
		Interval:         time.Minute,
		Timeout:          time.Minute,
		FailureThreshold: 5,
		SuccessThreshold: 2,
	}
}

// CircuitBreaker implements the circuit breaker pattern
type CircuitBreaker struct {
	config       *Config
	state        State
	failures     uint32
	successes    uint32
	requests     uint32
	lastFailTime time.Time
	mu           sync.RWMutex
}

// NewCircuitBreaker creates a new circuit breaker
func NewCircuitBreaker(config *Config) *CircuitBreaker {
	if config == nil {
		config = DefaultConfig()
	}

	return &CircuitBreaker{
		config: config,
		state:  StateClosed,
	}
}

// Call executes the given function if the circuit breaker allows it
func (cb *CircuitBreaker) Call(fn func() error) error {
	if err := cb.beforeRequest(); err != nil {
		return err
	}

	err := fn()
	cb.afterRequest(err == nil)

	return err
}

// beforeRequest checks if the request is allowed
func (cb *CircuitBreaker) beforeRequest() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	switch cb.state {
	case StateOpen:
		// Check if timeout has passed
		if time.Since(cb.lastFailTime) > cb.config.Timeout {
			cb.setState(StateHalfOpen)
			return nil
		}
		return ErrCircuitOpen

	case StateHalfOpen:
		// Check if max requests exceeded
		if cb.requests >= cb.config.MaxRequests {
			return ErrTooManyRequests
		}
		cb.requests++
		return nil

	default: // StateClosed
		return nil
	}
}

// afterRequest updates the circuit breaker state based on the result
func (cb *CircuitBreaker) afterRequest(success bool) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if success {
		cb.onSuccess()
	} else {
		cb.onFailure()
	}
}

// onSuccess handles successful requests
func (cb *CircuitBreaker) onSuccess() {
	switch cb.state {
	case StateHalfOpen:
		cb.successes++
		if cb.successes >= cb.config.SuccessThreshold {
			cb.setState(StateClosed)
		}

	case StateClosed:
		cb.failures = 0
	}
}

// onFailure handles failed requests
func (cb *CircuitBreaker) onFailure() {
	cb.lastFailTime = time.Now()

	switch cb.state {
	case StateHalfOpen:
		cb.setState(StateOpen)

	case StateClosed:
		cb.failures++
		if cb.failures >= cb.config.FailureThreshold {
			cb.setState(StateOpen)
		}
	}
}

// setState changes the circuit breaker state
func (cb *CircuitBreaker) setState(state State) {
	cb.state = state
	cb.failures = 0
	cb.successes = 0
	cb.requests = 0
}

// GetState returns the current state
func (cb *CircuitBreaker) GetState() State {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}

// Reset resets the circuit breaker to closed state
func (cb *CircuitBreaker) Reset() {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	cb.setState(StateClosed)
}

// Stats returns current statistics
func (cb *CircuitBreaker) Stats() map[string]interface{} {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	return map[string]interface{}{
		"state":     cb.state.String(),
		"failures":  cb.failures,
		"successes": cb.successes,
		"requests":  cb.requests,
	}
}
