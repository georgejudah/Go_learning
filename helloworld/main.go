package main

//types of packages
// Executable - Generates a file which we can run
// Reusable - Code used as 'helpers'
// name of the package defines whether you are building an executable or reusable helpers package

import "fmt"

// format package

func main() {
	// if package is going to be main, func should be main as well
	fmt.Println("Hi There!")
}

// how do we run the code in our project
// 1. go run - compiles and executes one or two files
// 2. go build - just compiles it
// 3. go fmt - formats all the code in each file in the current directory
// 4. go install - Compiles and installs a package
// 5. go get - Downloads the raw sourch code of someone else's package
// go test - Runs any tests associated with the current project

// how the main.go file organized
// package declaration
// import statements for packages
// declare functions, tell Got to do things
