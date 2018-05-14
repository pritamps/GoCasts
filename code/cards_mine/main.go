package main

import (
	"fmt"
)

func main() {
	fmt.Println(sayHello())

	// Slice of type string
	cards := []string{"Ace of Diamonds"}
	cards = append(cards, "Six of Spades")
	// fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func sayHello() string {
	return "Starting...\n"
}
