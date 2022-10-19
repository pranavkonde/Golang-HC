package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// encoding := converting data to valid json

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}
func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.lco", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "learncodeonline.lco", "lpo123", []string{"web", "js"}},
		{" Bootcamp", 399, "learncodeonline.lco", "qwe123", nil},
	}
	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") // marshal indent will indent it
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
