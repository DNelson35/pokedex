package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	pokecache "github.com/DNelson35/pokedex/internal"
)

var baseUrl = "https://pokeapi.co/api/v2/location-area"
type config struct {
	Next string  
	Prev string  
	Cache *pokecache.Cache
}
type PokemonLocationData struct {
	Count      int         `json:"count"`
	Next string `json:"next"`
	Prev string `json:"previous"`
	Results    []struct { 
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandExit (cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp (cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, com := range getCommands() {
		fmt.Printf("%v: %v\n", com.name, com.description)
	}
	for _, com := range getComandsWithArgs(){
		fmt.Printf("%v: %v\n", com.name, com.description)
	}
	return nil
}

func commandMap (cfg *config) error {
	var resp *http.Response
	var err error
	
	data, exist := cfg.Cache.Entry[cfg.Next]
	if exist {
		var locationData PokemonLocationData
		if err := json.Unmarshal(data.Val, &locationData); err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %v", err)
		}

		for _, res := range locationData.Results {
			fmt.Println(res.Name)
		}

		cfg.Next = locationData.Next
		cfg.Prev = locationData.Prev

		return nil
	}

	if cfg.Next != ""{
		resp, err = http.Get(cfg.Next)
	}else {
		resp, err = http.Get(baseUrl)
	}

  
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code: %v", resp.Status)
	}
	if err != nil {
		return err
	}

	var locationData PokemonLocationData

	if err = json.Unmarshal(body, &locationData); err != nil {
		return err
	}

	cfg.Next = locationData.Next
	cfg.Prev = locationData.Prev
	cfg.Cache.Add(resp.Request.URL.String(), body)

	for _, res := range locationData.Results{
		fmt.Println(res.Name)
	}
	return nil
}

func commandMapb (cfg *config) error{
	if cfg.Prev == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	data, exist := cfg.Cache.Entry[cfg.Prev]
	if exist {
		var locationData PokemonLocationData
		if err := json.Unmarshal(data.Val, &locationData); err != nil {
			return fmt.Errorf("failed to unmarshal cached data: %v", err)
		}

		for _, res := range locationData.Results {
			fmt.Println(res.Name)
		}
		cfg.Next = locationData.Next
		cfg.Prev = locationData.Prev
		return nil
	}

	resp, err := http.Get(cfg.Prev)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status code: %v", resp.Status)
	}

	var locationData PokemonLocationData

	if err = json.Unmarshal(body, &locationData); err != nil {
		return err
	}
	cfg.Next = locationData.Next
	cfg.Prev = locationData.Prev
	cfg.Cache.Add(resp.Request.URL.String(), body)

	for _, res := range locationData.Results {
		fmt.Println(res.Name)
	}
	return nil
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaEncounter struct {
	PokemonEncountersList []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func commandExplore(loc string) error{
	endpoint := baseUrl + "/" + loc
	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		return fmt.Errorf("Status code: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var encounters LocationAreaEncounter
	if err = json.Unmarshal(body, &encounters); err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", loc)
	fmt.Println("Found Pokemon:")
	for _, pokeinfo := range encounters.PokemonEncountersList {
		fmt.Printf("- %v\n",pokeinfo.Pokemon.Name)
	}

	return nil
}

