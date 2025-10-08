package models

import (
	"errors"
	"time"
)

// User represents a blog user
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Never expose password in JSON
	FullName  string    `json:"full_name"`
	Bio       string    `json:"bio,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Role      string    `json:"role"` // admin, editor, author, subscriber
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// RegisterRequest represents a user registration request
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	User         *User  `json:"user"`
}

// UpdateUserRequest represents a request to update user profile
type UpdateUserRequest struct {
	FullName *string `json:"full_name,omitempty"`
	Bio      *string `json:"bio,omitempty"`
	Avatar   *string `json:"avatar,omitempty"`
}

// ChangePasswordRequest represents a password change request
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

// Validate validates the registration request
func (r *RegisterRequest) Validate() error {
	if r.Username == "" {
		return errors.New("username is required")
	}
	if len(r.Username) < 3 || len(r.Username) > 30 {
		return errors.New("username must be between 3 and 30 characters")
	}
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if r.FullName == "" {
		return errors.New("full name is required")
	}
	return nil
}

// Validate validates the login request
func (r *LoginRequest) Validate() error {
	if r.Email == "" {
		return errors.New("email is required")
	}
	if r.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

// Validate validates the change password request
func (r *ChangePasswordRequest) Validate() error {
	if r.CurrentPassword == "" {
		return errors.New("current password is required")
	}
	if r.NewPassword == "" {
		return errors.New("new password is required")
	}
	if len(r.NewPassword) < 8 {
		return errors.New("new password must be at least 8 characters")
	}
	if r.CurrentPassword == r.NewPassword {
		return errors.New("new password must be different from current password")
	}
	return nil
}

// HasRole checks if user has a specific role
func (u *User) HasRole(role string) bool {
	return u.Role == role
}

// IsAdmin checks if user is an admin
func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

// CanEdit checks if user can edit content
func (u *User) CanEdit() bool {
	return u.Role == "admin" || u.Role == "editor" || u.Role == "author"
}
