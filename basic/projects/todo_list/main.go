package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Task represents a single todo item
type Task struct {
	ID          int
	Description string
	Completed   bool
	CreatedAt   time.Time
}

// TodoList manages a collection of tasks
type TodoList struct {
	tasks  []Task
	nextID int
}

// NewTodoList creates a new todo list
func NewTodoList() *TodoList {
	return &TodoList{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

// AddTask adds a new task to the list
func (tl *TodoList) AddTask(description string) {
	task := Task{
		ID:          tl.nextID,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}
	tl.tasks = append(tl.tasks, task)
	tl.nextID++
	fmt.Printf("✓ Task added: #%d - %s\n", task.ID, task.Description)
}

// ListTasks displays all tasks
func (tl *TodoList) ListTasks() {
	if len(tl.tasks) == 0 {
		fmt.Println("\n📝 No tasks yet. Add some tasks to get started!")
		return
	}

	fmt.Println("\n📝 Your Tasks:")
	fmt.Println(strings.Repeat("=", 70))

	pendingCount := 0
	completedCount := 0

	for _, task := range tl.tasks {
		status := "⬜"
		if task.Completed {
			status = "✅"
			completedCount++
		} else {
			pendingCount++
		}

		fmt.Printf("%s [%d] %s\n", status, task.ID, task.Description)
	}

	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("Total: %d | Pending: %d | Completed: %d\n",
		len(tl.tasks), pendingCount, completedCount)
}

// CompleteTask marks a task as completed
func (tl *TodoList) CompleteTask(id int) {
	for i := range tl.tasks {
		if tl.tasks[i].ID == id {
			if tl.tasks[i].Completed {
				fmt.Printf("⚠️  Task #%d is already completed\n", id)
			} else {
				tl.tasks[i].Completed = true
				fmt.Printf("✓ Task #%d marked as completed!\n", id)
			}
			return
		}
	}
	fmt.Printf("❌ Task #%d not found\n", id)
}

// DeleteTask removes a task from the list
func (tl *TodoList) DeleteTask(id int) {
	for i, task := range tl.tasks {
		if task.ID == id {
			tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
			fmt.Printf("✓ Task #%d deleted\n", id)
			return
		}
	}
	fmt.Printf("❌ Task #%d not found\n", id)
}

// ClearCompleted removes all completed tasks
func (tl *TodoList) ClearCompleted() {
	newTasks := make([]Task, 0)
	count := 0

	for _, task := range tl.tasks {
		if !task.Completed {
			newTasks = append(newTasks, task)
		} else {
			count++
		}
	}

	tl.tasks = newTasks
	fmt.Printf("✓ Cleared %d completed task(s)\n", count)
}

// ShowPending shows only pending tasks
func (tl *TodoList) ShowPending() {
	fmt.Println("\n⏳ Pending Tasks:")
	fmt.Println(strings.Repeat("=", 70))

	count := 0
	for _, task := range tl.tasks {
		if !task.Completed {
			fmt.Printf("⬜ [%d] %s\n", task.ID, task.Description)
			count++
		}
	}

	if count == 0 {
		fmt.Println("No pending tasks! 🎉")
	}
	fmt.Println(strings.Repeat("=", 70))
}

func printBanner() {
	fmt.Println("\n╔════════════════════════════════════════════════╗")
	fmt.Println("║                                                ║")
	fmt.Println("║           📝  Go Todo List  📝                 ║")
	fmt.Println("║                                                ║")
	fmt.Println("╚════════════════════════════════════════════════╝")
}

func printMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📋 Menu:")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("  [1] Add Task")
	fmt.Println("  [2] List All Tasks")
	fmt.Println("  [3] Complete Task")
	fmt.Println("  [4] Delete Task")
	fmt.Println("  [5] Show Pending Tasks")
	fmt.Println("  [6] Clear Completed Tasks")
	fmt.Println("  [q] Quit")
	fmt.Println(strings.Repeat("=", 50))
}

func getInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func getTaskID(reader *bufio.Reader) (int, error) {
	input := getInput(reader, "Enter task ID: ")
	return strconv.Atoi(input)
}

func main() {
	todoList := NewTodoList()
	reader := bufio.NewReader(os.Stdin)

	printBanner()

	// Add some sample tasks
	fmt.Println("\n💡 Adding sample tasks...")
	todoList.AddTask("Learn Go basics")
	todoList.AddTask("Build a todo list app")
	todoList.AddTask("Practice concurrency")

	for {
		printMenu()
		choice := getInput(reader, "\n👉 Enter your choice: ")

		switch choice {
		case "1":
			description := getInput(reader, "Enter task description: ")
			if description != "" {
				todoList.AddTask(description)
			} else {
				fmt.Println("❌ Task description cannot be empty")
			}

		case "2":
			todoList.ListTasks()

		case "3":
			id, err := getTaskID(reader)
			if err != nil {
				fmt.Println("❌ Invalid task ID")
			} else {
				todoList.CompleteTask(id)
			}

		case "4":
			id, err := getTaskID(reader)
			if err != nil {
				fmt.Println("❌ Invalid task ID")
			} else {
				todoList.DeleteTask(id)
			}

		case "5":
			todoList.ShowPending()

		case "6":
			todoList.ClearCompleted()

		case "q", "Q", "quit", "exit":
			fmt.Println("\n👋 Thanks for using Go Todo List! Goodbye!")
			return

		default:
			fmt.Println("\n❌ Invalid choice. Please try again.")
		}
	}
}
