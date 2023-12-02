// delivery/http/fact_handler.go

package http

import (
	"github.com/gin-gonic/gin"
	"MainProject/internal/entity"
	"MainProject/internal/service"
	"net/http"
)

// FactHandler represents the HTTP handler for managing Fact entities
type FactHandler struct {
	FactService *service.FactService
}


// NewFactHandler creates a new FactHandler instance
func NewFactHandler(fs *service.FactService) *FactHandler {
	return &FactHandler{
		FactService: fs,
	}
}

// Assuming you have necessary imports and FactService initialized

// CreateFactHandler handles the creation of a new Fact entity
func (fh *FactHandler) CreateFact(c *gin.Context) {
	var newFact entity.Fact
	if err := c.BindJSON(&newFact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdFact, err := fh.FactService.CreateFact(&newFact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdFact)
}

// UpdateFactByIDHandler handles the update of a Fact entity by its ID
func (fh *FactHandler) UpdateFactByID(c *gin.Context) {
	factID := c.Param("id")

	var updatedFact entity.Fact
	if err := c.BindJSON(&updatedFact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := fh.FactService.UpdateFactByID(factID, updatedFact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeleteFactByID handles the deletion of a Fact entity by its ID
func (fh *FactHandler) DeleteFactByID(c *gin.Context) {
	factID := c.Param("id")

	err := fh.FactService.DeleteFactByID(factID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fact deleted successfully"})
}

// GetFactByID retrieves a Fact entity by its ID
func (fh *FactHandler) GetFactByID(c *gin.Context) {
	factID := c.Param("id")

	fact, err := fh.FactService.GetFactByID(factID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fact)
}

// GetAllFacts retrieves all Fact entities
func (fh *FactHandler) GetAllFacts(c *gin.Context) {
	facts, err := fh.FactService.GetAllFacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, facts)
}
func (fh *FactHandler) GetAllFactGroupNames(c *gin.Context) {
	factGroupNames, err := fh.FactService.GetAllFactGroupNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, factGroupNames)
}
// delivery/http/fact_handler.go

// GetAllDistinctMappedEntities retrieves all distinct values for MappedEntity
func (fh *FactHandler) GetAllDistinctMappedEntities(c *gin.Context) {
	mappedEntities, err := fh.FactService.GetDistinctMappedEntities()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mappedEntities)
}
