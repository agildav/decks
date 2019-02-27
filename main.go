package main

func main() {
	cards := newDeckFromFile("./my_cards.txt")

	cards.shuffle()
	cards.print()
}
