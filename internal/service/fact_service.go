// service/fact_service.go

package service

import (
	"MainProject/internal/entity"
	"MainProject/internal/repository"
)

// FactService represents the service for managing Fact entities
type FactService struct {
	FactRepo repository.FactRepository
}

// NewFactService creates a new FactService instance
func NewFactService(factRepo repository.FactRepository) *FactService {
	return &FactService{
		FactRepo: factRepo,
	}
}

// CreateFact creates a new Fact entity
func (fs *FactService) CreateFact(newFact *entity.Fact) (*entity.Fact, error) {
	createdFact, err := fs.FactRepo.Create(*newFact)
	if err != nil {
		return nil, err
	}
	return &createdFact, nil
}

// UpdateFactByID updates a Fact entity by its ID
func (fs *FactService) UpdateFactByID(id string, updatedFact entity.Fact) (*entity.Fact, error) {
	// Perform the update operation in the repository
	updated, err := fs.FactRepo.UpdateByID(id, updatedFact)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}
// GetAllFactGroupNames retrieves all unique FactGroupNames
func (fs *FactService) GetAllFactGroupNames() ([]string, error) {
	return fs.FactRepo.GetAllFactGroupNames()
}
// service/fact_service.go

// GetDistinctMappedEntities retrieves all distinct values for MappedEntity
func (fs *FactService) GetDistinctMappedEntities() ([]string, error) {
	return fs.FactRepo.GetAllMappedEntities()
}



// DeleteFactByID deletes a Fact entity by its ID
func (fs *FactService) DeleteFactByID(id string) error {
	return fs.FactRepo.DeleteByID(id)
}

// GetFactByID retrieves a Fact entity by its ID
func (fs *FactService) GetFactByID(id string) (entity.Fact, error) {
	return fs.FactRepo.GetByID(id)
}

// GetAllFacts retrieves all Fact entities
func (fs *FactService) GetAllFacts() ([]entity.Fact, error) {
	return fs.FactRepo.GetAll()
}

