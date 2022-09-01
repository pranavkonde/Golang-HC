package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This need to go in file"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err) // will terminate the program
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err) // will terminate the program
	}

	fmt.Println("Length is", length)
	file.Close()
	readFile("./mylcogofile.txt")

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text inside file is", string(databyte))

}
