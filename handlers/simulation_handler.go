package handlers

import (
	"net/http"
	"simulator-api/domain/entities"
	"simulator-api/domain/usecase"
	"github.com/gin-gonic/gin"
)

// SimulationHandler handles HTTP requests and delegates to usecases.
type SimulationHandler struct {
	uc *usecase.SimulationUseCase
}

func NewSimulationHandler(uc *usecase.SimulationUseCase) *SimulationHandler {
	return &SimulationHandler{uc: uc}
}

func (h *SimulationHandler) CalculateSimulation(c *gin.Context) {
	var input entities.FinancialIndependence
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	result, err := h.uc.SaveFinanceIndependence(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to process simulation"})
		return
	}

	c.JSON(http.StatusOK, result)
}