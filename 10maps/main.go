package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Maps in Go")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println(languages)
	fmt.Println("short of RB is",languages["RB"])

	delete(languages, "RB")
	fmt.Println(languages)

	// looping over maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}