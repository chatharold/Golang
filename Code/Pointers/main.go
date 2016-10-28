package main

import (
	"fmt"
)

func main() {

	var a = 30

	fmt.Println()
	fmt.Println("Address of a =", &a)
	fmt.Println()

	aPointer := &a
	fmt.Println("Pointers address =", aPointer)
	fmt.Println()

	if &a == aPointer {
		fmt.Println("They have the same address")
	} else {
		fmt.Println("Not the same")
	}
	fmt.Println()

	fmt.Println("'A' value is =", a)
	fmt.Println()

	fmt.Println("Pointers value is also =", *aPointer)
	fmt.Println()

	*aPointer = 40
	fmt.Println("Assigning a value to the pointer =", *aPointer)
	fmt.Println()

	/*
		Address of a = 0xc82000a268

		Pointers address = 0xc82000a268

		They have the same address

		'A' value is = 30

		Pointers value is also = 30

		Assigning a value to the pointer = 40
	*/
}
