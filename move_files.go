package main

import (
	"fmt"
	"io"

	"strings"
)


func Move(input string , reader io.Reader,writer io.Writer){
	inputs:=strings.Split(input, " ")
	switch len(inputs){
	case 3 :
		sourceDestination:=inputs[1]
		copyDestination:=inputs[2]
		CopyFiles(sourceDestination,copyDestination,writer)
		Delete(sourceDestination,move,reader,writer)
	default:
		fmt.Fprintln(writer,"Invalid arguments")

	}

}