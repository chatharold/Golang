// Copyright 2016 Harold Ramos

package main

import "fmt"

func main() {
	var name interface{} = "harold"
	nameValue, cond := name.(string) // nameValue gets name(value) && cond(if is a string or not)
	if cond {
		fmt.Println("String", nameValue) // String harold
	} else {
		fmt.Println("Number")
	}
}
