package main

import "fmt"

// Customer represents information about a customer.
type Customer struct {
	Name  string
	Email string
}

func printTitle(title string) {
	fmt.Println(title)
}

func main() {
	var title string
	var author string
	var edition int
	var copies int
	var inStock bool
	var royaltyPercentage float64
	var specialOffer bool
	var discountPercentage float64
	var x int
	title = "For the Love of Go"
	author = "Jon Arundel"
	edition = 1
	copies = 99
	inStock = true
	royaltyPercentage = 12.5
	specialOffer = true
	discountPercentage = 25
	printTitle(title)
	fmt.Println(author)
	fmt.Println(edition)
	fmt.Println(copies)
	fmt.Println(inStock)
	fmt.Println(royaltyPercentage)
	fmt.Println(specialOffer)
	fmt.Println(discountPercentage)
	fmt.Println(x)
}
