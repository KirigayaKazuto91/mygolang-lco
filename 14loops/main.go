package main

import "fmt"

func main()  {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days{
	// 	fmt.Printf("Index: %v, Day: %v\n", index, day)
	// }

	rVal := 1

	for rVal < 10 {

		if rVal == 2 {
			goto lco
		}

		if rVal == 5 {
			break
		}

		fmt.Println("Value is: ", rVal)
		rVal++
	}

	lco:
	fmt.Println("Jumping to LCO")
}