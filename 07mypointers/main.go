package main

import "fmt"

func main()  {
	fmt.Println("Pointers in Go")

	// var ptr *int
	// fmt.Println("value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("value of actual pointer is", ptr)
	fmt.Println("value of actual pointer is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("value of actual pointer is", *ptr)
}