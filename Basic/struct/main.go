package main

import "fmt"

type Person struct {
	name string
	age  int
	add  string
}

func main() {
	// Example 1
	var person1 Person
	person1.name = "santosh"
	person1.age = 20
	person1.add = "test Address"

	fmt.Println(person1)
	fmt.Println(person1.name)

	// Example 2
	person2 := Person{
		name: "naresh",
		age:  19,
		add:  "Test Add",
	}
	fmt.Println(person2)

	// Example 3
	// var person3 = new(Person)
	person3 := new(Person)
	person3.name = "rajeshwari"
	person3.age = 22
	person3.add = "Test Address"

	fmt.Println(person3)
}
