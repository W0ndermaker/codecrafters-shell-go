package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("input error:", err)
		}

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		command := args[0]
		args = args[1:]

		switch command {
		case "echo":
			fmt.Println(strings.Join(args, " "))
		case "type":
			err = typeCommand(args)
			if err != nil {
				fmt.Printf("{type} command error: %v\n", err)
			}
		case "pwd":
			path, err := os.Getwd()
			if err != nil {
				fmt.Printf("pwd command error: %v\n", err)
				continue
			}
			fmt.Println(path)

		case "exit":
			os.Exit(0)
		default:
			// поиск внешней команды
			_, err := exec.LookPath(command)
			if err != nil {
				fmt.Printf("%v: command not found\n", command)
			} else {
				cmd := exec.Command(command, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			}
		}

	}

}

func typeCommand(args []string) error {
	builtinCommands := map[string]interface{}{
		"echo": nil,
		"exit": nil,
		"type": nil,
		"pwd":  nil,
	}

	for _, arg := range args {
		if _, ok := builtinCommands[arg]; ok {
			fmt.Printf("%v is a shell builtin\n", arg)
			continue
		} else {
			path, err := exec.LookPath(arg)
			if errors.Is(err, exec.ErrDot) {
				err = nil
			}
			if err != nil {
				fmt.Printf("%v: not found\n", arg)
				continue
			}
			fmt.Printf("%v is %v\n", arg, path)
		}

	}
	return nil
}
