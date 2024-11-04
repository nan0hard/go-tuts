package main

import "fmt"

func main() {
	// Make a slice
	nums := []int{10, 50, 20, 40, 0, 100}

	// iteration using for loop
	for i := 0; i < len(nums); i++ {
		fmt.Printf("At index %v and it's value is %v \n", i, nums[i])
	}

	// using range in array
	for ind, num := range nums {
		fmt.Println(ind, num) // gives two values index and it's number else use underscore(_)
	}

	hm := map[int]string{0: "nitish", 1: "Nitish", 2: "Satyendra", 3: "Kumar", 4: "Singh"} // Unordered Map

	for key, val := range hm {
		fmt.Println(key, val)
	}

	for key := range hm { // If single value then it will print key/iterate on keys
		fmt.Println(key)
	}

	// Iterating on particular string

	for i, c := range "Nitish Singh" {
		fmt.Println(i, string(c)) // by default c returns rune/unicode value kind of ascii value only but differnt and i is not an index, it's a size if the unicode character is > 255 then it will take 2 bytes and in this  case it will increment i by 2
	}

}
