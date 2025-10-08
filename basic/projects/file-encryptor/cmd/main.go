package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DimaJoyti/go-pro/basic/projects/file-encryptor/internal/crypto"
	"github.com/DimaJoyti/go-pro/basic/projects/file-encryptor/internal/progress"
	"github.com/DimaJoyti/go-pro/basic/projects/file-encryptor/pkg/encryptor"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "encrypt":
		handleEncrypt()
	case "decrypt":
		handleDecrypt()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func handleEncrypt() {
	inputFile := getFlag("--input", "")
	outputFile := getFlag("--output", "")
	password := getFlag("--password", "")
	deleteOriginal := hasFlag("--delete")

	if inputFile == "" {
		fmt.Println("Error: --input flag is required")
		os.Exit(1)
	}

	// Auto-generate output filename if not provided
	if outputFile == "" {
		outputFile = inputFile + ".enc"
	}

	// Get password if not provided
	if password == "" {
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)
	}

	// Confirm password
	fmt.Print("Confirm password: ")
	var confirmPassword string
	fmt.Scanln(&confirmPassword)

	if password != confirmPassword {
		fmt.Println("Error: Passwords do not match")
		os.Exit(1)
	}

	// Get file info
	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🔐 File Encryption Tool")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("\nEncrypting: %s\n", inputFile)
	fmt.Printf("Size: %s\n\n", progress.FormatBytes(fileInfo.Size()))

	// Create encryptor
	config := encryptor.DefaultConfig()
	config.DeleteOriginal = deleteOriginal
	enc := crypto.NewEncryptor(config)

	// Create progress bar
	bar := progress.NewBar(fileInfo.Size(), "Encrypting")

	// Encrypt file
	startTime := time.Now()
	result, err := enc.EncryptFile(inputFile, outputFile, password, func(current, total int64) {
		bar.Update(current)
	})

	if err != nil {
		fmt.Printf("\n\n❌ Encryption failed: %v\n", err)
		os.Exit(1)
	}

	bar.Finish()
	duration := time.Since(startTime)

	fmt.Println("\n✓ Encryption complete!")
	fmt.Printf("  Input:  %s (%s)\n", result.InputFile, progress.FormatBytes(result.OriginalSize))
	fmt.Printf("  Output: %s (%s)\n", result.OutputFile, progress.FormatBytes(result.EncryptedSize))
	fmt.Printf("  Time:   %s\n", progress.FormatDuration(duration))

	if deleteOriginal {
		fmt.Printf("  ⚠️  Original file deleted\n")
	}
}

func handleDecrypt() {
	inputFile := getFlag("--input", "")
	outputFile := getFlag("--output", "")
	password := getFlag("--password", "")

	if inputFile == "" {
		fmt.Println("Error: --input flag is required")
		os.Exit(1)
	}

	// Auto-generate output filename if not provided
	if outputFile == "" {
		if strings.HasSuffix(inputFile, ".enc") {
			outputFile = strings.TrimSuffix(inputFile, ".enc")
		} else {
			outputFile = inputFile + ".dec"
		}
	}

	// Get password if not provided
	if password == "" {
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)
	}

	// Get file info
	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n🔓 File Decryption Tool")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("\nDecrypting: %s\n", inputFile)
	fmt.Printf("Size: %s\n\n", progress.FormatBytes(fileInfo.Size()))

	// Create encryptor
	config := encryptor.DefaultConfig()
	enc := crypto.NewEncryptor(config)

	// Create progress bar
	bar := progress.NewBar(fileInfo.Size(), "Decrypting")

	// Decrypt file
	startTime := time.Now()
	result, err := enc.DecryptFile(inputFile, outputFile, password, func(current, total int64) {
		bar.Update(current)
	})

	if err != nil {
		fmt.Printf("\n\n❌ Decryption failed: %v\n", err)
		fmt.Println("   Possible reasons:")
		fmt.Println("   - Incorrect password")
		fmt.Println("   - Corrupted file")
		fmt.Println("   - Invalid encrypted file")
		os.Exit(1)
	}

	bar.Finish()
	duration := time.Since(startTime)

	fmt.Println("\n✓ Decryption complete!")
	fmt.Printf("  Input:  %s\n", result.InputFile)
	fmt.Printf("  Output: %s (%s)\n", result.OutputFile, progress.FormatBytes(result.DecryptedSize))
	fmt.Printf("  Time:   %s\n", progress.FormatDuration(duration))
}

func getFlag(flag, defaultValue string) string {
	for i, arg := range os.Args {
		if arg == flag && i+1 < len(os.Args) {
			return os.Args[i+1]
		}
	}
	return defaultValue
}

func hasFlag(flag string) bool {
	for _, arg := range os.Args {
		if arg == flag {
			return true
		}
	}
	return false
}

func printUsage() {
	fmt.Println("🔐 File Encryption Tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  encrypt <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  encrypt     Encrypt a file")
	fmt.Println("  decrypt     Decrypt a file")
	fmt.Println("  help        Show this help message")
	fmt.Println()
	fmt.Println("Encrypt Options:")
	fmt.Println("  --input <file>      Input file to encrypt (required)")
	fmt.Println("  --output <file>     Output encrypted file (default: input.enc)")
	fmt.Println("  --password <pass>   Password for encryption (prompted if not provided)")
	fmt.Println("  --delete            Delete original file after encryption")
	fmt.Println()
	fmt.Println("Decrypt Options:")
	fmt.Println("  --input <file>      Input file to decrypt (required)")
	fmt.Println("  --output <file>     Output decrypted file (default: removes .enc)")
	fmt.Println("  --password <pass>   Password for decryption (prompted if not provided)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  # Encrypt a file")
	fmt.Println("  encrypt encrypt --input document.pdf")
	fmt.Println()
	fmt.Println("  # Encrypt with custom output")
	fmt.Println("  encrypt encrypt --input secret.txt --output secret.encrypted")
	fmt.Println()
	fmt.Println("  # Encrypt and delete original")
	fmt.Println("  encrypt encrypt --input data.zip --delete")
	fmt.Println()
	fmt.Println("  # Decrypt a file")
	fmt.Println("  encrypt decrypt --input document.pdf.enc")
	fmt.Println()
	fmt.Println("Security:")
	fmt.Println("  - Algorithm: AES-256-GCM")
	fmt.Println("  - Key Derivation: PBKDF2 with 100,000 iterations")
	fmt.Println("  - Salt: 32 bytes (random per file)")
	fmt.Println("  - Minimum password length: 8 characters")
	fmt.Println()
	fmt.Println("⚠️  Important:")
	fmt.Println("  - Never lose your password (files cannot be recovered)")
	fmt.Println("  - Use strong passwords (12+ characters recommended)")
	fmt.Println("  - Backup important files before encryption")
}
