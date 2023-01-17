package main

func main() {
	// assigning a slice, can contain only same type of data
	// cards := []string{"Ace of diamonds", newCard()}
	cards := deck{"Ace of diamonds", newCard()}
	// append creates a new slice, does not actually append to the current slice
	cards = append(cards, "Six of Spades")

	cards.print()
	// for loop
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// assigning variables
	// var card string = "Ace of Spades"
	// Static type language
	// new var, name of the card, only string will be assigned to this variable = Assign the value
	// alternate way of assigning variables similar to above
	// card := "Ace of Spades"
	// := syntax only when you're initialisation of the variable

	// updating the variable
	// card = "Five of diamonds"

	// if package is going to be main, func should be main as well

}

// string mentioned in the function declaration says explicitly what data type we are going to return
func newCard() string {
	return "Five of diamonds"
}
