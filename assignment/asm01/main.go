package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Printf("Need at least 3 arguments passed, has only: %d", len(args))
		return
	}

	code := args[len(args)-1]

	if !isValidCountryCode(code) {
		fmt.Printf("Invalid country code: %s\n", code)
		return
	}

	rawName := args[0 : len(args)-1]
	fullName := getFullNameFromCountryCode(rawName, code)

	fmt.Printf("Output: %s", fullName)
}

func isValidCountryCode(code string) bool {
	switch code {
	case "VN", "US":
		return true
	default:
		return false
	}
}

func getFullNameFromCountryCode(rawName []string, code string) string {
	firstName, lastName := rawName[0], rawName[1]
	middleName := " "
	fullName := ""

	if len(rawName) > 2 {
		middleName = " " + strings.Join(rawName[2:], " ") + " "
	}

	switch code {
	case "US":
		fullName = firstName + middleName + lastName
	case "VN":
		fullName = lastName + middleName + firstName
	}
	return fullName
}
