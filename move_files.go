package main

import "strings"


func Move(input string){
	inputs:=strings.Split(input, " ")
	sourceDestination:=inputs[1]
	copyDestination:=inputs[2]
	CopyFiles(sourceDestination,copyDestination)
	Delete(sourceDestination,move)
}