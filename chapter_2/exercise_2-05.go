package main

import (
	"fmt"
)

// ISBN validator
func validISBN(input string) bool {
	return false
}

// ISBN check digit generator
func generateValidISBN(input string) string {
	return ""
}

func main() {
	isbn := "XXX"
	fmt.Println("XXX is a valid ISBN number:", validISBN(isbn))

	incompleteIsbn := "YYY"
	fmt.Println("Valid ISBN for", incompleteIsbn, "is",
		generateValidISBN(incompleteIsbn))
}
