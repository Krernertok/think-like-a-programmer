package main

import (
	"fmt"
)

// ISBN validator
func validISBN(input string) bool {
	digits, err := getISBNDigits(input)

	if err != nil || len(digits) != 13 {
		return false
	}

	sum := getISBNDigitSum(digits)

	return sum%10 == 0
}

// Generates an ISBN check digit from 12 digits
func getISBNCheckDigit(input string) (int, error) {
	digits, err := getISBNDigits(input)

	if err != nil {
		return 0, err
	}

	if len(digits) != 12 {
		err := fmt.Errorf("incorrect amount of digits in %s to calculate check digit", input)
		return 0, err
	}

	sum := getISBNDigitSum(digits)

	check := 10 - (sum % 10)
	if check == 10 {
		check = 0
	}

	return check, nil
}

// getISBNDigits returns a slice of the digits from a string consisting of
// numerals and hyphens
func getISBNDigits(input string) ([]int, error) {
	digits := []int{}

	for _, r := range input {
		if r >= '0' && r <= '9' {
			digits = append(digits, int(r-'0'))
		} else if r == '-' {
			continue
		} else {
			err := fmt.Errorf("invalid characters in input string: %s", input)
			return nil, err
		}
	}

	return digits, nil
}

// getISBNDigitSum returns the sum of digits for determining a valid ISBN or
// to calculate an ISBN check digit
func getISBNDigitSum(digits []int) int {
	total := 0
	for i, d := range digits {
		if i%2 == 0 {
			total += d
		} else {
			total += 3 * d
		}
	}
	return total
}

func main() {
	valid := "978-0131103627"
	invalid := "978-0131103629"
	invalid2 := "978-A1311X3629"

	fmt.Println(valid, "is a valid ISBN:", validISBN(valid))
	fmt.Println(invalid, "is a valid ISBN:", validISBN(invalid))
	fmt.Println(invalid2, "is a valid ISBN:", validISBN(invalid2))

	fmt.Println()

	incomplete1 := "978-013110362"
	checkDigit, err := getISBNCheckDigit(incomplete1)
	fmt.Println("Check digit for", incomplete1, "is", checkDigit, "Error:", err)

	incomplete2 := "978-3-16-148410"
	checkDigit, err = getISBNCheckDigit(incomplete2)
	fmt.Println("Check digit for", incomplete2, "is", checkDigit, "Error:", err)

	incomplete3 := "978-01311036"
	checkDigit, err = getISBNCheckDigit(incomplete3)
	fmt.Println("Check digit for", incomplete3, "is", checkDigit, "Error:", err)

	incomplete4 := "978-0131A036"
	checkDigit, err = getISBNCheckDigit(incomplete4)
	fmt.Println("Check digit for", incomplete4, "is", checkDigit, "Error:", err)

}
