package main

import (
	"log"
	"os"
	"simulator-api/data/database"
	datarepo "simulator-api/data/repository"
	"simulator-api/domain/usecase"
	"simulator-api/handlers"
	"simulator-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	// Inicializar banco de dados
	if err := database.Init("simulator.db"); err != nil {
		log.Fatalf("❌ Erro ao inicializar banco de dados: %v\n", err)
	}
	defer database.Close()

	// Dependency injection
	repo := datarepo.NewSimulationRepositoryImpl()
	uc := usecase.NewSimulationUseCase(repo)
	handler := handlers.NewSimulationHandler(uc)

	router := gin.Default()
	router.Use(cors.Default())
	routes.RegisterRoutes(router, handler)

	// Porta para produção (Railway) ou desenvolvimento local
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor rodando na porta %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}



