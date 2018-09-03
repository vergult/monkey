package main

import (
	"fmt"
	"os"
	"os/user"

	"gitlab.com/vergult/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the Monkey programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
