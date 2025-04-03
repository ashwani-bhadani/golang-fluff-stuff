package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in golang")

	//var ptr *int //stores integer but string can be the data type
	//fmt.Println("value of pointer : ", ptr) //if pointer is not initialised, its value is nil

	myNumber := 23
	var numptr = &myNumber //here i'm not just creating a pointer but one that is referencing to some data in memory
	//thus referene means &<varname>

	fmt.Println("value of pointer ref: ", numptr) //will give memory address directly
	fmt.Println("value of pointer: ", *numptr)    //asteric shows whats inside this pointer
	//*numptr means the actual value inside it so you can perform computations over it

	*numptr = *numptr * 2            //this will ensure/guarantee whatever computation is actually performed on the values not their references
	fmt.Println("Result: ", numptr)  //see memory address remains same
	fmt.Println("Result: ", *numptr) //will print the stored value
	fmt.Printf("Type of pointer %T \n", numptr)
	fmt.Printf("Type of pointer %T \n", *numptr)

	/*
		-> since pointers is direct memory address reference so you are guaranteed actual
		 value is being passed not corrupted by improper references
	*/
}
