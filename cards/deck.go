package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Code that describes what a deck is and how it works
// Create a new type of deck which is a slice of string

// this is a custom type in deck
type deck []string

// not oops- no Oops in go, we are rather doing something similar here creating a newDeck function with a custom type deck and based on the type, assigning it's own functions
func newDeck() deck {
	// cards := deck{"Ace of Spades", "Two of Spades"}
	cards := deck{}
	cardSuits := []string{"Spades", "diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// for using _ instead of a var char - it suppresses the warning that the variable is unused/us saying that we don't care about that variable
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// function which loops throught the deck of cards

// receiver on a function/reciver functions - which means any variable of type deck gets access to print() function
// d is the actual copy of the deck(instance of deck or reference of deck)
// by convention we use 1 letter or lettter - for the instance of deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	//returning multiple values - of type (deck, deck)
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//["red", "yellow", "blue"] = "red,yellow,blue"
	return strings.Join([]string(d), ",")
}

// saves the deck of cards to a file. Takes input param - filename
func (d deck) saveToFile(filename string) error {
	err := os.WriteFile(filename, []byte(d.toString()), 0666) //read and write file permission
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func newDeckFromFile(filename string) deck {
	data_byte, err := os.ReadFile(filename)
	//if no data, err might return nil
	if err != nil {
		//Option #1 - log the error and return a call to newDeck() - we are giving them a new deck incase no data from the harddrive
		//Option #2 - log the error and entirely quit the program
		log.Fatal(err)
		//Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.
		os.Exit(1)
	}
	//convert byte type to string type
	data_string := string(data_byte)
	//string slice ie deck
	data_string_slice := strings.Split(data_string, ",")
	return deck(data_string_slice)
}

// shuffle a deck of cards
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		// newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1)
		//swapping
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
