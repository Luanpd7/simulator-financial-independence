package repositories

import "simulator-api/domain/entities"

// SimulationRepository defines persistence operations for simulation records.
// Business rules (calculations) must live in usecases, not here.
type SimulationRepository interface {
	Save(record entities.SimulationRecord) error
}