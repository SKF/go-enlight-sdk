package models

type Node struct {
	BaseNode
	*AssetNode
	*MeasurementPoint
	*InspectionPoint
}
