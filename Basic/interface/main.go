package main

import "fmt"

// Define an interface
type Animal interface {
	Speak() string
}

// Define a type Dog
type Dog struct {
}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Define a type Cat
type Cat struct {
}

// Implement the Speak method for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Function that takes an Animal and prints its sound
func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	MakeAnimalSpeak(dog) // Outputs: Woof!
	MakeAnimalSpeak(cat) // Outputs: Meow!
}
