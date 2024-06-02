package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	investmentYears := 10.0
	expectedReturnRate := 5.5

	fmt.Print("What is the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("What is the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("What is the investment years: ")
	fmt.Scan(&investmentYears)
	fmt.Println()

	fmt.Println("===========================================================================")

	futureValue, futureRealValue := calculateInvestment(inflationRate, investmentAmount, investmentYears, expectedReturnRate)
	valueMsg := fmt.Sprintf("The future value is $%v", futureValue)
	fmt.Println(valueMsg)
	realValueMsg, _ := fmt.Printf("The future value considering inflation rate of %v is $%v", inflationRate, futureRealValue)
	fmt.Println(realValueMsg)

	fmt.Println("===========================================================================")
}

func calculateInvestment(inflationRate, investmentAmount, investmentYears, expectedReturnRate float64) (float64, float64) {
	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, investmentYears)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, investmentYears)

	return futureValue, futureRealValue
}
