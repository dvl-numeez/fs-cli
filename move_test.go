package main

import (
	"bytes"
	"os"
	"testing"
)

//TODO
//There is been issue with copy functionality which caused move to fail
// We handle nil pointer deference issue in copy functionalities
//That will hopefully solve all the testing issues with the move functionality

func TestMove(t *testing.T) {
	t.Run("Testing Move function",func(t *testing.T) {
		sourceDestination:="/Users/numeezbaloch17/Documents/fs-cli/move_test.txt"
		moveDestination:="/Users/numeezbaloch17/Documents/fs-cli/move_dir"
		input:="move "+sourceDestination+" "+moveDestination
		output:=bytes.Buffer{}
		Move(input,os.Stdin,&output)
		_,err:=os.Open(moveDestination+"/move_test.txt")
		if err !=nil{
			t.Error("File did not move to the move destination")
		}

		_,err = os.Open(sourceDestination)
		if err == nil{
			t.Error("File did not get removed from the source destination")
		}
		//To keep the move test file back to its place so every time this test can run 
		reverseInput:="move "+moveDestination+"/move_test.txt "+"/Users/numeezbaloch17/Documents/fs-cli"
		Move(reverseInput,os.Stdin,os.Stdout)

	t.Run("Passing invalid arguments",func(t *testing.T) {
		inputs:=[]string{""," ","move abc def ghi",}
		for _,input:= range inputs{
			output:=bytes.Buffer{}
			Move(input,os.Stdin,&output)
			expected:=`Invalid arguments
`
			if output.String()!=expected{
				t.Errorf("Expected : %s Actual : %s",expected,output.String())
			}
		}
	})

		

	
	})
}