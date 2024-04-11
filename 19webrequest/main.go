package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
const url = "https://www.google.com"

func main()  {
	fmt.Println("Hello, Welcome To Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is: %T\n", response)

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(data)

	fmt.Println("Content of the URL is: \n", content)
}