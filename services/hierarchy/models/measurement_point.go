package models

type MeasurementPoint struct {
	Bearing         int32           `json:"bearing"`
	Angular         Orientation     `json:"orientation"`
	MeasurementType MeasurementType `json:"measurementType"`
	Shaft           string          `json:"shaft"`
	ShaftSide       ShaftSide       `json:"shaftSide"`
	FixedSpeedRPM   float64         `json:"FixedSpeedRPM"`
}

type Orientation string

const (
	Axial      Orientation = "axial"
	Radial     Orientation = "radial"
	Horizontal Orientation = "horizontal"
	Vertical   Orientation = "vertical"
)

type MeasurementType string

const (
	Displacement MeasurementType = "displacement"
	Acceleration MeasurementType = "acceleration"
	Velocity     MeasurementType = "velocity"
	Temperature  MeasurementType = "temperature"
	DCGAP        MeasurementType = "dc_gap"
	AMPLPHASE    MeasurementType = "ampl_phase"
	BOV          MeasurementType = "bov"
	Speed        MeasurementType = "speed"
	Envelope3    MeasurementType = "envelope_3"
	Envelope2    MeasurementType = "envelope_2"
)

type ShaftSide string

const (
	DE  ShaftSide = "de"
	NDE ShaftSide = "nde"
)
