package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func start(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "<h1>Server is Running ...</h1")
	req.ParseForm()       // map[]
	fmt.Println(req.Form) // creates a pointer to the map passing values
	temp, err := template.ParseFiles("folder.gtpl")
	temp.Execute(res, nil)
	if err != nil {
		log.Fatalln("Template", err)
	}
	for k, v := range req.Form {
		fmt.Println(k, v)
		file := os.Mkdir(`/Users/chatharold/Desktop/`+k+``, 0711)
		if file != nil {
			log.Fatalln(err)
		}
	}

}
func main() {
	log.Println("Running")
	http.HandleFunc("/", start)
	http.ListenAndServe(":8000", nil)
}
