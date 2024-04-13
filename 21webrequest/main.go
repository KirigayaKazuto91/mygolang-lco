package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Welcome to Web Request in Golang!")
	// GetRequest()
	PostRequest()
}

func GetRequest() {
 const url = "http://localhost:8000/get"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Status Code Is: ", response.StatusCode)
	fmt.Println("Content Length Is: ", response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	bytecount, _ := responseString.Write(content)

	fmt.Println("Bytecount is :", bytecount)
	fmt.Println(responseString.String())


	
	// fmt.Println(string(content))

	defer response.Body.Close()
}

func PostRequest(){
	const url = "http://localhost:8000/post"

	Body := strings.NewReader(`{"name": "John Doe", "age": 25, "email": "John Doe"}`)

	response, err := http.Post(url, "application/json", Body)	
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Status Code Is: ", response.StatusCode)
	fmt.Println("Content Length Is: ", response.ContentLength)
	content , _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response Body Is: ", string(content))
}