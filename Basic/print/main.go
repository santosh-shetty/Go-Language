package main

import "fmt"

// =========== Declaration ==========
const (
	first = iota + 1
	second
	third
)

// =========== Main Function ============
func main() {
	fmt.Println("!--------- Hello World-------!")
	fmt.Println(first, second, third)
	data(first, second, third)
}

// =========== Custom Function ============
func data(a int, b int, c int) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// =========== Custom Function ============
