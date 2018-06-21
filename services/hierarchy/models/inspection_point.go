package models

import (
	"fmt"
)

type InspectionPoint struct {
	ValueType   ValueType `json:"valueType"`
	NumericUnit string    `json:"unit"`
}

type ValueType string

const (
	ValueTypeNumeric ValueType = "numeric"
)

var valueTypes = []ValueType{
	ValueTypeNumeric,
}

func (t ValueType) String() string {
	return string(t)
}

func (t ValueType) Validate() error {
	for _, valueType := range valueTypes {
		if t == valueType {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid value type", t)
}

func (p InspectionPoint) Validate() error {
	if err := p.ValueType.Validate(); err != nil {
		return err
	}

	if p.ValueType == ValueTypeNumeric && p.NumericUnit == "" {
		return fmt.Errorf("ValueType is numeric, numeric unit cannot be empty string")
	}

	return nil
}
