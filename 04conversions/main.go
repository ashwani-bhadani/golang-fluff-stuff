package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza in range 1 & 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	//converting rating
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //strconv to parse string to int float bool
	//common way for comma , ok syntax is you move the program checking if there is error or not
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1) //need tp trim the trailing chars, use strings package
	}

}
