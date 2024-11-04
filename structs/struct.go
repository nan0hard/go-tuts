package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	price     float32
	status    string
	createdAt time.Time // nanosecond precision
}

// Constructor
func newOrder(id string, price float32, status string) *order {
	myOrder := order{
		id:     id,
		price:  price,
		status: status,
	}
	return &myOrder
}

func (o *order) changeOrderStatus(status string) {
	o.status = status
}

func (o order) getPrice() float32 {
	return o.price
}

func mainas() {
	order1 := order{
		id: "order1", price: 2500.25, status: "Pending",
	}

	order1.createdAt = time.Now()

	fmt.Println("Order1 struct", order1)
	order1.changeOrderStatus("Delivered")
	fmt.Println("Order1 struct", order1)

	order2 := newOrder("1", 32, "Pending")
	fmt.Println(order2)

	// Inline struct
	order3 := struct {
		id   string
		name string
	}{"1", "Nitish"}

	fmt.Println(order3)

}
