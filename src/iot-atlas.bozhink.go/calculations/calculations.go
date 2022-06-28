package calculations

import "math"

// AbsoluteZeroCf is the absolute zero represented in °C
const AbsoluteZeroCf float32 = -273.15

// AbsoluteZeroCd is the absolute zero represented in °C
const AbsoluteZeroCd float64 = -273.15

// See https://en.wikipedia.org/wiki/Dew_point

// Dew point constants
const af float32 = 6.1121 // mbar == hPa
const ad float64 = 6.1121 // mbar == hPa

const bf float32 = 18.678 // NA
const bd float64 = 18.678 // NA

const cf float32 = 257.14 // °C
const cd float64 = 257.14 // °C

const df float32 = 234.5 // °C
const dd float64 = 234.5 // °C

// Heat index constants
const c1f float32 = -8.78469475556
const c1d float64 = -8.78469475556

const c2f float32 = 1.61139411
const c2d float64 = 1.61139411

const c3f float32 = 2.33854883889
const c3d float64 = 2.33854883889

const c4f float32 = -0.14611605
const c4d float64 = -0.14611605

const c5f float32 = -0.012308094
const c5d float64 = -0.012308094

const c6f float32 = -0.016424827778
const c6d float64 = -0.016424827778

const c7f float32 = 0.002211732
const c7d float64 = 0.002211732

const c8f float32 = 0.00072546
const c8d float64 = 0.00072546

const c9f float32 = -0.000003582
const c9d float64 = -0.000003582

func gamma(t float64, h float64) float64 {
	if h <= 0.0 || h > 100.0 {
		return 0
	}

	return math.Log(h/100.0) + (bd*t)/(cd+t)
}

// GetDewPoint returns the temperature of the dew point in °C.
func GetDewPoint(t float64, h float64) float64 {
	g := gamma(t, h)
	return cd * g / (bd - g)
}

// GetSaturatedVaporPressure returns the saturated water vapor pressure in mbar (hPa).
func GetSaturatedVaporPressure(t float64, h float64) float64 {
	return ad * math.Exp((bd*t)/(cd+t))
}

// GetActualVaporPressure returns the actual vapor pressure in mbar (hPa).
func GetActualVaporPressure(t float64, h float64) float64 {
	return (h / 100.0) * GetSaturatedVaporPressure(t, h)
}

// GetHeatIndex returns the heat index.
func GetHeatIndex(t float64, h float64) float64 {
	t2 := t * t
	h2 := h * h
	return c1d + c2d*t + c3d*h + c4d*t*h + c5d*t2 + c6d*h2 + c7d*t2*h + c8d*t*h2 + c9d*t2*h2
}
