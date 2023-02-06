package main

func main() {
	// assigning a slice, can contain only same type of data
	// cards := []string{"Ace of diamonds", newCard()}
	//new type of deck, check the deck.go file
	// cards := deck{"Ace of diamonds", newCard()}
	// append creates a new slice, does not actually append to the current slice
	// cards = append(cards, "Six of Spades")
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")

	//shuffle and print
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// hand, remainingDeck := deal(cards, 5)
	// cards.print()
	// hand.print()
	// remainingDeck.print()

	// for loop
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// updating the variable
	// card = "Five of diamonds"

	// if package is going to be main, func should be main as well

}

// string mentioned in the function declaration says explicitly what data type we are going to return
func newCard() string {
	return "Five of diamonds"
}
