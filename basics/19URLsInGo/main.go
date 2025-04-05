package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://lbsan27joao:4590/learncode?coursename=golang&tutorialmode=playground&istrialbased=y"

func main() {
	fmt.Println("Welcome to URLs in GO")
	fmt.Println(myUrl)

	//parsing the URL
	result, err := url.Parse(myUrl)
	errorHandler(err)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparams := result.Query() //a good way to store all query params
	fmt.Printf("the query params are: %v and of type %T\n", qparams, qparams)
	//O/P: the query params are: map[coursename:[golang] istrialbased:[y] tutorialmode:[playground]]
	//  and of type url.Values and means they are key-value types

	fmt.Println(qparams["coursename"]) //since K-V, string-string,
	fmt.Println(qparams["istrialbased"])

	// can also use for loop in range
	for _, val := range qparams {
		fmt.Println("Param is: ", val) //order is not guaranteed
		/* OUTPUT
		Param is:  [playground]
		Param is:  [y]
		Param is:  [golang]
		*/
	}

	//if all info is in chunks, creating an URL, URL is struct & case-sensitive
	//***IMP*** we do not want to pass a copy of URL but a reference (&url) of it
	partsOfUrl := &url.URL{
		Scheme:   "tcp",
		Host:     "localhost:61616",
		Path:     "activemq",
		RawQuery: "user=admin",
	}
	var anotherURL string = partsOfUrl.String()
	fmt.Println(anotherURL)

}

// Since errorHandler starts with a lowercase e, it is unexported (private) to the package it is defined in.
func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
