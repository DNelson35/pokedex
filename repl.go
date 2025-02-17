package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

// strings.Fields would have gave the same result


func cleanInput(txt string) []string {
	var wordArr []string
	if len(txt) == 0 {
		return []string{}
	}
  textArr :=  strings.Split(strings.ToLower(txt), " ")

	for _, word := range textArr {
		if word != ""{
			wordArr = append(wordArr, word)
		}
	}

	return wordArr
}

type cliCommands struct {
	name				string
	description string
	callback 		func(*config) error
}



func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := &config{}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}


func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name: 			 "map",
			description: "Display locations on map",
			callback:  	 commandMap,
		},
		"mapb": {
			name: 			 "mapb",
			description: "display last set of locations",
			callback: commandMapb,
		},
	}
}

