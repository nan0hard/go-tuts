package main

import "fmt"

// Sending data to go routine using channel
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing Number: ", num)
// 		time.Sleep(time.Second)
// 	}
// }

// Receive data from the go routine
func sum(result chan int, num1 int, num2 int) {
	res := num1 + num2
	result <- res
}

func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("Processing")
	// done <- true // May be it won't run if there is some issue with fmt.Println() which will never happen so will use defer instead. just to be sure.
}

func main() {

	isDone := make(chan bool)

	go task(isDone)

	<-isDone // blocking code

	// result := make(chan int)

	// go sum(result, 1, 2)

	// res := <-result
	// fmt.Print(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "Nitish is a good boy"
	// msg := <-messageChan

	// fmt.Println(msg)
}
