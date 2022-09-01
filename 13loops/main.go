package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, day)
	// }

	rougueValue := 1
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}
		if rougueValue == 5 {
			// break
			rougueValue++
			continue
		}
		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

	//labels
lco:
	fmt.Println("Jumping")
}
