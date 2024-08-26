package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)


func Copy(input string, writer io.Writer){
	inputs:=strings.Split(input, " ")
	switch len(inputs){
	case 3 :
		sourceDestination:=inputs[1]
		copyDestination:=inputs[2]
		_,sourceErr:=os.Open(sourceDestination)
		if sourceErr!=nil{
			fmt.Fprintln(writer,sourceErr)
		}
		_,copyErr:=os.Open(copyDestination)
		if copyErr!=nil{
			fmt.Fprintln(writer,copyErr)
		}
		if sourceErr==nil && copyErr==nil{
			CopyFiles(sourceDestination,copyDestination,writer)
		}
		
	default:
		fmt.Fprintln(writer,"Invalid arguments")

	}
}

func CopyFiles(sourceDestination,copyDestination string, writer io.Writer){
	fileInfo,err:=os.Stat(sourceDestination)
	if err!=nil{
		fmt.Fprintln(writer,"Error occured : ",err)	
		
	}
	nameOfFile:=fileInfo.Name()
	if fileInfo.IsDir(){
		err:=os.Mkdir(filepath.Join(copyDestination,nameOfFile),os.ModePerm)
		if err!=nil{
			fmt.Fprintln(writer,"Error occured : ",err)
		}
		files,err:=os.ReadDir(sourceDestination)
		if err!=nil{
			fmt.Fprintln(writer,"Error occured : ",err)
		}
		for _,file := range files {
			source:=filepath.Join(sourceDestination,file.Name())
			copy:=filepath.Join(copyDestination,nameOfFile)
			CopyFiles(source,copy,writer)
		}


	}else{
	data,err:= os.ReadFile(sourceDestination)
	if err!=nil{
		fmt.Fprintln(writer,"Error occured : ",err)
	}
	path:=filepath.Join(copyDestination,nameOfFile)
	err=os.WriteFile(path,data,os.ModePerm)
	if err!=nil{
		fmt.Fprintln(writer,"Error occured : ",err)
	}
}

}