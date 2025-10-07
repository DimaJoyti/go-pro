package testutil

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go-pro-backend/internal/cache"
	"go-pro-backend/internal/domain"
)

// MockCacheManager is a mock implementation of cache.CacheManager
type MockCacheManager struct {
	mu    sync.RWMutex
	data  map[string]interface{}
	calls map[string]int
}

// NewMockCacheManager creates a new mock cache manager
func NewMockCacheManager() *MockCacheManager {
	return &MockCacheManager{
		data:  make(map[string]interface{}),
		calls: make(map[string]int),
	}
}

func (m *MockCacheManager) Get(ctx context.Context, key string) (interface{}, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.calls["Get"]++

	if val, ok := m.data[key]; ok {
		return val, nil
	}
	return nil, cache.ErrCacheMiss
}

func (m *MockCacheManager) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls["Set"]++

	m.data[key] = value
	return nil
}

func (m *MockCacheManager) Delete(ctx context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls["Delete"]++

	delete(m.data, key)
	return nil
}

func (m *MockCacheManager) Exists(ctx context.Context, key string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.calls["Exists"]++

	_, ok := m.data[key]
	return ok, nil
}

func (m *MockCacheManager) GetCallCount(method string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.calls[method]
}

func (m *MockCacheManager) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = make(map[string]interface{})
	m.calls = make(map[string]int)
}

// Implement remaining CacheManager interface methods
func (m *MockCacheManager) GetSession(ctx context.Context, sessionID string) (*cache.SessionData, error) {
	val, err := m.Get(ctx, "session:"+sessionID)
	if err != nil {
		return nil, err
	}
	return val.(*cache.SessionData), nil
}

func (m *MockCacheManager) SetSession(ctx context.Context, sessionID string, session *cache.SessionData, expiration time.Duration) error {
	return m.Set(ctx, "session:"+sessionID, session, expiration)
}

func (m *MockCacheManager) DeleteSession(ctx context.Context, sessionID string) error {
	return m.Delete(ctx, "session:"+sessionID)
}

func (m *MockCacheManager) RefreshSession(ctx context.Context, sessionID string, expiration time.Duration) error {
	return nil
}

func (m *MockCacheManager) CleanupExpiredSessions(ctx context.Context) error {
	return nil
}

func (m *MockCacheManager) AcquireLock(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return true, nil
}

func (m *MockCacheManager) ReleaseLock(ctx context.Context, key string) error {
	return nil
}

func (m *MockCacheManager) ExtendLock(ctx context.Context, key string, ttl time.Duration) error {
	return nil
}

func (m *MockCacheManager) Allow(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	return true, nil
}

func (m *MockCacheManager) AllowN(ctx context.Context, key string, limit int, window time.Duration, n int) (bool, error) {
	return true, nil
}

func (m *MockCacheManager) Remaining(ctx context.Context, key string, limit int, window time.Duration) (int, error) {
	return limit, nil
}

func (m *MockCacheManager) ResetRateLimit(ctx context.Context, key string) error {
	return m.Delete(ctx, key)
}

func (m *MockCacheManager) Publish(ctx context.Context, channel string, message interface{}) error {
	return nil
}

func (m *MockCacheManager) Subscribe(ctx context.Context, channel string, handler func(string, []byte)) error {
	return nil
}

func (m *MockCacheManager) Unsubscribe(ctx context.Context, channel string) error {
	return nil
}

func (m *MockCacheManager) Health(ctx context.Context) error {
	return nil
}

func (m *MockCacheManager) Close() error {
	return nil
}

// MockCourseRepository is a mock implementation of repository.CourseRepository
type MockCourseRepository struct {
	mu      sync.RWMutex
	courses map[string]*domain.Course
	calls   map[string]int
}

// NewMockCourseRepository creates a new mock course repository
func NewMockCourseRepository() *MockCourseRepository {
	return &MockCourseRepository{
		courses: make(map[string]*domain.Course),
		calls:   make(map[string]int),
	}
}

func (m *MockCourseRepository) Create(ctx context.Context, course *domain.Course) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls["Create"]++

	m.courses[course.ID] = course
	return nil
}

func (m *MockCourseRepository) GetByID(ctx context.Context, id string) (*domain.Course, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.calls["GetByID"]++

	if course, ok := m.courses[id]; ok {
		return course, nil
	}
	return nil, fmt.Errorf("course not found")
}

func (m *MockCourseRepository) GetAll(ctx context.Context, req *domain.PaginationRequest) ([]*domain.Course, int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	m.calls["GetAll"]++

	courses := make([]*domain.Course, 0, len(m.courses))
	for _, course := range m.courses {
		courses = append(courses, course)
	}

	return courses, int64(len(courses)), nil
}

func (m *MockCourseRepository) Update(ctx context.Context, course *domain.Course) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls["Update"]++

	if _, ok := m.courses[course.ID]; !ok {
		return fmt.Errorf("course not found")
	}

	m.courses[course.ID] = course
	return nil
}

func (m *MockCourseRepository) Delete(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls["Delete"]++

	if _, ok := m.courses[id]; !ok {
		return fmt.Errorf("course not found")
	}

	delete(m.courses, id)
	return nil
}

func (m *MockCourseRepository) GetCallCount(method string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.calls[method]
}

func (m *MockCourseRepository) ResetCalls() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = make(map[string]int)
}

func (m *MockCourseRepository) AddCourse(course *domain.Course) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.courses[course.ID] = course
}

// MockMessagingService is a mock implementation of messaging service
type MockMessagingService struct {
	mu        sync.RWMutex
	published []MockMessage
	calls     map[string]int
}

type MockMessage struct {
	Topic   string
	Key     string
	Payload interface{}
}

func NewMockMessagingService() *MockMessagingService {
	return &MockMessagingService{
		published: make([]MockMessage, 0),
		calls:     make(map[string]int),
	}
}

func (m *MockMessagingService) PublishUserEvent(ctx context.Context, eventType string, user *domain.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls["PublishUserEvent"]++

	m.published = append(m.published, MockMessage{
		Topic:   "user-events",
		Key:     user.ID,
		Payload: user,
	})
	return nil
}

func (m *MockMessagingService) PublishCourseEvent(ctx context.Context, eventType string, course *domain.Course) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls["PublishCourseEvent"]++

	m.published = append(m.published, MockMessage{
		Topic:   "course-events",
		Key:     course.ID,
		Payload: course,
	})
	return nil
}

func (m *MockMessagingService) GetPublishedMessages() []MockMessage {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return append([]MockMessage{}, m.published...)
}

func (m *MockMessagingService) GetCallCount(method string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.calls[method]
}

func (m *MockMessagingService) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.published = make([]MockMessage, 0)
	m.calls = make(map[string]int)
}
