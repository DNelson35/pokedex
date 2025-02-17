package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"time"
	"github.com/DNelson35/pokedex/internal"
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

type cliCommandsWithArgs struct {
	name				string
	description string
	callback		func(string) error
}



func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := &config{
		Cache: pokecache.NewCache(5 * time.Second),
	}
	
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
		}else {
			commandWithArgs, exists := getComandsWithArgs()[commandName]
			arg := words[1]
			if exists {
				err := commandWithArgs.callback(arg)
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

func getComandsWithArgs() map[string]cliCommandsWithArgs {
	return map[string]cliCommandsWithArgs{
		"explore": {
			name: "explore",
			description: "prints pokemon in a location",
			callback: commandExplore,
		},
	}
}

