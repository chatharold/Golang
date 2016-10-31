// Copyright 2017 Harold Ramos.
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

func serve(res http.ResponseWriter, req *http.Request) {
	// Template
       	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err, "Creating template...")
	}	
	// Form
	req.ParseForm()

	// Getting fields
	user := fmt.Sprint(req.FormValue("username"))
	pass := fmt.Sprint(req.FormValue("password"))
	
	// Print fields when they have data
	if user != "" && pass != "" {
	fmt.Println("Username =", user, "-", "Password =", pass)
	}
	temp.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8000", nil)
}

