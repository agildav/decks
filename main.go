package main

import "fmt"

func main() {
	cards := newDeck()

	cardsString := cards.toString()

	fmt.Println(cardsString)
}
