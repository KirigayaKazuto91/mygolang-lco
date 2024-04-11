package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://courses.learncodeonline.in/learn/search?show=all&type=100&search=golang"

func main()  {
	fmt.Println("Welcome to Handling Url in Golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Port: ", result.Port())
	fmt.Println("Query: ", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type: %T\n", qparams)

	fmt.Println(qparams)

	for key, val := range qparams {
		fmt.Printf("Key: %v, Value: %v\n", key, val)
	}

	partsOfUrl := &url.URL {
		Scheme: "https",
		Host: "courses.learncodeonline.in",
		Path: "/learn/search",
		RawQuery: "show=all&type=100&search=golang",
	}

	fmt.Println(partsOfUrl.String())
}