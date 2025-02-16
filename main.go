package main
import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan(){
			input := scanner.Text()
			result := cleanInput(input)
			if len(result) <= 0 {
				fmt.Println("no input detected")
			} else {
				fmt.Printf("Your command was: %v\n", result[0])
			}
		}
	}
}

