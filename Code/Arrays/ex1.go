// Arrays
package main

import "fmt"

func main() {
	aFirst := [4]int{2, 3, 4, 5}
	fmt.Println(aFirst) // [2 3 4 5]

	aSecond := make([]int, 30)
	aSecond[2] = 20
	fmt.Println(aSecond) // [0 0 20 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

}