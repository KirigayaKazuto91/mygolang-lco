package main

import "fmt"

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func main()  {
	fmt.Println("Welcome to Structs in Go")

	kasa := User{"Kasa", "kasa@go.dev", true, 25}
	fmt.Println(kasa)
	fmt.Printf("User is %+v\n", kasa)
	fmt.Printf("Name is %v and Email is %v\n", kasa.Name, kasa.Email)
}