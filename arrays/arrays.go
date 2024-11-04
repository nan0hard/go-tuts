package main

import "fmt"

func main() {
	// var arr [4]int

	// Length of an array
	// fmt.Println(len(arr))

	// arr[0] = 1 // By default all numbers are 0
	// fmt.Println(arr)

	// String array
	var names [3]string // By default the entries will be ""{empty string}
	names[1] = "Nitish is a good boy!!"
	fmt.Println(names)

	// Initialization
	roll_No := [3]int{1, 2, 3}
	fmt.Println(roll_No)

	// 2D array initialization
	nums := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(len(nums[0]))

}
