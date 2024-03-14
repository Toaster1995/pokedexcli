package main

import "fmt"

func commandHelp() error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	cmds := getCliCommands()
	for i := range cmds {
		fmt.Printf("%v: %v\n", cmds[i].name, cmds[i].description)
	}
	return nil
}
