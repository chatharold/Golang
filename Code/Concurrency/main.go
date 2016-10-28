package main

import (
	"fmt"
)

func fA(A chan string) {

	A <- "channel...A"

}

func fB(B chan string) {

	B <- "channel...B"

}

func main() {

	A := make(chan string)
	B := make(chan string)

	go fA(A)
	outA := <-A
	fmt.Println(outA)

	go fB(B)
	outB := <-B
	fmt.Println(outB)

	var input string
	fmt.Scanln(&input)

}
