package models

import (
	"fmt"
)

type MeasurementPoint struct {
	Bearing         int32           `json:"bearing"`
	Angular         Orientation     `json:"orientation"`
	MeasurementType MeasurementType `json:"measurementType"`
	Shaft           string          `json:"shaft"`
	ShaftSide       ShaftSide       `json:"shaftSide"`
}

type Orientation string

const (
	Axial      Orientation = "axial"
	Radial     Orientation = "radial"
	Horizontal Orientation = "horizontal"
	Vertical   Orientation = "vertical"
)

var orientations = []Orientation{
	Axial, Radial, Horizontal, Vertical,
}

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

var measurementTypes = []MeasurementType{
	Displacement, Acceleration, Velocity, Temperature, DCGAP, AMPLPHASE, BOV, Speed, Envelope3, Envelope2,
}

type ShaftSide string

const (
	DE  ShaftSide = "de"
	NDE ShaftSide = "nde"
)

var shaftSides = []ShaftSide{
	DE, NDE,
}

func (mt MeasurementType) String() string {
	return string(mt)
}
func (o Orientation) String() string {
	return string(o)
}
func (ss ShaftSide) String() string {
	return string(ss)
}

func (mt MeasurementType) Validate() error {
	for _, measurementType := range measurementTypes {
		if mt == measurementType {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valide measurement type", mt)
}

func (o Orientation) Validate() error {
	for _, orentation := range orientations {
		if o == orentation {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valide orientation type", o)
}

func (ss ShaftSide) Validate() error {
	for _, shiftside := range shaftSides {
		if ss == shiftside {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid ShaftSide", ss)
}

func (meas MeasurementPoint) Validate() error {

	if meas.Bearing == 0 {
		return fmt.Errorf("Bearing cannot be zero")
	}

	if err := meas.Angular.Validate(); err != nil {
		return err
	}

	if err := meas.MeasurementType.Validate(); err != nil {
		return err
	}

	if meas.Shaft == "" {
		return fmt.Errorf("Shaft cannot be empty string")
	}

	if err := meas.ShaftSide.Validate(); err != nil {
		return err
	}

	return nil
}
