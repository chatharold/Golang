package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tempDir  *template.Template // Directory where all html pages are located.
	tempHome *template.Template // Static page
)

// Template Page
func templatePage() {
	tempDir, err := template.ParseGlob("./templates/*.html") // All pages
	if err != nil {
		log.Fatalln("Parsing template", err)
		os.Exit(1)
	}
	tempHome = tempDir.Lookup("home.html") // Static template
}

// Login Page
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Login page:</h1>")
}

//  Start Handler
func Start(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tempHome.Execute(w, nil)
	}
}

func main() {
	templatePage() // Executing template
	log.Println("App is running on port: 8000")
	http.HandleFunc("/login", Login) // Login handler
	http.HandleFunc("/", Start)      // Start handler
	http.ListenAndServe(":8000", nil)
}
