-- Seed Lessons for GO-PRO Course
-- This script populates the lessons table with all 20 lessons

-- First, ensure we have a default course
INSERT INTO gopro.courses (id, slug, title, description, level, status, created_at, updated_at)
VALUES (
    'course-go-pro-2024',
    'go-programming-mastery',
    'GO-PRO: Complete Go Programming Mastery',
    'Master Go programming from basics to production-ready microservices. Learn idiomatic Go, concurrency patterns, testing, and build real-world applications.',
    'beginner',
    'published',
    NOW(),
    NOW()
) ON CONFLICT (slug) DO NOTHING;

-- Phase 1: Foundations (Lessons 1-5)

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-01-basics',
    'course-go-pro-2024',
    'go-syntax-and-basic-types',
    'Go Syntax and Basic Types',
    'Learn the fundamental syntax of Go and work with basic data types including integers, floats, strings, booleans, and constants.',
    'Introduction to Go programming language, basic types, variable declarations, constants, and the iota identifier.',
    1,
    180,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-02-variables-functions',
    'course-go-pro-2024',
    'variables-constants-and-functions',
    'Variables, Constants, and Functions',
    'Master variable declarations, understand scope and visibility, and create functions with various parameter and return patterns.',
    'Variable declarations (var, :=), scope rules, function definitions, multiple return values, named returns, and variadic functions.',
    2,
    240,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-03-control-structures',
    'course-go-pro-2024',
    'control-structures-and-loops',
    'Control Structures and Loops',
    'Learn control flow with if/else, switch statements, and various loop patterns. Master defer statements and error handling basics.',
    'If/else statements, switch statements, for loops (all variants), break and continue, defer statements.',
    3,
    180,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-04-data-structures',
    'course-go-pro-2024',
    'arrays-slices-and-maps',
    'Arrays, Slices, and Maps',
    'Work with Go''s built-in data structures: arrays, slices, and maps. Understand their internals and best practices.',
    'Arrays declaration and usage, slices creation and manipulation, maps operations, range loops, memory considerations.',
    4,
    300,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-05-pointers-memory',
    'course-go-pro-2024',
    'pointers-and-memory-management',
    'Pointers and Memory Management',
    'Understand pointers, memory allocation, and Go''s garbage collection. Learn when to use pointers vs values.',
    'Pointer basics and syntax, address and dereference operators, memory allocation with new and make, garbage collection basics.',
    5,
    240,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

-- Phase 2: Intermediate (Lessons 6-10)

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-06-structs-methods',
    'course-go-pro-2024',
    'structs-and-methods',
    'Structs and Methods',
    'Define custom types with structs and attach methods. Understand value vs pointer receivers and method sets.',
    'Struct definition and initialization, anonymous structs and fields, method definitions and receivers, value vs pointer receivers.',
    6,
    300,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-07-interfaces',
    'course-go-pro-2024',
    'interfaces-and-polymorphism',
    'Interfaces and Polymorphism',
    'Master Go''s interface system for polymorphism and abstraction. Learn interface composition and common patterns.',
    'Interface definition and implementation, empty interface and type assertions, type switches, interface composition.',
    7,
    360,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-08-error-handling',
    'course-go-pro-2024',
    'error-handling-patterns',
    'Error Handling Patterns',
    'Implement robust error handling with custom errors, wrapping, and recovery. Learn Go''s error handling best practices.',
    'Error interface and custom errors, error wrapping and unwrapping, panic and recover, error handling best practices.',
    8,
    240,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-09-goroutines-channels',
    'course-go-pro-2024',
    'goroutines-and-channels',
    'Goroutines and Channels',
    'Learn concurrent programming with goroutines and channels. Master buffered channels, select statements, and synchronization.',
    'Goroutine creation and lifecycle, channel creation and operations, buffered vs unbuffered channels, select statements.',
    9,
    420,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-10-packages-modules',
    'course-go-pro-2024',
    'packages-and-modules',
    'Packages and Modules',
    'Organize code into packages and manage dependencies with Go modules. Learn package design and versioning.',
    'Package organization and naming, import statements, Go modules and versioning, dependency management, package documentation.',
    10,
    300,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

-- Phase 3: Advanced (Lessons 11-15)

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-11-advanced-concurrency',
    'course-go-pro-2024',
    'advanced-concurrency-patterns',
    'Advanced Concurrency Patterns',
    'Master advanced concurrency patterns including worker pools, pipelines, and context-based cancellation.',
    'Worker pools and fan-out/fan-in, pipeline patterns, context package for cancellation, sync package primitives, race condition detection.',
    11,
    480,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-12-testing-benchmarking',
    'course-go-pro-2024',
    'testing-and-benchmarking',
    'Testing and Benchmarking',
    'Write comprehensive tests with table-driven patterns, benchmarks, and test coverage analysis.',
    'Unit testing with testing package, table-driven tests, test helpers and setup/teardown, benchmarking and profiling, test coverage analysis.',
    12,
    360,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-13-http-rest-apis',
    'course-go-pro-2024',
    'http-servers-and-rest-apis',
    'HTTP Servers and REST APIs',
    'Build production-ready REST APIs with proper routing, middleware, authentication, and validation.',
    'HTTP server basics with net/http, routing and middleware, JSON handling and validation, authentication and authorization, API design best practices.',
    13,
    480,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-14-database-integration',
    'course-go-pro-2024',
    'database-integration',
    'Database Integration',
    'Integrate databases with connection pooling, transactions, and migrations. Learn repository patterns.',
    'Database/sql package, connection pooling and management, prepared statements and transactions, ORM alternatives and patterns, migration strategies.',
    14,
    420,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-15-microservices',
    'course-go-pro-2024',
    'microservices-architecture',
    'Microservices Architecture',
    'Design and build microservices with proper communication patterns, configuration, and monitoring.',
    'Microservices design principles, service communication patterns, configuration management, logging and monitoring, deployment strategies.',
    15,
    540,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

-- Phase 4: Expert (Lessons 16-20)

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-16-performance-optimization',
    'course-go-pro-2024',
    'performance-optimization-and-profiling',
    'Performance Optimization and Profiling',
    'Profile and optimize Go applications for maximum performance. Learn memory optimization and CPU profiling.',
    'Profiling Go applications, identifying performance bottlenecks, memory optimization, CPU optimization, benchmarking, caching strategies.',
    16,
    420,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-17-security-practices',
    'course-go-pro-2024',
    'security-best-practices',
    'Security Best Practices',
    'Implement authentication, authorization, and secure coding practices. Prevent common vulnerabilities.',
    'Authentication and authorization, TLS and secure communications, input validation, SQL injection prevention, rate limiting, cryptographic best practices.',
    17,
    360,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-18-deployment-devops',
    'course-go-pro-2024',
    'deployment-and-devops',
    'Deployment and DevOps',
    'Deploy Go applications with Docker, Kubernetes, and CI/CD pipelines. Master DevOps practices.',
    'Building Go applications, Docker containerization, Kubernetes deployment, CI/CD pipelines, monitoring and logging, configuration management.',
    18,
    480,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-19-design-patterns',
    'course-go-pro-2024',
    'advanced-design-patterns',
    'Advanced Design Patterns',
    'Implement sophisticated design patterns and architectural patterns. Use generics and functional programming.',
    'Advanced design patterns, functional programming concepts, generics (Go 1.18+), plugin architectures, event-driven systems, domain-driven design.',
    19,
    420,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

INSERT INTO gopro.lessons (id, course_id, slug, title, description, content, lesson_order, estimated_duration_minutes, status, created_at, updated_at)
VALUES (
    'lesson-20-production-systems',
    'course-go-pro-2024',
    'building-production-systems',
    'Building Production Systems',
    'Design and build complete production-ready systems. Apply all concepts in real-world scenarios.',
    'Production system architecture, comprehensive observability, scalability and maintainability, production challenges, industry best practices.',
    20,
    600,
    'published',
    NOW(),
    NOW()
) ON CONFLICT (course_id, slug) DO UPDATE SET
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    estimated_duration_minutes = EXCLUDED.estimated_duration_minutes,
    updated_at = NOW();

-- Summary
SELECT
    COUNT(*) as total_lessons,
    SUM(estimated_duration_minutes) as total_minutes,
    ROUND(SUM(estimated_duration_minutes) / 60.0, 1) as total_hours
FROM gopro.lessons
WHERE course_id = 'course-go-pro-2024' AND status = 'published';

