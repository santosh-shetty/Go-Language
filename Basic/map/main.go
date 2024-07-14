package main

import (
	"fmt"
)

func main() {

	// Example 1
	studGrades := make(map[string]int)

	studGrades["santosh"] = 90
	studGrades["ram"] = 100
	studGrades["sham"] = 30
	studGrades["hari"] = 60
	studGrades["Viaya"] = 80

	delete(studGrades, "ram")
	for name, marks := range studGrades {
		fmt.Println(name, " marks is ", marks)
	}

	marks, exist := studGrades["naresh"]
	fmt.Println("Naresh Marks : ", marks)
	fmt.Println("Naresh Exist : ", exist)

	// Example 2
	studData := map[string]string{
		"name":    "santosh",
		"surname": "shetty",
		"add":     "test Address",
	}

	fmt.Println(studData)
	for _, v := range studData {
		fmt.Println(v)
	}

}
