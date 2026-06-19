package main

import (
	"log"
	"simulator-api/data/database"
	"simulator-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar banco de dados
	if err := database.Init("simulator.db"); err != nil {
		log.Fatalf("❌ Erro ao inicializar banco de dados: %v\n", err)
	}
	defer database.Close()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.Run(":8080")
}