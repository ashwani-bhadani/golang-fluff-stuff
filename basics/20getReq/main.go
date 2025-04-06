package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to GET requests in GO")
	PerformGetRequest()
}

func PerformGetRequest() { //modify for this to handle url as parameter
	const jsonUrl = "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(jsonUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Response status code: ", response.StatusCode)
	fmt.Println("TLS Server name : ", response.TLS.ServerName)

	var contentStr strings.Builder //response is still inside content string
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := contentStr.Write(content)

	fmt.Println("Byte count is: ", byteCount)
	fmt.Println(contentStr.String()) //this is a better way as library is more powerful, just need to use strings.Builder
	fmt.Println("Content Length: ", len(string(content)))
}
