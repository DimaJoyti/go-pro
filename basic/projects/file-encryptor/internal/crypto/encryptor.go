package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/DimaJoyti/go-pro/basic/projects/file-encryptor/pkg/encryptor"
)

var (
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
)

// Encryptor handles file encryption and decryption
type Encryptor struct {
	config *encryptor.EncryptionConfig
}

// NewEncryptor creates a new encryptor with the given configuration
func NewEncryptor(config *encryptor.EncryptionConfig) *Encryptor {
	if config == nil {
		config = encryptor.DefaultConfig()
	}
	return &Encryptor{config: config}
}

// EncryptFile encrypts a file with the given password
func (e *Encryptor) EncryptFile(inputPath, outputPath, password string, progress encryptor.ProgressCallback) (*encryptor.EncryptionResult, error) {
	// Validate password
	if err := ValidatePassword(password); err != nil {
		return nil, err
	}

	// Open input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer inputFile.Close()

	// Get file info
	fileInfo, err := inputFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("getting file info: %w", err)
	}
	fileSize := fileInfo.Size()

	// Generate salt
	salt, err := GenerateSalt(e.config.SaltSize)
	if err != nil {
		return nil, fmt.Errorf("generating salt: %w", err)
	}

	// Derive key from password
	key := DeriveKey(password, salt, e.config.KeySize, e.config.Iterations)
	defer SecureWipe(key)

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("creating cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("creating GCM: %w", err)
	}

	// Generate nonce
	nonce, err := GenerateNonce(gcm.NonceSize())
	if err != nil {
		return nil, fmt.Errorf("generating nonce: %w", err)
	}

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return nil, fmt.Errorf("creating output file: %w", err)
	}
	defer outputFile.Close()

	// Write salt and nonce to output file
	if _, err := outputFile.Write(salt); err != nil {
		return nil, fmt.Errorf("writing salt: %w", err)
	}
	if _, err := outputFile.Write(nonce); err != nil {
		return nil, fmt.Errorf("writing nonce: %w", err)
	}

	// Read and encrypt file in chunks
	buffer := make([]byte, e.config.ChunkSize)
	var totalRead int64

	for {
		n, err := inputFile.Read(buffer)
		if err != nil && err != io.EOF {
			return nil, fmt.Errorf("reading input file: %w", err)
		}
		if n == 0 {
			break
		}

		// Encrypt chunk
		ciphertext := gcm.Seal(nil, nonce, buffer[:n], nil)

		// Write encrypted chunk
		if _, err := outputFile.Write(ciphertext); err != nil {
			return nil, fmt.Errorf("writing encrypted data: %w", err)
		}

		totalRead += int64(n)

		// Report progress
		if progress != nil && e.config.ShowProgress {
			progress(totalRead, fileSize)
		}

		if err == io.EOF {
			break
		}
	}

	// Get output file size
	outputInfo, err := outputFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("getting output file info: %w", err)
	}

	// Delete original file if requested
	if e.config.DeleteOriginal {
		if err := os.Remove(inputPath); err != nil {
			return nil, fmt.Errorf("deleting original file: %w", err)
		}
	}

	return &encryptor.EncryptionResult{
		InputFile:     inputPath,
		OutputFile:    outputPath,
		OriginalSize:  fileSize,
		EncryptedSize: outputInfo.Size(),
		Success:       true,
	}, nil
}

// DecryptFile decrypts a file with the given password
func (e *Encryptor) DecryptFile(inputPath, outputPath, password string, progress encryptor.ProgressCallback) (*encryptor.DecryptionResult, error) {
	// Validate password
	if err := ValidatePassword(password); err != nil {
		return nil, err
	}

	// Open input file
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, fmt.Errorf("opening input file: %w", err)
	}
	defer inputFile.Close()

	// Get file info
	fileInfo, err := inputFile.Stat()
	if err != nil {
		return nil, fmt.Errorf("getting file info: %w", err)
	}
	fileSize := fileInfo.Size()

	// Read salt
	salt := make([]byte, e.config.SaltSize)
	if _, err := io.ReadFull(inputFile, salt); err != nil {
		return nil, fmt.Errorf("reading salt: %w", err)
	}

	// Derive key from password
	key := DeriveKey(password, salt, e.config.KeySize, e.config.Iterations)
	defer SecureWipe(key)

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("creating cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("creating GCM: %w", err)
	}

	// Read nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(inputFile, nonce); err != nil {
		return nil, fmt.Errorf("reading nonce: %w", err)
	}

	// Create output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return nil, fmt.Errorf("creating output file: %w", err)
	}
	defer outputFile.Close()

	// Read and decrypt file
	ciphertext, err := io.ReadAll(inputFile)
	if err != nil {
		return nil, fmt.Errorf("reading encrypted data: %w", err)
	}

	// Decrypt
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	// Write decrypted data
	if _, err := outputFile.Write(plaintext); err != nil {
		return nil, fmt.Errorf("writing decrypted data: %w", err)
	}

	// Report progress
	if progress != nil && e.config.ShowProgress {
		progress(fileSize, fileSize)
	}

	return &encryptor.DecryptionResult{
		InputFile:     inputPath,
		OutputFile:    outputPath,
		DecryptedSize: int64(len(plaintext)),
		Success:       true,
	}, nil
}
