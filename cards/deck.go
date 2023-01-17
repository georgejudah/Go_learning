package main

import "fmt"

// Code that describes what a deck is and how it works
// Create a new type of deck which is a slice of string

type deck []string

// function which loops throught the deck of cards

// receiver on a function - which means any variable of type deck gets access to print() function
// d is the actual copy of the deck(instance of deck or reference of deck)
// by convention we use 1 letter or lettter - for the instance of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
