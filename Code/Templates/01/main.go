package main

import (
	"fmt"
)

func main() {
	// Passing name variable
	name := "Harold Ramos"

	// Template
	tpl := `
	<!doctype html>
	<html>
    	<head></head>
	<body>
	<h1> Welcome: ` + name + `</h1>
	</body>
	</html>
	`

	// Output
	fmt.Println(tpl)

	/*
			Result:
			<!doctype html>
		        <html>
		        <head></head>
		        <body>
		        <h1> Welcome: Harold Ramos</h1>
		        </body>
		        </html>
	*/
}
