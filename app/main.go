package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// bufio.Reader has peek() method that
func main() {
	// REPL
	for {
		fmt.Print("$ ")

		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal("input error:", err)
		}

		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		}

		fmt.Println(command + ": command not found")
	}
}
