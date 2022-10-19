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
	// EncodeJson()
	DecodeJson()
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

// data coming from web is in byte format
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "learncodeonline.lco",
		"Tags": ["web-dev","js"]
	}
	`)
	var lcoCourses course
	checkvalid := json.Valid(jsonDataFromWeb)
	if checkvalid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourses)
		fmt.Printf("%#v\n", lcoCourses)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
