package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tempDir   *template.Template
	tempHome  *template.Template
	tempLogin *template.Template
)

// Locates all the files and routes them
func Directories() {
	tempDir, err := template.ParseGlob("./templates/*.html") // All files inside directory
	if err != nil {
		log.Fatalln("Creating err", err)
		os.Exit(1)
	}
	tempHome = tempDir.Lookup("home.html") 	 // Looks in the directory for file: home.html
	tempLogin = tempDir.Lookup("login.html") // Looks in the directory for file: login.html
}
// login handler
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tempLogin.Execute(w, nil) // calls tempLogin file called login.html
	}
}
// start handler
func start(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tempHome.Execute(w, nil) // calls tempHome file called home.html
	}
}

func main() {
	Directories()
	log.Println("App is running ...")
	http.HandleFunc("/login", login)
	http.HandleFunc("/", start)
	http.ListenAndServe(":8000", nil)
}
