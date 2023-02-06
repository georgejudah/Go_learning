package main

import (
	"fmt"
)

func main() {
	//way 1- declaration
	// var colors map[string]string

	//way2
	// colors := make(map[string]string)
	// keys are string and values are string as well

	//another way
	colors := map[string]string{
		"red":   "#FF0000",
		"blue":  "#0000FF",
		"white": "#fffff",
	}

	//add key value pair
	// colors["white"] = "#fffff"
	//
	//accessing individual keys
	// structName.white

	//delete
	// delete(colors, "white")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(input_map map[string]string) {
	for color, hex := range input_map {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
