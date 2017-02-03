package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue // Skips the number 5
		} else if i == 9 {
			break
		}

		fmt.Println(i)
	}
}
