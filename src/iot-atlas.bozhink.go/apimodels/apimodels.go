package apimodels

// ReadingRequestModel - reading request model.
type ReadingRequestModel struct {
	Sensor      string  `json:"sensor,omitempty" bson:"sensor,omitempty"`
	Humidity    float32 `json:"humidity,omitempty" bson:"humidity,omitempty"`
	Temperature float32 `json:"temperature,omitempty" bson:"temperature,omitempty"`
	HeatIndex   float32 `json:"heatindex,omitempty" bson:"heatindex,omitempty"`
	DewPoint    float32 `json:"dewpoint,omitempty" bson:"dewpoint,omitempty"`
	Pressure    float32 `json:"pressure,omitempty" bson:"pressure,omitempty"`
	Altitude    float32 `json:"altitude,omitempty" bson:"altitude,omitempty"`
	Illuminance float32 `json:"illuminance,omitempty" bson:"illuminance,omitempty"`
	UVA         float32 `json:"uva,omitempty" bson:"uva,omitempty"`
	UVB         float32 `json:"uvb,omitempty" bson:"uvb,omitempty"`
	UVIndex     float32 `json:"uvIndex,omitempty" bson:"uvIndex,omitempty"`
}

// EventRequestModel - event request model.
type EventRequestModel struct {
	Sender   string                `json:"sender,omitempty" bson:"sender,omitempty"`
	Event    string                `json:"event,omitempty" bson:"event,omitempty"`
	Readings []ReadingRequestModel `json:"readings,omitempty" bson:"readings,omitempty"`
	Version  string                `json:"version,omitempty" bson:"version,omitempty"`
}
