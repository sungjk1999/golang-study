package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// if - else if - else
	amountStolen := 64650
	if amountStolen > 1000000 {
		fmt.Println("We've hit the jackpot!")
	} else if amountStolen >= 5000 {
		fmt.Println("Think of all the candy we can buy!")
	} else {
		fmt.Println("Why did we even do this?")
	}

	// switch
	name := "H. J. Simp"
	switch name {
	case "Butch":
		fmt.Println("Head to Robbers Roost.")
	case "Bonnie":
		fmt.Println("Stay put in Joplin.")
	default:
		fmt.Println("Just hide!")
	}

	// Scoped Short Declaration Statement
	if success := true; success {
		fmt.Println("We're rich!")
	} else {
		fmt.Println("Where did we go wrong?")
	}

	// Randomizing
	amountLeft := rand.Intn(10000)
	fmt.Println("amountLeft is: ", amountLeft)

	// Random Seeding
	rand.Seed(time.Now().UnixNano())
	newLeft := rand.Intn(10000)
	fmt.Println("amountLeft is: ", newLeft)
}
