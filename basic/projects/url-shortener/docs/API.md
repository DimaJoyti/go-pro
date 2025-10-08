# URL Shortener API Documentation

## Base URL
```
http://localhost:8080
```

## Endpoints

### 1. Shorten URL

Create a shortened URL.

**Endpoint:** `POST /api/shorten`

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "url": "https://example.com/very/long/url/path",
  "custom_code": "mycode"  // optional
}
```

**Response:** `201 Created`
```json
{
  "short_code": "mycode",
  "short_url": "http://localhost:8080/mycode",
  "original_url": "https://example.com/very/long/url/path",
  "created_at": "2024-01-15T10:30:00Z"
}
```

**Error Responses:**

`400 Bad Request` - Invalid URL or code
```json
{
  "error": "Bad Request",
  "message": "invalid URL",
  "code": 400
}
```

`409 Conflict` - Short code already exists
```json
{
  "error": "Conflict",
  "message": "Short code already exists",
  "code": 409
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{
    "url": "https://github.com/DimaJoyti/go-pro",
    "custom_code": "gopro"
  }'
```

---

### 2. Redirect to Original URL

Redirect to the original URL using the short code.

**Endpoint:** `GET /:code`

**Response:** `301 Moved Permanently`

Redirects to the original URL and tracks analytics.

**Error Responses:**

`404 Not Found` - Short code not found
```json
{
  "error": "Not Found",
  "message": "Short URL not found",
  "code": 404
}
```

**Example:**
```bash
curl -L http://localhost:8080/gopro
```

---

### 3. Get URL Statistics

Retrieve statistics for a shortened URL.

**Endpoint:** `GET /api/stats/:code`

**Response:** `200 OK`
```json
{
  "short_code": "gopro",
  "original_url": "https://github.com/DimaJoyti/go-pro",
  "clicks": 42,
  "created_at": "2024-01-15T10:30:00Z",
  "last_accessed": "2024-01-15T15:45:00Z",
  "analytics": {
    "total_clicks": 42,
    "referrers": {
      "https://google.com": 20,
      "https://twitter.com": 15,
      "direct": 7
    },
    "user_agents": {
      "Mozilla/5.0": 30,
      "Chrome/120.0": 12
    },
    "click_history": [
      {
        "timestamp": "2024-01-15T15:45:00Z",
        "referrer": "https://google.com",
        "user_agent": "Mozilla/5.0",
        "ip_address": "192.168.1.1"
      }
    ]
  }
}
```

**Error Responses:**

`404 Not Found` - Short code not found

**Example:**
```bash
curl http://localhost:8080/api/stats/gopro
```

---

### 4. List All URLs

Retrieve all shortened URLs (admin endpoint).

**Endpoint:** `GET /api/urls`

**Response:** `200 OK`
```json
[
  {
    "short_code": "gopro",
    "original_url": "https://github.com/DimaJoyti/go-pro",
    "created_at": "2024-01-15T10:30:00Z",
    "last_accessed": "2024-01-15T15:45:00Z",
    "clicks": 42,
    "analytics": { ... }
  },
  {
    "short_code": "abc123",
    "original_url": "https://example.com",
    "created_at": "2024-01-15T11:00:00Z",
    "clicks": 10,
    "analytics": { ... }
  }
]
```

**Example:**
```bash
curl http://localhost:8080/api/urls
```

---

### 5. Health Check

Check if the service is running.

**Endpoint:** `GET /health`

**Response:** `200 OK`
```json
{
  "status": "healthy",
  "timestamp": "2024-01-15T10:30:00Z",
  "version": "1.0.0"
}
```

**Example:**
```bash
curl http://localhost:8080/health
```

---

## Error Codes

| Code | Description |
|------|-------------|
| 200  | Success |
| 201  | Created |
| 301  | Moved Permanently (redirect) |
| 400  | Bad Request |
| 404  | Not Found |
| 405  | Method Not Allowed |
| 409  | Conflict |
| 500  | Internal Server Error |

---

## Rate Limiting

Currently, there is no rate limiting implemented. In production, consider adding:
- Per-IP rate limiting
- Per-user rate limiting (with authentication)
- Global rate limiting

---

## Authentication

Currently, the API is open and does not require authentication. For production use, consider adding:
- API key authentication
- JWT tokens
- OAuth 2.0

---

## CORS

CORS is enabled for all origins (`*`). In production, configure specific allowed origins.

---

## Examples

### Complete Workflow

```bash
# 1. Shorten a URL
RESPONSE=$(curl -s -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://github.com/DimaJoyti/go-pro"}')

# Extract short code
SHORT_CODE=$(echo $RESPONSE | jq -r '.short_code')

# 2. Access the shortened URL
curl -L http://localhost:8080/$SHORT_CODE

# 3. Check statistics
curl http://localhost:8080/api/stats/$SHORT_CODE

# 4. List all URLs
curl http://localhost:8080/api/urls
```

### Using with JavaScript

```javascript
// Shorten URL
async function shortenURL(url, customCode = null) {
  const response = await fetch('http://localhost:8080/api/shorten', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      url: url,
      custom_code: customCode
    })
  });
  
  return await response.json();
}

// Get statistics
async function getStats(code) {
  const response = await fetch(`http://localhost:8080/api/stats/${code}`);
  return await response.json();
}

// Usage
const result = await shortenURL('https://example.com', 'mycode');
console.log(result.short_url);

const stats = await getStats('mycode');
console.log(`Clicks: ${stats.clicks}`);
```

### Using with Python

```python
import requests

# Shorten URL
def shorten_url(url, custom_code=None):
    payload = {'url': url}
    if custom_code:
        payload['custom_code'] = custom_code
    
    response = requests.post(
        'http://localhost:8080/api/shorten',
        json=payload
    )
    return response.json()

# Get statistics
def get_stats(code):
    response = requests.get(f'http://localhost:8080/api/stats/{code}')
    return response.json()

# Usage
result = shorten_url('https://example.com', 'mycode')
print(f"Short URL: {result['short_url']}")

stats = get_stats('mycode')
print(f"Clicks: {stats['clicks']}")
```

---

## Best Practices

1. **Always validate URLs** before shortening
2. **Use HTTPS** in production
3. **Implement rate limiting** to prevent abuse
4. **Add authentication** for admin endpoints
5. **Monitor analytics** for suspicious activity
6. **Set up logging** for debugging
7. **Use custom codes** for branded links
8. **Implement link expiration** for temporary URLs

---

## Future Enhancements

- [ ] QR code generation
- [ ] Link expiration
- [ ] Password-protected links
- [ ] Custom domains
- [ ] Bulk URL shortening
- [ ] Analytics dashboard
- [ ] Export statistics
- [ ] Link preview
- [ ] UTM parameter support

