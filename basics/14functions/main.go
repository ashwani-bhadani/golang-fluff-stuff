package main

import "fmt"

func main() {
	fmt.Println("understanding functions in golang")
	// greeter //this means passing reference of a funciton
	greeter()
	fmt.Println("adding two nos.: ", adder(23, 45))
	fmt.Println("adding lots of nos.: ", proAdder(23, 45, 12, 34, 32, 32, 43, 12))
	// result, _ := adderMulti(89, 98)
	result, message := adderMulti(89, 98)
	fmt.Printf("%v result is: %v \n", message, result)
}

// lamdas & immediately executable funcs also there
func greeter() {
	fmt.Println("Hello from golang !!")
}

// Generics allow you to write functions, types, and data structures that can work with any type,
// while maintaining type safety. Available in Go starting from Go 1.18 (released in March 2022).
func adder(intVal1 int, intVal2 int) int { //need to declare type of params & also return type
	return intVal1 + intVal2
}

// multiple return values
func adderMulti(intVal1 int, intVal2 int) (int, string) { //since multiple, enclose in parantheses
	return intVal1 + intVal2, "Addition Completed !"
}

// "..." is called veriadic function
func proAdder(values ...int) int { //have any amount of input params, values is now a slice
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
