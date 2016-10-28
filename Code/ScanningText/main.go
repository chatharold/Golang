package main

import (
	"fmt"
	"github.com/goscan"
)

func main() {
	fmt.Print("Enter full name: ")
	name := goscan.Text()
	fmt.Println(name)
}
