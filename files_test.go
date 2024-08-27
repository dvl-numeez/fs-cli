package main

import (
	"bytes"
	"testing"
)


func TestListFiles(t *testing.T) {
	directory:="/Users/numeezbaloch17/Documents/fs-cli/test_dir"
	output:=bytes.Buffer{}
	ListAllFiles(directory,&output)
	expected:=
`go.txt
readme.md
swift.txt
`
	if output.String()!=expected{
		t.Errorf("Actual : %s Wanted : %s",output.String(),expected)
	}
t.Run("Listing files with extension",func(t *testing.T){
	directory:="/Users/numeezbaloch17/Documents/fs-cli/test_dir"
	output:=bytes.Buffer{}
	ListAllFilesWithExtension(directory,".md",&output)
	expected:=
`readme.md
`
	if output.String()!=expected{
		t.Errorf("Actual : %s Wanted : %s",output.String(),expected)
	}

})

t.Run("Testing `ListFiles function`",func(t *testing.T){
	cases:=[]struct{
		command string
		output string
	}{
		{
			"list /Users/numeezbaloch17/Documents/fs-cli/test_dir",
`go.txt
readme.md
swift.txt
`,
		},
		{
			"list /Users/numeezbaloch17/Documents/fs-cli/test_dir .md",
`readme.md
`,
		},
{
	"list /Users/numeezbaloch17/Documents/fs-cli/test_dir .md .txt",
`Invalid number of arguments
`,	
},

}
for _,c:=range cases{
	output:= bytes.Buffer{}
	ListFiles(c.command,&output)
	if output.String()!=c.output{
		t.Errorf("Expected : %s Actual : %s",c.output,output.String())
	}

}

})
// This is working but it seems formatting issues which make test fails
t.Run("Passing a broken directory",func(t *testing.T){
	directory:="/Users/numeezbaloch17/Documents/fs-cli/wrong_dir"
	output:=bytes.Buffer{}
	
	expected:=`Error occured :  open /Users/numeezbaloch17/Documents/fs-cli/wrong_dir: no such file or directory
`
	ListAllFiles(directory,&output)
	if output.String()!=expected{
		t.Errorf("Actual : %s Wanted : %s",output.String(),expected)
	}
})

t.Run("Passing a broken directory to ListAllFilesWithExtension",func(t *testing.T){
	directory:="/Users/numeezbaloch17/Documents/fs-cli/wrong_dir"
	output:=bytes.Buffer{}
	
	expected:=`Error occured :  open /Users/numeezbaloch17/Documents/fs-cli/wrong_dir: no such file or directory
`
	ListAllFilesWithExtension(directory,".md",&output)
	if output.String()!=expected{
		t.Errorf("Actual : %s Wanted : %s",output.String(),expected)
	}
})
}


