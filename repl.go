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

func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan(){
			result := cleanInput(scanner.Text())
			if len(result) == 0 {
				fmt.Println("no input detected")
			} else {
				fmt.Printf("Your command was: %v\n", result[0])
			}
		}
	}
}