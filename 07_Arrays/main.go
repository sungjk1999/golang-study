package main

import "fmt"

func changeLastElement(slice []int, element int) {
	// Original slice value be changed
	slice[len(slice)-1] = element
}

func appendLastElement(slice []int, element int) {
	// Not work!
	slice = append(slice, element)
}

func main() {

	// Explict
	var arr1 = [5]int{1, 2, 3, 4, 5} // or
	var arr2 = [...]int{1, 2, 3, 4, 5}

	// Implict
	arr3 := [1]string{"Apple"}
	fmt.Println(arr1, arr2, arr3)

	// Explict
	var slice1 = []int{1, 2, 3}

	// Implict
	slice2 := []string{"Samsung"}
	fmt.Println(slice1, slice2)

	changeLastElement(slice1, 5)
	fmt.Println(slice1)

	appendLastElement(slice1, 7) // Not work.
	fmt.Println(slice1)

	slice1 = append(slice1, 7)
	fmt.Println(slice1)
}
