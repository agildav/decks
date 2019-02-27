package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

// Deal a hand
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Print the deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Convert a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save the deck to a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Load a deck from a file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffle a deck
func (d deck) shuffle() {
	// Generate a new source and seed for random
	newSrc := rand.NewSource(time.Now().UnixNano())
	r := rand.New(newSrc)

	for i := range d {
		newPos := r.Intn(len(d))

		d[i], d[newPos] = d[newPos], d[i]
	}
}
