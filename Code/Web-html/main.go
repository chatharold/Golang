// Adding two numbers using html \\
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func server(res http.ResponseWriter, req *http.Request) {

	// ***** Creating template  ***** \\

	temp, err := template.ParseFiles("index.gtpl")
	if err != nil {
		log.Fatalln(err)
	}

	// ***** Form info ***** \\

	req.ParseForm()

	var numF, numS int

	for index, value := range req.Form {
		fmt.Println()
		fmt.Println("---------------------")
		fmt.Println("Input=", index, value)
		fmt.Println("---------------------")

		switch index { // ***** Reads the index value (first - second) ***** \\
		case "first":
			strF := value[0]
			numF, _ = strconv.Atoi(strF)
			fmt.Println("Number =", numF)
		case "second":
			strS := value[0]
			numS, _ = strconv.Atoi(strS)
			fmt.Println("Number =", numS)
		}
	}

	temp.Execute(res, numF+numS)
}

func main() {
	log.Println("App is Running...")
	http.HandleFunc("/", server)
	http.ListenAndServe(":8000", nil)
}
