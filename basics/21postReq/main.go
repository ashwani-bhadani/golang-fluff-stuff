package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to post in golang")
	// PostJsonReq()
	sendFormDataPostReq()
}

func sendFormDataPostReq() {
	const myUrl = "http://localhost:8080/postform"
	//form data fake, posting any data can be accessed by url package
	data := url.Values{}
	data.Add("firstname", "Ashwani")
	data.Add("lastname", "Bhadani")
	data.Add("email", "ashwani@gmail.com")
	data.Add("phone", "3245125621")

	response, err := http.PostForm(myUrl, data) //post data in wwww-url-encoded-form
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("form repsonse is: ", string(content))
}

func PostJsonReq() {
	const myUrl = "http://localhost:8080/postmessage"

	//json payload, we can use new reader `` & create any data format
	requestBody := strings.NewReader(`
		{
			"username": "ashwani",
			"age": 26
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}
