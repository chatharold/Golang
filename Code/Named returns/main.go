package main

import "fmt"

func calc(a, b int) (add, mul, div, sub int) {

	add = a + b
	mul = a * b
	div = a / b
	sub = a - b
	return
}

func main() {

	sum, mul, div, sub := calc(10, 2)
	fmt.Println("Sum =", sum)
	fmt.Println("Mul =", mul)
	fmt.Println("Div =", div)
	fmt.Println("Sub =", sub)
}
