package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome To Switch Case in Golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of Dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and You Can Move 1 Block")
	case 2:
		fmt.Println("Dice Value is 2 and You Can Move 2 Blocks")
		fallthrough
	case 3:
		fmt.Println("Dice Value is 3 and You Can Move 3 Blocks")
	case 4:
		fmt.Println("Dice Value is 4 and You Can Move 4 Blocks")
	case 5:
		fmt.Println("Dice Value is 5 and You Can Move 5 Blocks")
	case 6:
		fmt.Println("Dice Value is 6 and You Can Move 6 Blocks and Roll the Dice Again")

		default:
			fmt.Println("Invalid Dice Value")
	}
}