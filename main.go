package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cmds := getCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		method := scanner.Text()
		if len(method) == 0 {
			continue
		}
		method = cleanInput(method)
		cmd, ok := cmds[method]
		if !ok {
			fmt.Println("Command not available")
			continue
		}
		error := cmd.callback()
		if error != nil {
			fmt.Println(error)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
	}
}

func cleanInput(text string) string {
	return strings.TrimSpace(strings.ToLower(text))
}
