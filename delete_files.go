package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
func Delete(input string,mode string, writer io.Writer){
	switch mode {
	case "delete":
		fmt.Println("Are you sure you want to delete the directory/file")
		fmt.Println("Press Y/N")
		reader:=bufio.NewScanner(os.Stdin)
		reader.Scan()
		choice:=reader.Text()
		switch strings.ToLower(choice){
		case "y":
			DeleteDirectory(input,mode,writer)
		case "n":
			fmt.Println("Delete operation cancelled")
		default:
		fmt.Println("Invalid input please enter Y or N according to your needs")
		
	}
	case "move":
			DeleteDirectory(input,mode,writer)

}
	


}
func DeleteDirectory(input string,mode string,writer io.Writer){
	switch mode{
	case "delete":
		inputs:=strings.Split(input, " ")
		length:=len(inputs)
		if length==2{

		deleteAllContentsOfDirectory(inputs[1],writer)
		}else{
			fmt.Println("Invalid number of arguments")
		}
	case "move":
		DeleteForMove(input)
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

func deleteAllContentsOfDirectory(directory string,writer io.Writer){
	result:=IsDirectory(directory)
	
	if result{
	
	files,err:=os.ReadDir(directory)
	if err!=nil{
		fmt.Println("Error occured 3: ",err)
	}
	for _,file:=range files{
		path:=directory+"/"+file.Name()
		isDirectory:=IsDirectory(path)
		if isDirectory{
			deleteAllContentsOfDirectory(path,writer)
		}
		os.Remove(path)
	}
	
}
os.Remove(directory)
}

func IsDirectory(directory string)bool{
	info,_:=os.Stat(directory)
	return info.IsDir()
}
func DeleteForMove(dir string){
	deleteAllContentsOfDirectory(dir,os.Stdin)
		err:=os.Remove(dir)
		if err!=nil{
			fmt.Println("Error occured  1 : ",err)
		}
}