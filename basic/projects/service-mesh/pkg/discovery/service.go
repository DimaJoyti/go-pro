package discovery

import (
	"fmt"
	"time"
)

// Service represents a service in the mesh
type Service struct {
	ID       string            `json:"id"`
	Name     string            `json:"name"`
	Address  string            `json:"address"`
	Port     int               `json:"port"`
	Tags     []string          `json:"tags"`
	Metadata map[string]string `json:"metadata"`
	Health   HealthStatus      `json:"health"`
}

// HealthStatus represents the health status of a service
type HealthStatus string

const (
	HealthPassing  HealthStatus = "passing"
	HealthWarning  HealthStatus = "warning"
	HealthCritical HealthStatus = "critical"
)

// ServiceRegistry manages service registration and discovery
type ServiceRegistry interface {
	Register(service *Service) error
	Deregister(serviceID string) error
	Discover(serviceName string) ([]*Service, error)
	HealthCheck(serviceID string) (HealthStatus, error)
	Watch(serviceName string) (<-chan []*Service, error)
}

// ServiceInstance represents a running instance of a service
type ServiceInstance struct {
	Service     *Service
	LastSeen    time.Time
	Connections int
	Load        float64
}

// NewService creates a new service
func NewService(name, address string, port int) *Service {
	return &Service{
		ID:       fmt.Sprintf("%s-%d", name, time.Now().UnixNano()),
		Name:     name,
		Address:  address,
		Port:     port,
		Tags:     []string{},
		Metadata: make(map[string]string),
		Health:   HealthPassing,
	}
}

// WithTags adds tags to the service
func (s *Service) WithTags(tags ...string) *Service {
	s.Tags = append(s.Tags, tags...)
	return s
}

// WithMetadata adds metadata to the service
func (s *Service) WithMetadata(key, value string) *Service {
	s.Metadata[key] = value
	return s
}

// GetEndpoint returns the full endpoint URL
func (s *Service) GetEndpoint() string {
	return fmt.Sprintf("%s:%d", s.Address, s.Port)
}

// IsHealthy checks if the service is healthy
func (s *Service) IsHealthy() bool {
	return s.Health == HealthPassing
}
