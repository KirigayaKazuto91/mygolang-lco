package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to Functions in Golang")

	result := adder (3, 5)
	fmt.Println("Result is:", result)

	proResult, myMessage := proAdder(3, 5, 7, 9)
	fmt.Println("Pro Result is:", proResult)
	fmt.Println("Message is:", myMessage)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder (values ...int)( int, string){
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum, "Pro Result"

}


func greeter() {
	fmt.Println("Hello World!")
}