package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var ExitCommandErr error = errors.New("Exit command")

func CommandHandler(command string, args []string) error {
	var err error
	switch command {
	case "echo":
		fmt.Println(strings.Join(args, " "))

	case "type":
		err := typeCommand(args)
		if err != nil {
			err = fmt.Errorf("{type} command error: %w\n", err)
		}

	case "pwd":
		path, err := os.Getwd()
		if err != nil {
			err = fmt.Errorf("{pwd} command error: %w\n", err)
		} else {
			fmt.Println(path)
		}

	case "cd":
		if len(args) == 1 {
			path := args[0]
			if args[0] == "~" {
				path = os.Getenv("HOME")
			}

			err := os.Chdir(path)
			if err != nil {
				dir := strings.Split(path, "\\")
				fmt.Printf("cd: %v: No such file or directory\n", dir[len(dir)-1])
			}
		}

		if len(args) > 1 {
			fmt.Println("cd: too many arguments")
		}

	case "exit":
		err = ExitCommandErr

	default:
		cmd := exec.Command(command, args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Printf("%v: command not found\n", command)

		}
	}
	return err
}

func typeCommand(args []string) error {
	var err error
	builtinCommands := map[string]interface{}{
		"echo": nil,
		"exit": nil,
		"type": nil,
		"pwd":  nil,
		"cd":   nil,
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
	return err
}
