# GO-PRO Lessons Architecture

## Overview
This document describes the architecture of the lessons system in the GO-PRO platform.

## System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         Frontend (Next.js)                       │
│                     http://localhost:3001                        │
└─────────────────────────────────────────────────────────────────┘
                              │
                              │ HTTP Requests
                              ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Backend API (Go)                            │
│                     http://localhost:8080                        │
│                                                                   │
│  Routes:                                                          │
│  GET /api/v1/curriculum          → Full curriculum               │
│  GET /api/v1/curriculum/lesson/{id} → Lesson details             │
└─────────────────────────────────────────────────────────────────┘
```

## Frontend Routing Structure

### Before Fix (❌ Broken)
```
/learn/lesson-1/page.tsx  → Redirects to /learn/lesson-1 (LOOP!)
/learn/lesson-[id]/page.tsx → Expects /learn/1, /learn/2, etc.

Links used:
- /learn/lesson-1  ❌
- /learn/lesson-2  ❌
- /learn/lesson-3  ❌
```

### After Fix (✅ Working)
```
/learn/[id]/page.tsx  → Dynamic route for all lessons

Links use:
- /learn/1  ✅
- /learn/2  ✅
- /learn/3  ✅
```

## Component Hierarchy

```
┌─────────────────────────────────────────────────────────────────┐
│                         App Routes                               │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  /                                                                │
│  └─ page.tsx (Homepage)                                          │
│     └─ "Start Learning" → /learn/1                               │
│                                                                   │
│  /curriculum                                                      │
│  └─ page.tsx (Curriculum Overview)                               │
│     └─ Lesson Cards → /learn/{id}                                │
│                                                                   │
│  /learn                                                           │
│  ├─ page.tsx (Learning Dashboard)                                │
│  │  └─ Recent Lessons → /learn/{id}                              │
│  │                                                                │
│  └─ [id]                                                          │
│     └─ page.tsx (Dynamic Lesson Page)                            │
│        ├─ Theory Tab                                              │
│        ├─ Practice Tab (CodeEditor)                               │
│        └─ Exercise Tab (CodeEditor + ExerciseSubmission)          │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘
```

## Lesson Page Components

```
┌─────────────────────────────────────────────────────────────────┐
│                    Lesson Page (/learn/[id])                     │
├─────────────────────────────────────────────────────────────────┤
│                                                                   │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │ Breadcrumb Navigation                                    │    │
│  │ Home > Curriculum > Lesson {id}                          │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │ Lesson Header                                            │    │
│  │ - Title                                                  │    │
│  │ - Description                                            │    │
│  │ - Difficulty Badge                                       │    │
│  │ - Duration, Objectives Count                             │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │ Tabs: [Theory] [Practice] [Exercise]                     │    │
│  ├─────────────────────────────────────────────────────────┤    │
│  │                                                           │    │
│  │ Theory Tab:                                               │    │
│  │ ┌───────────────────────────────────────────────────┐   │    │
│  │ │ Learning Objectives                                │   │    │
│  │ │ - Objective 1                                      │   │    │
│  │ │ - Objective 2                                      │   │    │
│  │ │ - ...                                              │   │    │
│  │ └───────────────────────────────────────────────────┘   │    │
│  │                                                           │    │
│  │ ┌───────────────────────────────────────────────────┐   │    │
│  │ │ Theory & Concepts                                  │   │    │
│  │ │ (Markdown-rendered content)                        │   │    │
│  │ └───────────────────────────────────────────────────┘   │    │
│  │                                                           │    │
│  │ Practice Tab:                                             │    │
│  │ ┌───────────────────────────────────────────────────┐   │    │
│  │ │ CodeEditor Component                               │   │    │
│  │ │ - Monaco Editor                                    │   │    │
│  │ │ - Run Code Button                                  │   │    │
│  │ │ - Show Solution Button                             │   │    │
│  │ │ - Output Panel                                     │   │    │
│  │ └───────────────────────────────────────────────────┘   │    │
│  │                                                           │    │
│  │ Exercise Tab:                                             │    │
│  │ ┌───────────────────────────────────────────────────┐   │    │
│  │ │ CodeEditor (left)  │  ExerciseSubmission (right)  │   │    │
│  │ │ - Exercise code    │  - Requirements              │   │    │
│  │ │ - Run button       │  - Submit button             │   │    │
│  │ │                    │  - Feedback display          │   │    │
│  │ └───────────────────────────────────────────────────┘   │    │
│  │                                                           │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │ Navigation                                               │    │
│  │ [← Previous Lesson]  [All Lessons]  [Next Lesson →]      │    │
│  └─────────────────────────────────────────────────────────┘    │
│                                                                   │
└─────────────────────────────────────────────────────────────────┘
```

## Data Flow

### 1. User Navigates to Lesson
```
User clicks lesson
    ↓
Router navigates to /learn/{id}
    ↓
LessonPage component mounts
    ↓
useEffect hook triggers
    ↓
Extract lessonId from params
```

### 2. Fetch Lesson Data
```
api.getLessonDetail(lessonId)
    ↓
HTTP GET /api/v1/curriculum/lesson/{id}
    ↓
Backend: curriculumService.GetLessonDetail()
    ↓
Returns LessonDetail JSON
    ↓
Frontend: Transform API response to component format
    ↓
setLessonData(transformedData)
```

### 3. Render Lesson
```
LessonData state updated
    ↓
Component re-renders
    ↓
Display lesson content in tabs
    ↓
User interacts with code editor, exercises, etc.
```

## API Response Structure

### GET /api/v1/curriculum/lesson/{id}

```json
{
  "success": true,
  "data": {
    "id": 1,
    "title": "Go Syntax and Basic Types",
    "description": "Learn the fundamental syntax...",
    "duration": "3-4 hours",
    "difficulty": "beginner",
    "phase": "Foundations",
    "objectives": [
      "Set up a Go development environment",
      "Understand Go's basic syntax...",
      ...
    ],
    "theory": "# Go Program Structure\n\n...",
    "code_example": "package main\n\nimport \"fmt\"...",
    "solution": "package main\n\nimport \"fmt\"...",
    "exercises": [
      {
        "id": "basic-variables",
        "title": "Variable Declaration Practice",
        "description": "Practice declaring variables...",
        "requirements": [...],
        "initial_code": "package main...",
        "solution": "package main..."
      }
    ],
    "next_lesson_id": 2,
    "prev_lesson_id": null
  },
  "message": "lesson detail retrieved successfully",
  "request_id": "...",
  "timestamp": "2025-10-07T..."
}
```

## Key Components

### 1. CodeEditor (`components/learning/code-editor.tsx`)
- Monaco Editor integration
- Syntax highlighting for Go
- Run code functionality (placeholder)
- Show/hide solution
- Copy code to clipboard
- Reset to initial code

### 2. ExerciseSubmission (`components/learning/exercise-submission.tsx`)
- Display exercise requirements
- Submit solution (placeholder)
- Show feedback
- Track previous submissions

### 3. LessonPage (`app/learn/[id]/page.tsx`)
- Main lesson container
- Tab management (Theory, Practice, Exercise)
- Lesson data fetching
- Navigation between lessons

## URL Patterns

### Correct Patterns ✅
```
Homepage:           /
Curriculum:         /curriculum
Learning Dashboard: /learn
Lesson 1:          /learn/1
Lesson 2:          /learn/2
Lesson N:          /learn/{N}
```

### Incorrect Patterns ❌ (Fixed)
```
/learn/lesson-1    ← Old pattern (removed)
/learn/lesson-2    ← Old pattern (removed)
/learn/lesson-N    ← Old pattern (removed)
```

## Backend Service

### Location
`backend/internal/service/curriculum.go`

### Key Functions
```go
// GetLessonDetail returns detailed information about a specific lesson
func (s *curriculumService) GetLessonDetail(ctx context.Context, lessonID int) (*domain.LessonDetail, error)
```

### Current Implementation
- Mock data for Lesson 1
- Returns hardcoded lesson content
- Includes theory, code examples, exercises
- Provides navigation (next/prev lesson IDs)

### Future Enhancement
- Load lessons from database
- Support for all 20 lessons
- Dynamic content management
- Progress tracking integration

## State Management

### Component State
```typescript
const [activeTab, setActiveTab] = useState("lesson");
const [lessonData, setLessonData] = useState<LessonData | null>(null);
const [loading, setLoading] = useState(true);
```

### Loading States
1. **Initial**: `loading = true`, `lessonData = null`
2. **Loading**: Fetching from API
3. **Success**: `loading = false`, `lessonData = {...}`
4. **Error**: `loading = false`, `lessonData = null`

## Error Handling

### Current Implementation
```typescript
try {
  const lesson = await api.getLessonDetail(lessonId);
  setLessonData(transformedData);
} catch (error) {
  console.error('Failed to load lesson:', error);
  setLessonData(null);
} finally {
  setLoading(false);
}
```

### Future Enhancements
- Display user-friendly error messages
- Retry mechanism
- Fallback content
- 404 handling for non-existent lessons

## Performance Considerations

### Current
- Client-side data fetching
- No caching
- Full page re-render on navigation

### Potential Optimizations
- Implement SWR or React Query for caching
- Prefetch next/previous lessons
- Code splitting for Monaco Editor
- Lazy load exercise components

## Accessibility

### Current Features
- Semantic HTML structure
- Keyboard navigation support
- ARIA labels on interactive elements
- Responsive design

### Future Enhancements
- Screen reader announcements
- Focus management
- Keyboard shortcuts
- High contrast mode support

