package main

//https://go.dev/doc/effective_go#defer
/*
Key Points:
	-when a func executes, it line by line, defer is used to delay the execution of a
	 function until the surrounding function returns.
	-on encountering defer keyword, that line is executed at the very end of func itself
	 invoked immediately before the return stmt
	-Deferred functions run in reverse order (LIFO – Last In, First Out).
	-Arguments in deferred functions are evaluated immediately, not when the deferred function runs.
	-Go will remember defer declaration and run it last, just before the function ends,
	 even if there’s an error or return statement.

It’s super useful for things you want to do at the end, like:
	-Closing a file
	-Unlocking a mutex
	-Releasing a database connection
*/

import "fmt"

func main() {
	defer fmt.Println("executing defer 1") //will be executed towards end (1st guy to move into defer stack, comes out last)
	defer fmt.Println("executing defer 2")
	defer fmt.Println("executing defer 3") //will be executed in LIFO i.e, reversed order (went last, comes out)
	fmt.Println("Hello in defer tutorial")
	myDefer() //this line does not get deferred compared to main function, this will be called 1st
}

func myDefer() {
	for i := 0; i <= 10; i++ {
		defer fmt.Printf("printing defered in for loop item: %v\n", i)
		fmt.Printf("Hello from for loop having defer with i : %v\n", i)
	}
}
