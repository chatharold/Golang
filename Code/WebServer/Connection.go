package main

import (
	"fmt"
	"net/http"
)

/*
	Go to localhost:3000
*/

func start(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Welcome")
}

func main() {

	http.HandleFunc("/", start)
	http.ListenAndServe(":3000", nil)

}
