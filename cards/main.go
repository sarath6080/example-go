package main

import "fmt"

func main() {

	cards := []string{"Ace of hearts", newCard()}
	cards = append(cards, "queen fo spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
