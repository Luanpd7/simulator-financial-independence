package handlers

import (
	"log"
	"simulator-api/data/repository"
	"simulator-api/domain/entities"
	"github.com/gin-gonic/gin"
)

func CalculateSimulation(c *gin.Context) {
	

	var input entities.FinancialIndependence
	if err := c.BindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}


	// Salvar no banco
	record := repository.SimulationRecord{
		CurrentAssets:       input.CurrentAssets,
		MonthlyContribution: input.MonthlyContribution,
		AnnualPercentage:    input.AnnualPercentage,
		CurrentAge:          input.CurrentAge,
		RetirementAge:       input.RetirementAge,
		TimeInYears:         input.TimeInYears,
		Inflation:           input.Inflation,
	}

	if err := repository.SaveSimulation(record); err != nil {
		log.Printf("⚠️  Erro ao salvar simulação no banco: %v\n", err)
	}

	c.JSON(200, gin.H{
		"savedAt":           record.CreatedAt,
	})
}