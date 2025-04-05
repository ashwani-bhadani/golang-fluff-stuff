package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang!")
	content := "I'm writing to a file from golang! again !"

	//we can work with files using os module
	//this will create in current dir, can face permission issue & not create file, handling error is must
	file, err := os.Create("./myGoFile.txt") //file creation is an OS operation
	//1st check for error comma-ok syntax usage
	if err != nil {
		panic(err) //shuts down the program & shows the error
	}

	//to write into files use io package
	length, err := io.WriteString(file, content) //this can also execute or fail, return length of string written\
	checkNilError(err)
	fmt.Println("SUCCESS writing to file, length is :", length)
	defer file.Close() //recommended way is to always use defer keyword

	readFile("./myGoFile.txt") // !! possible issue as defer file writer has not closed yet !!
}

func checkNilError(err error) {
	if err != nil {
		panic(err) //shuts down the program & shows the error
	}
}

/*
----------------Reading the file----------------
***** ioutil has been deprecated since Go 1.16*****
ioutil Function		Replacement

ioutil.ReadFile		os.ReadFile
ioutil.WriteFile	os.WriteFile
ioutil.ReadAll		io.ReadAll
ioutil.NopCloser	io.NopCloser
ioutil.Discard		io.Discard
ioutil.TempFile		os.CreateTemp
ioutil.TempDir		os.MkdirTemp
*/
func readFile(fileNameWithPath string) {
	//for reading there is separate utility of I/O
	//remember files are read in bytestream, here databyte
	dataByte, err := ioutil.ReadFile(fileNameWithPath) //can throw up errors
	checkNilError(err)
	fmt.Println("text data inside the file is : ", string(dataByte))
}
