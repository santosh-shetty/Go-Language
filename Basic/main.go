package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Enter Your Name : ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("Let's Start Quiz ", name)

	// Question 1
	fmt.Println("A. Which of the following colors contain equal amounts of RBG?")
	fmt.Println("	1 - White")
	fmt.Println("	2 - Gray")
	fmt.Println("	3 - Black")
	fmt.Println("	4 - All of the above")

	// reader := bufio.NewReader(os.Stdin)
	var ans1 int
	fmt.Scan(&ans1)

	// Question 2
	fmt.Println("B. What is the correct syntax to write an HTML comment?")
	fmt.Println("	1 - <!-- Comment -->")
	fmt.Println("	2 - // Comment")
	fmt.Println("	3 - # Comment")
	fmt.Println("	4 - /* Comment */")
	var ans2 int
	fmt.Scan(&ans2)

	// Check the answer
	rightAns := 0
	if ans1 == 4 {
		rightAns++
	}

	if ans2 == 1 {
		rightAns++

	}

	fmt.Println(rightAns, " Answers is correct Out of 2")
}
