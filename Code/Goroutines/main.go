package main

import (
	"fmt"
	"runtime"
)

// Go routines
func say(str string) {
	for i := 0; i <= 5; i++ {
		runtime.Gosched()
		fmt.Println(str)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*
hello
world
hello
world
hello
world
hello
world
hello
world
hello
world
*/
