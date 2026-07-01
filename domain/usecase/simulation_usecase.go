package usecase

import (
	"fmt"
	"math"
	"simulator-api/domain/entities"
	domainRepo "simulator-api/domain/repositories"
	"time"
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
	years = calculeYearsToRetirement(input)

	fmt.Print("input.CurrentAssets-", input.CurrentAssets)
	fmt.Print("input.MonthlyContribution-", input.MonthlyContribution)

	/// Calculate total contributed and total interest earned
	totalContributed := resultTotalContributed(input, years)

	/// Calculate final amount using the finalAmount function
	finalAmount := finalAmount(input, years)

	// Calculate inflation adjusted final amount
	inflationAdj := adjustForInflation(finalAmount, input.Inflation, years)

	// Calculate total interest earned
	totalInterestEarned := resultTotalEarnings(finalAmount, totalContributed)

	// Calculate annual real interest rate
	annualRealInterest := annualRealInterest(input.AnnualPercentage, input.Inflation)

	// Calculate month real interest rates
	monthRealInterest := monthRealInterest(annualRealInterest)

	result := entities.SimulationResult{
		FinalAmount:             finalAmount,
		InflationAdjustedAmount: inflationAdj,
		YearsToRetirement:       years,
		RealRateYear:            annualRealInterest,
		RealRateMonth:           monthRealInterest,
		TotalContributed:        totalContributed,
		TotalInterestEarned:     totalInterestEarned,
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

func finalAmount(input entities.FinancialIndependence, years int64) float64 {
	monthlyRate := annualToMonthlyRate(input.AnnualPercentage)
	months := totalMonths(years)

	
	if monthlyRate == 0 {
		return input.CurrentAssets + input.MonthlyContribution*months
	}

	growthFactor := math.Pow(1+monthlyRate, months)

	currentAssetsValue := input.CurrentAssets * growthFactor

	contributionsValue := input.MonthlyContribution *
		((growthFactor - 1) / monthlyRate)

	return currentAssetsValue + contributionsValue
}

func annualToMonthlyRate(annualRate float64) float64 {
	const percent = 100.0

	annualRateDecimal := annualRate / percent
	return math.Pow(1+annualRateDecimal, 1.0/12.0) - 1
}

func adjustForInflation(value float64, inflationAnnual float64, years int64) float64 {
	return value / math.Pow(1+inflationAnnual/100.0, float64(years))
}

func calculeYearsToRetirement(input entities.FinancialIndependence) int64 {
	var years int64
	if input.TimeInYears > 0 {
		years = input.TimeInYears
	} else {
		years = input.RetirementAge - input.CurrentAge
		if years < 0 {
			years = 0
		}
	}
	return years
}

func resultTotalEarnings(finalAmount float64, totalContributed float64) float64 {
	return finalAmount - totalContributed
}

func resultTotalContributed(input entities.FinancialIndependence, years int64) float64 {
	months := totalMonths(years)
	return input.CurrentAssets + input.MonthlyContribution*months
}

func totalMonths(year int64) float64 {
	return float64(year * 12)
}

func returnRateMonth(annualRate float64) float64 {
	const percent = 100.0

	annualRateDecimal := annualRate / percent
	monthlyRate := math.Pow(1+annualRateDecimal, 1.0/12.0) - 1

	return monthlyRate
}

func annualRealInterest(annualInterest float64, inflation float64) float64 {
	const percent = 100.0

	annualRate := annualInterest / percent
	inflationRate := inflation / percent

	realRate := (1+annualRate)/(1+inflationRate) - 1

	return realRate * percent
}

func monthRealInterest(annualRealInterest float64) float64 {
	const percent = 100.0

	annualRate := annualRealInterest / percent
	monthlyRate := math.Pow(1+annualRate, 1.0/12.0) - 1

	return monthlyRate * percent
}
