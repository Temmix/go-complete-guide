package scannerconversion

import (
	"bufio"
	"strconv"

	"fmt"
)

func ConvertScannerTextToFloat(scanner *bufio.Scanner) ([]float64, error) {
	var pricesList []float64
	for scanner.Scan() {
		floatPrice, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			printError(err)
			return nil, err
		}
		pricesList = append(pricesList, floatPrice)

		err = scanner.Err()
		if err != nil {
			printError(err)
			return nil, err
		}
	}
	return pricesList, nil
}

func printError(err error) {
	fmt.Println("Conversion of price to float failed")
	fmt.Println(err)
}
