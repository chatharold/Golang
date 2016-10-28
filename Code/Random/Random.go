package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	Random number in this case from 0 till 9
*/

func main() {

	start := rand.NewSource(time.Now().UnixNano())

	pass := rand.New(start)

	out := pass.Intn(10)

	fmt.Println(out)
}
