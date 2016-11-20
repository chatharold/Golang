package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Reflect type
	name := "juan"
	fmt.Println(reflect.TypeOf(name))
}
