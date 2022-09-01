package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes,myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Result is: ", proRes)
	fmt.Println("Pro message is:",myMessage)
}
func greeter() {
	fmt.Println("Namstey from golang")
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func proAdder(values ...int) int ,string{ //values are passed as slices
	total := 0
	for _, val := range values {
		total += val
	}
	return total,"Hi Pro"
}
