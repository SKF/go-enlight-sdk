package models

type InspectionPoint struct {
	ValueType   ValueType `json:"valueType"`
	NumericUnit string    `json:"unit"`
}

type ValueType string
