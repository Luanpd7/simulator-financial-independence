package entities
import "time"

type FinancialIndependence struct {
	CurrentAssets       float64 `json:"currentAssets"`
	MonthlyContribution float64 `json:"monthlyContribution"`
	AnnualPercentage    float64 `json:"annualPercentage"`
	CurrentAge          int64   `json:"currentAge"`
	RetirementAge       int64   `json:"retirementAge"`
	TimeInYears         int64   `json:"timeInYears"`
	Inflation           float64 `json:"inflation"`
}


type SimulationRecord struct {
	//Registrado automaticamente quando a simulação é salva no banco
	CreatedAt time.Time

	//Atualizado automaticamente toda vez que o registro for atualizado
	UpdatedAt time.Time

	//Valor do patrimônio atual
	CurrentAssets float64

	//Aporte mensal
	MonthlyContribution float64

	//Percentual de retorno anual esperado
	AnnualPercentage float64

	//Idade atual
	CurrentAge int64

	//Idade de aposentadoria
	RetirementAge int64

	//Opcioanl tempo em anos para a aposentadoria, caso o usuário queira calcular com base nisso ao invés da idade de aposentadoria
	TimeInYears int64

	//Inflação
	Inflation float64

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