package main

import "fmt"

// in golang since we have structs instead of classes, we define funcs in structs & call them methods
func main() {
	fmt.Println("Welcome to methods in golang!")
	ashwani := User{"Ashwani", "ashwani@gmail.com", 26, true}
	ashwani.GetStatus() //using  a method
	ashwani.generateEmail()
	fmt.Printf("did the user details changed? %+v", ashwani) //so the method did not change the value in struct
	//but passes along a copy. so functions/methods pass along a copy of object, use pointers to verify
	//if you want changes in the original object pass a pointer of it/reference of it
}

type User struct {
	Name   string
	Email  string
	Age    int32
	Status bool
}

func (user User) GetStatus() { //keep 1st letter caps for public if you want to export it
	fmt.Println("Is user active: ", user.Status)
}

func (user User) generateEmail() {
	user.Email = "tester@go.dev"
	fmt.Println("new Email is : ", user.Email)
}
