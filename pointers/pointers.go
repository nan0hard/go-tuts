package main

import "fmt"

// Simple by value
func numChange(num int) {
	num = 5
	fmt.Println("In numChange by value", num)
}

// pass by reference
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum by reference", *num)
}

func main() {
	num := 1
	numChange(num)
	fmt.Println("In main", num)

	changeNum(&num)
	fmt.Println("After changeNum ref", num)
}
