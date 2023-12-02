package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"MainProject/internal/handler"
	"MainProject/internal/repository"
	"MainProject/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").
		SetMaxPoolSize(10).
		SetMinPoolSize(5)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	factCollection := client.Database("your_database_name").Collection("your_collection_name")

	// Initialize repository, service, and HTTP handler
	factRepo := repository.NewFactRepository(factCollection)
factService := service.NewFactService(*factRepo)
factHandler := http.NewFactHandler(factService)

	router := gin.Default()

	// Routes
	router.POST("/facts", factHandler.CreateFact)
	router.PUT("/facts/:id", factHandler.UpdateFactByID)
	router.DELETE("/facts/:id", factHandler.DeleteFactByID)
	router.GET("/facts/:id", factHandler.GetFactByID)
	router.GET("/facts", factHandler.GetAllFacts)
	router.GET("/factgroupnames", factHandler.GetAllFactGroupNames)
	router.GET("/factmappedentity", factHandler.GetAllDistinctMappedEntities)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
