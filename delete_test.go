package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Run("Testing delete directory function", func(t *testing.T) {
		output := bytes.Buffer{}
		dir := "/Users/numeezbaloch17/Documents/fs-cli/test_delete_dir"
		file := dir + "/" + "check.txt"
		
		err := os.WriteFile(file, []byte("This is delete test"), os.ModePerm)
		if err != nil {
			t.Error(err)
		}
		input := "/Users/numeezbaloch17/Documents/fs-cli/test_delete_dir/check.txt"
		deleteAllContentsOfDirectory(input, &output)
		_, err = os.Open(file)
		if err == nil {
			t.Error("Expecting an error but did not get it")
		}

	})
	t.Run("Deleting a directory", func(t *testing.T) {
		output := bytes.Buffer{}
		dir := "/Users/numeezbaloch17/Documents/fs-cli/test_delete_directory"
		files := []string{dir + "/" + "check.txt", dir + "/" + "file1.txt", dir + "/" + "file2.txt", dir + "/" + "file3.txt"}

		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			t.Error(err)
		}
		for _, file := range files {
			err = os.WriteFile(file, []byte("This is delete test"), os.ModePerm)
			if err != nil {
				t.Error(err)
			}
		}
		deleteAllContentsOfDirectory(dir, &output)
		_, err = os.Stat(dir)
		isNotExist := os.IsNotExist(err)
		if !isNotExist {
			t.Errorf("Expected value %t Got : %t", true, isNotExist)
		}
	})
	t.Run("Check whether the directory is empty or not",func(t *testing.T){
		cases:=[]struct{
			dir string
			result bool
		}{
			{"/Users/numeezbaloch17/Documents/fs-cli/bin",false},
			{"/Users/numeezbaloch17/Documents/fs-cli/copy_test_dir",false},
			{"/Users/numeezbaloch17/Documents/fs-cli/empty_dir",true},
			{"/Users/numeezbaloch17/Documents/fs-cli/test_dir",false},
		}

		for _,c:=range cases{
			actual,err:=IsDirectoryEmpty(c.dir)
			if err!=nil{
				t.Error(err)
			}
			if actual!=c.result{
				t.Errorf("Actual : %v Expected : %v",c.result,actual)
			}

		}
	})

}

func TestDeleteDirectory(t *testing.T) {
	t.Run("Checking Delete Directory function",func(t *testing.T) {
		dir:="/Users/numeezbaloch17/Documents/fs-cli/helper.txt"
		os.WriteFile(dir,[]byte("This is the test file for delete functionality"),os.ModePerm)
		input:="delete /Users/numeezbaloch17/Documents/fs-cli/helper.txt"
		output:=bytes.Buffer{}
		DeleteDirectory(input,delete,&output)
		_,err:=os.Open(dir)
		if err==nil{
			t.Error("Expecting an error but did not get it")
		}
	})
	t.Run("Passing invalid directory",func(t *testing.T){
		input:="delete /Users/numeezbaloch17/Documents/fs-cli/helper.txt .md"
		output:=bytes.Buffer{}
		DeleteDirectory(input,delete,&output)
		expected:=`Invalid number of arguments
`
		if expected!=output.String(){
			t.Errorf("Expected : %s Actual : %s",expected,output.String())
		}
	})
}


func TestDeleteForMove(t *testing.T) {
	t.Run("Checking the function",func(t *testing.T){
		dir:="/Users/numeezbaloch17/Documents/fs-cli/helper.txt"
		os.WriteFile(dir,[]byte("This is the test file for delete functionality"),os.ModePerm)
		output:=bytes.Buffer{}
		DeleteForMove(dir,&output)
		_,err:=os.Open(dir)
		if err==nil{
			t.Error("Expecting an error but did not get it")
		}
	})
}

func TestDeleteFunction(t *testing.T) {
	basePath:="/Users/numeezbaloch17/Documents/fs-cli/test_delete_dir"
	data:="This is "
	files:=[]string{"file1.txt","file2.txt","file3.txt","file4.txt","file5.txt","file6.txt","file7.txt"}
	for _,file:= range files{
		err:=os.WriteFile(basePath+"/"+file,[]byte(data+file),os.ModePerm)
		if err!=nil{
			t.Error(err)
		}
	}
	cases:=[]struct{
		dir string
		input string
		output string
	}{
		{basePath+"/"+"file1.txt","y\n",""},
		{basePath+"/"+"file2.txt","y\n",""},
		{basePath+"/"+"file3.txt","y\n",""},
		{basePath+"/"+"file4.txt","y\n",""},
		{basePath+"/"+"file5.txt","n\n","Delete operation cancelled"},
		{basePath+"/"+"file6.txt","n\n","Delete operation cancelled"},
		{basePath+"/"+"file7.txt","x\n","Invalid input please enter Y or N according to your needs"},
	}

	for _,c:=range cases{
		output:=bytes.Buffer{}
		reader:=strings.NewReader(c.input)
		Delete("delete "+c.dir,delete,reader,&output)
		if output.String()!=c.output{
			t.Errorf("Expected : %s Actual : %s",c.output,output.String())
		}
	}
	

}