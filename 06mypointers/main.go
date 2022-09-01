package main

import "fmt"

func main() {
	fmt.Println("welcome to a class of pointers")
	// var ptr *int
	// fmt.Println(ptr)
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	// var ptr1 *int=&myNumber
	// fmt.Println(pt)
	// *ptr=*ptr+2
}
