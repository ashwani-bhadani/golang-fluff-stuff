package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//there is no inheritance in golang, no super keyword or parent as these makes code complex to read, structs is classes

	ashwani := User{"Ashwani", "ashwani@gmail.com", true, 26}
	fmt.Println(ashwani)
	fmt.Printf("Type of ashwani is : %T\n", ashwani)
	fmt.Printf("Name is : %v\nEmail is: %v\n", ashwani.Name, ashwani.Email) //%v is just for value
	//fmt.Printf("Name is : %+v\nEmail is: %+v\n", ashwani.Name, ashwani.Email) //!!! will not give details
	//%+v is to log struct type, will give field name, structure too, not just the values, details too
	fmt.Printf("details of user : %+v\n", ashwani)

}

type User struct {
	//User is a class so needs to be exported out
	//first letter as caps says these are public
	Name         string
	Email        string
	ActiveStatus bool
	Age          int
}
