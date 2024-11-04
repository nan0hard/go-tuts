package main

import "fmt"

// Variadic means any number

func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	nums := []int{1, 5, 8}
	fmt.Println(sum(nums...))
}
