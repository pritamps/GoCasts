package main

import (
	"os"
	"testing"
)

// Create new deck and ensure it has correct number of cards
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spaces, but got %s", d[0])
	}
}

func TestFileIO(t *testing.T) {
	// Just in case some files were left over from a previous test
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != len(d) {
		t.Errorf("Unexpected number of cards in new Deck")
	}

	os.Remove("_deckTesting")
}
