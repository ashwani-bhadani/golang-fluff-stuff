package main

import "fmt"

func main() {
	fmt.Println("golang control flow if-else")
	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Server usage low"
	} else if loginCount > 10 && loginCount < 20 {
		result = "Server usage moderate"
	} else {
		result = "Server usage high"
	}

	fmt.Println(result)

	//you can create vars adhoc & use
	if 9%2 == 0 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}

	//you can directly initialise & use a var, put 1 ';' & keep checking conditions, used in webdevelopment
	if num := 3; num > 10 && num < 20 { //like value comes from web request, assign it to var & check it on the go
		result = "Server usage moderate"
	} else {
		result = "Server usage low or high"
	}
	fmt.Println(result)

	// if err != nil {}

}
