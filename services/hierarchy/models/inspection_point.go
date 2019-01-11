package models

type InspectionPoint struct {
	ValueType   ValueType `json:"valueType"`
	NumericUnit string    `json:"unit"`
	Answers     Answers   `json:"answers"`

	VisualizationType     VisualizationType `json:"visualizationType"`
	VisualizationMinValue string            `json:"visualizationMinValue"`
	VisualizationMaxValue string            `json:"visualizationMaxValue"`
}

type ValueType string

const (
	ValueTypeNumeric      ValueType = "numeric"
	ValueTypeSingleChoice ValueType = "single_choice"
	ValueTypeMultiChoice  ValueType = "multi_choice"
)

type Answers []string

// VisualizationType - defines visualization type when value type is numeric
type VisualizationType string

// Constants for VisualizationType
const (
	VisualizationTypeNone          VisualizationType = "visualization_none"
	VisualizationTypeCircularGauge VisualizationType = "visualization_circular_gauge"
	VisualizationTypeLevelGauge    VisualizationType = "visualization_level_gauge"
)
