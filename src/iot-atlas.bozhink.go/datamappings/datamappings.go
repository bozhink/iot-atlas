package datamappings

import (
	"time"

	"iot-atlas.bozhink.go/apimodels"
	"iot-atlas.bozhink.go/calculations"
	"iot-atlas.bozhink.go/datamodels"
)

// GetDTO returns event data model.
func GetDTO(e *apimodels.EventRequestModel) *datamodels.EventDataModel {
	if e == nil {
		return nil
	}

	event := *e
	if event.Readings == nil {
		return nil
	}

	n := len(event.Readings)
	if n < 1 {
		return nil
	}

	k := 0
	numberOfValidReadings := 0
	validReadings := make([]int, n)

	for i, reading := range event.Readings {
		if len(reading.Sensor) > 0 && reading.Temperature > calculations.AbsoluteZeroCf && reading.Humidity >= 0 && reading.Humidity <= 100 {
			numberOfValidReadings++
			validReadings[k] = i
			k++
		}
	}

	if numberOfValidReadings < 1 || numberOfValidReadings > n {
		return nil
	}

	eventDTO := datamodels.EventDataModel{
		Sender:   event.Sender,
		Event:    event.Event,
		Version:  event.Version,
		Date:     time.Now(),
		Readings: make([]datamodels.ReadingDataModel, numberOfValidReadings),
	}

	for i := 0; i < numberOfValidReadings; i++ {
		reading := event.Readings[validReadings[i]]
		t := (float64)(reading.Temperature)
		h := (float64)(reading.Humidity)

		readingDTO := datamodels.ReadingDataModel{
			Sensor:      reading.Sensor,
			Humidity:    reading.Humidity,
			Temperature: reading.Temperature,
			HeatIndex:   reading.HeatIndex,
			DewPoint:    reading.DewPoint,
			Pressure:    reading.Pressure,
			Altitude:    reading.Altitude,
			Illuminance: reading.Illuminance,
			UVA:         reading.UVA,
			UVB:         reading.UVB,
			UVIndex:     reading.UVIndex,
			DP:          (float32)((int)(calculations.GetDewPoint(t, h)*100)) / 100.0,
			Ps:          (float32)((int)(calculations.GetSaturatedVaporPressure(t, h)*100)) / 100.0,
			Pa:          (float32)((int)(calculations.GetActualVaporPressure(t, h)*100)) / 100.0,
			HI:          (float32)((int)(calculations.GetHeatIndex(t, h)*100)) / 100.0,
		}

		eventDTO.Readings[i] = readingDTO
	}

	return &eventDTO
}
