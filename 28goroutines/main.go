// parallelism := doing multiple task at same time
// concurrency := doing one task at one time

// Thread - Managed by OS, fixedstack-1MB
// Goroutines - Managed by Go runtime, Flexible stack-2KB

//Goroutine - Do not communicate by sharing memory, instead share memory by communicating

package main

import (
	"fmt"
	"time"
)

func main() {
	go greeter("Hello") //if sleep is not given than only world will be printed
	greeter("World")

}
func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println(s)
	}
}
