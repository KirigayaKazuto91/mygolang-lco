package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to module in GoLang!")
	greet()
	r := mux.NewRouter()
	r.HandleFunc("/", serve).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
	
}

func greet() {
	fmt.Println("Hello, World!")
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World!</h1>"))

	fmt.Println("Serving...")
}