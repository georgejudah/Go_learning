package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting the program")
	// remember no bracket around for loop
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for i := 0; i <= 10; i++
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println("The number", number, "is even")
		} else {
			fmt.Println("The number", number, "is odd")
		}

	}

}
