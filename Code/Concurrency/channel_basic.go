package main

import "fmt"

// Creating a function that takes a channel and returns a string to the channel
func say(ch chan string) {
	ch <- "channel outside of main"
}

func main() {
	ch := make(chan string, 1)     // Creating a string channel with a buffer size of 1
	ch <- "channel inside of main" // Sending a string to the channel
	fmt.Println(<-ch)              // Printing the channel

	go say(ch)        // Calling the func say(takes a channel) returns a string
	fmt.Println(<-ch) // Printing the channel

	/*
		channel inside of main
		channel outside of main
	*/
}
