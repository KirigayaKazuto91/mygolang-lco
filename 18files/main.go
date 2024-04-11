package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Hello, Welcome To File Handling")

	content := "This Need to go in a file - Learning Golang"

	file, err := os.Create("./myFile.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length of the file is: ", length)

	defer file.Close()

	readFile("./myFile.txt")
} 

func readFile(filename string){
	data, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Data in the file is: \n", data)
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}