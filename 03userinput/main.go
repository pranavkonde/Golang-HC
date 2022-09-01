package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)
	fmt.Println("Enter the rating for our Pizza:")

	//comma ok || err err

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T\n", input)
	fmt.Println(err)
}
