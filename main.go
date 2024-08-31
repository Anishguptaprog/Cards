package main

func main() {
	// card := newCard()
	// fmt.Println(card)
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 10)
	// hand.print()
	// remainingCards.print()

	// cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
