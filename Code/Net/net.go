package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// IP ... stdout the ip info
func IP() string {
	ip, err := exec.Command("ifconfig").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(ip)
}

// HomePage ...handles home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, IP())
}

// PageHandler ... does the routing
func PageHandler() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Scanner ... scans the stdin
func Scanner() {
	// Scans from terminal
	var command string
	fmt.Print("Enter command: ")
	fmt.Scan(&command)
	fmt.Print("Printing ifconfig in localhost:8000")

	// Decision ifconfig
	switch command {
	case "ifconfig":
		PageHandler()
	default:
		fmt.Println("not found")
	}
}

func main() {
	Scanner()
}
