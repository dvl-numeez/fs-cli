package main

import (
	"fmt"
	"os"
	"strings"
)


func Copy(input string){
	inputs:=strings.Split(input, " ")
	sourceDestination:=inputs[1]
	copyDestination:=inputs[2]
	data,err:= os.ReadFile(sourceDestination)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}
	err=os.WriteFile(copyDestination,data,0644)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}

}