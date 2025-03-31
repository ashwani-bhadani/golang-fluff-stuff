package main

import "fmt"

const LoginName string = "sleepy walrus"

//in go const uppercase 1st letter means that variable is public, accessible everywhere

func main() {
	fmt.Println(LoginName)
	fmt.Printf("Type is : %T \n", LoginName)
	//var username string = "ashwani"
	//fmt.Println(username)
	//fmt.Printf("variable is of type: %T \n", username)

	name := "Ashwani"   // Go infers type as string implicitly
	age := 28           // Go infers type as int
	isDeveloper := true // Go infers type as bool
	/*Only works inside functions (not at package level).
		Cannot be used for reassigning (only declares new variables).
		    if num := 10; num%2 == 0 {
	        fmt.Println(num, "is even")
	    } else {
	        fmt.Println(num, "is odd")
	    }
	    // fmt.Println(num) // ❌ Error: num is not accessible here
	*/

	//implicit type declaration, now cannot change type later
	var website = "https://go.dev/ref/spec"
	fmt.Println(website)

	//no var style using walrus operator, allowed only inside methods not in global scopes, follow a syntax
	numberOfUserLimit := 70
	fmt.Println(numberOfUserLimit)
	//numberOfUserLimit = 89.120 //cannot use 89.120 (untyped float constant 89.12) as int value in assignment (truncated)

	if isDeveloper {
		fmt.Printf("My name is %s and I'm %d years old and I am a developer.\n", name, age)
	} else {
		fmt.Printf("My name is %s and I'm %d years old and I am not a developer.\n", name, age)
	}

	/*
			var isLoggedIn bool = true
			fmt.Println(isLoggedIn)
			fmt.Printf("type of  logged in is : %T \n", isLoggedIn)

		var smallVal uint64 = 256
		fmt.Println(smallVal)
		fmt.Printf("Remember unint has diff sizes best for OS/server, here it is : %T \n", smallVal)
	*/

	var floatVal float64 = 255.455572787847478343748 //more bit num means more precision
	fmt.Println(floatVal)
	fmt.Printf("Remember unint has diff sizes best for OS/server, here it is : %T \n", floatVal)

	//default values & aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Default value is 0, here it is : %T \n", anotherVar)

}
