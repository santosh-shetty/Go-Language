package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Add  string `json:"address"`
}

func main() {
	person1 := Person{
		Name: "santosh",
		Age:  18,
		Add:  "test address",
	}
	fmt.Println("Normal person 1  data : ", person1)

	// Marshal Json
	res, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error During json Marshal : ", err)
	}
	fmt.Println("Marshell person 1  data : ", string(res))

	// unMarshal Json
	person2 := `{
		"name": "naresh",
		"age": 16,
		"address": "test address"
	}`

	var person2Data Person
	err1 := json.Unmarshal([]byte(person2), &person2Data)
	if err1 != nil {
		fmt.Println("Error During json UnMarshal : ", err1)
	}
	fmt.Println("unMarshell person  2 data : ", person2Data)

}
