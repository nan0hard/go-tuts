package main

import (
	"fmt"
)

func main() {
	// simple switch - break is required here, default is optional

	// i := 10

	// switch i {
	// case 1:
	// 	fmt.Println("one")

	// case 2:
	// 	fmt.Println("two")

	// default:
	// 	fmt.Println("other", i)
	// }

	// Multiconditional Switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's a weekend")
	// default:
	// 	fmt.Println("Workday!!")
	// }

	// type Switch
	typeof := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It's and Integer")
		case bool:
			fmt.Println("It's a boolean")
		case string:
			fmt.Println("It's a string")

		default:
			fmt.Println("Other type")
		}
	}

	typeof(1.2)
}
