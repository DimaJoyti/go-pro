package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/DimaJoyti/go-pro/basic/projects/blog-engine/internal/models"
)

// PostRepository handles post data access
type PostRepository interface {
	Create(ctx context.Context, post *models.Post) error
	GetByID(ctx context.Context, id int64) (*models.Post, error)
	GetBySlug(ctx context.Context, slug string) (*models.Post, error)
	List(ctx context.Context, filter *models.PostFilter) ([]*models.Post, error)
	Update(ctx context.Context, post *models.Post) error
	Delete(ctx context.Context, id int64) error
	IncrementViewCount(ctx context.Context, id int64) error
}

type postRepository struct {
	db *sql.DB
}

// NewPostRepository creates a new post repository
func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

// Create creates a new post
func (r *postRepository) Create(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO posts (title, slug, content, html_content, excerpt, author_id, category_id, status, published_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(
		ctx, query,
		post.Title, post.Slug, post.Content, post.HTMLContent, post.Excerpt,
		post.AuthorID, post.CategoryID, post.Status, post.PublishedAt,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)

	return err
}

// GetByID retrieves a post by ID
func (r *postRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	query := `
		SELECT p.id, p.title, p.slug, p.content, p.html_content, p.excerpt,
		       p.author_id, p.category_id, p.status, p.view_count,
		       p.published_at, p.created_at, p.updated_at,
		       u.id, u.username, u.email, u.full_name, u.avatar, u.role
		FROM posts p
		LEFT JOIN users u ON p.author_id = u.id
		WHERE p.id = $1
	`

	post := &models.Post{Author: &models.User{}}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID, &post.Title, &post.Slug, &post.Content, &post.HTMLContent, &post.Excerpt,
		&post.AuthorID, &post.CategoryID, &post.Status, &post.ViewCount,
		&post.PublishedAt, &post.CreatedAt, &post.UpdatedAt,
		&post.Author.ID, &post.Author.Username, &post.Author.Email,
		&post.Author.FullName, &post.Author.Avatar, &post.Author.Role,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("post not found")
	}

	return post, err
}

// GetBySlug retrieves a post by slug
func (r *postRepository) GetBySlug(ctx context.Context, slug string) (*models.Post, error) {
	query := `
		SELECT p.id, p.title, p.slug, p.content, p.html_content, p.excerpt,
		       p.author_id, p.category_id, p.status, p.view_count,
		       p.published_at, p.created_at, p.updated_at,
		       u.id, u.username, u.email, u.full_name, u.avatar, u.role
		FROM posts p
		LEFT JOIN users u ON p.author_id = u.id
		WHERE p.slug = $1
	`

	post := &models.Post{Author: &models.User{}}
	err := r.db.QueryRowContext(ctx, query, slug).Scan(
		&post.ID, &post.Title, &post.Slug, &post.Content, &post.HTMLContent, &post.Excerpt,
		&post.AuthorID, &post.CategoryID, &post.Status, &post.ViewCount,
		&post.PublishedAt, &post.CreatedAt, &post.UpdatedAt,
		&post.Author.ID, &post.Author.Username, &post.Author.Email,
		&post.Author.FullName, &post.Author.Avatar, &post.Author.Role,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("post not found")
	}

	return post, err
}

// List retrieves posts with filters
func (r *postRepository) List(ctx context.Context, filter *models.PostFilter) ([]*models.Post, error) {
	query := `
		SELECT p.id, p.title, p.slug, p.excerpt, p.author_id, p.status,
		       p.view_count, p.published_at, p.created_at,
		       u.username, u.full_name, u.avatar
		FROM posts p
		LEFT JOIN users u ON p.author_id = u.id
		WHERE 1=1
	`

	args := []interface{}{}
	argCount := 1

	if filter.AuthorID != nil {
		query += fmt.Sprintf(" AND p.author_id = $%d", argCount)
		args = append(args, *filter.AuthorID)
		argCount++
	}

	if filter.CategoryID != nil {
		query += fmt.Sprintf(" AND p.category_id = $%d", argCount)
		args = append(args, *filter.CategoryID)
		argCount++
	}

	if filter.Status != "" {
		query += fmt.Sprintf(" AND p.status = $%d", argCount)
		args = append(args, filter.Status)
		argCount++
	}

	if filter.Search != "" {
		query += fmt.Sprintf(" AND (p.title ILIKE $%d OR p.content ILIKE $%d)", argCount, argCount)
		args = append(args, "%"+filter.Search+"%")
		argCount++
	}

	// Order by
	if filter.OrderBy != "" {
		query += " ORDER BY " + filter.OrderBy
	} else {
		query += " ORDER BY p.created_at DESC"
	}

	// Pagination
	if filter.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argCount)
		args = append(args, filter.Limit)
		argCount++
	}

	if filter.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argCount)
		args = append(args, filter.Offset)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []*models.Post{}
	for rows.Next() {
		post := &models.Post{Author: &models.User{}}
		err := rows.Scan(
			&post.ID, &post.Title, &post.Slug, &post.Excerpt, &post.AuthorID,
			&post.Status, &post.ViewCount, &post.PublishedAt, &post.CreatedAt,
			&post.Author.Username, &post.Author.FullName, &post.Author.Avatar,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, rows.Err()
}

// Update updates a post
func (r *postRepository) Update(ctx context.Context, post *models.Post) error {
	query := `
		UPDATE posts
		SET title = $1, slug = $2, content = $3, html_content = $4, excerpt = $5,
		    category_id = $6, status = $7, published_at = $8, updated_at = NOW()
		WHERE id = $9
	`

	_, err := r.db.ExecContext(
		ctx, query,
		post.Title, post.Slug, post.Content, post.HTMLContent, post.Excerpt,
		post.CategoryID, post.Status, post.PublishedAt, post.ID,
	)

	return err
}

// Delete deletes a post
func (r *postRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// IncrementViewCount increments the view count of a post
func (r *postRepository) IncrementViewCount(ctx context.Context, id int64) error {
	query := `UPDATE posts SET view_count = view_count + 1 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// GenerateSlug generates a URL-friendly slug from a title
func GenerateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove special characters (basic implementation)
	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, slug)
	return slug
}
