package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const defaultAccountBalance = 1000
const accountBalanceFilename = "balance.txt"

func main() {
	accountBalance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----")
	}

	for loop := true; loop; {
		fmt.Println("Welcome to Go Bank!")
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check balance.")
		fmt.Println("2. Deposit money.")
		fmt.Println("3. Withdraw money.")
		fmt.Println("4. Quit.")

		fmt.Print("Your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Amount to deposit: ")

			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0!")
				continue
			}

			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New account balance:", accountBalance)
		case 3:
			fmt.Print("Amount to withdraw: ")

			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0!")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have!")
				continue
			}

			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New account balance:", accountBalance)
		default:
			fmt.Println("Thanks for vising Go Bank!")
			loop = false
		}
		fmt.Println()
	}
}

func writeBalanceToFile(balance float64) {
	balanceAsStr := fmt.Sprint(balance)

	os.WriteFile(accountBalanceFilename, []byte(balanceAsStr), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFilename)
	if err != nil {
		return defaultAccountBalance, errors.New("Failed to find balance file.")
	}
	balanceAsStr := string(data)

	balanceAsFloat64, err := strconv.ParseFloat(balanceAsStr, 64)
	if err != nil {
		return defaultAccountBalance, errors.New("Failed to parse balance value.")
	}

	return balanceAsFloat64, nil
}
