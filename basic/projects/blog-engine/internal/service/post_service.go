package service

import (
	"context"
	"fmt"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/models"
	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/repository"
	"github.com/russross/blackfriday/v2"
)

// PostService handles post business logic
type PostService interface {
	CreatePost(ctx context.Context, req *models.CreatePostRequest, authorID int64) (*models.Post, error)
	GetPost(ctx context.Context, id int64) (*models.Post, error)
	GetPostBySlug(ctx context.Context, slug string) (*models.Post, error)
	ListPosts(ctx context.Context, filter *models.PostFilter) ([]*models.Post, error)
	UpdatePost(ctx context.Context, id int64, req *models.UpdatePostRequest, userID int64) (*models.Post, error)
	DeletePost(ctx context.Context, id int64, userID int64) error
	PublishPost(ctx context.Context, id int64, userID int64) error
	IncrementViewCount(ctx context.Context, id int64) error
}

type postService struct {
	postRepo repository.PostRepository
	userRepo repository.UserRepository
}

// NewPostService creates a new post service
func NewPostService(postRepo repository.PostRepository, userRepo repository.UserRepository) PostService {
	return &postService{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

// CreatePost creates a new post
func (s *postService) CreatePost(ctx context.Context, req *models.CreatePostRequest, authorID int64) (*models.Post, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Generate slug from title
	slug := repository.GenerateSlug(req.Title)

	// Convert markdown to HTML
	htmlContent := string(blackfriday.Run([]byte(req.Content)))

	// Generate excerpt (first 200 characters)
	excerpt := req.Content
	if len(excerpt) > 200 {
		excerpt = excerpt[:200] + "..."
	}

	post := &models.Post{
		Title:       req.Title,
		Slug:        slug,
		Content:     req.Content,
		HTMLContent: htmlContent,
		Excerpt:     excerpt,
		AuthorID:    authorID,
		CategoryID:  req.CategoryID,
		Status:      req.Status,
	}

	// Set published_at if status is published
	if req.Status == "published" {
		now := time.Now()
		post.PublishedAt = &now
	}

	// Create post
	if err := s.postRepo.Create(ctx, post); err != nil {
		return nil, fmt.Errorf("creating post: %w", err)
	}

	return post, nil
}

// GetPost retrieves a post by ID
func (s *postService) GetPost(ctx context.Context, id int64) (*models.Post, error) {
	return s.postRepo.GetByID(ctx, id)
}

// GetPostBySlug retrieves a post by slug
func (s *postService) GetPostBySlug(ctx context.Context, slug string) (*models.Post, error) {
	return s.postRepo.GetBySlug(ctx, slug)
}

// ListPosts retrieves posts with filters
func (s *postService) ListPosts(ctx context.Context, filter *models.PostFilter) ([]*models.Post, error) {
	return s.postRepo.List(ctx, filter)
}

// UpdatePost updates a post
func (s *postService) UpdatePost(ctx context.Context, id int64, req *models.UpdatePostRequest, userID int64) (*models.Post, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Get existing post
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check authorization (only author or admin can update)
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if post.AuthorID != userID && !user.IsAdmin() {
		return nil, fmt.Errorf("unauthorized to update this post")
	}

	// Update fields
	if req.Title != nil {
		post.Title = *req.Title
		post.Slug = repository.GenerateSlug(*req.Title)
	}

	if req.Content != nil {
		post.Content = *req.Content
		post.HTMLContent = string(blackfriday.Run([]byte(*req.Content)))

		// Update excerpt
		excerpt := *req.Content
		if len(excerpt) > 200 {
			excerpt = excerpt[:200] + "..."
		}
		post.Excerpt = excerpt
	}

	if req.CategoryID != nil {
		post.CategoryID = req.CategoryID
	}

	if req.Status != nil {
		post.Status = *req.Status

		// Set published_at if changing to published
		if *req.Status == "published" && post.PublishedAt == nil {
			now := time.Now()
			post.PublishedAt = &now
		}
	}

	// Update post
	if err := s.postRepo.Update(ctx, post); err != nil {
		return nil, fmt.Errorf("updating post: %w", err)
	}

	return post, nil
}

// DeletePost deletes a post
func (s *postService) DeletePost(ctx context.Context, id int64, userID int64) error {
	// Get existing post
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Check authorization
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if post.AuthorID != userID && !user.IsAdmin() {
		return fmt.Errorf("unauthorized to delete this post")
	}

	return s.postRepo.Delete(ctx, id)
}

// PublishPost publishes a draft post
func (s *postService) PublishPost(ctx context.Context, id int64, userID int64) error {
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Check authorization
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if post.AuthorID != userID && !user.IsAdmin() {
		return fmt.Errorf("unauthorized to publish this post")
	}

	post.Status = "published"
	now := time.Now()
	post.PublishedAt = &now

	return s.postRepo.Update(ctx, post)
}

// IncrementViewCount increments the view count of a post
func (s *postService) IncrementViewCount(ctx context.Context, id int64) error {
	return s.postRepo.IncrementViewCount(ctx, id)
}
