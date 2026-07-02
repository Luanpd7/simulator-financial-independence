package main

import (
	"log"
	"fmt"
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



	
    var ponteiroValue *int

	value := 1

	ponteiroValue = &value



	fmt.Println("Value:", value) 
	fmt.Println("Ponteiro:", ponteiroValue) 
	fmt.Println("Usando o valor daquele ponteiro:", *ponteiroValue) 


	defer database.Close()

	// Dependency injection
	repo := datarepo.NewSimulationRepositoryImpl()
	uc := usecase.NewSimulationUseCase(repo)
	handler := handlers.NewSimulationHandler(uc)

	router := gin.Default()
	router.Use(cors.Default())
	routes.RegisterRoutes(router, handler)
	router.Run(":8080")
	
}