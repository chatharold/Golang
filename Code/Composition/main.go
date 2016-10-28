// Copyright by Harold Ramos 2016.
package main

import (
	"fmt"
)

type computer struct {
	model string
	year int
	on string
	off string
}
	
func (c computer) On() string {
	c.on = "ON"
	return c.on
}

func (c computer) Off() string {
	c.off = "OFF"
	return c.off
}

type Switch interface {
	On() string
	Off() string
}

func flipOn(s Switch) string {
	return s.On()
}

func main() {
	apple := computer{"Macbook", 2016, "", ""}
	fmt.Println("Model", apple.model)
	fmt.Println("Year", apple.year)

	fmt.Println("Turned", apple.On())
	fmt.Println("Turned", apple.Off())
	
	fmt.Println("Calling the interface", Switch.Off(apple)) // Interface - method(Object)
	fmt.Println("Calling the function", flipOn(apple)) // func - object  creates a function making the code more abstract.
}