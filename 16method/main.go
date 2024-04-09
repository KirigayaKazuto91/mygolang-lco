package main

import "fmt"

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) getStatus() {
	fmt.Println("Is user active? : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "newemail"
	fmt.Println("Email of this user is : ", u.Email)
}

func main()  {
	fmt.Println("Welcome to Structs in Go")

	kasa := User{"Kasa", "kasa@go.dev", true, 25}
	fmt.Println(kasa)
	fmt.Printf("User is %+v\n", kasa)
	fmt.Printf("Name is %v and Email is %v\n", kasa.Name, kasa.Email)

	kasa.getStatus()

	kasa.NewMail()
}