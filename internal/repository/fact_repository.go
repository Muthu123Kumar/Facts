// repository/fact_repository.go

package repository

import (
	"context"
	"time"
	

	"MainProject/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// FactRepository represents the repository for managing Fact entities
type FactRepository struct {
	Collection *mongo.Collection
	Context    context.Context
}

// NewFactRepository creates a new FactRepository instance
func NewFactRepository(collection *mongo.Collection) *FactRepository {
	return &FactRepository{
		Collection: collection,
		Context:    context.TODO(), // You can use an appropriate context here
	}
}

// Create creates a new Fact entity in the database
func (fr *FactRepository) Create(newFact entity.Fact) (entity.Fact, error) {
	newFact.FactId = primitive.NewObjectID()
	newFact.CreatedAt = time.Now().UTC()
	newFact.Version = 1

	_, err := fr.Collection.InsertOne(fr.Context, newFact)
	if err != nil {
		return entity.Fact{}, err
	}

	return newFact, nil
}


// UpdateByID updates a Fact entity by its ID
// UpdateByID updates a Fact entity by its ID
func (fr *FactRepository) UpdateByID(objID string, updatedFact entity.Fact) (entity.Fact, error) {
	
	// Get existing fact details
	existingFact, err := fr.GetByID(objID)
	if err != nil {
		return entity.Fact{}, err
	}

	// Maintain existing CreatedAt value
	updatedFact.CreatedAt = existingFact.CreatedAt

	// Prepare the update query
	update := bson.M{
		"$set": bson.M{
			"name":             updatedFact.Name,
			"factType":         updatedFact.FactType,
			"keyMomentEnabled": updatedFact.KeyMoment,
			"factGroupId":      updatedFact.FactGroupId,
			"factGroupName":    updatedFact.FactGroupName,
			"description":      updatedFact.Description,
			"mediaType":        updatedFact.MediaType,
			"query":            updatedFact.Query,
			"mappedEntity":     updatedFact.MappedEntity,
			"status":           updatedFact.Status,
			"updatedAt":        time.Now().UTC(), // Update updatedAt to current time
			"version":          existingFact.Version + 1, // Increment version by 1
		},
	}

	// Perform the update operation
	_, err = fr.Collection.UpdateOne(fr.Context, bson.M{"_id": objID}, update)
	if err != nil {
		return entity.Fact{}, err
	}

	// Return the updated Fact entity
	updatedFact.FactId = existingFact.FactId // Set the ID of the updated fact
	updatedFact.Version = existingFact.Version + 1 // Update the version
	updatedFact.UpdatedAt = time.Now().UTC() // Set the updated time

	return updatedFact, nil
}




// DeleteByID deletes a Fact entity by its ID
func (fr *FactRepository) DeleteByID(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = fr.Collection.DeleteOne(fr.Context, bson.M{"_id": objID})
	return err
}

// GetByID retrieves a Fact entity by its ID
func (fr *FactRepository) GetByID(id string) (entity.Fact, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entity.Fact{}, err
	}

	var fact entity.Fact
	err = fr.Collection.FindOne(fr.Context, bson.M{"_id": objID}).Decode(&fact)
	if err != nil {
		return entity.Fact{}, err
	}

	return fact, nil
}

// GetAll retrieves all Fact entities
func (fr *FactRepository) GetAll() ([]entity.Fact, error) {
	cursor, err := fr.Collection.Find(fr.Context, bson.M{})
	if err != nil {
		return nil, err
	}

	var facts []entity.Fact
	if err := cursor.All(fr.Context, &facts); err != nil {
		return nil, err
	}

	return facts, nil
}


// GetAllFactGroupNames retrieves all distinct values for the "factGroupName" field
func (fr *FactRepository) GetAllFactGroupNames() ([]string, error) {
	distinctValues, err := fr.Collection.Distinct(fr.Context, "factGroupName", bson.D{})
	if err != nil {
		return nil, err
	}

	var factGroupNames []string
	for _, value := range distinctValues {
		if groupName, ok := value.(string); ok {
			factGroupNames = append(factGroupNames, groupName)
		}
	}

	return factGroupNames, nil
}

// GetAllMappedEntities retrieves all distinct values for the "MappedEntity" field
func (fr *FactRepository) GetAllMappedEntities() ([]string, error) {
	distinctValues, err := fr.Collection.Distinct(fr.Context, "mappedEntity", bson.M{})
	if err != nil {
		return nil, err
	}

	var mappedEntities []string
	for _, value := range distinctValues {
		if entity, ok := value.(string); ok {
			mappedEntities = append(mappedEntities, entity)
		}
	}

	return mappedEntities, nil
}


