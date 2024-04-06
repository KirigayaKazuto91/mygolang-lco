package main

import "fmt"

func main()  {
	fmt.Println("Welcome to Arrays in Go")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is", fruitList)
	fmt.Println("Lenght of the array is", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Onion"}
	fmt.Println("Veg list is", vegList)

}