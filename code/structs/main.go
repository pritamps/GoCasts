package main

import (
	"fmt"
)

// Define struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	// Note: ORDERING IS IMPORTANT HERE!! Sucks, but such is life.
	me := person{"Pritam", "Sukumar"}
	fmt.Println("New struct initialized with:", me.firstName, me.lastName)

	// Much better way to define it
	gautam := person{firstName: "Gautam", lastName: "Sukumar"}
	fmt.Println("New struct initialized with:", gautam.firstName, gautam.lastName)

	// Can directly print the things
	fmt.Println("Printing struct directly", me)

	// Define a var -- things are AUTOMATICALLY ASSIGNED (strings to "", and so on)
	var randomDude person
	randomDude.firstName = "Pavitra"
	randomDude.lastName = "Prabhakar"
	fmt.Printf("%+v", randomDude)
}
