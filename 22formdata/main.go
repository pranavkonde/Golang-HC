package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	PerformPostFormJsonRequest()
}

func PerformPostFormJsonRequest() {
	const myurl = "http://localhost:1111/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "hitesh")
	data.Add("lastname", "konde")
	data.Add("email", "hitesh@dev.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
