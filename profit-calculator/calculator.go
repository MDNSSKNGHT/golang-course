package main

import "fmt"

func main() {
	revenue := promptForFloatingValue("Revenue")
	expenses := promptForFloatingValue("Expenses")
	taxRate := promptForFloatingValue("Tax Rate")

	earningsBeforeTax, earningsAfterTax, ratio := calculateEarningsAndRatio(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func promptForFloatingValue(prompt string) float64 {
	var value float64

	fmt.Printf("%s: ", prompt)
	fmt.Scan(&value)

	return value
}

func calculateEarningsAndRatio(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (100 - taxRate) / 100
	ratio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, ratio
}
