package entities

type FinancialIndependence struct {
	CurrentAssets       float64 `json:"currentAssets"`
	MonthlyContribution float64 `json:"monthlyContribution"`
	AnnualPercentage    float64 `json:"annualPercentage"`
	CurrentAge          int64   `json:"currentAge"`
	RetirementAge       int64   `json:"retirementAge"`
	TimeInYears         int64   `json:"timeInYears"`
	Inflation           float64 `json:"inflation"`
}


