package main

import (
	"fmt"
	"testing"
)

// Test if a deck gets created
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Have all 16 cards?
	const nCrds int = 16
	if len(d) != nCrds {
		errMsg := fmt.Sprintf(":: expected length %d, got %d ::\n", nCrds, len(d))
		t.Errorf(colorError(errMsg))
	}

	// Have the 7th card named Four of Diamonds?
	const card string = "Four of Diamonds"
	const idx int = 7
	if d[idx] != card {
		errMsg := fmt.Sprintf(":: expected %s, got %s ::\n", card, (d[idx]))
		t.Errorf(colorError(errMsg))
	}
}
