package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// Read in a file
	content, err1 := ioutil.ReadFile("example.txt")

	if err1 != nil {
		fmt.Println("Error while write in a file : ", err1)
	}
	fmt.Printf("%s", content)
}
