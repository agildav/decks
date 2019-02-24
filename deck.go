package main

import "fmt"

// Create type deck, slice of cards
type deck []string

// Create a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
