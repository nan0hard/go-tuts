package main

import (
	"os"
)

func main() {
	// f, err := os.Create("example1.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()
	// f.WriteString("Nitish is a good boy.!")
	// f.WriteString("Nitish is a bad boy!!")

	// contents, err := os.ReadFile("example1.txt")
	// fmt.Println(string(contents))

	// // Another method for writing -> Also appends
	// bytes := []byte("Nitish is a man!!")
	// f.Write(bytes)

	// Read and write to another file
	// srcFile, err := os.Open("example1.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer srcFile.Close()

	// destFile, err := os.Create("dest.txt")

	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(srcFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	err1 := writer.WriteByte(b)
	// 	if err1 != nil {
	// 		panic(err1)
	// 	}
	// }

	// writer.Flush()

	// Delete a file
	e := os.Remove("dest.txt")
	if e != nil {
		panic(e)
	}
}
