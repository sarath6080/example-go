package main

func main() {
	cards := deck{"Ace of hearts", newCard()}
	cards = append(cards, "queen fo spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
