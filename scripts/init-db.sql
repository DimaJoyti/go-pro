-- GO-PRO Learning Platform Database Initialization Script
-- This script sets up the initial database schema and seed data

-- Create database extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

-- Create schemas
CREATE SCHEMA IF NOT EXISTS gopro;
CREATE SCHEMA IF NOT EXISTS audit;

-- Set search path
SET search_path TO gopro, public;

-- Create enum types
CREATE TYPE difficulty_level AS ENUM ('beginner', 'intermediate', 'advanced');
CREATE TYPE lesson_status AS ENUM ('draft', 'published', 'archived');
CREATE TYPE user_role AS ENUM ('student', 'instructor', 'admin');

-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    role user_role DEFAULT 'student',
    is_active BOOLEAN DEFAULT TRUE,
    email_verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    last_login_at TIMESTAMP WITH TIME ZONE
);

-- Courses table
CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    slug VARCHAR(100) UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    long_description TEXT,
    difficulty difficulty_level DEFAULT 'beginner',
    estimated_duration_hours INTEGER,
    is_published BOOLEAN DEFAULT FALSE,
    instructor_id UUID REFERENCES users(id),
    thumbnail_url VARCHAR(500),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Lessons table
CREATE TABLE lessons (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
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

-- Exercises table
CREATE TABLE exercises (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    lesson_id UUID REFERENCES lessons(id) ON DELETE CASCADE,
    slug VARCHAR(100) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    instructions TEXT NOT NULL,
    starter_code TEXT,
    solution_code TEXT,
    test_cases JSONB,
    difficulty difficulty_level DEFAULT 'beginner',
    max_attempts INTEGER DEFAULT 0, -- 0 means unlimited
    time_limit_seconds INTEGER,
    exercise_order INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(lesson_id, slug),
    UNIQUE(lesson_id, exercise_order)
);

-- User progress for lessons
CREATE TABLE user_lesson_progress (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    lesson_id UUID REFERENCES lessons(id) ON DELETE CASCADE,
    status VARCHAR(20) DEFAULT 'not_started', -- not_started, in_progress, completed
    started_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    time_spent_seconds INTEGER DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, lesson_id)
);

-- User progress for exercises
CREATE TABLE user_exercise_progress (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    exercise_id UUID REFERENCES exercises(id) ON DELETE CASCADE,
    status VARCHAR(20) DEFAULT 'not_started', -- not_started, in_progress, completed
    attempts INTEGER DEFAULT 0,
    best_score INTEGER DEFAULT 0,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(user_id, exercise_id)
);

-- Exercise submissions
CREATE TABLE exercise_submissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    exercise_id UUID REFERENCES exercises(id) ON DELETE CASCADE,
    code TEXT NOT NULL,
    language VARCHAR(20) DEFAULT 'go',
    status VARCHAR(20) DEFAULT 'pending', -- pending, running, completed, failed
    score INTEGER DEFAULT 0,
    max_score INTEGER DEFAULT 100,
    execution_time_ms INTEGER,
    memory_usage_kb INTEGER,
    test_results JSONB,
    error_message TEXT,
    submitted_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- User course enrollments
CREATE TABLE course_enrollments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    course_id UUID REFERENCES courses(id) ON DELETE CASCADE,
    enrolled_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    completed_at TIMESTAMP WITH TIME ZONE,
    progress_percentage DECIMAL(5,2) DEFAULT 0.00,
    UNIQUE(user_id, course_id)
);

-- Audit log table
CREATE TABLE audit.activity_log (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES gopro.users(id),
    action VARCHAR(100) NOT NULL,
    resource_type VARCHAR(50),
    resource_id UUID,
    old_values JSONB,
    new_values JSONB,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes for better performance
CREATE INDEX idx_courses_slug ON courses(slug);
CREATE INDEX idx_courses_published ON courses(is_published);
CREATE INDEX idx_courses_instructor ON courses(instructor_id);
CREATE INDEX idx_lessons_course ON lessons(course_id);
CREATE INDEX idx_lessons_order ON lessons(course_id, lesson_order);
CREATE INDEX idx_exercises_lesson ON exercises(lesson_id);
CREATE INDEX idx_exercises_order ON exercises(lesson_id, exercise_order);
CREATE INDEX idx_user_lesson_progress_user ON user_lesson_progress(user_id);
CREATE INDEX idx_user_lesson_progress_lesson ON user_lesson_progress(lesson_id);
CREATE INDEX idx_user_exercise_progress_user ON user_exercise_progress(user_id);
CREATE INDEX idx_user_exercise_progress_exercise ON user_exercise_progress(exercise_id);
CREATE INDEX idx_exercise_submissions_user ON exercise_submissions(user_id);
CREATE INDEX idx_exercise_submissions_exercise ON exercise_submissions(exercise_id);
CREATE INDEX idx_exercise_submissions_status ON exercise_submissions(status);
CREATE INDEX idx_course_enrollments_user ON course_enrollments(user_id);
CREATE INDEX idx_course_enrollments_course ON course_enrollments(course_id);

-- Full-text search indexes
CREATE INDEX idx_courses_search ON courses USING gin(to_tsvector('english', title || ' ' || description));
CREATE INDEX idx_lessons_search ON lessons USING gin(to_tsvector('english', title || ' ' || description));

-- Create functions for updating timestamps
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create triggers for automatic timestamp updates
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_courses_updated_at BEFORE UPDATE ON courses
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_lessons_updated_at BEFORE UPDATE ON lessons
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_exercises_updated_at BEFORE UPDATE ON exercises
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_lesson_progress_updated_at BEFORE UPDATE ON user_lesson_progress
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_user_exercise_progress_updated_at BEFORE UPDATE ON user_exercise_progress
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Insert seed data

-- Demo instructor user
INSERT INTO users (id, username, email, password_hash, first_name, last_name, role, email_verified)
VALUES (
    '550e8400-e29b-41d4-a716-446655440000',
    'instructor_demo',
    'instructor@gopro.dev',
    '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPj2L1q9RZF9u', -- password: "demo123"
    'Demo',
    'Instructor',
    'instructor',
    TRUE
);

-- Demo student user
INSERT INTO users (id, username, email, password_hash, first_name, last_name, role, email_verified)
VALUES (
    '550e8400-e29b-41d4-a716-446655440001',
    'student_demo',
    'student@gopro.dev',
    '$2a$12$LQv3c1yqBWVHxkd0LHAkCOYz6TtxMQJqhN8/LewdBPj2L1q9RZF9u', -- password: "demo123"
    'Demo',
    'Student',
    'student',
    TRUE
);

-- GO-PRO course
INSERT INTO courses (id, slug, title, description, long_description, difficulty, estimated_duration_hours, is_published, instructor_id)
VALUES (
    '550e8400-e29b-41d4-a716-446655440002',
    'go-pro',
    'GO-PRO: Complete Go Programming Mastery',
    'Master Go programming from basics to advanced microservices',
    'This comprehensive course takes you from Go beginner to professional developer. Learn syntax, concurrency, web development, microservices, and best practices through hands-on exercises and real-world projects.',
    'beginner',
    40,
    TRUE,
    '550e8400-e29b-41d4-a716-446655440000'
);

-- Course lessons
INSERT INTO lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status)
VALUES
    (
        '550e8400-e29b-41d4-a716-446655440003',
        '550e8400-e29b-41d4-a716-446655440002',
        'go-syntax-basics',
        'Go Syntax and Basic Types',
        'Learn Go''s fundamental syntax and work with basic data types',
        'Introduction to Go programming language, basic types, constants, and iota. Learn about variables, zero values, and type conversions.',
        1,
        60,
        'published'
    ),
    (
        '550e8400-e29b-41d4-a716-446655440004',
        '550e8400-e29b-41d4-a716-446655440002',
        'variables-functions',
        'Variables, Constants, and Functions',
        'Master variable declarations, constants, and function patterns',
        'Variable scope, function declarations, multiple returns, and variadic functions. Learn about named returns and function types.',
        2,
        75,
        'published'
    ),
    (
        '550e8400-e29b-41d4-a716-446655440005',
        '550e8400-e29b-41d4-a716-446655440002',
        'control-structures',
        'Control Structures and Loops',
        'Learn if/else, switch statements, and various loop patterns',
        'Conditional statements, switch cases, for loops, and control flow. Understand range loops and loop control statements.',
        3,
        90,
        'published'
    );

-- Course exercises
INSERT INTO exercises (id, lesson_id, slug, title, description, instructions, starter_code, difficulty, exercise_order, test_cases)
VALUES
    (
        '550e8400-e29b-41d4-a716-446655440006',
        '550e8400-e29b-41d4-a716-446655440003',
        'basic-types-practice',
        'Basic Types Practice',
        'Work with Go''s basic data types and type conversions',
        'Create variables of different types and perform type conversions. Implement functions that work with integers, floats, strings, and booleans.',
        'package main

import "fmt"

// TODO: Implement the following functions

func addNumbers(a, b int) int {
    // Add two integers and return the result
    return 0
}

func convertToString(num int) string {
    // Convert integer to string
    return ""
}

func main() {
    fmt.Println("Basic types practice")
}',
        'beginner',
        1,
        '[{"input": [5, 3], "expected": 8, "description": "Add 5 + 3"}, {"input": [42], "expected": "42", "description": "Convert 42 to string"}]'
    ),
    (
        '550e8400-e29b-41d4-a716-446655440007',
        '550e8400-e29b-41d4-a716-446655440003',
        'constants-iota',
        'Constants and iota',
        'Practice with constants and enumerated values using iota',
        'Create constant declarations using iota for enumerated values. Implement functions that work with constant values.',
        'package main

import "fmt"

// TODO: Define constants using iota
const (
    // Define weekdays starting from Sunday = 0
)

func getDayName(day int) string {
    // Return the day name for the given day number
    return ""
}

func main() {
    fmt.Println("Constants and iota practice")
}',
        'beginner',
        2,
        '[{"input": [0], "expected": "Sunday", "description": "Get Sunday for day 0"}, {"input": [1], "expected": "Monday", "description": "Get Monday for day 1"}]'
    );

-- Course enrollment for demo student
INSERT INTO course_enrollments (user_id, course_id, progress_percentage)
VALUES (
    '550e8400-e29b-41d4-a716-446655440001',
    '550e8400-e29b-41d4-a716-446655440002',
    25.00
);

-- Progress tracking for demo student
INSERT INTO user_lesson_progress (user_id, lesson_id, status, started_at, completed_at, time_spent_seconds)
VALUES
    (
        '550e8400-e29b-41d4-a716-446655440001',
        '550e8400-e29b-41d4-a716-446655440003',
        'completed',
        NOW() - INTERVAL '2 days',
        NOW() - INTERVAL '1 day',
        3600
    ),
    (
        '550e8400-e29b-41d4-a716-446655440001',
        '550e8400-e29b-41d4-a716-446655440004',
        'in_progress',
        NOW() - INTERVAL '1 day',
        NULL,
        1800
    );

-- Exercise progress for demo student
INSERT INTO user_exercise_progress (user_id, exercise_id, status, attempts, best_score, completed_at)
VALUES
    (
        '550e8400-e29b-41d4-a716-446655440001',
        '550e8400-e29b-41d4-a716-446655440006',
        'completed',
        2,
        95,
        NOW() - INTERVAL '1 day'
    ),
    (
        '550e8400-e29b-41d4-a716-446655440001',
        '550e8400-e29b-41d4-a716-446655440007',
        'in_progress',
        1,
        60,
        NULL
    );

-- Create views for easier querying

-- Course overview with instructor information
CREATE VIEW course_overview AS
SELECT
    c.id,
    c.slug,
    c.title,
    c.description,
    c.difficulty,
    c.estimated_duration_hours,
    c.is_published,
    u.username as instructor_username,
    u.first_name || ' ' || u.last_name as instructor_name,
    COUNT(l.id) as lesson_count,
    c.created_at,
    c.updated_at
FROM courses c
LEFT JOIN users u ON c.instructor_id = u.id
LEFT JOIN lessons l ON c.id = l.course_id AND l.status = 'published'
GROUP BY c.id, u.username, u.first_name, u.last_name;

-- User progress summary
CREATE VIEW user_progress_summary AS
SELECT
    u.id as user_id,
    u.username,
    c.id as course_id,
    c.title as course_title,
    ce.enrolled_at,
    ce.progress_percentage,
    COUNT(DISTINCT ulp.lesson_id) as lessons_started,
    COUNT(DISTINCT CASE WHEN ulp.status = 'completed' THEN ulp.lesson_id END) as lessons_completed,
    COUNT(DISTINCT l.id) as total_lessons,
    COUNT(DISTINCT uep.exercise_id) as exercises_attempted,
    COUNT(DISTINCT CASE WHEN uep.status = 'completed' THEN uep.exercise_id END) as exercises_completed
FROM users u
JOIN course_enrollments ce ON u.id = ce.user_id
JOIN courses c ON ce.course_id = c.id
LEFT JOIN lessons l ON c.id = l.course_id AND l.status = 'published'
LEFT JOIN user_lesson_progress ulp ON u.id = ulp.user_id AND l.id = ulp.lesson_id
LEFT JOIN exercises e ON l.id = e.lesson_id
LEFT JOIN user_exercise_progress uep ON u.id = uep.user_id AND e.id = uep.exercise_id
GROUP BY u.id, u.username, c.id, c.title, ce.enrolled_at, ce.progress_percentage;

-- Grant permissions
GRANT ALL PRIVILEGES ON SCHEMA gopro TO gopro_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA gopro TO gopro_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA gopro TO gopro_user;
GRANT ALL PRIVILEGES ON SCHEMA audit TO gopro_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA audit TO gopro_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA audit TO gopro_user;

-- Create stored procedures for common operations

-- Function to enroll user in course
CREATE OR REPLACE FUNCTION enroll_user_in_course(
    p_user_id UUID,
    p_course_id UUID
)
RETURNS BOOLEAN AS $$
BEGIN
    INSERT INTO course_enrollments (user_id, course_id)
    VALUES (p_user_id, p_course_id)
    ON CONFLICT (user_id, course_id) DO NOTHING;

    RETURN FOUND;
END;
$$ LANGUAGE plpgsql;

-- Function to update lesson progress
CREATE OR REPLACE FUNCTION update_lesson_progress(
    p_user_id UUID,
    p_lesson_id UUID,
    p_status VARCHAR(20),
    p_time_spent INTEGER DEFAULT 0
)
RETURNS VOID AS $$
BEGIN
    INSERT INTO user_lesson_progress (user_id, lesson_id, status, started_at, completed_at, time_spent_seconds)
    VALUES (
        p_user_id,
        p_lesson_id,
        p_status,
        CASE WHEN p_status IN ('in_progress', 'completed') THEN COALESCE((SELECT started_at FROM user_lesson_progress WHERE user_id = p_user_id AND lesson_id = p_lesson_id), NOW()) END,
        CASE WHEN p_status = 'completed' THEN NOW() END,
        p_time_spent
    )
    ON CONFLICT (user_id, lesson_id)
    DO UPDATE SET
        status = p_status,
        completed_at = CASE WHEN p_status = 'completed' THEN NOW() ELSE user_lesson_progress.completed_at END,
        time_spent_seconds = user_lesson_progress.time_spent_seconds + p_time_spent,
        updated_at = NOW();
END;
$$ LANGUAGE plpgsql;

COMMIT;