package main

import "fmt"

func main() {

	// Println
	fmt.Println("Let's first see how", "the Println() method works.")
	fmt.Println("Notice that each statement adds a newline for us.")
	fmt.Println("There's also a default space", "between the string arguments.")

	// Print
	fmt.Print("Print", "is", "different")
	fmt.Print("See?")

	// Printf
	animal1 := "cat"
	animal2 := "dog"
	fmt.Printf("Are you a %v or a %v person?", animal1, animal2)

	// Sprintln
	step1, step2 := "Breathe in...", "Breathe out..."
	meditation := fmt.Sprintln(step1, step2)
	fmt.Println(meditation)

	// Sprintf
	template := "I wish I had a %v."
	pet := "puppy"
	wish := fmt.Sprintf(template, pet)
	fmt.Println(wish)

	// Scan
	fmt.Println("What would you like for lunch?")
	var food string
	fmt.Scan(&food)
	fmt.Printf("Sure, we can have %v for lunch.", food)
}
