package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (100 - taxRate) / 100
	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Printf("Earnings Before Tax: %f\n", earningsBeforeTax)
	fmt.Printf("Earnings After Tax: %f\n", earningsAfterTax)
	fmt.Printf("Ratio: %f\n", ratio)
}
