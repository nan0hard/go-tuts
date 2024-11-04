package main

import "fmt"

type OrderStatus int
type OrderStatus2 string

const (
	Received OrderStatus = iota
	Pending
	Processing
	In_Transit
	Delivered
)

const (
	Received2   OrderStatus2 = "received"
	Pending2                 = "pending"
	Processing2              = "processing"
	In_Transit2              = "in-Transit"
	Delivered2               = "delivered"
)

func changeOrderStatus2(status OrderStatus2) {
	fmt.Println("Order status changed to:", status)
}

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	changeOrderStatus(Delivered)
	changeOrderStatus2(Processing2)
}
