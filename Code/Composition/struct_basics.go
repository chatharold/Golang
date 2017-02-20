package main

import "fmt"

// Computer struct ... A Struct contains fields
type Computer struct {
	model  string
	year   int
	action string
}

// Info function ... Functions receives an input and deliver and output
func Info(c Computer) (string, int) {
	return c.model, c.year
}

// On method ... A method does an action
func (c Computer) On() string {
	c.action = "ON"
	return c.action
}

// Turn interface ... An interface is a collection of methods
type Turn interface {
	On() string
}

// Creating a function attached to an interface
func pressButton(t Turn) string {
	return t.On()
}

func main() {
	// Passing values to the struct
	apple := Computer{"MAC", 2017, ""}

	// Printing values from the struct
	fmt.Println("The computer's model is a:", apple.model)
	fmt.Println("It was created in:", apple.year) // Calls the  values from a struct: object.field <- using a dot notation

	// Calling the Info function another way of resuming the info
	fmt.Print("Information about the apple computer: ")
	fmt.Println(Info(apple)) // Call the function and pass the object: function(object)

	// Calling the method
	fmt.Println("The computer is:", apple.On()) // Calls the method object.Method <- using a dot notation

	// Calling the interface
	fmt.Println("Turn the computer: ", Turn.On(apple)) // Calls the interface: interface.Method(object)

	// Calling the function attached to the interface
	fmt.Println("Press button and turn: ", pressButton(apple)) // Calls the function to interface: function(object)

	/*
	  1- For calling Structs and Methods use dot notation
	  2- For calling functions use nameFunction(object)
	  3- For calling interfaces use nameInterface.Method(object)
	*/
}
