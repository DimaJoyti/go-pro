# Todo List Project

A simple command-line todo list application built with Go.

## Features

- ✅ Add new tasks
- ✅ List all tasks
- ✅ Mark tasks as completed
- ✅ Delete tasks
- ✅ Show only pending tasks
- ✅ Clear all completed tasks
- ✅ Task statistics (total, pending, completed)
- ✅ Interactive menu system
- ✅ Clean, emoji-enhanced interface

## How to Run

```bash
cd basic/projects/todo_list
go run main.go
```

## Usage Example

```
╔════════════════════════════════════════════════╗
║                                                ║
║           📝  Go Todo List  📝                 ║
║                                                ║
╚════════════════════════════════════════════════╝

💡 Adding sample tasks...
✓ Task added: #1 - Learn Go basics
✓ Task added: #2 - Build a todo list app
✓ Task added: #3 - Practice concurrency

==================================================
📋 Menu:
==================================================
  [1] Add Task
  [2] List All Tasks
  [3] Complete Task
  [4] Delete Task
  [5] Show Pending Tasks
  [6] Clear Completed Tasks
  [q] Quit
==================================================

👉 Enter your choice: 2

📝 Your Tasks:
======================================================================
⬜ [1] Learn Go basics
⬜ [2] Build a todo list app
⬜ [3] Practice concurrency
======================================================================
Total: 3 | Pending: 3 | Completed: 0
```

## Operations

### 1. Add Task
Add a new task to your todo list with a description.

### 2. List All Tasks
Display all tasks with their status (pending ⬜ or completed ✅).

### 3. Complete Task
Mark a task as completed by entering its ID.

### 4. Delete Task
Remove a task from the list by entering its ID.

### 5. Show Pending Tasks
Display only tasks that are not yet completed.

### 6. Clear Completed Tasks
Remove all completed tasks from the list.

## Code Structure

- `Task` struct - Represents a single todo item with ID, description, completion status, and creation time
- `TodoList` struct - Manages the collection of tasks
- `AddTask` - Add a new task
- `ListTasks` - Display all tasks with statistics
- `CompleteTask` - Mark a task as done
- `DeleteTask` - Remove a task
- `ClearCompleted` - Remove all completed tasks
- `ShowPending` - Show only pending tasks

## Learning Objectives

- Struct design and methods
- Slice manipulation (add, remove, filter)
- User input handling
- String formatting and display
- Time handling
- Interactive CLI applications
- State management

## Future Enhancements

- Save tasks to a file (persistence)
- Edit task descriptions
- Set task priorities
- Add due dates
- Search and filter tasks
- Task categories/tags

