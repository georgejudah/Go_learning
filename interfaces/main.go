package main

//here, this code will not run. But you can understand that there is a similar logic between the English Bot and Spanish Bot
//life would be much easier if the englishBot and the spanishBot can be passed to a single function instead of having two

// what if you have inputs of different datatypes. Would that mean you have to create different functions for different data types??
import "fmt"

type englishBot struct{}
type SpanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting
	return "Hi There!"
}

func (SpanishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting
	return "Hola"
}

func printGreeting(eb englishBot) {
	fmt.Println((eb.getGreeting()))
}

func printGreeting(eb SpanishBot) {
	fmt.Println((sb.getGreeting()))
}
