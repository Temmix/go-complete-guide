package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err == nil {
		return balance
	}
	return 0
}

func writeBalanceToFile(balance float64) error {
	balanceText := fmt.Sprint(balance)
	return os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	fmt.Println("Welcome to Structure Bank Plc")
	fmt.Println()

	// Using infinite loops
	for {
		var accountBalance = getBalanceFromFile()
		var choice int
		fmt.Println("1. Print your bank account balance")
		fmt.Println("2. Make a deposit")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Print("What operation will you like to perform, options above: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("******************************************************")
			fmt.Println("Your account balance is", accountBalance)
			fmt.Println("******************************************************")

		case 2:
			var depositAmount float64
			fmt.Print("How much do you want to deposit:")
			fmt.Scan(&depositAmount)

			fmt.Println("******************************************************")
			if depositAmount <= 0 {
				fmt.Println("Deposit amount should be greater than 0")
				fmt.Println("******************************************************")
				continue
			}
			accountBalance += depositAmount
			err := writeBalanceToFile(accountBalance)
			if err == nil {
				fmt.Println("Your new account balance is", accountBalance)
				fmt.Println("******************************************************")
			}

		case 3:
			var withdrawalAmount float64
			fmt.Print("How much do you want to withdrawal:")
			fmt.Scan(&withdrawalAmount)

			fmt.Println("******************************************************")
			if withdrawalAmount <= 0 {
				fmt.Println("Withdrawal amount should be greater than 0")
				fmt.Println("******************************************************")
				continue
			}
			if withdrawalAmount > accountBalance {
				fmt.Println("Withdrawal amount can not be more than the account balance")
				fmt.Println("******************************************************")
				continue
			}
			accountBalance -= withdrawalAmount
			err := writeBalanceToFile(accountBalance)
			if err == nil {
				fmt.Println("Your new account balance is", accountBalance)
				fmt.Println("******************************************************")
			}
		default:
			fmt.Println("Goodbye and thanks for banking with us")
			return
		}
	}
}
