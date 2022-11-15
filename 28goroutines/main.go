// parallelism := doing multiple task at same time
// concurrency := doing one task at one time

// Thread - Managed by OS, fixedstack-1MB
// Goroutines - Managed by Go runtime, Flexible stack-2KB

//Goroutine - Do not communicate by sharing memory, instead share memory by communicating

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	// go greeter("Hello") //if sleep is not given than only world will be printed
	// greeter("World")
	websiteList := []string{
		"https://lco/dev",
		"https://github.com",
		"https://fb.com",
	}
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1) // 1=no. of goroutines
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(5 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
