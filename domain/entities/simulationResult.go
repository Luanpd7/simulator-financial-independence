package entities

type SimulationResult struct {
    //Patrimônio final
	FinalAmount float64

	//Patrimônio corrigido pela inflação
	InflationAdjustedAmount float64

	//Anos até a aposentadoria
	YearsToRetirement int64

	//Juros real por ano descontando a inflação
	RealRateYear float64

	//Juros real por mês descontando a inflação
	RealRateMonth float64

	//Quanto saiu do bolso
	TotalContributed float64

	//Quanto rendeu no total
	TotalInterestEarned float64
}