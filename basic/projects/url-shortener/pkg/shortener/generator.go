package shortener

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"strings"
)

const (
	// DefaultCodeLength is the default length for generated codes
	DefaultCodeLength = 6

	// Charset for generating short codes
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Generator handles short code generation
type Generator struct {
	length int
}

// NewGenerator creates a new code generator
func NewGenerator(length int) *Generator {
	if length < 3 {
		length = DefaultCodeLength
	}
	return &Generator{
		length: length,
	}
}

// Generate creates a random short code
func (g *Generator) Generate() (string, error) {
	return g.GenerateWithLength(g.length)
}

// GenerateWithLength creates a random short code with specific length
func (g *Generator) GenerateWithLength(length int) (string, error) {
	code := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range code {
		num, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		code[i] = charset[num.Int64()]
	}

	return string(code), nil
}

// GenerateFromURL creates a short code based on the URL
// This provides a deterministic code for the same URL
func (g *Generator) GenerateFromURL(url string) string {
	// Use base64 encoding of URL hash
	encoded := base64.URLEncoding.EncodeToString([]byte(url))

	// Clean up the encoded string
	encoded = strings.ReplaceAll(encoded, "=", "")
	encoded = strings.ReplaceAll(encoded, "+", "")
	encoded = strings.ReplaceAll(encoded, "/", "")

	// Take first N characters
	if len(encoded) > g.length {
		return encoded[:g.length]
	}

	return encoded
}

// GenerateMultiple creates multiple unique short codes
func (g *Generator) GenerateMultiple(count int) ([]string, error) {
	codes := make([]string, count)
	seen := make(map[string]bool)

	for i := 0; i < count; i++ {
		for {
			code, err := g.Generate()
			if err != nil {
				return nil, err
			}

			// Ensure uniqueness in this batch
			if !seen[code] {
				codes[i] = code
				seen[code] = true
				break
			}
		}
	}

	return codes, nil
}

// IsValid checks if a code is valid
func IsValid(code string) bool {
	if len(code) < 3 || len(code) > 20 {
		return false
	}

	for _, char := range code {
		if !strings.ContainsRune(charset+"-_", char) {
			return false
		}
	}

	return true
}

// Sanitize cleans a custom code
func Sanitize(code string) string {
	// Remove invalid characters
	var result strings.Builder
	for _, char := range code {
		if strings.ContainsRune(charset+"-_", char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}
