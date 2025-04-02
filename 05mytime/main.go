package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 Monday 00:57:03.384592"))
	fmt.Println(presentTime.Format("01-02-2006 Monday")) //this is the way in go you have to specify formatting of date

	createdDate := time.Date(2020, time.September, 10, 23, 43, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday")) //you have to pass monday here to set the format

	//go env tells you the formats it can create program in, an executable for any Os (_for MAC it is GOOS="darwin")
	//use go build will compile after finding the main.go file even tough your are in any OS, it will give windows file
	//%GOOS="windows" go build | GOOS="linux" go build
	/*
		Common Go Environment Variables:
			GOROOT – Path to the Go installation.
			GOPATH – Path for Go workspaces (where src, bin, and pkg reside).
			GOBIN – Directory where compiled binaries are stored.
			GOOS – Target operating system (e.g., linux, darwin, windows).
			GOARCH – Target architecture (e.g., amd64, arm64).
			GOMODCACHE – Directory for Go module cache.

		Set an environment variable (temporarily for the session):
			go env -w GO111MODULE=on
	*/
}
