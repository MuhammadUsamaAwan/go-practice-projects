package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/healthcheck", heathCheckHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started on port 5000")

	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v\n", err)
		return
	}
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name: %s\n", name)
}

func heathCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello from Go!")
}
