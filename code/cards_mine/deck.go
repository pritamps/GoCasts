package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new custom type
type deck []string

// Create and return a deck of playing cards
func newDeck() deck {
	cards := deck{}

	// Define suits and values
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Create a receiver function of type 'deck'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Create a receiver function to change to string
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) shuffle() {

	// Create RNG with source
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// Swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Handle error
		fmt.Println("==== ERROR ERROR ERROR ====")
		fmt.Println(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
