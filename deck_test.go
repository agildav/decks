package main

import (
	"fmt"
	"os"
	"testing"
)

// Validate a deck
func validateDeck(d deck, t *testing.T) {
	// Have all 16 cards?
	const nCrds = 16
	if len(d) != nCrds {
		errMsg := fmt.Sprintf(":: expected length %d, got %d ::\n", nCrds, len(d))
		t.Errorf(colorError(errMsg))
	}

	// Have the 7th card named Four of Diamonds?
	const card = "Four of Diamonds"
	const idx int = 7
	if d[idx] != card {
		errMsg := fmt.Sprintf(":: expected %s, got %s ::\n", card, (d[idx]))
		t.Errorf(colorError(errMsg))
	}
}

// Test if a deck gets created
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Validate the deck
	validateDeck(d, t)
}

// Test if a deck gets saved / load - to / from a file
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	const filename = "./_deck_test.txt"
	// Remove previous file
	err := os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	newDeck := newDeckFromFile(filename)

	// Validate the deck
	validateDeck(newDeck, t)

	// Remove file
	err = os.Remove(filename)
	if err != nil {
		errMsg := fmt.Sprintf(":: could not remove file %s ::\n", filename)
		t.Errorf(colorError(errMsg))
	}
}
