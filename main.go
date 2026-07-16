package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"simulator-api/data/database"
	datarepo "simulator-api/data/repository"
	"simulator-api/domain/usecase"
	"simulator-api/handlers"
	"simulator-api/routes"

)

func main() {

	// Carrega variáveis de ambiente do arquivo .env antes de iniciar a conexão com o PostgreSQL
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

    if err := database.Init(); err != nil {
        log.Fatalf("Erro ao conectar ao banco: %v", err)
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
