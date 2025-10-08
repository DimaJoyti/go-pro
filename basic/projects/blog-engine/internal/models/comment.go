package models

import (
	"errors"
	"time"
)

// Comment represents a blog post comment
type Comment struct {
	ID        int64     `json:"id"`
	PostID    int64     `json:"post_id"`
	UserID    int64     `json:"user_id"`
	User      *User     `json:"user,omitempty"`
	ParentID  *int64    `json:"parent_id,omitempty"` // For nested comments
	Content   string    `json:"content"`
	Status    string    `json:"status"` // pending, approved, spam
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateCommentRequest represents a request to create a comment
type CreateCommentRequest struct {
	PostID   int64  `json:"post_id"`
	ParentID *int64 `json:"parent_id,omitempty"`
	Content  string `json:"content"`
}

// UpdateCommentRequest represents a request to update a comment
type UpdateCommentRequest struct {
	Content *string `json:"content,omitempty"`
	Status  *string `json:"status,omitempty"`
}

// Validate validates the create comment request
func (r *CreateCommentRequest) Validate() error {
	if r.PostID == 0 {
		return errors.New("post_id is required")
	}
	if r.Content == "" {
		return errors.New("content is required")
	}
	if len(r.Content) > 1000 {
		return errors.New("content must be less than 1000 characters")
	}
	return nil
}

// Validate validates the update comment request
func (r *UpdateCommentRequest) Validate() error {
	if r.Content != nil && *r.Content == "" {
		return errors.New("content cannot be empty")
	}
	if r.Content != nil && len(*r.Content) > 1000 {
		return errors.New("content must be less than 1000 characters")
	}
	if r.Status != nil && *r.Status != "pending" && *r.Status != "approved" && *r.Status != "spam" {
		return errors.New("status must be 'pending', 'approved', or 'spam'")
	}
	return nil
}
