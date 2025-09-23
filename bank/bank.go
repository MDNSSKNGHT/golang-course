package main

import "fmt"

func main() {
	accountBalance := 1000.0

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
			fmt.Println("Balance updated! New account balance:", accountBalance)
		default:
			fmt.Println("Thanks for vising Go Bank!")
			loop = false
		}
		fmt.Println()
	}
}
