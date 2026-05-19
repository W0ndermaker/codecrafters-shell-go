package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// bufio.Reader has peek() method that
func main() {
	fmt.Print("$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal("input error:", err)
	}
	fmt.Println(command[:len(command)-1] + ": command not found")
}
