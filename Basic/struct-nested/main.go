package main

import "fmt"

type Person struct {
	name string
	age  int
	add  Address
}

type Address struct {
	city    string
	state   string
	pincode int
	area    Area
}

type Area struct {
	street       string
	buildingName string
	roomNo       int
}

func main() {
	// Example 1
	var person1 Person
	person1.name = "santosh"
	person1.age = 20

	person1.add.city = "mumbai"
	person1.add.state = "maharastra"
	person1.add.pincode = 401105

	person1.add.area.street = "Thakur house street"
	person1.add.area.buildingName = "Thakur"
	person1.add.area.roomNo = 103

	fmt.Println(person1)

	// Example 2
	person2 := Person{
		name: "naresh",
		age:  10,

		add: Address{
			city:    "mumbai",
			state:   "Maharastra",
			pincode: 4022023,

			area: Area{
				street:       "Thakur complex Stree",
				buildingName: "Thakur building",
				roomNo:       3214,
			},
		},
	}

	fmt.Println(person2)
}
