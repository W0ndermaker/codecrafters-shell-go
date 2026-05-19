package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal("input error:", err)
		}

		args := strings.Fields(input)
		command := args[0]
		args = args[1:]

		switch command {
		case "echo":
			fmt.Println(strings.Join(args, " "))
		case "type":
			typeCommand(args)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(command + ": command not found")
		}

	}
}

func typeCommand(args []string) {
	builtinCommands := map[string]interface{}{
		"echo": nil,
		"exit": nil,
		"type": nil,
	}
	for _, arg := range args {
		trimmed := strings.TrimSpace(arg)
		if _, ok := builtinCommands[trimmed]; !ok {
			fmt.Printf("%v: not found\n", arg)
			continue
		}
		fmt.Printf("%v is a shell builtin\n", arg)
	}
}
