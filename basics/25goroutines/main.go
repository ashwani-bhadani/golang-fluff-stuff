package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// normally took 9 secs, with go-routines took 1 sec
// on go routine size can be as small as 2 kb
var waitgrp sync.WaitGroup //to make main method wait for routines to finish, "usually must be a pointer"
//go runtime manages go routine, but OS manages threads? 5 routines can try to write on same memory
//who is managing the variables being used & locks on it preventing inconsistency
//mutex :  mutual exclusion lock, can be read & write level mutex too, locks memory while one go-routines uses it
//there is RW read-write mutex, locks for read, until some routines starts writing, stops the read process until write is done, then allows read again

var mut sync.Mutex //usually should be pointer
var globalCounter int = 0

func main() {
	start := time.Now()
	//go routine is simple created by adding a word 'go'
	// go greeter("hellow") //this will print world , we fired a thread but not waited to comeback
	// & print, meanwhile main thread completed its execution
	// greeter("world")
	websites := []string{
		"https://dummyjson.com/products",
		"https://jsonplaceholder.typicode.com/posts",
		"https://dummyjson.com/carts",
		"https://dummyjson.com/users",
		"https://jsonplaceholder.typicode.com/posts/1/comments",
		"https://dummyjson.com/posts",
		"https://dummyjson.com/comments",
		"https://dummyjson.com/quotes",
		"https://jsonplaceholder.typicode.com/comments?postId=1",
		"https://dummyjson.com/todos",
		"https://dummyjson.com/recipes",
		"https://dummyjson.com/image",
		"https://jsonplaceholder.typicode.com/comments?postId=1",
		"https://dummyjson.com/auth",
	}

	for _, web := range websites { //we do not need to use index
		/*
			calling waitgrp.Add(1) after the actual function call — this can cause a race condition or
			panic if waitgrp.Done() gets called before Add(1).
			Move the waitgrp.Add(1) before starting the goroutine.
			Othreise will get neagtive waitgroup counter,
			Actually run getStatusCode(web) in a goroutine using the go keyword.
		*/
		waitgrp.Add(1)        //just to keep track how mainy thread went in & expect completion
		go getStatusCode(web) //this will be slower normally, as fetching data, using 'go'
	}
	waitgrp.Wait() //this part of wait group goes only at the end of main to not let main thread finish & stop
	totalExecTime := time.Since(start)
	fmt.Println("Total execution time is: ", totalExecTime)
}

func greeter(s string) {
	for i := 0; i <= 6; i++ {
		time.Sleep(3 * time.Millisecond) //diff tread comes & goes back & recieves it. we use packages 'sync' for this
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer waitgrp.Done() //this part signifies thread is completed
	start := time.Now()

	result, err := http.Get(endpoint)
	if err != nil {
		log.Fatal("error in GET")
	}

	mut.Lock() //can also have read-write locks
	globalCounter += 1
	modifiedEndpoint := "Entry: " + strconv.Itoa(globalCounter) + " " + endpoint
	mut.Unlock() //programmers responsibility to lock/unlock

	elapsedTime := time.Since(start)
	fmt.Printf("%d status code for %s & iteration took time: %v\n", result.StatusCode, modifiedEndpoint, elapsedTime)
}
