// Harold Ramos

package main

import (
	"fmt"
)

func modify(ten *int) {
	*ten = 30 // changing from 10 to 30
}

func main() {
	ten := 10                        // value = 10
	fmt.Println("From main", ten)    // printing value = 10
	modify(&ten)                     // passing the address of 10
	fmt.Println("From pointer", ten) // printing value = 10 after modifying value to 30
}
