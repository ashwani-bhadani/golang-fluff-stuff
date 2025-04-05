package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps in golang")
	//new keyword initialises as zero started value, can give errors, make does not give zero started value
	langs := make(map[string]string) //map here is hashtable or key value, can use make for maps just like slices
	langs["JS"] = "javascript"
	langs["RB"] = "Ruby"
	langs["JA"] = "Java"
	langs["RU"] = "Rust"
	langs["PY"] = "Python" //adding values to map

	fmt.Println("list of all languages: ", langs)
	fmt.Println("key JS value : ", langs["JS"]) //getting value for key

	delete(langs, "RB") //delete based on key, can use delete for slices too!
	fmt.Println("list of all languages: ", langs)

	//looping through maps, like forEach in java
	for key, value := range langs {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	//using comma-ok syntax as you don't care about key, walrus operator helps you in that
	for _, value := range langs {
		fmt.Printf("value is %v\n", value) //now you cannot use keys
	}

}
