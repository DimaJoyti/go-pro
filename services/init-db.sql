-- GO-PRO Microservices Database Initialization
-- This script creates separate schemas for each microservice

-- Create schemas
CREATE SCHEMA IF NOT EXISTS users;
CREATE SCHEMA IF NOT EXISTS courses;
CREATE SCHEMA IF NOT EXISTS progress;
CREATE SCHEMA IF NOT EXISTS notifications;

-- ============================================================================
-- USER SERVICE SCHEMA
-- ============================================================================

-- Users table
CREATE TABLE IF NOT EXISTS users.users (
    id VARCHAR(255) PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'student',
    is_active BOOLEAN NOT NULL DEFAULT true,
    email_verified BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_login_at TIMESTAMP
);

-- User profiles table
CREATE TABLE IF NOT EXISTS users.profiles (
    user_id VARCHAR(255) PRIMARY KEY REFERENCES users.users(id) ON DELETE CASCADE,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    bio TEXT,
    avatar_url VARCHAR(500),
    timezone VARCHAR(50),
    language VARCHAR(10) DEFAULT 'en',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Sessions table
CREATE TABLE IF NOT EXISTS users.sessions (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL REFERENCES users.users(id) ON DELETE CASCADE,
    token VARCHAR(500) NOT NULL,
    ip_address VARCHAR(45),
    user_agent TEXT,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for user service
CREATE INDEX IF NOT EXISTS idx_users_email ON users.users(email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users.users(username);
CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON users.sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_token ON users.sessions(token);

-- ============================================================================
-- COURSE SERVICE SCHEMA
-- ============================================================================

-- Courses table
CREATE TABLE IF NOT EXISTS courses.courses (
    id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    difficulty VARCHAR(50),
    duration_hours INTEGER,
    is_published BOOLEAN NOT NULL DEFAULT false,
    created_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Lessons table
CREATE TABLE IF NOT EXISTS courses.lessons (
    id VARCHAR(255) PRIMARY KEY,
    course_id VARCHAR(255) NOT NULL REFERENCES courses.courses(id) ON DELETE CASCADE,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    content TEXT,
    order_index INTEGER NOT NULL,
    duration_minutes INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Exercises table
CREATE TABLE IF NOT EXISTS courses.exercises (
    id VARCHAR(255) PRIMARY KEY,
    lesson_id VARCHAR(255) NOT NULL REFERENCES courses.lessons(id) ON DELETE CASCADE,
    title VARCHAR(500) NOT NULL,
    description TEXT,
    exercise_type VARCHAR(50) NOT NULL,
    difficulty VARCHAR(50),
    points INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Course tags table
CREATE TABLE IF NOT EXISTS courses.tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Course-Tag relationship
CREATE TABLE IF NOT EXISTS courses.course_tags (
    course_id VARCHAR(255) NOT NULL REFERENCES courses.courses(id) ON DELETE CASCADE,
    tag_id INTEGER NOT NULL REFERENCES courses.tags(id) ON DELETE CASCADE,
    PRIMARY KEY (course_id, tag_id)
);

-- Indexes for course service
CREATE INDEX IF NOT EXISTS idx_courses_created_by ON courses.courses(created_by);
CREATE INDEX IF NOT EXISTS idx_courses_is_published ON courses.courses(is_published);
CREATE INDEX IF NOT EXISTS idx_lessons_course_id ON courses.lessons(course_id);
CREATE INDEX IF NOT EXISTS idx_lessons_order ON courses.lessons(course_id, order_index);
CREATE INDEX IF NOT EXISTS idx_exercises_lesson_id ON courses.exercises(lesson_id);

-- ============================================================================
-- PROGRESS SERVICE SCHEMA
-- ============================================================================

-- User progress table
CREATE TABLE IF NOT EXISTS progress.user_progress (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    lesson_id VARCHAR(255) NOT NULL,
    course_id VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'not_started',
    completion_rate DECIMAL(5,2) DEFAULT 0.00,
    score DECIMAL(5,2),
    time_spent_minutes INTEGER DEFAULT 0,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, lesson_id)
);

-- Course enrollments table
CREATE TABLE IF NOT EXISTS progress.enrollments (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    course_id VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    enrolled_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP,
    UNIQUE(user_id, course_id)
);

-- Achievements table
CREATE TABLE IF NOT EXISTS progress.achievements (
    id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    icon_url VARCHAR(500),
    points INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- User achievements table
CREATE TABLE IF NOT EXISTS progress.user_achievements (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    achievement_id VARCHAR(255) NOT NULL REFERENCES progress.achievements(id),
    unlocked_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, achievement_id)
);

-- Indexes for progress service
CREATE INDEX IF NOT EXISTS idx_progress_user_id ON progress.user_progress(user_id);
CREATE INDEX IF NOT EXISTS idx_progress_course_id ON progress.user_progress(course_id);
CREATE INDEX IF NOT EXISTS idx_progress_status ON progress.user_progress(status);
CREATE INDEX IF NOT EXISTS idx_enrollments_user_id ON progress.enrollments(user_id);
CREATE INDEX IF NOT EXISTS idx_enrollments_course_id ON progress.enrollments(course_id);
CREATE INDEX IF NOT EXISTS idx_user_achievements_user_id ON progress.user_achievements(user_id);

-- ============================================================================
-- NOTIFICATION SERVICE SCHEMA
-- ============================================================================

-- Notifications table
CREATE TABLE IF NOT EXISTS notifications.notifications (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL,
    channel VARCHAR(50) NOT NULL,
    subject VARCHAR(500),
    body TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    sent_at TIMESTAMP,
    read_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Notification preferences table
CREATE TABLE IF NOT EXISTS notifications.preferences (
    user_id VARCHAR(255) PRIMARY KEY,
    email_enabled BOOLEAN DEFAULT true,
    push_enabled BOOLEAN DEFAULT true,
    sms_enabled BOOLEAN DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for notification service
CREATE INDEX IF NOT EXISTS idx_notifications_user_id ON notifications.notifications(user_id);
CREATE INDEX IF NOT EXISTS idx_notifications_status ON notifications.notifications(status);
CREATE INDEX IF NOT EXISTS idx_notifications_created_at ON notifications.notifications(created_at);

-- ============================================================================
-- GRANT PERMISSIONS
-- ============================================================================

-- Grant permissions to gopro user
GRANT ALL PRIVILEGES ON SCHEMA users TO gopro;
GRANT ALL PRIVILEGES ON SCHEMA courses TO gopro;
GRANT ALL PRIVILEGES ON SCHEMA progress TO gopro;
GRANT ALL PRIVILEGES ON SCHEMA notifications TO gopro;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA users TO gopro;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA courses TO gopro;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA progress TO gopro;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA notifications TO gopro;

GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA users TO gopro;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA courses TO gopro;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA progress TO gopro;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA notifications TO gopro;

