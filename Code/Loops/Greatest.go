// Harold Ramos

package main

import "fmt"

func main() {
	// ... Slice of ints ... //
	values := []int{
		1, 59, 0,
		2, 5, 7,
	}
	// ... Get the first digit ... //
	least := values[0]
	// ... Loop through all the numbers
	// ... and if you find the smallest print it ... //
	for _, loop := range values {
		if loop < least {
			least = loop
		}
	}
	fmt.Println(least) // 0
}
