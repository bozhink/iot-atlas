package datamappings

import (
	"time"

	"../apimodels"
	"../calculations"
	"../datamodels"
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

	eventDTO := datamodels.EventDataModel{
		Sender:   event.Sender,
		Event:    event.Event,
		Version:  event.Version,
		Date:     time.Now(),
		Readings: make([]datamodels.ReadingDataModel, n),
	}

	for i := 0; i < n; i++ {
		reading := event.Readings[i]
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
			DP:          (float32)((int)(calculations.GetDewPoint(t, h)*100)) / 100.0,
			Ps:          (float32)((int)(calculations.GetSaturatedVaporPressure(t, h)*100)) / 100.0,
			Pa:          (float32)((int)(calculations.GetActualVaporPressure(t, h)*100)) / 100.0,
			HI:          (float32)((int)(calculations.GetHeatIndex(t, h)*100)) / 100.0,
		}

		if reading.Temperature < calculations.AbsoluteZeroCf {
			readingDTO.Temperature = calculations.AbsoluteZeroCf
			readingDTO.DewPoint = calculations.AbsoluteZeroCf
			readingDTO.HeatIndex = calculations.AbsoluteZeroCf
		}

		if reading.Humidity < 0 || reading.Humidity > 100 {
			readingDTO.Humidity = 0
			readingDTO.DewPoint = calculations.AbsoluteZeroCf
			readingDTO.HeatIndex = calculations.AbsoluteZeroCf
		}

		eventDTO.Readings[i] = readingDTO
	}

	return &eventDTO
}
