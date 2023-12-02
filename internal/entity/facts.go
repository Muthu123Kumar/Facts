// entity/fact.go

package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fact represents the Fact entity
type Fact struct {
	FactId            primitive.ObjectID `json:"FactId,omitempty" bson:"_id,omitempty"`
	FactType      string             `json:"factType" bson:"factType"`
	Name          string             `json:"name" bson:"name"`
	KeyMoment     string             `json:"keyMomentEnabled" bson:"keyMomentEnabled"`
	FactGroupId   string             `json:"factGroupId" bson:"factGroupId"`
	FactGroupName string             `json:"factGroupName" bson:"factGroupName"`
	Description   string             `json:"description" bson:"description"`
	MediaType     []string           `json:"mediaType" bson:"mediaType"`
	Query         string             `json:"query" bson:"query"`
	MappedEntity  string             `json:"mappedEntity" bson:"mappedEntity"`
	Status        string             `json:"status" bson:"status"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt" bson:"updatedAt"`
	Version       int                `json:"version" bson:"version"`
}
