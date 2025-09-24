package main

import (
	"errors"
	"fmt"
	"os"
)

const profitCalculatorFilename = "calculations.txt"

func main() {
	revenue, err := promptForFloatingValue("Revenue")
	if err != nil {
		printError(err)
		return
	}

	expenses, err := promptForFloatingValue("Expenses")
	if err != nil {
		printError(err)
		return
	}

	taxRate, err := promptForFloatingValue("Tax Rate")
	if err != nil {
		printError(err)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateEarningsAndRatio(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
	writeCalculationsToFile(earningsBeforeTax, earningsAfterTax, ratio)
}

func printError(err error) {
	fmt.Println("-----")
	fmt.Printf("ERROR: %v\n", err)
	fmt.Println("-----")
}

func writeCalculationsToFile(earningsBeforeTax, earningsAfterTax, ratio float64) {
	line1 := fmt.Sprintf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	line2 := fmt.Sprintf("Earnings After Tax: %.2f\n", earningsAfterTax)
	line3 := fmt.Sprintf("Ratio: %.2f\n", ratio)
	fileContents := fmt.Sprintf("%s%s%s", line1, line2, line3)

	os.WriteFile(profitCalculatorFilename, []byte(fileContents), 0644)
}

func promptForFloatingValue(prompt string) (float64, error) {
	var value float64

	fmt.Printf("%s: ", prompt)
	fmt.Scan(&value)

	if value <= 0 {
		return 0, errors.New("Value must be greater than 0.")
	}

	return value, nil
}

func calculateEarningsAndRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (100 - taxRate) / 100
	ratio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
