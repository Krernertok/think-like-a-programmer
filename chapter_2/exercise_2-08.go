package main

import (
	"fmt"
	"strings"
)

func main() {
	var number string
	var baseFrom int
	var baseTo int

	fmt.Println("The number you want to convert followed by its base (2-16), ex. 101 2:")
	fmt.Scanln(&number, &baseFrom)

	if baseFrom > 16 || baseFrom < 0 {
		fmt.Println("Base must be between 2 and 16.")
		return
	}

	fmt.Println("The base you want to convert the number to (2-16):")
	fmt.Scanln(&baseTo)

	if baseTo > 16 || baseTo < 0 {
		fmt.Println("Base must be between 2 and 16.")
		return
	}

	numberInt, err := convertStringToInt(number, baseFrom)

	if err != nil {
		fmt.Println(err)
		return
	}

	convertedNumber, err := convertIntToString(numberInt, baseTo)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s (base %d) converted to base %d: %s\n", number, baseFrom, baseTo, convertedNumber)
}

// convertIntegerToString returns the binary representation of an integer as a string
func convertIntToString(i, base int) (string, error) {
	binary := []string{}

	for i > 0 {
		digit, err := getDigitString(i % base)
		if err != nil {
			return "", err
		}
		binary = append([]string{digit}, binary...)
		i /= base
	}

	return strings.Join(binary, ""), nil
}

// getDigitString converts the remainder (int) to a string
func getDigitString(remainder int) (string, error) {
	digits := []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
	}

	if remainder < 0 || remainder >= len(digits) {
		return "", fmt.Errorf("Not a valid remainder: %v", remainder)
	}

	return digits[remainder], nil
}

// getDigitInt converts the digit rune to an integer
func getDigitInt(digit rune) (int, error) {
	digits := map[rune]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'a': 10,
		'b': 11,
		'c': 12,
		'd': 13,
		'e': 14,
		'f': 15,
	}

	value, ok := digits[digit]

	if !ok {
		return 0, fmt.Errorf("Not a valid digit: %v", string(digit))
	}

	return value, nil
}

// getDecimal returns an integer from a binary number represented by a string
func convertStringToInt(b string, base int) (int, error) {
	decimal := 0

	for i, v := range b {
		digit, err := getDigitInt(v)

		if digit > base-1 {
			return 0, fmt.Errorf("Invalid digit in number %s base %d", b, base)
		}

		if err != nil {
			return 0, err
		}

		exp := len(b) - i - 1
		if digit > 0 {
			decimal += digit * power(base, exp)
		}
	}

	return decimal, nil
}

// power calculates an integer x to its nth power
func power(x, n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		return 0
	}

	result := x
	for i := 1; i < n; i++ {
		result *= x
	}
	return result
}
