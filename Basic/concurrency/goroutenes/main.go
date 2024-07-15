package main

import (
	"fmt"
	"time"
)

func printText(txt string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(txt, " - ", i)
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	go printText("Text 1")
	printText("Text 2")
	time.Sleep(1 * time.Second)
}
