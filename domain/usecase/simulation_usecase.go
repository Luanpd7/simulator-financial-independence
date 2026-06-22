package usecase

import (
	"math"
	"time"

	"simulator-api/domain/entities"
	domainRepo "simulator-api/domain/repositories"
)

// SimulationUseCase contains business rules for financial independence simulation.
type SimulationUseCase struct {
	repository domainRepo.SimulationRepository
}

func NewSimulationUseCase(repo domainRepo.SimulationRepository) *SimulationUseCase {
	return &SimulationUseCase{repository: repo}
}

// SaveFinanceIndependence executes the simulation (business rules) and persists the record.
func (uc *SimulationUseCase) SaveFinanceIndependence(input entities.FinancialIndependence) (entities.SimulationResult, error) {
	// Determine years to retirement
	var years int64
	if input.TimeInYears > 0 {
		years = input.TimeInYears
	} else {
		years = input.RetirementAge - input.CurrentAge
		if years < 0 {
			years = 0
		}
	}

	months := float64(years * 12)
	monthlyRate := (input.AnnualPercentage / 100.0) / 12.0
	var finalAmount float64
    
	if monthlyRate != 0 {
		finalAmount = input.CurrentAssets * math.Pow(1+monthlyRate, months)
		fvAnnuity := input.MonthlyContribution * ((math.Pow(1+monthlyRate, months) - 1) / monthlyRate)
		finalAmount += fvAnnuity
	} else {
		finalAmount = input.CurrentAssets + input.MonthlyContribution*months
	}
    
	inflationAdj := finalAmount / math.Pow(1+(input.Inflation/100.0), float64(years))
	totalContributed := input.CurrentAssets + input.MonthlyContribution*months
	totalInterest := finalAmount - totalContributed
      
	result := entities.SimulationResult{
		FinalAmount:             finalAmount,
		InflationAdjustedAmount: inflationAdj,
		YearsToRetirement:       years,
		RealRateYear:            input.AnnualPercentage - input.Inflation,
		RealRateMonth:           monthlyRate - (input.Inflation/100.0)/12.0,
		TotalContributed:        totalContributed,
		TotalInterestEarned:     totalInterest,
	}

	// Build record to persist
	record := entities.SimulationRecord{
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
		CurrentAssets:           input.CurrentAssets,
		MonthlyContribution:     input.MonthlyContribution,
		AnnualPercentage:        input.AnnualPercentage,
		CurrentAge:              input.CurrentAge,
		RetirementAge:           input.RetirementAge,
		TimeInYears:             input.TimeInYears,
		Inflation:               input.Inflation,
		FinalAmount:             result.FinalAmount,
		InflationAdjustedAmount: result.InflationAdjustedAmount,
		YearsToRetirement:       result.YearsToRetirement,
		RealRateYear:            result.RealRateYear,
		RealRateMonth:           result.RealRateMonth,
		TotalContributed:        result.TotalContributed,
		TotalInterestEarned:     result.TotalInterestEarned,
	}

	// Persist using repository (data layer implementation injected via DI)
	if err := uc.repository.Save(record); err != nil {
		return entities.SimulationResult{}, err
	}

	return result, nil
}
