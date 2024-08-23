package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
const(
	delete = "delete"
	move = "move"
)

func startRepl(){
	reader:=bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("> ")
		reader.Scan()
		input:=reader.Text()
		words:=strings.Split(input, " ")
		command:=strings.ToLower(words[0])
		switch command{
		case "list":
			ListFiles(input)
		case "delete":
			Delete(input,delete)
		case "copy":
			Copy(input)
		case "move":
			Move(input)
		case "exit":
			Exit()
		case "clear":
			ClearScreen()

		default:
			fmt.Printf("%s command does not exists\n",command)
		}
	}
}

