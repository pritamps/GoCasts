package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

// Define struct
type person struct {
	firstName string
	lastName  string
	// Nested struct
	contactInfo
}

func main() {
	// Note: ORDERING IS IMPORTANT HERE!! Sucks, but such is life.
	me := person{"Pritam", "Sukumar", contactInfo{"pritamps@gmail.com", 560011}}

	fmt.Println("New struct initialized with:", me.firstName, me.lastName)

	// Much better way to define it
	gautam := person{
		firstName: "Gautam",
		lastName:  "Sukumar",
		contactInfo: contactInfo{
			email:   "gautamps@gmail.com",
			zipCode: 7901},
	}
	fmt.Println("New struct initialized with:", gautam.firstName, gautam.lastName, gautam.contactInfo)

	// Can directly print the things
	fmt.Println("Printing struct directly", me)

	// Define a var -- things are AUTOMATICALLY ASSIGNED (strings to "", and so on)
	var randomDude person
	randomDude.firstName = "Pavitra"
	randomDude.lastName = "Prabhakar"
	randomDude.contactInfo = contactInfo{
		email:   "xyz@gmail.com",
		zipCode: 560011}
	randomDude.print()
	// Pass reference to randomDude
	randomDude.updateFirstName("Peter")
	randomDude.print()
}

// Receiver on structs
func (p person) print() {
	fmt.Printf("\nFirst name: %s,\nLast Name: %s,\nContact Info: %v", p.firstName, p.lastName, p.contactInfo)
}

func (p *person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}
