package main

import (
	"fmt"
	"time" // Go's built-in time package
)

func main() {
	fmt.Println("Welcome to time manager")

	// Get the current time
	presentTime := time.Now()

	// Format the time for better readability
	formattedTime := presentTime.Format("2006-01-02 15:04:05") // YYYY-MM-DD HH:MM:SS
	fmt.Println("Current time:", formattedTime)
}
