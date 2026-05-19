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

		args := strings.Split(strings.TrimSpace(input), " ")
		command := args[0]
		args = args[1:]

		switch command {

		case "echo":
			fmt.Println(strings.Join(args, " "))
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(command + ": command not found")
		}

	}
}
