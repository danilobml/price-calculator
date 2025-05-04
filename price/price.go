package price

import "fmt"

type TaxedPricesMap map[string]float64

type TaxIncludedPriceJob struct {
	TaxRate           float64        `json:"taxRate"`
	InputPrices       []float64      `json:"inputPrices"`
	TaxIncludedPrices TaxedPricesMap `json:"taxIncludedFiles"`
}

func NewTaxIncludedPriceJob(taxRate float64, inputPrices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
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

func GenerateJobs(taxRates, prices []float64) []TaxIncludedPriceJob {
	var jobs []TaxIncludedPriceJob
	for _, taxRate := range taxRates {
		newJob := NewTaxIncludedPriceJob(taxRate, prices)
		newJob.Process()
		jobs = append(jobs, *newJob)
	}

	return jobs
}
