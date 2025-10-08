package loadbalancer

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/DimaJoyti/go-pro/basic/projects/service-mesh/pkg/discovery"
)

var (
	ErrNoHealthyInstances = errors.New("no healthy instances available")
)

// LoadBalancer selects service instances for requests
type LoadBalancer interface {
	Select(instances []*discovery.Service) (*discovery.Service, error)
	UpdateStats(instance *discovery.Service, success bool)
}

// RoundRobinBalancer implements round-robin load balancing
type RoundRobinBalancer struct {
	counter uint64
}

// NewRoundRobinBalancer creates a new round-robin balancer
func NewRoundRobinBalancer() *RoundRobinBalancer {
	return &RoundRobinBalancer{}
}

// Select selects the next instance using round-robin
func (b *RoundRobinBalancer) Select(instances []*discovery.Service) (*discovery.Service, error) {
	if len(instances) == 0 {
		return nil, ErrNoHealthyInstances
	}

	// Filter healthy instances
	healthy := make([]*discovery.Service, 0)
	for _, instance := range instances {
		if instance.IsHealthy() {
			healthy = append(healthy, instance)
		}
	}

	if len(healthy) == 0 {
		return nil, ErrNoHealthyInstances
	}

	// Round-robin selection
	index := atomic.AddUint64(&b.counter, 1) % uint64(len(healthy))
	return healthy[index], nil
}

// UpdateStats updates statistics (no-op for round-robin)
func (b *RoundRobinBalancer) UpdateStats(instance *discovery.Service, success bool) {
	// No stats needed for round-robin
}

// LeastConnectionsBalancer implements least connections load balancing
type LeastConnectionsBalancer struct {
	connections map[string]int
	mu          sync.RWMutex
}

// NewLeastConnectionsBalancer creates a new least connections balancer
func NewLeastConnectionsBalancer() *LeastConnectionsBalancer {
	return &LeastConnectionsBalancer{
		connections: make(map[string]int),
	}
}

// Select selects the instance with least connections
func (b *LeastConnectionsBalancer) Select(instances []*discovery.Service) (*discovery.Service, error) {
	if len(instances) == 0 {
		return nil, ErrNoHealthyInstances
	}

	b.mu.RLock()
	defer b.mu.RUnlock()

	var selected *discovery.Service
	minConnections := int(^uint(0) >> 1) // Max int

	for _, instance := range instances {
		if !instance.IsHealthy() {
			continue
		}

		connections := b.connections[instance.ID]
		if connections < minConnections {
			minConnections = connections
			selected = instance
		}
	}

	if selected == nil {
		return nil, ErrNoHealthyInstances
	}

	return selected, nil
}

// UpdateStats updates connection statistics
func (b *LeastConnectionsBalancer) UpdateStats(instance *discovery.Service, success bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if success {
		b.connections[instance.ID]++
	} else {
		if b.connections[instance.ID] > 0 {
			b.connections[instance.ID]--
		}
	}
}

// WeightedBalancer implements weighted load balancing
type WeightedBalancer struct {
	weights map[string]int
	mu      sync.RWMutex
}

// NewWeightedBalancer creates a new weighted balancer
func NewWeightedBalancer() *WeightedBalancer {
	return &WeightedBalancer{
		weights: make(map[string]int),
	}
}

// SetWeight sets the weight for a service instance
func (b *WeightedBalancer) SetWeight(instanceID string, weight int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.weights[instanceID] = weight
}

// Select selects an instance based on weights
func (b *WeightedBalancer) Select(instances []*discovery.Service) (*discovery.Service, error) {
	if len(instances) == 0 {
		return nil, ErrNoHealthyInstances
	}

	b.mu.RLock()
	defer b.mu.RUnlock()

	// Calculate total weight
	totalWeight := 0
	for _, instance := range instances {
		if instance.IsHealthy() {
			weight := b.weights[instance.ID]
			if weight == 0 {
				weight = 1 // Default weight
			}
			totalWeight += weight
		}
	}

	if totalWeight == 0 {
		return nil, ErrNoHealthyInstances
	}

	// Simple weighted selection (can be improved with better algorithm)
	// For now, just return first healthy instance
	for _, instance := range instances {
		if instance.IsHealthy() {
			return instance, nil
		}
	}

	return nil, ErrNoHealthyInstances
}

// UpdateStats updates statistics (no-op for weighted)
func (b *WeightedBalancer) UpdateStats(instance *discovery.Service, success bool) {
	// Could adjust weights based on success/failure
}
