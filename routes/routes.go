package routes

import (
	"simulator-api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, h *handlers.SimulationHandler) {
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API funcionando"})
	})

	router.POST("/simulation", h.CalculateSimulation)
}