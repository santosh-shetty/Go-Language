package main

import (
	"fmt"
)

func multiplyFun(ch1 chan int) {
	value := <-ch1
	fmt.Println("Received value:", value)
}

func PrintGreet(num int, ch2 chan int) {
	fmt.Println("Run Number : ", num)
	ch2 <- num
}
func main() {

	// // Example  1
	// ch1 := make(chan int)
	// fmt.Println("Before Function Call")
	// go multiplyFun(ch1)
	// ch1 <- 10
	// fmt.Println("After Function Call")
	// close(ch1)
	// // time.Sleep(1 * time.Second)

	// Example 2

	ch2 := make(chan int)
	for i := 1; i <= 5; i++ {
		go PrintGreet(i, ch2)
		v := <-ch2
		fmt.Println("Done:", v)
	}
	close(ch2)

}
