package main

import "fmt"

func update(yourmap map[string]int, key string, element int) {
	yourmap[key] = element // map is reference type.
}

func main() {
	// Empty map
	mymap := make(map[string]int)
	customers := map[string]int{}

	// Non-empty map
	employees := map[string]int{
		"John":  1001,
		"Ezira": 1002,
		"Emily": 1003,
	}

	// Add or Update
	update(mymap, "Apple", 23)
	customers["John"] = 36

	// Access
	value, status := employees["John"]
	fmt.Println(value, status)

	// Delete
	delete(employees, "John")
	fmt.Println(mymap, customers, employees)
}
