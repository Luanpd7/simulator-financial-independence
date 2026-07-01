package entities

type SimulationResult struct {
    //Patrimônio final
	FinalAmount float64 `json:"finalAmount"`

	//Patrimônio corrigido pela inflação
	InflationAdjustedAmount float64 `json:"inflationAdjustedAmount"`

	//Anos até a aposentadoria
	YearsToRetirement int64 `json:"yearsToRetirement"`

	//Juros real por ano descontando a inflação
	RealRateYear float64 `json:"realRateYear"`

	//Juros real por mês descontando a inflação
	RealRateMonth float64 `json:"realRateMonth"`

	//Quanto saiu do bolso
	TotalContributed float64 `json:"totalContributed"`

	//Quanto rendeu no total
	TotalInterestEarned float64 `json:"totalInterestEarned"`
}