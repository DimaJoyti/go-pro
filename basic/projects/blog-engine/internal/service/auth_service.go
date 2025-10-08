package service

import (
	"context"
	"fmt"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/auth"
	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/models"
	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/repository"
)

// AuthService handles authentication business logic
type AuthService interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error)
	Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*models.LoginResponse, error)
	ChangePassword(ctx context.Context, userID int64, req *models.ChangePasswordRequest) error
}

type authService struct {
	userRepo   repository.UserRepository
	jwtManager *auth.JWTManager
}

// NewAuthService creates a new auth service
func NewAuthService(userRepo repository.UserRepository, jwtManager *auth.JWTManager) AuthService {
	return &authService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

// Register registers a new user
func (s *authService) Register(ctx context.Context, req *models.RegisterRequest) (*models.User, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Check if email already exists
	existingUser, _ := s.userRepo.GetByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("email already registered")
	}

	// Check if username already exists
	existingUser, _ = s.userRepo.GetByUsername(ctx, req.Username)
	if existingUser != nil {
		return nil, fmt.Errorf("username already taken")
	}

	// Hash password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}

	// Create user
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
		Role:     "subscriber", // Default role
		IsActive: true,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("creating user: %w", err)
	}

	// Clear password before returning
	user.Password = ""

	return user, nil
}

// Login authenticates a user and returns tokens
func (s *authService) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Get user by email
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, fmt.Errorf("account is inactive")
	}

	// Check password
	if !auth.CheckPassword(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate tokens
	token, err := s.jwtManager.GenerateToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("generating token: %w", err)
	}

	refreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("generating refresh token: %w", err)
	}

	// Clear password before returning
	user.Password = ""

	return &models.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		User:         user,
	}, nil
}

// RefreshToken generates new tokens from a refresh token
func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (*models.LoginResponse, error) {
	// Validate refresh token
	claims, err := s.jwtManager.ValidateToken(refreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token")
	}

	// Get user
	user, err := s.userRepo.GetByID(ctx, claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, fmt.Errorf("account is inactive")
	}

	// Generate new tokens
	token, err := s.jwtManager.GenerateToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("generating token: %w", err)
	}

	newRefreshToken, err := s.jwtManager.GenerateRefreshToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("generating refresh token: %w", err)
	}

	// Clear password before returning
	user.Password = ""

	return &models.LoginResponse{
		Token:        token,
		RefreshToken: newRefreshToken,
		User:         user,
	}, nil
}

// ChangePassword changes a user's password
func (s *authService) ChangePassword(ctx context.Context, userID int64, req *models.ChangePasswordRequest) error {
	// Validate request
	if err := req.Validate(); err != nil {
		return err
	}

	// Get user
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	// Check current password
	if !auth.CheckPassword(req.CurrentPassword, user.Password) {
		return fmt.Errorf("current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := auth.HashPassword(req.NewPassword)
	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}

	// Update password
	user.Password = hashedPassword
	user.UpdatedAt = time.Now()

	return s.userRepo.Update(ctx, user)
}
