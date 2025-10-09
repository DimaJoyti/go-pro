// This is a standalone example - run with: go run orders.go
// Note: Rename mainOrders to main to run this file standalone
package main

import (
	"errors"
	"fmt"
	"time"
)

type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusApproved  OrderStatus = "APPROVED"
	StatusShipped   OrderStatus = "SHIPPED"
	StatusDelivered OrderStatus = "DELIVERED"
	StatusCancelled OrderStatus = "CANCELLED"
)

type DetailedOrder struct {
	ID        string
	UserID    string
	Items     []OrderItem
	Total     float64
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	ProductID string
	Quantity  int
	Price     float64
}

var (
	ErrInvalidOrder     = errors.New("invalid order")
	ErrInvalidStatus    = errors.New("invalid status transition")
	ErrOrderNotFound    = errors.New("order not found")
	ErrInsufficientItem = errors.New("insufficient item quantity")
)

func NewOrder(userID string, items []OrderItem) (*DetailedOrder, error) {
	if len(items) == 0 {
		return nil, ErrInvalidOrder
	}

	total := 0.0
	for _, item := range items {
		if item.Quantity <= 0 || item.Price < 0 {
			return nil, ErrInvalidOrder
		}
		total += float64(item.Quantity) * item.Price
	}

	return &DetailedOrder{
		ID:        generateOrderID(),
		UserID:    userID,
		Items:     items,
		Total:     total,
		Status:    StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (o *DetailedOrder) UpdateStatus(newStatus OrderStatus) error {
	if !isValidTransition(o.Status, newStatus) {
		return ErrInvalidStatus
	}

	o.Status = newStatus
	o.UpdatedAt = time.Now()
	return nil
}

func isValidTransition(current, new OrderStatus) bool {
	transitions := map[OrderStatus][]OrderStatus{
		StatusPending:  {StatusApproved, StatusCancelled},
		StatusApproved: {StatusShipped, StatusCancelled},
		StatusShipped:  {StatusDelivered},
	}

	validNext, exists := transitions[current]
	if !exists {
		return false
	}

	for _, status := range validNext {
		if status == new {
			return true
		}
	}
	return false
}

func generateOrderID() string {
	return time.Now().Format("20060102-150405")
}

func mainOrders() {
	// Example usage of the order system
	fmt.Println("Order Management System")
	fmt.Println("======================")

	// Create order items
	items := []OrderItem{
		{ProductID: "PROD-001", Quantity: 2, Price: 29.99},
		{ProductID: "PROD-002", Quantity: 1, Price: 49.99},
	}

	// Create a new order
	order, err := NewOrder("USER-123", items)
	if err != nil {
		fmt.Println("Error creating order:", err)
		return
	}

	fmt.Printf("Created order: %s\n", order.ID)
	fmt.Printf("User ID: %s\n", order.UserID)
	fmt.Printf("Total: $%.2f\n", order.Total)
	fmt.Printf("Status: %s\n", order.Status)

	// Update order status
	err = order.UpdateStatus(StatusApproved)
	if err != nil {
		fmt.Println("Error updating status:", err)
		return
	}
	fmt.Printf("Updated status to: %s\n", order.Status)

	// Try to ship the order
	err = order.UpdateStatus(StatusShipped)
	if err != nil {
		fmt.Println("Error updating status:", err)
		return
	}
	fmt.Printf("Updated status to: %s\n", order.Status)
}
