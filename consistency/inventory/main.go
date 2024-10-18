package main

import (
	"fmt"
	"sync"
	"time"
)

// Product represents an item in the store
type Product struct {
	ID          int
	Description string
	Inventory   int
	mutex       sync.Mutex
}

// UpdateDescription updates the product description (eventually consistent)
func (p *Product) UpdateDescription(desc string) {
	go func() {
		// Simulate propagation delay
		time.Sleep(2 * time.Second)
		p.Description = desc
		fmt.Printf("Description updated to: %s\n", p.Description)
	}()
}

// Purchase attempts to buy the product (strongly consistent)
func (p *Product) Purchase(quantity int) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.Inventory >= quantity {
		p.Inventory -= quantity
		fmt.Printf("Purchased %d of product %d. Remaining inventory: %d\n", quantity, p.ID, p.Inventory)
	} else {
		fmt.Println("Not enough inventory to complete purchase.")
	}
}

func main() {
	product := &Product{
		ID:          1,
		Description: "Original Description",
		Inventory:   5,
	}

	// Update description (eventually consistent)
	product.UpdateDescription("New Description")

	// Purchase product (strongly consistent)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		product.Purchase(3)
	}()
	go func() {
		defer wg.Done()
		product.Purchase(3)
	}()
	wg.Wait()

	fmt.Printf("Final Inventory: %d\n", product.Inventory)
	// Wait to see the description update
	time.Sleep(3 * time.Second)
	fmt.Printf("Final Description: %s\n", product.Description)
}
