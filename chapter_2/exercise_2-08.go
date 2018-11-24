package main

import (
	"fmt"
	"strings"
)

func main() {

}

// convertIntegerToString returns the binary representation of an integer as a string
func convertIntegerToString(i, base int) (string, error) {
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
	switch remainder {
	case 0:
		return "0", nil
	case 1:
		return "1", nil
	case 2:
		return "2", nil
	case 3:
		return "3", nil
	case 4:
		return "4", nil
	case 5:
		return "5", nil
	case 6:
		return "6", nil
	case 7:
		return "7", nil
	case 8:
		return "8", nil
	case 9:
		return "9", nil
	case 10:
		return "a", nil
	case 11:
		return "b", nil
	case 12:
		return "c", nil
	case 13:
		return "d", nil
	case 14:
		return "e", nil
	case 15:
		return "f", nil
	}

	return "", fmt.Errorf("Not a valid remainder: %v", remainder)
}

// getDigitInt converts the digit rune to an integer
func getDigitInt(digit rune) (int, error) {
	switch digit {
	case '0':
		return 0, nil
	case '1':
		return 1, nil
	case '2':
		return 2, nil
	case '3':
		return 3, nil
	case '4':
		return 4, nil
	case '5':
		return 5, nil
	case '6':
		return 6, nil
	case '7':
		return 7, nil
	case '8':
		return 8, nil
	case '9':
		return 9, nil
	case 'a':
		return 10, nil
	case 'b':
		return 11, nil
	case 'c':
		return 12, nil
	case 'd':
		return 13, nil
	case 'e':
		return 14, nil
	case 'f':
		return 15, nil
	}
	return 0, fmt.Errorf("Not a valid digit: %v", digit)
}

// getDecimal returns an integer from a binary number represented by a string
func convertStringToInt(b string, base int) (int, error) {
	decimal := 0

	for i, v := range b {
		digit, err := getDigitInt(v)

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
