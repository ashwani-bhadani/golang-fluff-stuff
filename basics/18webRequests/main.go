package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
---------IMP for web dev---------
https://pkg.go.dev/net/http#Response
it is programmers responsiblity to close a response once reading is done
	// Close records whether the header directed that the connection be
	// closed after reading Body. The value is advice for clients: neither
	// ReadResponse nor Response.Write ever closes a connection.  ******IMP******
	Close bool
*/

const jsonPlaceholder = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("learning web requests handling")

	response, err := http.Get(jsonPlaceholder)
	errorHandler(err)

	fmt.Printf("Response type is: %T\n", response) //note the actual object is returned by http
	// not the copy, but the original reference to the object

	defer response.Body.Close() //callers repsonsibility to close connection
	//if you close before processing will get file already closed erro

	//majority of reading will be done by io-utils //ioutil.ReadAll() deprecated in Go 1.16
	databytes, err := io.ReadAll(response.Body)
	errorHandler(err)

	//1. can convert data byte into content
	content := string(databytes)
	fmt.Printf("Fetched the 1st post from JSON Placeholder: %v\n", content)

	//2. prettyfy the JSON:
	var prettyJSON bytes.Buffer                        //remember databytes is byte array & prettyJSON is byte buffer, convert to string !
	err = json.Indent(&prettyJSON, databytes, "", " ") //passing the actual reference to object, not a copy
	errorHandler(err)
	fmt.Println("Fetched 1st post & beautify the JSON: \n", prettyJSON.String())

	//3.unmarshalling it into a struct, also use json package
	var firstPost Post
	err = json.Unmarshal(databytes, &firstPost) //this &post actually passes the reference to object & not its copy
	errorHandler(err)
	fmt.Printf("Post title is %s\n", firstPost.Title)

}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Nody   string `json:"body"`
}
