package main

import (
	"fmt"
)

/*
Slices are a more flexible alternative to arrays, under the hood arrays only.
 They are built on top of arrays and mostly used, powerful & abstraction layer are called slices.
*/

func main() {
	fmt.Println("Welcome to slices")

	//same syntax as array but with no initialization, we can add as many & memory expands
	//in slices if we use this syntax we need to initialize it as well
	var fruitList = []string{"apple", "orange", "banana", "grapes"}
	fmt.Printf("type of fruitlist %T\n", fruitList)
	fmt.Println(fruitList)

	//in slices we can't add by positon but use append method & specify which slice what to add
	fruitList = append(fruitList, "mango", "sapota", "musk melon")
	fmt.Println(fruitList)

	//we use colon syntax in square to separate our slices, make parts off it
	//fruitList = append(fruitList[1:]) //this is uneding , 1st item is position 0
	fruitList = append(fruitList[1:3]) //range is non-inclusive so pos #3 item will not be picked, pos #1 & #2
	//ruitList = append(fruitList[:3]) //start from default zero pos#
	fmt.Println(fruitList)

	highScore := make([]int, 4) //decalre what type 1st & then the size of it
	highScore[0] = 34
	highScore[1] = 44
	highScore[2] = 54
	highScore[3] = 64
	//highScore[4] = 74 //this will give out of bound error but if u use append, append function reallocates the memory
	var slicePtr = &highScore
	//To explicitly print the memory address of slicePtr, use the %p format specifier with fmt.Printf:
	//mt.Println("Now memory location ", slicePtr) // lexer will auto format the data & print the list instead
	fmt.Printf("previously memory location: %p\n", slicePtr)
	highScore = append(highScore, 66, 67, 68, 69) //entire memory allocation happened again (proved by checking pointer)
	//to print the actual address of data inside slice
	fmt.Printf("New address of slice data: %p\n", &(*slicePtr)[0])
	fmt.Println(highScore)

	//-----------------how to remove value from slice based on index-----------------------------------//
	var courses = []string{"golang", "react.js", "springboot", "java", "javascript"}
	fmt.Println(courses) //we can use append to add or even remove values
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) //now delete is available
	//append reallocates the memory, we take portion of slice since index is non-inclusive, it'll be removed
	fmt.Println(courses)

}
