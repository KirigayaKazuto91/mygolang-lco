package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"`
	Price int `json:"price"`
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json Data Creation in Golang!")

	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCrouses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js", "react"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js", "react", "node"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	// package this data into json

	finalJson, err := json.MarshalIndent(lcoCrouses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"price": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev", "js", "react"]
		}
	`)

	var lcoCourse course
	
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("Course: %+v\n", lcoCourse)

	} else {
		fmt.Println("JSON was not valid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("Data: %#v\n", myOnlineData) 

	for k, v := range myOnlineData {
		fmt.Printf("Key is: %v and value is: %v and type is %T\n", k, v, v)
	}
	
}

