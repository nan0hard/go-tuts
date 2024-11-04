package main

import (
	"fmt"
	"os"
)

func main() {
	// cont, err := os.Open("text.txt")

	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }

	// fileInfo, err := cont.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Name of the file:", fileInfo.Name())
	// fmt.Println("isDirectory:", fileInfo.IsDir())
	// fmt.Println("File Size:", fileInfo.Size())

	// Reading a file
	// f, err := os.Open("text.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 14)
	// d, err := f.Read(buf)

	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < d; i++ {
	// 	fmt.Println(d, string(buf[i]))
	// }

	// fmt.Println(d, string(buf))

	// Read a file - Easy mode

	// data, err := os.ReadFile("text.txt") // only use for small files
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// Read a folder/files
	files, err := os.Open("../")

	if err != nil {
		panic(err)
	}

	defer files.Close()

	dir, err := files.ReadDir(-1)

	for _, val := range dir {
		fmt.Println(val.Name(), val.IsDir())
	}
}
