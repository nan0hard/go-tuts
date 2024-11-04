package main

import (
	"fmt"
	"time"
)

// Receive only channel
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	// Following two operations are not permitted.
	// emailChan <- "Nitish"
	// <-done
	for email := range emailChan {
		fmt.Println("Sending email to: ", email)
		time.Sleep(time.Second)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() { chan1 <- 10 }()
	go func() { chan2 <- "Hello" }()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Channel 1 value received", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("Channel 2 value received", chan2Val)

		}
	}
}
