package main

import "fmt"

// Const/var can be declared outside main

const hero string = "Superman"

var villain = "Joker"

const (
	port = 8080
	host = "localhost"
)

func main() {
	const firstName string = "Nitish"
	fmt.Println(hero, villain, port, host)
}
