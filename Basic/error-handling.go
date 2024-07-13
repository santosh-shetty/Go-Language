package main

import "fmt"

func divideFunc(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cant be zero!")
	}
	return a / b, nil
}

func main() {
	// Example 1
	ans1, err1 := divideFunc(10, 0)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("Divide Ans is : ", ans1)

	// Example 2
	ans2, err2 := divideFunc(10, 2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("Divide Ans is : ", ans2)

	// Example 1
	ans3, _ := divideFunc(20, 10)
	fmt.Println("Divide Ans is : ", ans3)
}
