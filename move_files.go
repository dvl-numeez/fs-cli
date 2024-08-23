package main

import (
	"io"
	
	"strings"
)


func Move(input string , writer io.Writer){
	inputs:=strings.Split(input, " ")
	sourceDestination:=inputs[1]
	copyDestination:=inputs[2]
	CopyFiles(sourceDestination,copyDestination,writer)
	Delete(sourceDestination,move,writer)
}