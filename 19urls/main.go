package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456"

func main() {
	fmt.Println("Welcome to handling URLs in Golang")
	fmt.Println(myurl)

	// Parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //lco.dev:3000
	fmt.Println(result.Path)     // /learn
	fmt.Println(result.Port())   // 3000
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=ghbj456

	qparams := result.Query()
	fmt.Printf("The type of query params are %T\n", qparams) // url.Values (i.e Key,Value pair)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Params is: ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
