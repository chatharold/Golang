package main

import "fmt"

func pass(v ...string) {
	fmt.Println(v)
	for _, get := range v {
		fmt.Println(get)
	}
}

func main() {
	pass("harold", "ramos")
}
