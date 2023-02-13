package main

import (
	"fmt"
	"io"
	"os"
)

//GOAL
// Create a program that reads the context of a text file and prints its contents to the terminal
//The file to open should be provided as an argument to the program when it is executed at the terminal.
// For example 'go run main.go myfile.txt' should open up the myfile.txt

//To read the arguements provided to a program, you can reference the variable 'os.Args', which is a
// slice of type string

// To open a file, check out the doc for the 'Open' function in the 'os' package
//what interfaces does the 'File' type implement?
//If the 'File' type implements the 'Reader' interface, you might be able to reuse that io.Copy function!

func main() {
	// Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	f, err := os.Open(argsWithoutProg) // for read access

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
