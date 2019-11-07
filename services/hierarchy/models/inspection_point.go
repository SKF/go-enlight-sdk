package models

import (
	"fmt"
	"strings"

	grpcapi "github.com/SKF/proto/hierarchy"
)

// InspectionPoint - holds parameters for inspection point
type InspectionPoint struct {
	// Type of value to record
	ValueType ValueType `json:"valueType" example:"numeric" enums:"numeric,single_choice,multi_choice,unknown"`
	// Unit of the value recorded, in case of numeric inspection
	NumericUnit string `json:"unit" example:"bar"`
	// Possible answers for single_choice and multi_choice inspections
	Answers Answers `json:"answers" example:"[\"first\", \"second\"]"`

	// Type of visualization in Enlight Centre
	VisualizationType     VisualizationType `json:"visualizationType" example:"visualization_circular_gauge" enums:"visualization_none,visualization_circular_gauge,visualization_level_gauge"`
	VisualizationMinValue string            `json:"visualizationMinValue" example:"3"`
	VisualizationMaxValue string            `json:"visualizationMaxValue" example:"13"`
}

type ValueType string

const (
	ValueTypeNumeric      ValueType = "numeric"
	ValueTypeSingleChoice ValueType = "single_choice"
	ValueTypeMultiChoice  ValueType = "multi_choice"
	ValueTypeUnknown      ValueType = "unknown"
)

var valueTypes = []ValueType{
	ValueTypeNumeric,
	ValueTypeSingleChoice,
	ValueTypeMultiChoice,
}

func ParseInspectionType(inspectionType string) ValueType {
	switch inspectionType {
	case "numeric":
		return ValueTypeNumeric
	case "single_choice":
		return ValueTypeSingleChoice
	case "multi_choice":
		return ValueTypeMultiChoice
	default:
		return ValueTypeUnknown
	}
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

type Answers []string

func (as Answers) Array() []string {
	return []string(as)
}

func (as Answers) Validate(valueType ValueType) error {
	if len(as) < 1 {
		return fmt.Errorf("ValueType is %s, there need to be at least 1 answers", valueType)
	}
	for _, answer := range as {
		if answer == "" {
			return fmt.Errorf("'%s' is not a valid answer", answer)
		}
	}
	return nil
}

// VisualizationType - defines visualization type when value type is numeric
type VisualizationType string

// Constants for VisualizationType
const (
	VisualizationTypeNone          VisualizationType = "visualization_none"
	VisualizationTypeCircularGauge VisualizationType = "visualization_circular_gauge"
	VisualizationTypeLevelGauge    VisualizationType = "visualization_level_gauge"
)

// Array fpr VisualizationType constants
var visualizationTypes = []VisualizationType{
	VisualizationTypeNone,
	VisualizationTypeCircularGauge,
	VisualizationTypeLevelGauge,
}

// String - stringifies VisualizationType
func (t VisualizationType) String() string {
	return string(t)
}

// Validate - validates VisualizationType
func (t VisualizationType) Validate() error {
	for _, visualizationType := range visualizationTypes {
		if t == visualizationType {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid visualization type", t)
}

func (p InspectionPoint) Validate() error {
	if err := p.ValueType.Validate(); err != nil {
		return err
	}

	if p.ValueType == ValueTypeNumeric {
		if p.NumericUnit == "" {
			return fmt.Errorf("ValueType is numeric, numeric unit cannot be empty string")
		}

		// Default to none
		if p.VisualizationType == "" {
			p.VisualizationType = VisualizationTypeNone
		}

		if err := p.VisualizationType.Validate(); err != nil {
			return err
		}

		if p.VisualizationType != VisualizationTypeNone && (p.VisualizationMinValue == "" || p.VisualizationMaxValue == "") {
			return fmt.Errorf("ValueType is numeric and VisualizationType: \"%s\", minValue: \"%s\", maxValue: \"%s\" - min/maxValue cannot be empty",
				p.VisualizationType.String(),
				p.VisualizationMinValue,
				p.VisualizationMaxValue)
		}
	}

	if p.ValueType == ValueTypeSingleChoice || p.ValueType == ValueTypeMultiChoice {
		if err := p.Answers.Validate(p.ValueType); err != nil {
			return err
		}
	}

	return nil
}

func (p InspectionPoint) ToGRPC() *grpcapi.InspectionPoint {
	return &grpcapi.InspectionPoint{
		ValueType:             grpcapi.ValueType(grpcapi.ValueType_value[strings.ToUpper(p.ValueType.String())]),
		NumericUnit:           p.NumericUnit,
		Answers:               p.Answers.Array(),
		VisualizationType:     grpcapi.VisualizationType(grpcapi.VisualizationType_value[strings.ToUpper(p.VisualizationType.String())]),
		VisualizationMinValue: p.VisualizationMinValue,
		VisualizationMaxValue: p.VisualizationMaxValue,
	}
}

func (p *InspectionPoint) FromGRPC(inspectPoint grpcapi.InspectionPoint) {
	p.ValueType = ValueType(strings.ToLower(inspectPoint.ValueType.String()))
	p.NumericUnit = inspectPoint.NumericUnit
	p.Answers = Answers(inspectPoint.Answers)
	p.VisualizationType = VisualizationType(strings.ToLower(inspectPoint.VisualizationType.String()))
	p.VisualizationMinValue = inspectPoint.VisualizationMinValue
	p.VisualizationMaxValue = inspectPoint.VisualizationMaxValue
}
