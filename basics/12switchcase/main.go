package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("control flow switch cases in go")
	//can get from math/rand or crypto/ran
	rand.Seed(time.Now().UnixNano()) //deprecated since go 1.20

	for num := 1; num <= 10; num++ {
		fmt.Printf("Rolling dice %v time.\n", num)
		diceNumber := rand.Intn(6) + 1 // since range 6 is not-inclusive
		// fmt.Println("Dice rolled is : ", diceNumber)

		switch diceNumber {
		case 1:
			fmt.Println("Dice value is 1 & you can  move 1 spot")
			fallthrough //we do not have automatic fallthrough in golang like old langs, declare explicitly
			//but it will execute only the case next after, not all the below ones
		case 2:
			fmt.Println("Dice value is 2 & you can move 2 spots")
		case 3:
			fmt.Println("Dice value is 3 & you can  move 3 spots")
		case 4:
			fmt.Println("Dice value is 4 & you can  move 4 spots")
		case 5:
			fmt.Println("Dice value is 5 & you can  move 5 spots")
		case 6:
			fmt.Println("Dice value is 6 & you can  move 6 spots or open")
		default:
			fmt.Println("This number is not possible")
		}
	}

}
