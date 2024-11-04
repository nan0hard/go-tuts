package main

import "fmt"

func add(a, b int) int { // concised from (a int, b int) to (a, b int)
	return a + b
}

func getLang() (string, string, string, string) {
	return "Java", "JavaScript", "TypeScript", "React"
}

func process(a int) int {
	return a
}

func main() {
	res := add(5, 8)
	fmt.Println(res)

	fmt.Println(getLang())
	// Or we can destructure it as well
	a, b, c, d := getLang()
	fmt.Println(a, c, b, d)
}
