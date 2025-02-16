package main

import "strings"

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
// strings.Fields would have gave the same result