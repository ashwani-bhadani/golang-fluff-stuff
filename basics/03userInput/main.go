package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "WELCOME to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	//comma ok , error ok syntax -> like try catch : treat problems/errors as true false
	// input, _ := reader.ReadString('\n') //in below 1st part is try & 2nd part is catch
	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)
	fmt.Printf("Type of rating is %T \n", input)
	fmt.Printf("Error occured & is of type %T", err)
}
