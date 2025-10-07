# GO-PRO API Security Documentation

## Overview

The GO-PRO Learning Platform API implements comprehensive security measures following OWASP best practices and industry standards. This document outlines all security features, configurations, and best practices.

## Security Features Implemented

### 1. Authentication & Authorization

#### JWT-Based Authentication
- **Access Tokens**: Short-lived tokens (15 minutes) for API access
- **Refresh Tokens**: Long-lived tokens (7 days) for token renewal
- **Token Validation**: HMAC-SHA256 signatures with secure secret keys
- **Claims Validation**: Issuer, audience, expiration, and custom claims validation

#### Role-Based Access Control (RBAC)
- **User Roles**: `user`, `admin`
- **Role Inheritance**: Admin users inherit all user permissions
- **Endpoint Protection**: Different endpoints require different role levels
- **Context Propagation**: User context available throughout request lifecycle

#### Default Test Accounts
```
Admin: admin@gopro.dev / admin123 (roles: admin, user)
User:  demo@gopro.dev / demo123   (roles: user)
```

### 2. API Security

#### Rate Limiting
- **Algorithm**: Token bucket implementation
- **Default Limits**: 100 requests per minute per IP
- **Burst Capacity**: 10 requests burst allowance
- **Automatic Cleanup**: Expired client entries removed periodically
- **IP Detection**: Real IP extraction from proxy headers

#### CORS (Cross-Origin Resource Sharing)
- **Allowed Origins**: Configurable whitelist (no wildcard in production)
- **Allowed Methods**: GET, POST, PUT, DELETE, OPTIONS
- **Allowed Headers**: Content-Type, Authorization, X-API-Key
- **Credentials**: Supported for authenticated requests
- **Preflight Caching**: 1-hour cache for OPTIONS requests

#### Input Validation & Sanitization
- **Content-Type Validation**: Only application/json accepted for POST/PUT
- **Request Size Limits**: 1MB maximum request body size
- **HTML Escaping**: Automatic XSS prevention through HTML encoding
- **Parameter Validation**: Alphanumeric, email, and UUID format validation
- **Field Length Limits**: Maximum field lengths enforced

### 3. Security Headers

#### Content Security Policy (CSP)
```
default-src 'self';
script-src 'self' 'unsafe-inline';
style-src 'self' 'unsafe-inline';
img-src 'self' data:;
font-src 'self';
connect-src 'self';
```

#### HTTP Strict Transport Security (HSTS)
- **Max Age**: 31536000 seconds (1 year)
- **Include Subdomains**: Enabled
- **Preload**: Enabled

#### Additional Security Headers
- `X-Content-Type-Options: nosniff`
- `X-Frame-Options: DENY`
- `X-XSS-Protection: 1; mode=block`
- `Referrer-Policy: strict-origin-when-cross-origin`
- `X-Permitted-Cross-Domain-Policies: none`

### 4. HTTPS & Transport Security

#### TLS Configuration
- **TLS Version**: TLS 1.2+ required
- **Certificate Management**: File-based certificate configuration
- **HTTP Redirect**: Optional HTTP to HTTPS redirection
- **HSTS**: Enforced for HTTPS connections

### 5. API Key Authentication

#### Admin Endpoints Protection
- **Header**: `X-API-Key` header required
- **Constant Time Comparison**: Prevents timing attacks
- **Combined Authentication**: API key + JWT + admin role required
- **Secure Storage**: API keys should be stored as environment variables

#### Protected Admin Endpoints
- `/api/v1/admin/stats` - Platform statistics
- `/api/v1/admin/users` - User management data

### 6. Password Security

#### Password Requirements
- **Minimum Length**: 8 characters
- **Maximum Length**: 128 characters
- **Character Requirements**: Uppercase, lowercase, and numeric characters required
- **Hashing**: bcrypt with default cost (currently 10)
- **No Storage**: Plain text passwords never stored

#### Account Security
- **Email Validation**: Proper email format validation
- **Account Status**: Active/inactive account management
- **Last Login Tracking**: Login timestamps recorded

### 7. Request Logging & Monitoring

#### Security Logging
- **Request Tracking**: Method, path, status, duration, client IP
- **User Context**: Authenticated user ID logged (when available)
- **Sensitive Data Protection**: Passwords, tokens, and API keys never logged
- **JSON Format**: Structured logging for analysis
- **Timestamp**: UTC timestamps for all log entries

#### Monitored Security Events
- Failed authentication attempts
- Rate limiting violations
- Invalid input validation
- Unauthorized access attempts
- Admin API key usage

## Security Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `JWT_SECRET` | JWT signing secret (min 32 chars) | Generated | Yes |
| `JWT_ISSUER` | JWT token issuer | go-pro-api | No |
| `JWT_AUDIENCE` | JWT token audience | go-pro-users | No |
| `CORS_ALLOWED_ORIGINS` | Comma-separated allowed origins | localhost only | No |
| `RATE_LIMIT_RPM` | Requests per minute per IP | 100 | No |
| `RATE_LIMIT_BURST` | Burst capacity | 10 | No |
| `ADMIN_API_KEY` | Admin API key | Generated | Yes |
| `HTTPS_ENABLED` | Enable HTTPS | false | No |
| `HTTPS_REDIRECT` | Redirect HTTP to HTTPS | false | No |
| `TLS_CERT_FILE` | TLS certificate file path | - | If HTTPS |
| `TLS_KEY_FILE` | TLS private key file path | - | If HTTPS |
| `LOG_LEVEL` | Logging level | INFO | No |
| `LOG_SENSITIVE_DATA` | Log sensitive data (DEV ONLY) | false | No |

### Security Hardening Checklist

#### Production Deployment
- [ ] Set strong `JWT_SECRET` (min 32 characters, cryptographically random)
- [ ] Configure production `ADMIN_API_KEY`
- [ ] Set specific `CORS_ALLOWED_ORIGINS` (no wildcards)
- [ ] Enable HTTPS with valid certificates
- [ ] Set appropriate rate limits for your use case
- [ ] Configure secure reverse proxy (nginx, Cloudflare, etc.)
- [ ] Enable request logging and monitoring
- [ ] Set up log aggregation and analysis
- [ ] Configure automated security scanning
- [ ] Implement database security (when migrating from in-memory)

#### Network Security
- [ ] Use WAF (Web Application Firewall)
- [ ] Implement DDoS protection
- [ ] Configure network segmentation
- [ ] Use VPC/private networks
- [ ] Implement intrusion detection
- [ ] Regular security audits and penetration testing

## API Endpoints Security

### Public Endpoints (No Authentication)
- `GET /` - API Documentation
- `GET /api/v1/health` - Health check with security status

### Authentication Endpoints
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Token refresh

### Protected Endpoints (JWT Required)
- `GET /api/v1/auth/profile` - User profile
- `GET /api/v1/courses` - List courses
- `GET /api/v1/courses/{id}` - Get course details
- `GET /api/v1/courses/{courseId}/lessons` - Get course lessons
- `GET /api/v1/lessons/{id}` - Get lesson details
- `GET /api/v1/exercises/{id}` - Get exercise details
- `POST /api/v1/exercises/{id}/submit` - Submit exercise solution
- `GET /api/v1/progress/{userId}` - Get user progress (own data or admin)
- `POST /api/v1/progress/{userId}/lesson/{lessonId}` - Update progress

### Admin Endpoints (JWT + API Key + Admin Role)
- `GET /api/v1/admin/stats` - Platform statistics
- `GET /api/v1/admin/users` - User statistics

## Usage Examples

### Authentication Flow
```bash
# 1. Login to get tokens
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "demo@gopro.dev", "password": "demo123"}'

# Response includes access_token and refresh_token

# 2. Use access token for protected endpoints
curl -X GET http://localhost:8080/api/v1/courses \
  -H "Authorization: Bearer <access_token>"

# 3. Refresh token when access token expires
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{"refresh_token": "<refresh_token>"}'
```

### Admin API Usage
```bash
# Admin endpoint requires API key + JWT + admin role
curl -X GET http://localhost:8080/api/v1/admin/stats \
  -H "Authorization: Bearer <admin_access_token>" \
  -H "X-API-Key: <admin_api_key>"
```

## Security Testing

### Automated Tests
The security implementation includes comprehensive test coverage:

- **Authentication Tests**: Token generation, validation, refresh
- **Authorization Tests**: Role-based access control
- **Middleware Tests**: Security headers, CORS, rate limiting
- **Input Validation Tests**: XSS prevention, size limits
- **Integration Tests**: End-to-end security flows

### Running Security Tests
```bash
cd backend
go test -v -run TestSecurity
go test -v security_test.go
```

### Manual Security Testing
```bash
# Test rate limiting
for i in {1..101}; do curl -s http://localhost:8080/api/v1/health; done

# Test CORS
curl -H "Origin: https://evil.com" http://localhost:8080/api/v1/health

# Test invalid tokens
curl -H "Authorization: Bearer invalid-token" http://localhost:8080/api/v1/courses

# Test input validation
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "<script>alert(1)</script>", "password": "test"}'
```

## Security Incident Response

### Monitoring & Alerting
- Monitor failed authentication attempts
- Track rate limiting violations
- Alert on admin API key usage
- Monitor unusual access patterns

### Incident Response Steps
1. **Identify**: Log analysis and threat detection
2. **Contain**: Revoke compromised tokens/keys
3. **Investigate**: Full audit trail analysis
4. **Remediate**: Fix vulnerabilities and strengthen security
5. **Learn**: Update security measures and documentation

## Security Updates & Maintenance

### Regular Security Tasks
- [ ] Review and rotate JWT secrets
- [ ] Update API keys
- [ ] Review CORS allowed origins
- [ ] Analyze security logs
- [ ] Update dependencies
- [ ] Security scan and audit
- [ ] Test security measures

### Security Monitoring
- Set up log aggregation (ELK stack, Splunk, etc.)
- Configure security alerts
- Regular vulnerability scanning
- Dependency security audits
- Penetration testing

## Compliance & Standards

### OWASP Compliance
- **A01 Broken Access Control**: Role-based access control implemented
- **A02 Cryptographic Failures**: Strong JWT secrets, bcrypt password hashing
- **A03 Injection**: Input validation and sanitization
- **A04 Insecure Design**: Security by design principles followed
- **A05 Security Misconfiguration**: Secure defaults and configuration
- **A06 Vulnerable Components**: Dependency management and updates
- **A07 Authentication Failures**: Strong authentication and session management
- **A08 Software Data Integrity**: Input validation and secure data handling
- **A09 Security Logging**: Comprehensive security logging implemented
- **A10 Server-Side Request Forgery**: Input validation prevents SSRF

### Additional Standards
- **PCI DSS**: Relevant controls for payment data (if applicable)
- **GDPR**: Data protection and privacy controls
- **SOC 2**: Security controls for service organizations

## Contact & Support

For security questions, vulnerabilities, or incidents:
- Create an issue in the repository
- Follow responsible disclosure practices
- Include detailed information about security issues

**Note**: Never include credentials, tokens, or sensitive data in issue reports.