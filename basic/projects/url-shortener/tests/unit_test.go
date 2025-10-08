package tests

import (
	"context"
	"testing"

	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/domain"
	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/repository"
	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/internal/service"
	"github.com/DimaJoyti/go-pro/basic/projects/url-shortener/pkg/shortener"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShortenRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request domain.ShortenRequest
		wantErr error
	}{
		{
			name: "valid request with http",
			request: domain.ShortenRequest{
				URL: "http://example.com",
			},
			wantErr: nil,
		},
		{
			name: "valid request with https",
			request: domain.ShortenRequest{
				URL: "https://example.com/path",
			},
			wantErr: nil,
		},
		{
			name: "valid request with custom code",
			request: domain.ShortenRequest{
				URL:        "https://example.com",
				CustomCode: "mycode",
			},
			wantErr: nil,
		},
		{
			name: "empty URL",
			request: domain.ShortenRequest{
				URL: "",
			},
			wantErr: domain.ErrEmptyURL,
		},
		{
			name: "invalid URL scheme",
			request: domain.ShortenRequest{
				URL: "ftp://example.com",
			},
			wantErr: domain.ErrInvalidURL,
		},
		{
			name: "custom code too short",
			request: domain.ShortenRequest{
				URL:        "https://example.com",
				CustomCode: "ab",
			},
			wantErr: domain.ErrCodeTooShort,
		},
		{
			name: "custom code too long",
			request: domain.ShortenRequest{
				URL:        "https://example.com",
				CustomCode: "this-is-a-very-long-code-that-exceeds-limit",
			},
			wantErr: domain.ErrCodeTooLong,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGenerator_Generate(t *testing.T) {
	gen := shortener.NewGenerator(6)

	t.Run("generates code of correct length", func(t *testing.T) {
		code, err := gen.Generate()
		require.NoError(t, err)
		assert.Len(t, code, 6)
	})

	t.Run("generates unique codes", func(t *testing.T) {
		codes := make(map[string]bool)
		for i := 0; i < 100; i++ {
			code, err := gen.Generate()
			require.NoError(t, err)
			assert.False(t, codes[code], "generated duplicate code")
			codes[code] = true
		}
	})

	t.Run("generates valid characters only", func(t *testing.T) {
		code, err := gen.Generate()
		require.NoError(t, err)
		assert.True(t, shortener.IsValid(code))
	})
}

func TestMemoryRepository(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewMemoryRepository()

	t.Run("save and find URL", func(t *testing.T) {
		url := domain.NewURL("test123", "https://example.com")

		err := repo.Save(ctx, url)
		require.NoError(t, err)

		found, err := repo.FindByCode(ctx, "test123")
		require.NoError(t, err)
		assert.Equal(t, url.ShortCode, found.ShortCode)
		assert.Equal(t, url.OriginalURL, found.OriginalURL)
	})

	t.Run("duplicate code returns error", func(t *testing.T) {
		repo.Clear()
		url := domain.NewURL("dup123", "https://example.com")

		err := repo.Save(ctx, url)
		require.NoError(t, err)

		err = repo.Save(ctx, url)
		assert.ErrorIs(t, err, domain.ErrCodeExists)
	})

	t.Run("find non-existent URL returns error", func(t *testing.T) {
		repo.Clear()
		_, err := repo.FindByCode(ctx, "nonexistent")
		assert.ErrorIs(t, err, domain.ErrURLNotFound)
	})

	t.Run("update URL", func(t *testing.T) {
		repo.Clear()
		url := domain.NewURL("upd123", "https://example.com")

		err := repo.Save(ctx, url)
		require.NoError(t, err)

		url.IncrementClicks("", "", "")
		err = repo.Update(ctx, url)
		require.NoError(t, err)

		found, err := repo.FindByCode(ctx, "upd123")
		require.NoError(t, err)
		assert.Equal(t, int64(1), found.Clicks)
	})

	t.Run("delete URL", func(t *testing.T) {
		repo.Clear()
		url := domain.NewURL("del123", "https://example.com")

		err := repo.Save(ctx, url)
		require.NoError(t, err)

		err = repo.Delete(ctx, "del123")
		require.NoError(t, err)

		_, err = repo.FindByCode(ctx, "del123")
		assert.ErrorIs(t, err, domain.ErrURLNotFound)
	})

	t.Run("exists check", func(t *testing.T) {
		repo.Clear()
		url := domain.NewURL("exists123", "https://example.com")

		exists, err := repo.Exists(ctx, "exists123")
		require.NoError(t, err)
		assert.False(t, exists)

		err = repo.Save(ctx, url)
		require.NoError(t, err)

		exists, err = repo.Exists(ctx, "exists123")
		require.NoError(t, err)
		assert.True(t, exists)
	})
}

func TestURLService(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewMemoryRepository()
	svc := service.NewURLService(repo, "http://localhost:8080")

	t.Run("shorten URL with auto-generated code", func(t *testing.T) {
		repo.Clear()
		req := &domain.ShortenRequest{
			URL: "https://example.com",
		}

		resp, err := svc.ShortenURL(ctx, req)
		require.NoError(t, err)
		assert.NotEmpty(t, resp.ShortCode)
		assert.Equal(t, req.URL, resp.OriginalURL)
		assert.Contains(t, resp.ShortURL, resp.ShortCode)
	})

	t.Run("shorten URL with custom code", func(t *testing.T) {
		repo.Clear()
		req := &domain.ShortenRequest{
			URL:        "https://example.com",
			CustomCode: "custom",
		}

		resp, err := svc.ShortenURL(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, "custom", resp.ShortCode)
	})

	t.Run("duplicate custom code returns error", func(t *testing.T) {
		repo.Clear()
		req := &domain.ShortenRequest{
			URL:        "https://example.com",
			CustomCode: "duplicate",
		}

		_, err := svc.ShortenURL(ctx, req)
		require.NoError(t, err)

		_, err = svc.ShortenURL(ctx, req)
		assert.ErrorIs(t, err, domain.ErrCodeExists)
	})

	t.Run("get original URL and track analytics", func(t *testing.T) {
		repo.Clear()
		req := &domain.ShortenRequest{
			URL:        "https://example.com",
			CustomCode: "track",
		}

		_, err := svc.ShortenURL(ctx, req)
		require.NoError(t, err)

		originalURL, err := svc.GetOriginalURL(ctx, "track", "https://google.com", "Mozilla/5.0", "127.0.0.1")
		require.NoError(t, err)
		assert.Equal(t, "https://example.com", originalURL)

		stats, err := svc.GetStats(ctx, "track")
		require.NoError(t, err)
		assert.Equal(t, int64(1), stats.Clicks)
	})
}

func TestURL_IncrementClicks(t *testing.T) {
	url := domain.NewURL("test", "https://example.com")

	t.Run("increment clicks updates count", func(t *testing.T) {
		url.IncrementClicks("https://google.com", "Mozilla/5.0", "127.0.0.1")
		assert.Equal(t, int64(1), url.Clicks)
		assert.Equal(t, int64(1), url.Analytics.TotalClicks)
	})

	t.Run("tracks referrers", func(t *testing.T) {
		url.IncrementClicks("https://google.com", "", "")
		assert.Equal(t, int64(2), url.Analytics.Referrers["https://google.com"])
	})

	t.Run("tracks user agents", func(t *testing.T) {
		url.IncrementClicks("", "Mozilla/5.0", "")
		assert.Equal(t, int64(2), url.Analytics.UserAgents["Mozilla/5.0"])
	})

	t.Run("maintains click history", func(t *testing.T) {
		assert.NotEmpty(t, url.Analytics.ClickHistory)
	})
}
