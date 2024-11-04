package main

import (
	"fmt"
	"time"
)

type customer struct {
	name        string
	phoneNumber int
}

type orderDetails struct {
	orderID   string
	status    string
	price     float32
	createdAt time.Time
	customer
}

func main() {
	newCustomer := customer{
		name:        "Nitish Singh",
		phoneNumber: 12345,
	}

	newOrder := orderDetails{
		orderID:  "Order1",
		status:   "Pending",
		price:    25,
		customer: newCustomer,
	}

	fmt.Println(newOrder.customer)
}
