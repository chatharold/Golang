package main

import "fmt"

type Products struct {
	name      string
	code      int
	available bool
	inventory int
	cost      float32
	from
}

type from struct {
	company string
}

func main() {
	car := Products{
		"toyota",
		12,
		true,
		4,
		123000,
		from{"japan"},
	}
	fmt.Println("Car model =", car.name)
	fmt.Println("Car bar code =", car.code)
	fmt.Println("Car available =", car.available)
	fmt.Println("Car left =", car.inventory)
	fmt.Println("Car price =", car.cost)
	fmt.Println("Car model =", car.from.company)
}
