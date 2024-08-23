package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
)


func ListFiles(input string ,writer io.Writer){
	inputs:=strings.Split(input," ")
	length:= len(inputs)
	switch length{
	case 2:
		ListAllFiles(inputs[1])
	case 3:
		ListAllFilesWithExtension(inputs[1],inputs[2])
	default:
		fmt.Println("Invalid number of arguments")
	}



}


func ListAllFiles(directory string){
	files,err:=os.ReadDir(directory)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}
	for _,file:= range files{
		fmt.Println(file.Name())
	}
}
func ListAllFilesWithExtension(directory,extension string){
	files,err:=os.ReadDir(directory)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}
	for _,file:= range files{
		if strings.Contains(file.Name(),extension){
			fmt.Println(file.Name())
		}
	}
}
func ListAllFilesWithFilter(directory,extension string){
	root:=os.DirFS(directory)
	results,err:=fs.Glob(root,extension)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}
	for _,res:=range results{
		fmt.Println(res)
	}
}