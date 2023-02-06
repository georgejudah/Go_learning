package main

import "fmt"

// define struct
// define what fields a person should have and also specify the data type of the different fields
type person struct {
	firstName string
	lastName  string
	// embedding struct
	//add info - you can specify the field name if you want - for the below contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//creating a new struct
	// way 2 - better way, struct doesn't get afftected if the order of the struct is modified
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	//way 1: alex := person{"Alex", "Anderson"}
	// fmt.Println(alex)
	//way 3
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// % + v will print out all the different field names and their values from Alex

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
		//make sure to use commas properly - every single line must have a comma, even if it is the last property declaration
	}
	//remeber - whenever you pass any data type into a func, Go creates a copu of each arguement,
	//and these copies are used inside of the function.
	//pointers - to update the updated values properly
	// & operator - give me the memory address of the value of the value this variable is pointing at
	// * operator - give me the value this memory address is pointing at
	jimPointer := &jim
	//jim Pointer points to the original address of the jim variable and not the copy
	jimPointer.updateName("Judah")
	jim.print()

}

func (pointerToPerson *person) updateName(newFirstname string) {
	// * in function def - it means we are working with a pointer to a person
	// turn address into value with *address
	// turn value into address with &value
	(*pointerToPerson).firstName = newFirstname
	// the above means we want to manuplate the value the pointer is referencing
}

// receiver function
func (p person) print() {
	fmt.Printf("%+v", p)
}
