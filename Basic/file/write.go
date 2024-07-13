package main

import (
	"fmt"
	"os"
)

func main() {
	// Create File
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while create file : ", err)
	}
	defer file.Close()
	// Write in a file
	content := "This is sample data"
	err1 := os.WriteFile("example.txt", []byte(content), 0644)

	if err1 != nil {
		fmt.Println("Error while write in a file : ", err1)
	}

	fmt.Print("Writing on file successfully")
}
