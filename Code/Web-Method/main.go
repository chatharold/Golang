// Copyright 2016 by Harold Ramos

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "<h1>Home page</h1>")
}

func login(res http.ResponseWriter, req *http.Request) {
	temp, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("Template...", err)
	}
	req.ParseForm()

	fmt.Println("Username: ", req.Form["username"])
	fmt.Println("Password: ", req.Form["password"])

	method := fmt.Sprint(req.Method)
	switch method {
	case "POST":
		fmt.Println("Post was used")
	case "GET":
		fmt.Println("Get was used")
	}
	temp.Execute(res, nil)
}

func main() {
	log.Println("App is running")
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8000", nil)
}
