# 🔐 File Encryption Tool

A secure command-line file encryption tool built with Go, featuring AES-256 encryption, password-based key derivation, and progress tracking for large files.

## 📋 Project Overview

This project demonstrates building a secure file encryption tool in Go with:
- AES-256-GCM encryption
- PBKDF2 key derivation
- Salt generation for security
- Progress bars for large files
- Directory encryption support
- Secure key management
- Cross-platform compatibility

## 🎯 Learning Objectives

- **Cryptography**: Implement AES encryption in Go
- **Security**: Password-based key derivation (PBKDF2)
- **File I/O**: Handle large files efficiently
- **CLI Development**: Build user-friendly command-line tools
- **Progress Tracking**: Real-time progress indicators
- **Error Handling**: Secure error handling practices

## 🚀 Features

### Core Features
- ✅ AES-256-GCM encryption
- ✅ Password-based encryption
- ✅ Salt generation (random per file)
- ✅ File and directory encryption
- ✅ Progress bars for large files
- ✅ Secure key derivation (PBKDF2)
- ✅ Metadata preservation

### Security Features
- ✅ Strong encryption (AES-256)
- ✅ Authenticated encryption (GCM mode)
- ✅ Random salt per file
- ✅ Secure password handling
- ✅ Memory wiping after use
- ✅ No password storage

## 📁 Project Structure

```
file-encryptor/
├── cmd/
│   └── main.go                 # CLI entry point
├── internal/
│   ├── crypto/
│   │   ├── encryptor.go        # Encryption logic
│   │   └── key.go              # Key derivation
│   └── progress/
│       └── bar.go              # Progress tracking
├── pkg/
│   └── encryptor/
│       └── models.go           # Data models
├── docs/
│   ├── SECURITY.md             # Security documentation
│   └── USAGE.md                # Usage examples
├── go.mod
├── Makefile
└── README.md
```

## 🔧 Installation

```bash
# Build
cd basic/projects/file-encryptor
go build -o encrypt ./cmd/main.go

# Install (optional)
sudo cp encrypt /usr/local/bin/
```

## 📖 Usage

### Encrypt a File

```bash
# Encrypt single file
./encrypt file --input document.pdf --output document.pdf.enc

# With custom password
./encrypt file --input secret.txt --password "MySecurePassword123"

# Encrypt and delete original
./encrypt file --input data.zip --delete
```

### Decrypt a File

```bash
# Decrypt file
./encrypt decrypt --input document.pdf.enc --output document.pdf

# Auto-detect output name
./encrypt decrypt --input secret.txt.enc
```

### Encrypt Directory

```bash
# Encrypt entire directory
./encrypt dir --input ./documents --output ./documents.enc

# Recursive encryption
./encrypt dir --input ./project --recursive
```

## 🎨 Example Output

```
🔐 File Encryption Tool

Encrypting: document.pdf
Size: 2.5 MB

[████████████████████████████████] 100% | 2.5 MB/2.5 MB | 1.2 MB/s

✓ Encryption complete!
  Input:  document.pdf
  Output: document.pdf.enc
  Time:   2.1s
```

## 🔐 Security

### Encryption Algorithm
- **Algorithm**: AES-256-GCM
- **Key Size**: 256 bits
- **Mode**: Galois/Counter Mode (authenticated encryption)

### Key Derivation
- **Function**: PBKDF2
- **Hash**: SHA-256
- **Iterations**: 100,000
- **Salt**: 32 bytes (random per file)

### File Format

```
[Salt: 32 bytes][Nonce: 12 bytes][Ciphertext][Auth Tag: 16 bytes]
```

## 🧪 Testing

```bash
# Run tests
go test ./...

# Test encryption/decryption
./encrypt file --input test.txt
./encrypt decrypt --input test.txt.enc
diff test.txt test.txt.dec
```

## 📚 Examples

### Batch Encryption

```bash
#!/bin/bash
for file in *.pdf; do
    ./encrypt file --input "$file" --delete
done
```

### Secure Backup

```bash
# Encrypt and backup
./encrypt dir --input ~/documents --output backup.enc
scp backup.enc user@server:/backups/
```

## 🎓 Learning Path

1. **Crypto Basics**: Study `internal/crypto/encryptor.go`
2. **Key Derivation**: Review `internal/crypto/key.go`
3. **Progress Tracking**: Examine `internal/progress/bar.go`
4. **CLI Interface**: Look at `cmd/main.go`

## ⚠️ Important Notes

- **Never lose your password** - Files cannot be recovered without it
- **Use strong passwords** - Minimum 12 characters recommended
- **Backup important files** - Before encryption
- **Test decryption** - Verify files can be decrypted
- **Secure deletion** - Use `shred` or similar for originals

## 🚀 Advanced Features

- [ ] Multiple encryption algorithms
- [ ] Public key encryption (RSA)
- [ ] File compression before encryption
- [ ] Encrypted archives
- [ ] Password strength meter
- [ ] Key file support
- [ ] Cloud storage integration

---

**Stay Secure! 🔐**

