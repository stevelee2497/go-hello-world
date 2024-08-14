package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")

	// Get a greeting message and print it.
	message, err := greetings.Hello("John")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Stark", "Lainester", "Baratheon"}
	messages, err := greetings.HelloEveryone(names)
	if err != nil {
		log.Fatal(err)
	}
	for name, message := range messages {
		fmt.Println(name, "=>", message)
	}
}
