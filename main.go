package main

import (
	"fmt"
	"log"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal("No user")
		return
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands!")
	fmt.Println("Type `exit` to exit")
	repl.Start(os.Stdin, os.Stdout)
}
