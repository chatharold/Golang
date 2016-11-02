// Copyright © 2017 Harold Ramos.
package main

import (
	"fmt"
)

func show(ch chan string) {
	// Sending multiple info to one channel
	ch <- "channel 1"
	ch <- "channel 2"
}

func main() {
	// Creating a channel
	ch := make(chan string)

	// Go routine
	go show(ch)

	// Prints out both channels
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
