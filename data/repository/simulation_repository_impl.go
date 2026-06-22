package repository

import (
	"log"
	"simulator-api/data/database"
	"simulator-api/domain/entities"
)

// SimulationRepositoryImpl implements persistence for simulation records.
type SimulationRepositoryImpl struct{}

func NewSimulationRepositoryImpl() *SimulationRepositoryImpl {
	return &SimulationRepositoryImpl{}
}

func (r *SimulationRepositoryImpl) Save(record entities.SimulationRecord) error {
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
