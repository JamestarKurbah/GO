package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// program to return multiple greeting message using slice

	// create a slice of names
	names := []string{"Glady", "Roys", "Neha"}
	// Get a greeting message and print it.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
