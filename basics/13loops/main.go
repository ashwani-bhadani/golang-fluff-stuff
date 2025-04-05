package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	// fmt.Println(days)

	//this is why prefer slices cause if you use arrays, even tough empty items there, len will be wrong at times, theres no ++d
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//this loop auto iterates over all the elements of a slice
	for i := range days {
		fmt.Println(days[i]) //i here unlike javascripts returns an index number
	}

	//for-each , in golang only loop is for, no other keyword
	//In Go, the range keyword must be used when iterating over slices, arrays, maps, etc.
	for index, day := range days {
		fmt.Printf("index is %v and day is %v\n", index+1, day)
	}

	var numVal int32 = 1
	for numVal <= 10 {
		// if numVal == 5 {
		// 	fmt.Println("will break")
		// 	break
		// 	fmt.Println("getting out of loop due to break")
		// }
		//example of continue
		// if numVal == 5 {
		// 	numVal++ //do this otherwise it will be infinite loop as counter never increments
		// 	continue
		// 	fmt.Println("will continue")
		// }

		if numVal == 5 {
			goto lco //using goto stmt to move on to any label declared
			//fmt.Println("loop execution stopped now") //unreachable
		}

		fmt.Println("number is ", numVal)
		numVal++
	}

	//goto statements, write lco & : after that to declare a goto (a label, just like var must be used if declared)
lco:
	fmt.Println("this is a goto statement :: loop execution stopped now")

}
