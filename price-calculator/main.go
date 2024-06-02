package main

import (
	"fmt"

	"temmix.com/price-calculator/filemanager"
	//"temmix.com/price-calculator/cmdmanager"
	"temmix.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.7}
	doneChans := make([]chan bool, len(taxRates))   // for concurrency
	errorChans := make([]chan error, len(taxRates)) // for concurrency

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
		// if err != nil {
		// 	fmt.Println("Could not process the job")
		// 	fmt.Println(err)
		// }
	}

	for index := range doneChans {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Completed the task: ", index+1)
		}
	}
}
