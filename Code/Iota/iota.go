// Copyright 2016 Harold Ramos.

package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 2 << (2 * 10)
)

func main() {
	fmt.Println(KB)        // Decimal 1024
	fmt.Printf("%b\n", KB) // Binary 10000000000
	fmt.Println(MB)        // Decimal 1048576
	fmt.Printf("%b\n", MB) // Binary 100000000000000000000
}
