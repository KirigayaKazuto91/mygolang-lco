package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Learning Course")
	fmt.Println("===============================")
	presentTime := time.Now()
	fmt.Println("Full Time Format")
	fmt.Println(presentTime)
	fmt.Println("===============================")
	fmt.Println("Time Format Show Year Only")
	fmt.Println(presentTime.Format("2006"))
	fmt.Println("Time Format Show Month Only(Numeric)")
	fmt.Println(presentTime.Format("01"))
	fmt.Println("Time Format Show Month Only(Alfabet)")
	fmt.Println(presentTime.Format("Jan"))
	fmt.Println("Time Format dd-Mmm-yyyy")
	fmt.Println(presentTime.Format("02-Jan-2006"))
	fmt.Println("===============================")
	fmt.Println("Time Format With time dd-Mmm-yyyy, hh-mm-ss")
	fmt.Println(presentTime.Format("02-Jan-2006 15:04:05"))
	fmt.Println("===============================")

	createdTime := time.Date(2023, time.December, 26, 0, 25, 0, 0, time.Local)
	fmt.Println("This is Created Time Format")
	fmt.Println(createdTime)
}