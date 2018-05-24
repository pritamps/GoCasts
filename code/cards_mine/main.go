package main

import (
	"fmt"
)

func main() {
	fmt.Println(sayHello())

	// Slice of type string
	cards := newDeck()

	// Two return values
	hand, remainingDeck := deal(cards, 5)

	fmt.Println("\n === The hand ===")
	hand.print()

	fmt.Println("\n === Remaining Cards ===")

	remainingDeck.print()

	fmt.Println("\n === Deck as string ===")
	fmt.Println(remainingDeck.toString())
	remainingDeck.saveToFile("deck.txt")

	fmt.Println("\n === Deck read from file ===")
	d := newDeckFromFile("deck.txt")
	d.print()

	fmt.Println("\n === Deck Shuffle ===")
	d.shuffle()
	d.print()
}

func sayHello() string {
	return "Starting...\n"
}
