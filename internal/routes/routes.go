package routes

import (
	"simulator-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET(
		"/test",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "API funcionando",
			})
		},
	)

	router.POST(
		"/simulation",
		handlers.CalculateSimulation,
	)


}