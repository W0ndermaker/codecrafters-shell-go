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
	for {
		fmt.Print("$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
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
		case "exit":
			os.Exit(0)
		default:
			_, err := exec.LookPath(command)
			if err != nil {
				fmt.Printf("%v: command not found\n", command)
			} else {
				cmd := exec.Command(command, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err = cmd.Run(); err != nil {
					fmt.Printf("command failed: %v\n", err)
				}
			}

		}

	}
}

func typeCommand(args []string) error {
	builtinCommands := map[string]interface{}{
		"echo": nil,
		"exit": nil,
		"type": nil,
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
