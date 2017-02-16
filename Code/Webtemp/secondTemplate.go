package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Templates ... Creates templates
var (
	dir, tempHome, tempLogin *template.Template
)

// Directory ... Location for all templates
func Directory() {
	dir, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatal("Parsing templates", err)
	}
	tempHome = dir.Lookup("index.html")
	tempLogin = dir.Lookup("login.html")
}

// HomePage ... home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tempHome.Execute(w, nil)
	}
}

// LoginPage ... login page
func LoginPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := fmt.Sprint(r.FormValue("user"))
	pass := fmt.Sprint(r.FormValue("password"))

	if user == "harold" && pass == "ramos" {
		tempLogin.Execute(w, nil)
	} else {
		tempHome.Execute(w, nil)
	}
}

func main() {
	Directory()
	log.Println("App is running in port :8000")
	http.HandleFunc("/login/", LoginPage)
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
