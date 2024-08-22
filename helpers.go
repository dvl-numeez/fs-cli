package main

import (
	"os"
	"os/exec"
	"runtime"
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

