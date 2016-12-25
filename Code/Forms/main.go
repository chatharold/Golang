// Copyright 2016 Harold Ramos

package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var syn sync.WaitGroup

var (
	num = 0
)

// Handler
func home(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln("Creating template", err)
	}

	// ... Getting the form ... //
	req.ParseForm()
	for _, getForm := range req.Form {
		value := fmt.Sprint(getForm[0])
		if value != "Send" && value != "" && req.Method == "POST" {
			t := time.Now().Hour() - 12
			m := time.Now().Minute()
			s := time.Now().Second()
			fmt.Println()
			fmt.Println("************************************************")
			fmt.Printf("Order# %d %s was ordered at: %d:%d:%d\n", num+1, value, t, m, s)
			fmt.Println("************************************************")
			// Incrementor
			num++
		}
	}
	tpl.Execute(res, nil)
}

// ... Web Server ... //
func serveWeb() {
	time.Sleep(time.Millisecond * 10)
	http.HandleFunc("/", home)
	log.Fatalln(http.ListenAndServe(":8000", nil))
	syn.Done()
}

// ... Network Server ... //
func serveNet() {
	time.Sleep(time.Millisecond * 20)
	ls, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalln("Listener...", err)
	}
	defer ls.Close()
	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatalln("Connecting ...", err)
		}
		io.WriteString(conn, "CONNECTED ...")
		conn.Close()
		syn.Done()
	}

}

func main() {
	time.Sleep(time.Millisecond * 30)
	log.Println("Server is running ...")
	syn.Add(2)
	go serveWeb()
	go serveNet()
	syn.Wait()
}
