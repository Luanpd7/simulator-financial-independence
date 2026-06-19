package repositories

import "simulator-api/domain/entities"

type SimulationRepository interface {
    Calculate(
        input entities.FinancialIndependence,
    ) entities.SimulationResult
}