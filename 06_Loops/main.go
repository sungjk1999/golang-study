package main

import "fmt"

func main() {

	menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

	fmt.Println("The menu:")
	for idx, food := range menu {
		fmt.Println(idx, food)
	}

	orders := map[string]string{
		"John":   "Cheeseburgers",
		"Janet":  "Hamburgers",
		"Jordan": "Fries",
	}

	orders["James"] = "Chicken Sandwich"

	fmt.Println("\nThe friend's orders:")

	for friend, order := range orders {
		fmt.Println(friend, ":", order)
	}
}
