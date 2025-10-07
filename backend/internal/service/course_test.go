package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go-pro-backend/internal/domain"
	"go-pro-backend/internal/testutil"
	"go-pro-backend/pkg/validator"
)

func TestCourseService_Create(t *testing.T) {
	tests := []struct {
		name        string
		course      *domain.Course
		wantErr     bool
		errContains string
	}{
		{
			name: "valid course",
			course: &domain.Course{
				ID:          "course-1",
				Title:       "Go Programming",
				Description: "Learn Go programming from basics to advanced",
				Lessons:     []string{},
			},
			wantErr: false,
		},
		{
			name: "missing title",
			course: &domain.Course{
				ID:          "course-2",
				Description: "Test description that is long enough",
			},
			wantErr:     true,
			errContains: "validation",
		},
		{
			name: "description too short",
			course: &domain.Course{
				ID:          "course-3",
				Title:       "Test Course",
				Description: "Short",
			},
			wantErr:     true,
			errContains: "validation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := testutil.NewMockCourseRepository()
			mockCache := testutil.NewMockCacheManager()
			mockMessaging := testutil.NewMockMessagingService()
			logger := testutil.NewTestLogger(t)
			v := validator.New()

			config := &Config{
				CourseRepo: mockRepo,
				Cache:      mockCache,
				Messaging:  mockMessaging,
				Logger:     logger,
				Validator:  v,
			}

			service := NewCourseService(config)
			ctx := context.Background()

			// Act
			err := service.Create(ctx, tt.course)

			// Assert
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, mockRepo.GetCallCount("Create"))
				assert.Equal(t, 1, mockCache.GetCallCount("Set"))
				assert.Equal(t, 1, mockMessaging.GetCallCount("PublishCourseEvent"))
			}
		})
	}
}

func TestCourseService_GetByID(t *testing.T) {
	tests := []struct {
		name        string
		courseID    string
		setupMock   func(*testutil.MockCourseRepository, *testutil.MockCacheManager)
		wantErr     bool
		errContains string
		checkCache  bool
	}{
		{
			name:     "course found in cache",
			courseID: "course-1",
			setupMock: func(repo *testutil.MockCourseRepository, cache *testutil.MockCacheManager) {
				course := testutil.CreateTestCourse("course-1", "Go Programming")
				cache.Set(context.Background(), "course:course-1", course, 0)
			},
			wantErr:    false,
			checkCache: true,
		},
		{
			name:     "course found in repository",
			courseID: "course-2",
			setupMock: func(repo *testutil.MockCourseRepository, cache *testutil.MockCacheManager) {
				course := testutil.CreateTestCourse("course-2", "Advanced Go")
				repo.AddCourse(course)
			},
			wantErr:    false,
			checkCache: false,
		},
		{
			name:     "course not found",
			courseID: "course-999",
			setupMock: func(repo *testutil.MockCourseRepository, cache *testutil.MockCacheManager) {
				// No setup - course doesn't exist
			},
			wantErr:     true,
			errContains: "not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := testutil.NewMockCourseRepository()
			mockCache := testutil.NewMockCacheManager()
			mockMessaging := testutil.NewMockMessagingService()
			logger := testutil.NewTestLogger(t)
			v := validator.New()

			tt.setupMock(mockRepo, mockCache)

			config := &Config{
				CourseRepo: mockRepo,
				Cache:      mockCache,
				Messaging:  mockMessaging,
				Logger:     logger,
				Validator:  v,
			}

			service := NewCourseService(config)
			ctx := context.Background()

			// Act
			course, err := service.GetByID(ctx, tt.courseID)

			// Assert
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				assert.Nil(t, course)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, course)
				assert.Equal(t, tt.courseID, course.ID)

				if tt.checkCache {
					// Should hit cache, not repository
					assert.Equal(t, 0, mockRepo.GetCallCount("GetByID"))
				} else {
					// Should hit repository and cache result
					assert.Equal(t, 1, mockRepo.GetCallCount("GetByID"))
					assert.Equal(t, 1, mockCache.GetCallCount("Set"))
				}
			}
		})
	}
}

func TestCourseService_Update(t *testing.T) {
	tests := []struct {
		name        string
		course      *domain.Course
		setupMock   func(*testutil.MockCourseRepository)
		wantErr     bool
		errContains string
	}{
		{
			name: "successful update",
			course: &domain.Course{
				ID:          "course-1",
				Title:       "Updated Go Programming",
				Description: "Updated description for the course",
				Lessons:     []string{},
			},
			setupMock: func(repo *testutil.MockCourseRepository) {
				course := testutil.CreateTestCourse("course-1", "Go Programming")
				repo.AddCourse(course)
			},
			wantErr: false,
		},
		{
			name: "course not found",
			course: &domain.Course{
				ID:          "course-999",
				Title:       "Non-existent Course",
				Description: "Description for non-existent course",
				Lessons:     []string{},
			},
			setupMock: func(repo *testutil.MockCourseRepository) {
				// No setup - course doesn't exist
			},
			wantErr:     true,
			errContains: "not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := testutil.NewMockCourseRepository()
			mockCache := testutil.NewMockCacheManager()
			mockMessaging := testutil.NewMockMessagingService()
			logger := testutil.NewTestLogger(t)
			v := validator.New()

			tt.setupMock(mockRepo)

			config := &Config{
				CourseRepo: mockRepo,
				Cache:      mockCache,
				Messaging:  mockMessaging,
				Logger:     logger,
				Validator:  v,
			}

			service := NewCourseService(config)
			ctx := context.Background()

			// Act
			err := service.Update(ctx, tt.course)

			// Assert
			if tt.wantErr {
				require.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				require.NoError(t, err)
				assert.Equal(t, 1, mockRepo.GetCallCount("Update"))
				assert.Equal(t, 1, mockCache.GetCallCount("Delete")) // Cache invalidation
				assert.Equal(t, 1, mockMessaging.GetCallCount("PublishCourseEvent"))
			}
		})
	}
}

func TestCourseService_Delete(t *testing.T) {
	// Arrange
	mockRepo := testutil.NewMockCourseRepository()
	mockCache := testutil.NewMockCacheManager()
	mockMessaging := testutil.NewMockMessagingService()
	logger := testutil.NewTestLogger(t)
	v := validator.New()

	course := testutil.CreateTestCourse("course-1", "Go Programming")
	mockRepo.AddCourse(course)

	config := &Config{
		CourseRepo: mockRepo,
		Cache:      mockCache,
		Messaging:  mockMessaging,
		Logger:     logger,
		Validator:  v,
	}

	service := NewCourseService(config)
	ctx := context.Background()

	// Act
	err := service.Delete(ctx, "course-1")

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 1, mockRepo.GetCallCount("Delete"))
	assert.Equal(t, 1, mockCache.GetCallCount("Delete"))
	assert.Equal(t, 1, mockMessaging.GetCallCount("PublishCourseEvent"))

	// Verify course is deleted
	_, err = mockRepo.GetByID(ctx, "course-1")
	assert.Error(t, err)
}

// Benchmark tests
func BenchmarkCourseService_Create(b *testing.B) {
	mockRepo := testutil.NewMockCourseRepository()
	mockCache := testutil.NewMockCacheManager()
	mockMessaging := testutil.NewMockMessagingService()
	logger := testutil.NewTestLogger(&testing.T{})
	v := validator.New()

	config := &Config{
		CourseRepo: mockRepo,
		Cache:      mockCache,
		Messaging:  mockMessaging,
		Logger:     logger,
		Validator:  v,
	}

	service := NewCourseService(config)
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		course := testutil.CreateTestCourse("course-bench", "Benchmark Course")
		_ = service.Create(ctx, course)
	}
}
