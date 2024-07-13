package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter your Name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Print("Your name is : ", name)
}
