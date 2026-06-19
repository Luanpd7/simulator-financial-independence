package repository

import (
	"log"
	"time"
	"simulator-api/data/database"
)

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

func SaveSimulation(record SimulationRecord) error {
	query := `
	INSERT INTO simulations (
		current_assets,
		monthly_contribution,
		annual_percentage,
		current_age,
		retirement_age,
		time_in_years,
		inflation,
		final_amount,
		inflation_adjusted_amount,
		years_to_retirement,
		real_rate_year,
		real_rate_month,
		total_contributed,
		total_interest_earned
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := database.DB.Exec(
		query,
		record.CurrentAssets,
		record.MonthlyContribution,
		record.AnnualPercentage,
		record.CurrentAge,
		record.RetirementAge,
		record.TimeInYears,
		record.Inflation,
		record.FinalAmount,
		record.InflationAdjustedAmount,
		record.YearsToRetirement,
		record.RealRateYear,
		record.RealRateMonth,
		record.TotalContributed,
		record.TotalInterestEarned,
	)

	if err != nil {
		log.Printf("❌ Erro ao salvar simulação: %v\n", err)
		return err
	}

	return nil
}




type SimulationRepositoryImpl struct{}
