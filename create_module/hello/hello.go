package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)


	// SINGLE
	// message, err := greetings.Hello("Ertan")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println(message)

	// MULTIPLE

	names := []string{"Ertan", "Behice", "Deniz"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
