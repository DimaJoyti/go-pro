package repository

import (
	"context"

	"go-pro-backend/internal/domain"
)

// CourseRepository defines the interface for course data operations
type CourseRepository interface {
	Create(ctx context.Context, course *domain.Course) error
	GetByID(ctx context.Context, id string) (*domain.Course, error)
	GetAll(ctx context.Context, pagination *domain.PaginationRequest) ([]*domain.Course, int64, error)
	Update(ctx context.Context, course *domain.Course) error
	Delete(ctx context.Context, id string) error
}

// LessonRepository defines the interface for lesson data operations
type LessonRepository interface {
	Create(ctx context.Context, lesson *domain.Lesson) error
	GetByID(ctx context.Context, id string) (*domain.Lesson, error)
	GetByCourseID(ctx context.Context, courseID string, pagination *domain.PaginationRequest) ([]*domain.Lesson, int64, error)
	GetAll(ctx context.Context, pagination *domain.PaginationRequest) ([]*domain.Lesson, int64, error)
	Update(ctx context.Context, lesson *domain.Lesson) error
	Delete(ctx context.Context, id string) error
}

// ExerciseRepository defines the interface for exercise data operations
type ExerciseRepository interface {
	Create(ctx context.Context, exercise *domain.Exercise) error
	GetByID(ctx context.Context, id string) (*domain.Exercise, error)
	GetByLessonID(ctx context.Context, lessonID string, pagination *domain.PaginationRequest) ([]*domain.Exercise, int64, error)
	GetAll(ctx context.Context, pagination *domain.PaginationRequest) ([]*domain.Exercise, int64, error)
	Update(ctx context.Context, exercise *domain.Exercise) error
	Delete(ctx context.Context, id string) error
}

// ProgressRepository defines the interface for progress data operations
type ProgressRepository interface {
	Create(ctx context.Context, progress *domain.Progress) error
	GetByID(ctx context.Context, id string) (*domain.Progress, error)
	GetByUserID(ctx context.Context, userID string, pagination *domain.PaginationRequest) ([]*domain.Progress, int64, error)
	GetByUserAndLesson(ctx context.Context, userID, lessonID string) (*domain.Progress, error)
	Update(ctx context.Context, progress *domain.Progress) error
	Delete(ctx context.Context, id string) error
}

// Repositories aggregates all repository interfaces
type Repositories struct {
	Course   CourseRepository
	Lesson   LessonRepository
	Exercise ExerciseRepository
	Progress ProgressRepository
}
