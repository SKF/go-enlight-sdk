package models

import (
	"fmt"
)

type Node struct {
	BaseNode
	*AssetNode
	*MeasurementPoint
	*InspectionPoint
}

func (node Node) Validate() (err error) {

	if err := node.BaseNode.Validate(); err != nil {
		return err
	}

	switch node.Type {
	case NodeTypeAsset:
		if node.AssetNode == nil {
			return fmt.Errorf("Node type is Asset, but data is nil")
		}
		if err = node.AssetNode.Validate(); err != nil {
			return err
		}
	case NodeTypeMeasurementPoint:
		if node.MeasurementPoint == nil {
			return fmt.Errorf("Node type is MeasurementPoint, but data is nil")
		}
		if err = node.MeasurementPoint.Validate(); err != nil {
			return err
		}
	case NodeTypeInspectionPoint:
		if node.InspectionPoint == nil {
			return fmt.Errorf("Node type is InspectionPoint, but data is nil")
		}
		if err = node.InspectionPoint.Validate(); err != nil {
			return err
		}
	}

	return nil
}
