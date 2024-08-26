package main

import (
	"bytes"
	"os"

	"testing"
)


func TestCopyFiles(t *testing.T) {
	t.Run("Copying a file",func(t *testing.T){
		sourceDestination:="/Users/numeezbaloch17/Documents/fs-cli/copy_test.txt"
		copyDestination:="/Users/numeezbaloch17/Documents/fs-cli/copy_test_dir"
		output:=bytes.Buffer{}
		CopyFiles(sourceDestination,copyDestination,&output)
		copyFile,err:=os.ReadFile(copyDestination+"/"+"copy_test.txt")
		if err!=nil{
			t.Error(err)
		}
		originalFile,err:=os.ReadFile(sourceDestination)
	if err!=nil{
		t.Error(err)
	}
	assertSameFile(t,copyFile,originalFile)
	})

	t.Run("Passing invalid arguments",func(t *testing.T){
		sourceDestination:="/Users/numeezbaloch17/Documents/fs-cli/copy_test.txt"
		copyDestination:="/Users/numeezbaloch17/Documents/fs-cli/copy_test_dir"
		output:=bytes.Buffer{}
		Copy(sourceDestination+copyDestination+" .md",&output)	
		expected:=
`Invalid arguments
`
		if output.String()!=expected{
			t.Errorf("Expected : %s Actual : %s",expected,output.String())
		}
	})
	
}

func TestCopy(t *testing.T){
	t.Run("Passing invalid arguments",func(t *testing.T){
		input:="copy move /Users/numeezbaloch17/Documents/fs-cli/test_move_dir /Users/numeezbaloch17/Documents/fs-cli/move_dir md"
		output:=bytes.Buffer{}
		Copy(input,&output)
		expected:=`Invalid arguments
`
		if output.String()!=expected{
			t.Errorf("Expected : %s Actual : %s",expected,output.String())
		}
	})
	t.Run("Checking the Copy function",func(t *testing.T){
		input:="copy /Users/numeezbaloch17/Documents/fs-cli/copy_test.txt /Users/numeezbaloch17/Documents/fs-cli/move_dir"
		output:=bytes.Buffer{}
		Copy(input,&output)
		_,err:=os.Open("/Users/numeezbaloch17/Documents/fs-cli/move_dir/copy_test.txt")
		if err!=nil{
			t.Error("Did not expecting an error but still got it it means file not copied")
		}
	})
}

func assertSameFile(t testing.TB,copyFile,originalFile []byte){
	t.Helper()
	if !bytes.Equal(copyFile,originalFile){
		t.Error("File/Directory is not copied into the copy destination")
	}
}


