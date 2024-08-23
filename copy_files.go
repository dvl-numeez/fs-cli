package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


func Copy(input string){
	inputs:=strings.Split(input, " ")
	sourceDestination:=inputs[1]
	copyDestination:=inputs[2]
	CopyFiles(sourceDestination,copyDestination)
}

func CopyFiles(sourceDestination,copyDestination string){
	fileInfo,err:=os.Stat(sourceDestination)
	if err!=nil{
		fmt.Println("Error occured : ",err)	
	}
	nameOfFile:=fileInfo.Name()
	if fileInfo.IsDir(){
		err:=os.Mkdir(filepath.Join(copyDestination,nameOfFile),os.ModePerm)
		if err!=nil{
			fmt.Println("Error occured : ",err)	
		}
		files,err:=os.ReadDir(sourceDestination)
		if err!=nil{
			fmt.Println("Error occured : ",err)	
		}
		for _,file := range files {
			source:=filepath.Join(sourceDestination,file.Name())
			copy:=filepath.Join(copyDestination,nameOfFile)
			CopyFiles(source,copy)
		}


	}else{
	data,err:= os.ReadFile(sourceDestination)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}
	path:=filepath.Join(copyDestination,nameOfFile)
	err=os.WriteFile(path,data,os.ModePerm)
	if err!=nil{
		fmt.Println("Error occured : ",err)
	}
}

}