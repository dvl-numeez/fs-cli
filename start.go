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
		input:=filterInput(reader.Text())
		words:=strings.Split(input, " ")
		command:=strings.ToLower(words[0])
		switch command{
		case "list":
			ListFiles(input,os.Stdin)
		case "delete":
			Delete(input,delete,os.Stdin,os.Stdout)
		case "copy":
			Copy(input,os.Stdout)
		case "move":
			Move(input,os.Stdin,os.Stdout)
		case "exit":
			Exit()
		case "clear":
			ClearScreen()

		default:
			fmt.Printf("%s command does not exists\n",command)
		}
	}
}

