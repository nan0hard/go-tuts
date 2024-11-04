package main

import (
	"fmt"
	"slices"
)

// slice - Dynamic array
// most used construct
// + useful methods

func main() {
	// Declaration
	// var nums []int // By default it will have nil(basically null) rather than 0 which is in arrays
	// fmt.Println(nums)
	// fmt.Println(nums == nil) // true

	// Initialization using make
	// var nums = make([]int, 2)
	// fmt.Println(nums == nil)
	// fmt.Println(nums, cap(nums), len(nums))

	// Initialization normally
	nums := []int{}
	fmt.Println(nums, len(nums), cap(nums))
	nums = append(nums, 1)
	nums = append(nums, 2)
	fmt.Println(nums, len(nums), cap(nums))
	nums = append(nums, 3)
	fmt.Println(nums, len(nums), cap(nums))

	// Copy from one array to another
	nums2 := make([]int, len(nums))
	copy(nums2, nums) // copy(destionation, source)
	fmt.Println(nums, nums2)

	// Slice operator
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 7)
	nums = append(nums, 8)
	fmt.Println(nums[0:7]) // excludes mentioned index at the end
	fmt.Println(nums[:6])
	fmt.Println(nums[3:]) // Prints till last including last one

	// slices package
	arr1 := []int{1, 2, 4, 8}
	arr2 := []int{1, 2, 4, 8}

	fmt.Println(slices.Equal(arr1, arr2))
}
