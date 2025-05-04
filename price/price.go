package price

import "fmt"

type TaxedPricesMap map[string]float64

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices TaxedPricesMap
}

func NewTaxIncludedPriceJob(taxRate float64, inputPrices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		InputPrices: inputPrices,
	}
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(TaxedPricesMap)
	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	job.TaxIncludedPrices = result
}
