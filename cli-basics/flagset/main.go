package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(2)
	}
	cmd := os.Args[1]

	switch cmd {
	case "greet":
		msg := "REMINDERS CLI - CLI BASICS"
		// Todo
		fmt.Printf("hello and welcome: %s\n", msg)
	case "help":
		fmt.Println("some help message")
	default:
		fmt.Printf("unknown command: %s\n", cmd)
	}
}
