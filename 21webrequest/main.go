package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main()  {
	fmt.Println("Welcome to Web Request in Golang!")
	GetRequest()
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