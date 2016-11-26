// Copyright 2016 Harold Ramos

package main

import "fmt"

func main() {
	// ranging
	num1 := []int{1, 2, 3, 4, 5}
	for _, pos := range num1 {
		fmt.Println(pos) // 1, 2, 3, 4, 5
	}
	// Divider
	fmt.Println("---------------")
	// iteraring
	for i := 0; i < len(num1); i++ {
		fmt.Println(num1[i]) // 1, 2, 3, 4, 5
	}
}
