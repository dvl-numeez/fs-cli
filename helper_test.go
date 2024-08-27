package main

import "testing"


func TestFilterInput(t *testing.T){
	cases:=[]struct{
		input string
		output string
	}{
		{"       copy           /Users/numeezbaloch17/Documents/minio-data        ","copy /Users/numeezbaloch17/Documents/minio-data"},
		{"             copy /Users/numeezbaloch17/Documents/minio-data","copy /Users/numeezbaloch17/Documents/minio-data"},
		{"copy /Users/numeezbaloch17/Documents/minio-data                ","copy /Users/numeezbaloch17/Documents/minio-data"},
		{"        copy       /Users/numeezbaloch17/Documents/minio-data           /Users/numeezbaloch17/Documents/minio-data                    ","copy /Users/numeezbaloch17/Documents/minio-data /Users/numeezbaloch17/Documents/minio-data"},

	}
	for _,c:=range cases{
		actual:=filterInput(c.input)
		if actual!=c.output{
			t.Errorf("Expected %s Got : %s",c.output,actual)
		}
	}
}
