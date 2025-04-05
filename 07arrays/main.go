//arrays are not that commonly used in golang

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Welcome to arrays in Go\n")
	/*
			Printf formats according to a format specifier and writes to standard output. It returns the number of bytes written and any write error encountered.
		    slices are more commonly used in golang, flexible and dynamic & easier to use
	*/

	var fruitList [4]string //must specify size during initialization

	fruitList[0] = "Apple"
	//fruitList[1] = "Peach" //will be replaced by blank space
	fruitList[2] = "Berries"
	fruitList[3] = "Melons"

	fmt.Println("printing array: ", fruitList)
	fmt.Println("printing array: ", len(fruitList)) //gives the amt of reserved memory not the item count

	var veggieList = [6]string{"potato", "carrot", "tomato", "cucumber"}
	fmt.Println("printing veggieList: ", veggieList)
}
