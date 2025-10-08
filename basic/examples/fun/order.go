// This is a standalone example - run with: go run order.go
// Note: Rename mainOrder to main to run this file standalone
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Order struct {
	ID     int
	Status string
	Price  float64
}

func mainOrder() {
	orders := generateOrders(10) // Generate 10 orders

	fmt.Print("Orders before sorting:\n")
	// Print unsorted orders
	for _, order := range orders {
		fmt.Printf("ID: %d, Status: %s, Price: %.2f\n", order.ID, order.Status, order.Price)
	}

	// Sort orders by price in ascending order
	sortOrders(orders)

	fmt.Print("\nOrders after sorting:\n")
	// Print sorted orders
	for _, order := range orders {
		fmt.Printf("ID: %d, Status: %s, Price: %.2f\n", order.ID, order.Status, order.Price)
	}
}

func processOrders(orders []*Order) {
	// Process orders
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "pending",
			Price:  rand.Float64() * 1000,
		}
	}
	return orders
}

// Sort orders by price in ascending order
func sortOrders(orders []*Order) {
	// Sort orders using built-in sort package
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].Price < orders[j].Price
	})
}
