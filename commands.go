package main

import (
	"os"
	"fmt"
)

func commandExit () error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp () error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, com := range getCommands() {
		fmt.Printf("%v: %v\n", com.name, com.description)
	}
	return nil
}

