package main

import (
	"fmt"
)

/*
	In this example we are modifying [a byte] in a string
	by converting it into a rune
	and then back into a string again.
*/

func main() {

	var name string = "harold" // Strings

	letter := []rune(name) // From string to rune

	letter[0] = 't' //  Modify rune

	out := string(letter) // From rune to string

	fmt.Println(out) // out

}
