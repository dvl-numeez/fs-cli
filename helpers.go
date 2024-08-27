package main

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)


func Exit(){
	os.Exit(0)
}
func ClearScreen() {
    if runtime.GOOS == "windows" {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    } else {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}


func filterInput(input string)string{
    results:=[]string{}
    instruction:=strings.TrimSpace(input)
    inputs:=strings.Split(instruction, " ")
    for _,i:= range inputs{
        if i!=""{
       results = append(results, i)
        }

    }
return  strings.Join(results, " ")
}
