package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Println("Currently running task:", id)
}

func main() {
	for i := 0; i < 10; i++ {
		// go task(i)

		go func(i int) {
			fmt.Println("Running task", i)
		}(i)
	}

	time.Sleep(time.Second * 1)
}
