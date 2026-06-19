package repositories

import "simulator-api/internal/domain/entities"

type SimulationRepository interface {
    Calculate(
        input entities.FinancialIndependence,
    ) entities.SimulationResult
}