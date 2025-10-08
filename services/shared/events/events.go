// Package events defines the event structures used across microservices
package events

import (
	"encoding/json"
	"time"
)

// EventType represents the type of event
type EventType string

const (
	// User Events
	EventUserCreated     EventType = "user.created"
	EventUserUpdated     EventType = "user.updated"
	EventUserDeleted     EventType = "user.deleted"
	EventUserLoggedIn    EventType = "user.logged_in"
	EventUserLoggedOut   EventType = "user.logged_out"
	EventPasswordChanged EventType = "user.password_changed"

	// Course Events
	EventCourseCreated EventType = "course.created"
	EventCourseUpdated EventType = "course.updated"
	EventCourseDeleted EventType = "course.deleted"
	EventLessonCreated EventType = "lesson.created"
	EventLessonUpdated EventType = "lesson.updated"
	EventLessonDeleted EventType = "lesson.deleted"

	// Progress Events
	EventProgressUpdated     EventType = "progress.updated"
	EventLessonStarted       EventType = "progress.lesson_started"
	EventLessonCompleted     EventType = "progress.lesson_completed"
	EventCourseStarted       EventType = "progress.course_started"
	EventCourseCompleted     EventType = "progress.course_completed"
	EventAchievementUnlocked EventType = "progress.achievement_unlocked"

	// Notification Events
	EventNotificationSent   EventType = "notification.sent"
	EventNotificationFailed EventType = "notification.failed"
)

// Event is the base event structure
type Event struct {
	ID            string                 `json:"id"`
	Type          EventType              `json:"type"`
	AggregateID   string                 `json:"aggregate_id"`
	AggregateType string                 `json:"aggregate_type"`
	Timestamp     time.Time              `json:"timestamp"`
	Version       int                    `json:"version"`
	Data          map[string]interface{} `json:"data"`
	Metadata      map[string]string      `json:"metadata"`
}

// NewEvent creates a new event
func NewEvent(eventType EventType, aggregateID, aggregateType string, data map[string]interface{}) *Event {
	return &Event{
		ID:            generateEventID(),
		Type:          eventType,
		AggregateID:   aggregateID,
		AggregateType: aggregateType,
		Timestamp:     time.Now().UTC(),
		Version:       1,
		Data:          data,
		Metadata:      make(map[string]string),
	}
}

// ToJSON converts the event to JSON
func (e *Event) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// FromJSON creates an event from JSON
func FromJSON(data []byte) (*Event, error) {
	var event Event
	err := json.Unmarshal(data, &event)
	return &event, err
}

// UserCreatedEvent represents a user creation event
type UserCreatedEvent struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// UserUpdatedEvent represents a user update event
type UserUpdatedEvent struct {
	UserID    string            `json:"user_id"`
	Changes   map[string]string `json:"changes"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// UserDeletedEvent represents a user deletion event
type UserDeletedEvent struct {
	UserID    string    `json:"user_id"`
	DeletedAt time.Time `json:"deleted_at"`
}

// UserLoggedInEvent represents a user login event
type UserLoggedInEvent struct {
	UserID    string    `json:"user_id"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	LoginAt   time.Time `json:"login_at"`
}

// CourseCreatedEvent represents a course creation event
type CourseCreatedEvent struct {
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
}

// CourseUpdatedEvent represents a course update event
type CourseUpdatedEvent struct {
	CourseID  string            `json:"course_id"`
	Changes   map[string]string `json:"changes"`
	UpdatedBy string            `json:"updated_by"`
	UpdatedAt time.Time         `json:"updated_at"`
}

// LessonCreatedEvent represents a lesson creation event
type LessonCreatedEvent struct {
	LessonID    string    `json:"lesson_id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	CreatedAt   time.Time `json:"created_at"`
}

// ProgressUpdatedEvent represents a progress update event
type ProgressUpdatedEvent struct {
	UserID         string    `json:"user_id"`
	LessonID       string    `json:"lesson_id"`
	CourseID       string    `json:"course_id"`
	Status         string    `json:"status"`
	CompletionRate float64   `json:"completion_rate"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// LessonCompletedEvent represents a lesson completion event
type LessonCompletedEvent struct {
	UserID      string    `json:"user_id"`
	LessonID    string    `json:"lesson_id"`
	CourseID    string    `json:"course_id"`
	Score       float64   `json:"score"`
	CompletedAt time.Time `json:"completed_at"`
}

// CourseCompletedEvent represents a course completion event
type CourseCompletedEvent struct {
	UserID       string    `json:"user_id"`
	CourseID     string    `json:"course_id"`
	FinalScore   float64   `json:"final_score"`
	TotalLessons int       `json:"total_lessons"`
	CompletedAt  time.Time `json:"completed_at"`
}

// AchievementUnlockedEvent represents an achievement unlock event
type AchievementUnlockedEvent struct {
	UserID        string    `json:"user_id"`
	AchievementID string    `json:"achievement_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	UnlockedAt    time.Time `json:"unlocked_at"`
}

// NotificationSentEvent represents a notification sent event
type NotificationSentEvent struct {
	NotificationID string    `json:"notification_id"`
	UserID         string    `json:"user_id"`
	Type           string    `json:"type"`
	Channel        string    `json:"channel"`
	SentAt         time.Time `json:"sent_at"`
}

// Helper function to generate event IDs
func generateEventID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// Helper function to generate random strings
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[time.Now().UnixNano()%int64(len(charset))]
	}
	return string(b)
}

// EventPublisher interface for publishing events
type EventPublisher interface {
	Publish(event *Event) error
	PublishBatch(events []*Event) error
}

// EventSubscriber interface for subscribing to events
type EventSubscriber interface {
	Subscribe(eventType EventType, handler EventHandler) error
	Unsubscribe(eventType EventType) error
}

// EventHandler is a function that handles events
type EventHandler func(event *Event) error

// EventStore interface for storing events
type EventStore interface {
	Save(event *Event) error
	GetByAggregateID(aggregateID string) ([]*Event, error)
	GetByType(eventType EventType, limit int) ([]*Event, error)
}
