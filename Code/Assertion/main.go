package main

import "fmt"

func main() {
	// Assertion: is ony used in interfaces
	// Definition forceful statement

	var name interface{} = "harold"
	_, check := name.(string)
	if check {
		fmt.Println("I am a string")
	} else {
		fmt.Println("something else")
	}
}
