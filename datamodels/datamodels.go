package datamodels

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ReadingDataModel - reading data model.
type ReadingDataModel struct {
	Sensor      string  `bson:"sensor,omitempty"`
	Humidity    float32 `bson:"humidity,omitempty"`
	Temperature float32 `bson:"temperature,omitempty"`
	HeatIndex   float32 `bson:"heatindex,omitempty"`
	DewPoint    float32 `bson:"dewpoint,omitempty"`
	Pressure    float32 `bson:"pressure,omitempty"`
	Altitude    float32 `bson:"altitude,omitempty"`
	DP          float32 `bson:"dp,omitempty"`
	Ps          float32 `bson:"ps,omitempty"`
	Pa          float32 `bson:"pa,omitempty"`
	HI          float32 `bson:"hi,omitempty"`
}

// EventDataModel - event data model.
type EventDataModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Sender   string             `bson:"sender,omitempty"`
	Event    string             `bson:"event,omitempty"`
	Date     time.Time          `bson:"date"`
	Readings []ReadingDataModel `bson:"readings"`
	Version  string             `bson:"version,omitempty"`
}
