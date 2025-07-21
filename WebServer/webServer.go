package main

import (
	"fmt"
	"net/http"
)

func main() {
	// handler
	http.HandleFunc("/", handler)
	http.HandleFunc("/Hello", handler2)
	http.ListenAndServe(":8080", nil)
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my home page!\n")
}
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my hello world page!\n")
}
