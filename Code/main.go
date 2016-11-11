// Copyright 2016 Harold Ramos.

package main

import "fmt"

func get(name ...string) {
	for _, all := range name {
		fmt.Println(all)
	}
}
func main() {
	slice := []string{"harold", "ramos"}
	get(slice...)
}
