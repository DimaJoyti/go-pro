// Package main implements a simple order management system
package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// OrderStatus represents the status of an order
type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusApproved  OrderStatus = "APPROVED"
	StatusShipped   OrderStatus = "SHIPPED"
	StatusDelivered OrderStatus = "DELIVERED"
	StatusCancelled OrderStatus = "CANCELLED"
)

// OrderItem represents an item in an order
type OrderItem struct {
	ProductID string
	Name      string
	Quantity  int
	Price     float64
}

// Order represents a customer order
type Order struct {
	ID        string
	UserID    string
	Items     []OrderItem
	Total     float64
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Errors
var (
	ErrInvalidOrder  = errors.New("invalid order")
	ErrInvalidStatus = errors.New("invalid status transition")
	ErrOrderNotFound = errors.New("order not found")
)

// NewOrder creates a new order
func NewOrder(userID string, items []OrderItem) (*Order, error) {
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

	return &Order{
		ID:        generateOrderID(),
		UserID:    userID,
		Items:     items,
		Total:     total,
		Status:    StatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// UpdateStatus updates the order status
func (o *Order) UpdateStatus(newStatus OrderStatus) error {
	if !isValidTransition(o.Status, newStatus) {
		return ErrInvalidStatus
	}

	o.Status = newStatus
	o.UpdatedAt = time.Now()
	return nil
}

// AddItem adds an item to the order
func (o *Order) AddItem(item OrderItem) error {
	if o.Status != StatusPending {
		return errors.New("cannot modify order after approval")
	}

	o.Items = append(o.Items, item)
	o.Total += float64(item.Quantity) * item.Price
	o.UpdatedAt = time.Now()
	return nil
}

// RemoveItem removes an item from the order
func (o *Order) RemoveItem(productID string) error {
	if o.Status != StatusPending {
		return errors.New("cannot modify order after approval")
	}

	for i, item := range o.Items {
		if item.ProductID == productID {
			o.Total -= float64(item.Quantity) * item.Price
			o.Items = append(o.Items[:i], o.Items[i+1:]...)
			o.UpdatedAt = time.Now()
			return nil
		}
	}

	return errors.New("item not found")
}

// String returns a string representation of the order
func (o *Order) String() string {
	return fmt.Sprintf("Order{ID: %s, UserID: %s, Total: $%.2f, Status: %s, Items: %d}",
		o.ID, o.UserID, o.Total, o.Status, len(o.Items))
}

// PrintDetails prints detailed order information
func (o *Order) PrintDetails() {
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Printf("║ Order ID: %-48s ║\n", o.ID)
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ User ID:     %-45s ║\n", o.UserID)
	fmt.Printf("║ Status:      %-45s ║\n", o.Status)
	fmt.Printf("║ Created:     %-45s ║\n", o.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("║ Updated:     %-45s ║\n", o.UpdatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║ Items:                                                     ║")

	for i, item := range o.Items {
		fmt.Printf("║ %2d. %-20s Qty: %3d  $%8.2f  $%8.2f ║\n",
			i+1, item.Name, item.Quantity, item.Price, float64(item.Quantity)*item.Price)
	}

	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Printf("║ Total:                                          $%9.2f ║\n", o.Total)
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
}

// OrderManager manages multiple orders
type OrderManager struct {
	orders map[string]*Order
}

// NewOrderManager creates a new order manager
func NewOrderManager() *OrderManager {
	return &OrderManager{
		orders: make(map[string]*Order),
	}
}

// CreateOrder creates and stores a new order
func (om *OrderManager) CreateOrder(userID string, items []OrderItem) (*Order, error) {
	order, err := NewOrder(userID, items)
	if err != nil {
		return nil, err
	}

	om.orders[order.ID] = order
	return order, nil
}

// GetOrder retrieves an order by ID
func (om *OrderManager) GetOrder(orderID string) (*Order, error) {
	order, exists := om.orders[orderID]
	if !exists {
		return nil, ErrOrderNotFound
	}
	return order, nil
}

// ListOrders returns all orders
func (om *OrderManager) ListOrders() []*Order {
	orders := make([]*Order, 0, len(om.orders))
	for _, order := range om.orders {
		orders = append(orders, order)
	}
	return orders
}

// Helper functions

func isValidTransition(current, new OrderStatus) bool {
	transitions := map[OrderStatus][]OrderStatus{
		StatusPending:   {StatusApproved, StatusCancelled},
		StatusApproved:  {StatusShipped, StatusCancelled},
		StatusShipped:   {StatusDelivered},
		StatusDelivered: {},
		StatusCancelled: {},
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
	return fmt.Sprintf("ORD-%d", time.Now().UnixNano())
}

// Main function - demonstrates the order system
func main() {
	fmt.Println("🛒 Order Management System Demo")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()

	// Create order manager
	manager := NewOrderManager()

	// Create sample orders
	fmt.Println("📝 Creating Orders...")
	fmt.Println()

	// Order 1
	items1 := []OrderItem{
		{ProductID: "P001", Name: "Laptop", Quantity: 1, Price: 999.99},
		{ProductID: "P002", Name: "Mouse", Quantity: 2, Price: 29.99},
	}
	order1, _ := manager.CreateOrder("USER-001", items1)
	order1.PrintDetails()
	fmt.Println()

	// Order 2
	items2 := []OrderItem{
		{ProductID: "P003", Name: "Keyboard", Quantity: 1, Price: 79.99},
		{ProductID: "P004", Name: "Monitor", Quantity: 2, Price: 299.99},
	}
	order2, _ := manager.CreateOrder("USER-002", items2)
	order2.PrintDetails()
	fmt.Println()

	// Update order status
	fmt.Println("📦 Processing Orders...")
	fmt.Println()

	order1.UpdateStatus(StatusApproved)
	fmt.Printf("✅ Order %s approved\n", order1.ID)

	order1.UpdateStatus(StatusShipped)
	fmt.Printf("🚚 Order %s shipped\n", order1.ID)

	order1.UpdateStatus(StatusDelivered)
	fmt.Printf("📬 Order %s delivered\n", order1.ID)
	fmt.Println()

	// List all orders
	fmt.Println("📋 All Orders:")
	fmt.Println()
	for _, order := range manager.ListOrders() {
		fmt.Printf("  • %s\n", order)
	}
	fmt.Println()

	// Error handling demo
	fmt.Println("❌ Error Handling Demo:")
	fmt.Println()

	// Try invalid status transition
	err := order1.UpdateStatus(StatusPending)
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	}

	// Try to modify delivered order
	err = order1.AddItem(OrderItem{ProductID: "P005", Name: "Cable", Quantity: 1, Price: 9.99})
	if err != nil {
		fmt.Printf("  Error: %v\n", err)
	}
	fmt.Println()

	fmt.Println("✨ Demo completed!")
}
