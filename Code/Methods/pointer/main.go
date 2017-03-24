// Method with a pointer gives you access to the fields inside the method p.Name
package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Info() {
	p.Name = "harold ramos"
}

func main() {
	harold := Person{}
	fmt.Println(harold.Name) // prints an empty value

	harold.Info()            // when calling the method Info() you are able to use the field values inside because of the pointer
	fmt.Println(harold.Name) // harold ramos
}
