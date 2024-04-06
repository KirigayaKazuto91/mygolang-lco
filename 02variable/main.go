package main

import "fmt"

// public variable
const LoginToken string = "asjdhjashdjashd"

func main() {
	var username string = "John Doe"
	fmt.Println("Hello, ", username)
	fmt.Printf("Variable is Type Of %T \n", username)

	var isLoggedIn bool = false
	fmt.Println("Is Logged In? ", isLoggedIn)
	fmt.Printf("Variable is Type Of %T \n", isLoggedIn)

	// shorthand variable
	anotherVariable := "I'm Another Variable"
	fmt.Println(anotherVariable)

	fmt.Println("Public Variable ", LoginToken)
}