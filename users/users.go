package main

import (
	"fmt"

	"example.com/users/user"
)

func main() {
	firstName := promptForUserData("Please enter your first name")
	lastName := promptForUserData("Please enter your last name")
	birthDate := promptForUserData("Please enter your birthdate (MM/DD/YYYY)")

	userVar := user.New(firstName, lastName, birthDate)
	userVar.PrintDetails()
}

func promptForUserData(prompt string) string {
	fmt.Printf("%s: ", prompt)

	var value string
	fmt.Scan(&value)

	return value
}
