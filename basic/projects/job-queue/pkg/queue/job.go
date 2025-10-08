package queue

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Priority levels for jobs
type Priority int

const (
	LowPriority Priority = iota
	NormalPriority
	HighPriority
	CriticalPriority
)

// JobStatus represents the status of a job
type JobStatus string

const (
	StatusPending   JobStatus = "pending"
	StatusRunning   JobStatus = "running"
	StatusCompleted JobStatus = "completed"
	StatusFailed    JobStatus = "failed"
	StatusRetrying  JobStatus = "retrying"
)

// Job represents a task to be executed
type Job struct {
	ID          string                 `json:"id"`
	Type        string                 `json:"type"`
	Payload     map[string]interface{} `json:"payload"`
	Priority    Priority               `json:"priority"`
	Status      JobStatus              `json:"status"`
	Attempts    int                    `json:"attempts"`
	MaxAttempts int                    `json:"max_attempts"`
	Error       string                 `json:"error,omitempty"`
	CreatedAt   time.Time              `json:"created_at"`
	StartedAt   *time.Time             `json:"started_at,omitempty"`
	CompletedAt *time.Time             `json:"completed_at,omitempty"`
	ScheduledAt *time.Time             `json:"scheduled_at,omitempty"`
}

// NewJob creates a new job
func NewJob(jobType string, payload map[string]interface{}) *Job {
	return &Job{
		ID:          uuid.New().String(),
		Type:        jobType,
		Payload:     payload,
		Priority:    NormalPriority,
		Status:      StatusPending,
		Attempts:    0,
		MaxAttempts: 3,
		CreatedAt:   time.Now(),
	}
}

// WithPriority sets the job priority
func (j *Job) WithPriority(priority Priority) *Job {
	j.Priority = priority
	return j
}

// WithMaxAttempts sets the maximum number of retry attempts
func (j *Job) WithMaxAttempts(maxAttempts int) *Job {
	j.MaxAttempts = maxAttempts
	return j
}

// WithSchedule schedules the job for later execution
func (j *Job) WithSchedule(scheduledAt time.Time) *Job {
	j.ScheduledAt = &scheduledAt
	return j
}

// ToJSON converts job to JSON
func (j *Job) ToJSON() ([]byte, error) {
	return json.Marshal(j)
}

// FromJSON creates a job from JSON
func FromJSON(data []byte) (*Job, error) {
	var job Job
	if err := json.Unmarshal(data, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

// MarkRunning marks the job as running
func (j *Job) MarkRunning() {
	j.Status = StatusRunning
	now := time.Now()
	j.StartedAt = &now
	j.Attempts++
}

// MarkCompleted marks the job as completed
func (j *Job) MarkCompleted() {
	j.Status = StatusCompleted
	now := time.Now()
	j.CompletedAt = &now
}

// MarkFailed marks the job as failed
func (j *Job) MarkFailed(err error) {
	if j.Attempts < j.MaxAttempts {
		j.Status = StatusRetrying
	} else {
		j.Status = StatusFailed
		now := time.Now()
		j.CompletedAt = &now
	}
	j.Error = err.Error()
}

// ShouldRetry checks if the job should be retried
func (j *Job) ShouldRetry() bool {
	return j.Attempts < j.MaxAttempts && j.Status == StatusRetrying
}

// GetRetryDelay calculates the delay before retry (exponential backoff)
func (j *Job) GetRetryDelay() time.Duration {
	// Exponential backoff: 2^attempts seconds
	delay := time.Duration(1<<uint(j.Attempts)) * time.Second
	if delay > 5*time.Minute {
		delay = 5 * time.Minute
	}
	return delay
}
