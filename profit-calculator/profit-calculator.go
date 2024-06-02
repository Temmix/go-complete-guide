package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("What is your Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("What is your Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("What is your tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf(`Expenses before tax is: %.2f`, ebt)
	fmt.Println()
	fmt.Printf("Your Profit is: %.2f", profit)
	fmt.Println()
	fmt.Printf("Ratio of Expenses before tax to profit is: %.2f\n", ratio)
}
