# 📚 Lesson Migration Guide

This guide explains how to populate the database with all 20 GO-PRO lessons.

## Overview

The GO-PRO course has been upgraded from 15 to 20 comprehensive lessons, covering everything from Go basics to production-ready systems. This migration adds:

- **Lessons 1-15**: Enhanced with modern practices, security, observability, and real-world examples
- **Lessons 16-20**: New expert-level content covering:
  - Performance Optimization and Profiling
  - Security Best Practices
  - Deployment and DevOps
  - Advanced Design Patterns
  - Building Production Systems

## Database Schema

The lessons are stored in the `gopro.lessons` table with the following structure:

```sql
CREATE TABLE gopro.lessons (
    id UUID PRIMARY KEY,
    course_id UUID REFERENCES gopro.courses(id),
    slug VARCHAR(100) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    content TEXT NOT NULL,
    lesson_order INTEGER NOT NULL,
    estimated_duration_minutes INTEGER,
    status lesson_status DEFAULT 'draft',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(course_id, slug),
    UNIQUE(course_id, lesson_order)
);
```

## Running the Migration

### Option 1: Using psql (Recommended)

```bash
# Navigate to the scripts directory
cd scripts

# Run the seed script
psql -h localhost -U postgres -d gopro -f seed-lessons.sql

# Verify the lessons were created
psql -h localhost -U postgres -d gopro -c "SELECT COUNT(*) FROM gopro.lessons;"
```

### Option 2: Using Docker

If you're running PostgreSQL in Docker:

```bash
# Copy the seed file to the container
docker cp scripts/seed-lessons.sql gopro-postgres:/tmp/

# Execute the seed script
docker exec -it gopro-postgres psql -U postgres -d gopro -f /tmp/seed-lessons.sql

# Verify
docker exec -it gopro-postgres psql -U postgres -d gopro -c "SELECT COUNT(*) FROM gopro.lessons;"
```

### Option 3: Using the Backend API

The backend can also seed lessons programmatically:

```bash
# Start the backend
cd backend
go run cmd/server/main.go

# Use the API to create lessons (requires admin authentication)
curl -X POST http://localhost:8080/api/v1/admin/seed-lessons \
  -H "Authorization: Bearer YOUR_ADMIN_TOKEN"
```

## Verification

After running the migration, verify the lessons:

```sql
-- Check total lessons
SELECT COUNT(*) as total_lessons FROM gopro.lessons;
-- Expected: 20

-- Check lesson distribution by phase
SELECT 
    CASE 
        WHEN lesson_order BETWEEN 1 AND 5 THEN 'Phase 1: Foundations'
        WHEN lesson_order BETWEEN 6 AND 10 THEN 'Phase 2: Intermediate'
        WHEN lesson_order BETWEEN 11 AND 15 THEN 'Phase 3: Advanced'
        WHEN lesson_order BETWEEN 16 AND 20 THEN 'Phase 4: Expert'
    END as phase,
    COUNT(*) as lesson_count
FROM gopro.lessons
GROUP BY phase
ORDER BY MIN(lesson_order);

-- Check total estimated duration
SELECT 
    SUM(estimated_duration_minutes) as total_minutes,
    ROUND(SUM(estimated_duration_minutes) / 60.0, 1) as total_hours
FROM gopro.lessons;
-- Expected: ~7,200 minutes (~120 hours)

-- List all lessons
SELECT 
    lesson_order,
    title,
    estimated_duration_minutes,
    status
FROM gopro.lessons
ORDER BY lesson_order;
```

## Lesson Summary

### Phase 1: Foundations (Lessons 1-5)
- **Total Duration**: ~18 hours
- **Focus**: Go basics, syntax, types, functions, control structures, data structures

### Phase 2: Intermediate (Lessons 6-10)
- **Total Duration**: ~25 hours
- **Focus**: Structs, interfaces, error handling, concurrency, packages

### Phase 3: Advanced (Lessons 11-15)
- **Total Duration**: ~38 hours
- **Focus**: Advanced concurrency, testing, HTTP/REST, databases, microservices

### Phase 4: Expert (Lessons 16-20)
- **Total Duration**: ~39 hours
- **Focus**: Performance, security, deployment, design patterns, production systems

**Total Course Duration**: ~120 hours of comprehensive content

## Updating Existing Lessons

The seed script uses `ON CONFLICT` clauses to safely update existing lessons:

```sql
ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();
```

This means you can run the script multiple times safely. It will:
- Create new lessons if they don't exist
- Update existing lessons with new metadata
- Preserve lesson IDs and relationships

## Rollback

If you need to rollback the migration:

```sql
-- Remove all lessons for the GO-PRO course
DELETE FROM gopro.lessons 
WHERE course_id = 'course-go-pro-2024';

-- Or remove specific lessons (16-20)
DELETE FROM gopro.lessons 
WHERE course_id = 'course-go-pro-2024' 
AND lesson_order BETWEEN 16 AND 20;
```

## Integration with Backend

The backend API automatically serves these lessons through the following endpoints:

```bash
# Get all lessons
GET /api/v1/lessons

# Get lessons for a course
GET /api/v1/courses/{courseId}/lessons

# Get a specific lesson
GET /api/v1/lessons/{lessonId}

# Get lesson with exercises
GET /api/v1/lessons/{lessonId}/exercises
```

## Next Steps

After populating the lessons:

1. **Update Frontend**: The frontend will automatically display all 20 lessons
2. **Create Exercises**: Add exercises for lessons 16-20 in the database
3. **Test API**: Verify all lesson endpoints return correct data
4. **Update Documentation**: Ensure all docs reference the new lesson count

## Troubleshooting

### Issue: "relation gopro.lessons does not exist"

**Solution**: Run the main database initialization script first:
```bash
psql -h localhost -U postgres -d gopro -f scripts/init-db.sql
```

### Issue: "duplicate key value violates unique constraint"

**Solution**: The script handles conflicts automatically. If you see this error, check if you're running an older version of the script.

### Issue: "course_id does not exist"

**Solution**: Ensure the default course exists. The seed script creates it, but you can manually create it:
```sql
INSERT INTO gopro.courses (id, slug, title, description, level, status)
VALUES (
    'course-go-pro-2024',
    'go-programming-mastery',
    'GO-PRO: Complete Go Programming Mastery',
    'Master Go programming from basics to production systems',
    'beginner',
    'published'
);
```

## Support

For issues or questions:
- Check the [main README](../README.md)
- Review the [course syllabus](syllabus.md)
- Examine the [backend documentation](../backend/README.md)

---

**Happy Learning!** 🚀

