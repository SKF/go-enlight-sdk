package models

import (
	"errors"
	"fmt"
	"strings"

	hierarchy_proto "github.com/SKF/proto/v2/hierarchy"
)

type LubricationPoint struct {
	// Type of lubricant used
	Lubricant string `json:"lubricant" example:"grease"`
	// Volume of lubricant, in the given unit
	Volume int32 `json:"lubricantVolume" example:"10"`
	// Unit that the volume is specified in
	Unit LubricantUnit `json:"lubricantUnit" swaggertype:"string" example:"cm3" enums:"gram,ounce,cm3,unknown"`
	// Interval between lubrication in days
	Interval int32 `json:"lubricateInterval" example:"5"`
	// ActivityAssetState the asset should be in during the lubrication activity
	ActivityAssetState LubricationActivityAssetState `json:"lubricationActivityAssetState" swaggertype:"string" example:"must_be_on" enums:"must_be_on,must_be_off"`
	// Instruction for lubrication activity
	Instructions string `json:"lubricateInstructions"`
}

type LubricantUnit string

const (
	Gram                  LubricantUnit = "gram"
	Ounce                 LubricantUnit = "ounce"
	CM3                   LubricantUnit = "cm3"
	Unknown_LubricantUnit LubricantUnit = "unknown"
)

var units = []LubricantUnit{
	Gram, Ounce, CM3,
}

func ParseUnit(unit string) (returnUnit LubricantUnit) {
	switch unit {
	case "gram":
		return Gram
	case "ounce":
		return Ounce
	case "cm3":
		return CM3
	default:
		return Unknown_LubricantUnit
	}
}

func (lu LubricantUnit) String() string {
	return string(lu)
}

func (lu LubricantUnit) Validate() error {
	for _, unit := range units {
		if lu == unit {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid Lubricant Unit", lu)
}

type LubricationActivityAssetState string

const (
	AssetMustBeOn     LubricationActivityAssetState = "must_be_on"
	AssetMustBeOff    LubricationActivityAssetState = "must_be_off"
	AssetStateUnknown LubricationActivityAssetState = ""
)

var activityAssetStates = []LubricationActivityAssetState{
	AssetMustBeOn, AssetMustBeOff, AssetStateUnknown,
}

func (laas LubricationActivityAssetState) Validate() error {
	for _, state := range activityAssetStates {
		if laas == state {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid lubrication activity asset state", laas)
}

func (laas LubricationActivityAssetState) String() string {
	return string(laas)
}

func (lp LubricationPoint) Validate() error {
	if lp.Lubricant == "" {
		return errors.New("Lubricant cannot be empty string")
	}

	if lp.Volume > 0 {
		if err := lp.Unit.Validate(); err != nil {
			return err
		}
	}

	if lp.Volume < 0 {
		return errors.New("Lubricant volume cannot be negative")
	}

	if lp.Interval < 0 {
		return errors.New("Lubricate interval cannot be negative")
	}

	if err := lp.ActivityAssetState.Validate(); err != nil {
		return err
	}

	return nil
}

func (p LubricationPoint) ToGRPC() *hierarchy_proto.LubricationPoint {
	return &hierarchy_proto.LubricationPoint{
		Lubricant: p.Lubricant,
		Volume:    p.Volume,
		Unit:      hierarchy_proto.LubricantUnit(hierarchy_proto.LubricantUnit_value[strings.ToUpper(p.Unit.String())]),
		Interval:  p.Interval,
	}
}

func (p *LubricationPoint) FromGRPC(lubePoint hierarchy_proto.LubricationPoint) {
	p.Lubricant = lubePoint.Lubricant
	p.Volume = lubePoint.Volume
	p.Unit = LubricantUnit(strings.ToLower(lubePoint.Unit.String()))
	p.Interval = lubePoint.Interval
}
