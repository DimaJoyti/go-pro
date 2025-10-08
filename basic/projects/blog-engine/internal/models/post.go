package models

import (
	"errors"
	"time"
)

// Post represents a blog post
type Post struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Content     string     `json:"content"`
	HTMLContent string     `json:"html_content"`
	Excerpt     string     `json:"excerpt"`
	AuthorID    int64      `json:"author_id"`
	Author      *User      `json:"author,omitempty"`
	CategoryID  *int64     `json:"category_id,omitempty"`
	Category    *Category  `json:"category,omitempty"`
	Tags        []Tag      `json:"tags,omitempty"`
	Status      string     `json:"status"` // draft, published, archived
	ViewCount   int64      `json:"view_count"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// Category represents a post category
type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description,omitempty"`
	PostCount   int       `json:"post_count"`
	CreatedAt   time.Time `json:"created_at"`
}

// Tag represents a post tag
type Tag struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	PostCount int       `json:"post_count"`
	CreatedAt time.Time `json:"created_at"`
}

// CreatePostRequest represents a request to create a post
type CreatePostRequest struct {
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	CategoryID *int64  `json:"category_id,omitempty"`
	Tags       []int64 `json:"tags,omitempty"`
	Status     string  `json:"status"`
}

// UpdatePostRequest represents a request to update a post
type UpdatePostRequest struct {
	Title      *string `json:"title,omitempty"`
	Content    *string `json:"content,omitempty"`
	CategoryID *int64  `json:"category_id,omitempty"`
	Tags       []int64 `json:"tags,omitempty"`
	Status     *string `json:"status,omitempty"`
}

// PostFilter represents filters for querying posts
type PostFilter struct {
	AuthorID   *int64
	CategoryID *int64
	TagID      *int64
	Status     string
	Search     string
	Limit      int
	Offset     int
	OrderBy    string
}

// Validate validates the create post request
func (r *CreatePostRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if len(r.Title) > 200 {
		return errors.New("title must be less than 200 characters")
	}
	if r.Content == "" {
		return errors.New("content is required")
	}
	if r.Status != "draft" && r.Status != "published" {
		return errors.New("status must be 'draft' or 'published'")
	}
	return nil
}

// Validate validates the update post request
func (r *UpdatePostRequest) Validate() error {
	if r.Title != nil && *r.Title == "" {
		return errors.New("title cannot be empty")
	}
	if r.Title != nil && len(*r.Title) > 200 {
		return errors.New("title must be less than 200 characters")
	}
	if r.Content != nil && *r.Content == "" {
		return errors.New("content cannot be empty")
	}
	if r.Status != nil && *r.Status != "draft" && *r.Status != "published" && *r.Status != "archived" {
		return errors.New("status must be 'draft', 'published', or 'archived'")
	}
	return nil
}
