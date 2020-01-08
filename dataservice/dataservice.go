package dataservice

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"../datamodels"
)

// RegisterEvent inserts event DTO into the MongoDB.
func RegisterEvent(client *mongo.Client, event *datamodels.EventDataModel) (interface{}, error) {
	if client == nil {
		return nil, errors.New("MongoDB client is nil")
	}

	if event == nil {
		return nil, errors.New("Event DTO is nil")
	}

	eventDTO := *event

	db := client.Database("iot")
	collection := db.Collection("events")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := collection.InsertOne(ctx, eventDTO)

	if cancel != nil {
		cancel()
	}

	return result, err
}
