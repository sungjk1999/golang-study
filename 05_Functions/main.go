package main

import "fmt"

func getLikesAndShares(postId int) (int, int) {
	var likesForPost, sharesForPost int
	switch postId {
	case 1:
		likesForPost = 5
		sharesForPost = 7
	case 2:
		likesForPost = 3
		sharesForPost = 11
	case 3:
		likesForPost = 22
		sharesForPost = 1
	case 4:
		likesForPost = 7
		sharesForPost = 9
	}
	fmt.Println("Likes: ", likesForPost, "Shares: ", sharesForPost)
	return likesForPost, sharesForPost
}

func calculateTaxes(revenue, deductions, credits float64) float64 {
	// Deferring Resolution
	defer fmt.Println("Taxes Calculated!")
	taxRate := .06143
	fmt.Println("Calculating Taxes")

	if deductions == 0 || credits == 0 {
		return revenue * taxRate
	}

	taxValue := (revenue - (deductions * credits)) * taxRate
	if taxValue >= 0 {
		return taxValue
	} else {
		return 0
	}
}

func main() {

	// Multiple return vals
	var likes, shares int
	likes, shares = getLikesAndShares(4)

	if likes > 5 {
		fmt.Println("Woohoo! We got some likes.")
	}
	if shares > 10 {
		fmt.Println("We went viral!")
	}

	// Deferring Resolution
	calculateTaxes(100, 10, 10)
}
