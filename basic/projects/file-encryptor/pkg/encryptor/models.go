package encryptor

import (
	"errors"
	"time"
)

// Errors
var (
	ErrInvalidPassword  = errors.New("invalid password")
	ErrFileNotFound     = errors.New("file not found")
	ErrInvalidFile      = errors.New("invalid encrypted file")
	ErrDecryptionFailed = errors.New("decryption failed")
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
)

// EncryptionConfig holds encryption configuration
type EncryptionConfig struct {
	KeySize        int   // Key size in bytes (32 for AES-256)
	SaltSize       int   // Salt size in bytes
	NonceSize      int   // Nonce size in bytes
	Iterations     int   // PBKDF2 iterations
	ChunkSize      int64 // Chunk size for large files
	ShowProgress   bool  // Show progress bar
	DeleteOriginal bool  // Delete original file after encryption
}

// DefaultConfig returns default encryption configuration
func DefaultConfig() *EncryptionConfig {
	return &EncryptionConfig{
		KeySize:        32,          // AES-256
		SaltSize:       32,          // 32 bytes salt
		NonceSize:      12,          // GCM standard nonce size
		Iterations:     100000,      // PBKDF2 iterations
		ChunkSize:      1024 * 1024, // 1MB chunks
		ShowProgress:   true,
		DeleteOriginal: false,
	}
}

// FileMetadata contains metadata about encrypted file
type FileMetadata struct {
	OriginalName  string    `json:"original_name"`
	OriginalSize  int64     `json:"original_size"`
	EncryptedAt   time.Time `json:"encrypted_at"`
	Algorithm     string    `json:"algorithm"`
	KeyDerivation string    `json:"key_derivation"`
}

// EncryptionResult contains the result of encryption operation
type EncryptionResult struct {
	InputFile     string        `json:"input_file"`
	OutputFile    string        `json:"output_file"`
	OriginalSize  int64         `json:"original_size"`
	EncryptedSize int64         `json:"encrypted_size"`
	Duration      time.Duration `json:"duration"`
	Success       bool          `json:"success"`
	Error         string        `json:"error,omitempty"`
}

// DecryptionResult contains the result of decryption operation
type DecryptionResult struct {
	InputFile     string        `json:"input_file"`
	OutputFile    string        `json:"output_file"`
	DecryptedSize int64         `json:"decrypted_size"`
	Duration      time.Duration `json:"duration"`
	Success       bool          `json:"success"`
	Error         string        `json:"error,omitempty"`
}

// ProgressCallback is called during encryption/decryption to report progress
type ProgressCallback func(current, total int64)
