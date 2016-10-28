package main

import (
	"fmt"
)

func main() {

	var name1 string = "harold"    // Explicit
	var name2 = 12                 // Implicit type not defined
	name3 := "car"                 // Short Hand
	num1, num2, num3 := 12, 13, 14 // Multiple short hand

	fmt.Println(name1, name2, name3)
	fmt.Println(num1, num2, num3)
}
