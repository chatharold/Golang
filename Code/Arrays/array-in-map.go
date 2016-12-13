// Harold Ramos

package main

import "fmt"

func main() {
	// ... Using an array inside a map ... //
	am := map[[2]int]string{
		[2]int{1, 2}: "hello",
	}
	fmt.Println(am) // map[[1 2]:hello]
}
