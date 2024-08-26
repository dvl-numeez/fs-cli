package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)


func ListFiles(input string ,writer io.Writer){
	inputs:=strings.Split(input," ")
	length:= len(inputs)
	switch length{
	case 2:
		ListAllFiles(inputs[1],writer)
	case 3:
		ListAllFilesWithExtension(inputs[1],inputs[2],writer)
	default:
		fmt.Fprintln(writer,"Invalid number of arguments")
	}
}


func ListAllFiles(directory string, writer io.Writer){
	files,err:=os.ReadDir(directory)
	if err!=nil{
		fmt.Fprintln(writer, "Error occured : ",err)
		
	}
	for _,file:= range files{
		fmt.Fprintln(writer,file.Name())
		
	}
}
func ListAllFilesWithExtension(directory,extension string,writer io.Writer){
	files,err:=os.ReadDir(directory)
	if err!=nil{
		fmt.Fprintln(writer,"Error occured : ",err)
	}
	for _,file:= range files{
		if strings.Contains(file.Name(),extension){
			fmt.Fprintln(writer,file.Name())
		}
	}
}

//Spare function
//This function can be used for searching with some modifications
// func ListAllFilesWithFilter(directory,extension string,writer io.Writer){
// 	root:=os.DirFS(directory)
// 	results,err:=fs.Glob(root,extension)
// 	if err!=nil{
// 		fmt.Fprintln(writer,"Error occured : ",err)
// 	}
// 	for _,res:=range results{
// 		fmt.Fprintln(writer,res)
// 	}
// }