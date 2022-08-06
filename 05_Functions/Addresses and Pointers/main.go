package main

import "fmt"

func brainwash(saying *string) {
	// Dereference saying below:
	*saying = "Beep Boop."
}

func main() {
	star := "Polaris"
	// Pointer of String
	var starAddress *string
	// Assign Address of `star`
	starAddress = &star
	// Dereferencing
	*starAddress = "Sirius"
	fmt.Println("The actual value of star is", star)

	// Changing Values in Different Scopes
	greeting := "Hello there!"

	// Call your brainwash() below:
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)
}
