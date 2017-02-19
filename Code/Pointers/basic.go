/*
  The basics about pointers:
  1- Variables contain a memory address and a value,
  in order to change the value you have to change the address first
  and you do so by referencing the address to a new variable.
*/

package main

import "fmt"

func main() {
	number := 1

	// Getting the address && value of the number variable
	fmt.Println("The number's address is =", &number)
	fmt.Println("The number's value is =", number)

	// Creating a pointer to number
	changed := &number // Gets the address of number && assigns it to changed
	*changed = 1000    // Changes the value of number

	// Printing the & of changed which is == & of number
	fmt.Println(&changed)
	fmt.Println(number) // Prints the value of number == 1000

	/*
		The number's address is = 0xc42000e220
		The number's value is = 1
		0xc42000c030
		1000
	*/
}
