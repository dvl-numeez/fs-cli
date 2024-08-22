package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
func Delete(input string){
	fmt.Println("Are you sure you want to delete the directory/file")
	fmt.Println("Press Y/N")
	reader:=bufio.NewScanner(os.Stdin)
	reader.Scan()
	choice:=reader.Text()
	switch strings.ToLower(choice){
	case "y":
		DeleteDirectory(input)
	case "n":
		fmt.Println("Delete operation cancelled")
	default:
		fmt.Println("Invalid input please enter Y or N according to your needs")

		
	}

}
func DeleteDirectory(input string){
	inputs:=strings.Split(input, " ")
	length:=len(inputs)
	if length==2{

		deleteAllContentsOfDirectory(inputs[1])
		err:=os.Remove(inputs[1])
		if err!=nil{
			fmt.Println("Error occured : ",err)
		}else{
			fmt.Println("Invalid number of arguments")
		}
	}

	
}

func  IsDirectoryEmpty(directory string)(bool,error){
	f, err := os.Open(directory)
    if err != nil {
        return false, err
    }
    defer f.Close()
	_, err = f.Readdirnames(1) 
    if err == io.EOF {
        return true, nil
    }
    return false, err 
}

func deleteAllContentsOfDirectory(directory string){
	
	files,err:=os.ReadDir(directory)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}
	for _,file:=range files{
		path:=directory+"/"+file.Name()
		isDirectory,err:=IsDirectory(path)
		if err!=nil{
			fmt.Println("Error occured : ",err)
		}
		if isDirectory{
			deleteAllContentsOfDirectory(path)
		}
		os.Remove(path)
	}
}

func IsDirectory(directory string)(bool,error){
	info,err:=os.Stat(directory)
	if err!=nil{
		return false,err
	}
	if info.IsDir(){
		return true,nil
	}
	return false,nil
}