package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to Slices in Go")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Manggo", "Banana")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 123
	highScore[1] = 999
	highScore[2] = 888
	highScore[3] = 787
	
	highScore = append(highScore, 862,865,234)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	
}