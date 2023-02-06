package main

//refactor to have a single function instead of two
import "fmt"

type bot interface {
	getGreeting() string
}

//by adding the above code, to whom ever it may concern, we have a new interface called bot. If you are a type
// in this program with a function called 'getGreeting' and you return a string then you are a honorary member
// of type 'bot'.

// Now that you're also an honorary member of type 'bot', you can now call this funtion called 'printGreeting'

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//listener function, if you're not using the instance inside the function, you need not have it
func (englishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting
	return "Hi There!"
}

//listener function
func (spanishBot) getGreeting() string {
	//VERY custom logic for generating a spanish greeting
	return "Hola"
}

// func printGreeting(eb englishBot) {
// 	fmt.Println((eb.getGreeting()))
// }

// func printGreeting(eb SpanishBot) {
// 	fmt.Println((sb.getGreeting()))
// }
