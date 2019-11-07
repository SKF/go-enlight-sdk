package models

import (
	"fmt"

	"github.com/SKF/go-utility/uuid"
	grpcapi "github.com/SKF/proto/hierarchy"
)

// MeasurementPoint describes an assets measurement points
type MeasurementPoint struct {
	// Bearing number on this asset
	Bearing int32 `json:"bearing" example:"1"`
	// Orientation of measurement
	Angular Orientation `json:"orientation" example:"vertical" enums:"axial,radial,horizontal,vertical,unknown"`
	// Type of measurement
	MeasurementType MeasurementType `json:"measurementType" example:"acceleration" enums:"displacement,acceleration,velocity,temperature,dc_gap,ampl_phase,box,speed,envelope_2,envelope_3,unknown"`
	// Identifier of shaft that this measurement point belongs to
	Shaft string `json:"shaft" example:"C"`
	// Which side of the given shaft this measurement point belongs to
	ShaftSide ShaftSide `json:"shaftSide" example:"nde" enums:"de,nde"`
	// Speed in RPM if this shaft has a fixed speed
	FixedSpeedRPM float64 `json:"fixedSpeedRpm,omitempty" example:"150"`
	// ID of measurement point location
	LocationId *uuid.UUID `json:"locationId,omitempty"`
	// Type of device used to take measurements on this point
	DADType string `json:"dadType,omitempty"`
}

// Orientation describes a measurement points orientation
type Orientation string

// Valid values of measurement points orientations
const (
	Axial              Orientation = "axial"
	Radial             Orientation = "radial"
	Horizontal         Orientation = "horizontal"
	Vertical           Orientation = "vertical"
	UnknownOrientation Orientation = "unknown"
)

var orientations = []Orientation{
	Axial, Radial, Horizontal, Vertical, UnknownOrientation,
}

// MeasurementType is measurement type unit
type MeasurementType string

// Valid measurement type values
const (
	Displacement           MeasurementType = "displacement"
	Acceleration           MeasurementType = "acceleration"
	Velocity               MeasurementType = "velocity"
	Temperature            MeasurementType = "temperature"
	DCGAP                  MeasurementType = "dc_gap"
	AMPLPHASE              MeasurementType = "ampl_phase"
	BOV                    MeasurementType = "bov"
	Speed                  MeasurementType = "speed"
	Envelope3              MeasurementType = "envelope_3"
	Envelope2              MeasurementType = "envelope_2"
	UnknownMeasurementType MeasurementType = "unknown"
)

var measurementTypes = []MeasurementType{
	Displacement, Acceleration, Velocity, Temperature, DCGAP, AMPLPHASE, BOV, Speed, Envelope3, Envelope2, UnknownMeasurementType,
}

// ShaftSide describes on what side of a shaft the measurement point is located
type ShaftSide string

// Valid shaft side values
const (
	DE               ShaftSide = "de"
	NDE              ShaftSide = "nde"
	UnknownShaftSide ShaftSide = "unknown"
)

var shaftSides = []ShaftSide{
	DE, NDE, UnknownShaftSide,
}

// ParseShaftSide takes a string and makes it a valid shaft side value
func ParseShaftSide(shaftSide string) ShaftSide {
	switch shaftSide {
	case "de":
		return DE
	case "nde":
		return NDE
	default:
		return UnknownShaftSide
	}
}

// ParseOrientation takes a string and makes it a valid orientation value
func ParseOrientation(orientation string) Orientation {
	switch orientation {
	case "axial":
		return Axial
	case "radial":
		return Radial
	case "horizontal":
		return Horizontal
	case "vertical":
		return Vertical
	default:
		return UnknownOrientation
	}
}

// ParseMeasurementType takes a string and makes it a valid measurement type value
func ParseMeasurementType(measurementType string) MeasurementType {
	switch measurementType {
	case "displacement":
		return Displacement
	case "acceleration":
		return Acceleration
	case "velocity":
		return Velocity
	case "temperature":
		return Temperature
	case "dc_gap":
		return DCGAP
	case "ampl_phase":
		return AMPLPHASE
	case "bov":
		return BOV
	case "speed":
		return Speed
	case "envelope_2":
		return Envelope2
	case "envelope_3":
		return Envelope3
	default:
		return UnknownMeasurementType
	}
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

// Validate - validates a MeasurementType
func (mt MeasurementType) Validate() error {
	for _, measurementType := range measurementTypes {
		if mt == measurementType {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid measurement type", mt)
}

// Validate - validates an Orientation
func (o Orientation) Validate() error {
	for _, orentation := range orientations {
		if o == orentation {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid orientation type", o)
}

// Validate - validates a ShaftSide
func (ss ShaftSide) Validate() error {
	for _, shiftside := range shaftSides {
		if ss == shiftside {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid ShaftSide", ss)
}

// Validate - validates a MeasurementPoint
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

	if err := meas.ShaftSide.Validate(); err != nil {
		return err
	}

	if meas.FixedSpeedRPM < 0 {
		return fmt.Errorf("FixedSpeedRPM cannot be negative")
	}

	if meas.LocationId != nil {
		if err := meas.LocationId.Validate(); err != nil {
			return fmt.Errorf("locationId is invalid: %s", meas.LocationId)
		}
	}

	return nil
}

// ToGRPC - converts a MeasurementPoint struct to grpcapi.MeasurementPoint
func (meas MeasurementPoint) ToGRPC() *grpcapi.MeasurementPoint {
	ret := grpcapi.MeasurementPoint{
		Bearing:         meas.Bearing,
		Angular:         meas.Angular.String(),
		MeasurementType: meas.MeasurementType.String(),
		Shaft:           meas.Shaft,
		ShaftSide:       meas.ShaftSide.String(),
		FixedSpeedRPM:   meas.FixedSpeedRPM,
		DadType:         meas.DADType,
	}
	if meas.LocationId != nil {
		ret.LocationId = meas.LocationId.String()
	}
	return &ret
}

// FromGRPC - converts to a MeasurementPoint from the gRPC MeasurementPoint struct
func (meas *MeasurementPoint) FromGRPC(measPoint grpcapi.MeasurementPoint) {
	meas.Bearing = measPoint.Bearing
	meas.Angular = ParseOrientation(measPoint.Angular)
	meas.MeasurementType = MeasurementType(measPoint.MeasurementType)
	meas.Shaft = measPoint.Shaft
	meas.ShaftSide = ParseShaftSide(measPoint.ShaftSide)
	meas.FixedSpeedRPM = measPoint.FixedSpeedRPM
	if measPoint.LocationId != "" {
		meas.LocationId = (*uuid.UUID)(&measPoint.LocationId)
	}
	meas.DADType = measPoint.DadType
}
