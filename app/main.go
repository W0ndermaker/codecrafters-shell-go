package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/shell-starter-go/app/commands"
	"github.com/google/shlex"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("input error:", err)
		}

		splitInput, _ := shlex.Split(input)

		command, args := splitInput[0], splitInput[1:]

		//_, args, _ := commands.ArgParser(input, reader)
		// if err != nil {
		// 	if errors.Is(err, commands.QuotesNumberdErr) {
		// 		fmt.Println("> ")
		// 		addInput := reader.ReadString()
		// 	}
		// }

		err = commands.CommandHandler(command, args)
		if err != nil {
			if errors.Is(err, commands.ExitCommandErr) {
				os.Exit(0)
			}
			fmt.Println(err)
		}
	}

}
