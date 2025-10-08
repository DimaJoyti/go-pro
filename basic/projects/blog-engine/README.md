# 📝 Blog Engine with CMS

A full-featured blog engine with REST API, authentication, Markdown support, and admin dashboard.

## 🎯 Features

- ✅ RESTful API
- ✅ JWT Authentication
- ✅ Markdown support
- ✅ Tag and category system
- ✅ Full-text search
- ✅ Admin dashboard
- ✅ PostgreSQL database
- ✅ Image uploads
- ✅ Comments system
- ✅ SEO optimization

## 🏗️ Architecture

```
blog-engine/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── routes.go
│   ├── auth/
│   │   ├── jwt.go
│   │   └── password.go
│   ├── models/
│   │   ├── post.go
│   │   ├── user.go
│   │   └── comment.go
│   ├── repository/
│   │   ├── post_repo.go
│   │   └── user_repo.go
│   └── service/
│       ├── post_service.go
│       └── auth_service.go
├── web/
│   ├── templates/
│   └── static/
├── migrations/
├── docs/
└── README.md
```

## 📖 API Endpoints

### Posts
- `GET /api/posts` - List all posts
- `GET /api/posts/:id` - Get single post
- `POST /api/posts` - Create post (auth required)
- `PUT /api/posts/:id` - Update post (auth required)
- `DELETE /api/posts/:id` - Delete post (auth required)

### Authentication
- `POST /api/auth/register` - Register user
- `POST /api/auth/login` - Login
- `POST /api/auth/refresh` - Refresh token

### Comments
- `GET /api/posts/:id/comments` - Get comments
- `POST /api/posts/:id/comments` - Add comment

## 🚀 Quick Start

```bash
# Setup database
createdb blogdb
psql blogdb < migrations/001_init.sql

# Set environment variables
export DATABASE_URL="postgres://localhost/blogdb"
export JWT_SECRET="your-secret-key"

# Run
go run cmd/server/main.go
```

## 🎓 Learning Objectives

- REST API design
- JWT authentication
- Database migrations
- Markdown processing
- Full-text search
- File uploads
- Admin interfaces

---

**Status**: Planned | **Difficulty**: Intermediate | **Time**: 8-12 hours

