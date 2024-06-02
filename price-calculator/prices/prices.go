package prices

import (
	"fmt"

	"temmix.com/price-calculator/iomanager"
	"temmix.com/price-calculator/scannerconversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	scanner, err := job.IOManager.Readfile()

	if err != nil {
		fmt.Println("opening file failed!")
		fmt.Println(err)
		return err
	}

	pricesList, err := scannerconversion.ConvertScannerTextToFloat(scanner)
	if err != nil {
		fmt.Println("Reading content file failed!")
		fmt.Println(err)
		return err
	}

	job.InputPrices = pricesList
	return nil
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadPrices()
	if err != nil {
		fmt.Println("Could not load prices!")
		errorChan <- err
		return
	}

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		priceValue := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = priceValue
	}
	//output := fmt.Sprintf("TaxRate of %v: %v", job.TaxRate, result)
	//fmt.Println(output)

	job.TaxIncludedPrices = result
	job.IOManager.WriteJSON(job)
	doneChan <- true
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
	}
}
