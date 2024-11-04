package main

import "fmt"

func main() {
	//  Declaring a map
	hm := make(map[int]string)

	// Setting a value in a map
	hm[1] = "Nitish"
	hm[0] = "Nano"

	fmt.Println(hm[0], hm[1])

	// If key is not present in map, then it will return a falsy value: For string: "", int: 0,  boolean: false
	fmt.Println(hm[10])

	hashMap := make(map[string]int)
	hashMap["Nitish"] = 100
	fmt.Println(hashMap["Nitish"], hashMap["Nano"], len(hashMap))

	// Delete a key
	delete(hm, 1)
	fmt.Println(hm, hashMap)

	// clear whole map
	clear(hm)
	fmt.Println(hm)

	// Initialize a map without make

	maps := map[int]string{1: "Nitish", 2: "Singh", 3: "Satyendra"}

	val, okay := maps[4]
	fmt.Println(val)

	// For checking if key is present or not
	if okay {
		fmt.Println("Present")
	} else {
		fmt.Println("Absent")
	}
}
