package models

import (
	"fmt"
)

type NodeSubType string

const (
	NodeSubTypeCompany            NodeSubType = "company"
	NodeSubTypeSite               NodeSubType = "site"
	NodeSubTypePlant              NodeSubType = "plant"
	NodeSubTypeShip               NodeSubType = "ship"
	NodeSubTypeSystem             NodeSubType = "system"
	NodeSubTypeFunctionalLocation NodeSubType = "functional_location"
	NodeSubTypeAsset              NodeSubType = "asset"
	NodeSubTypeMeasurementPoint   NodeSubType = "measurement_point"
	NodeSubTypeInspectionPoint    NodeSubType = "inspection_point"
)

var allNodeSubTypees = []NodeSubType{
	NodeSubTypeCompany, NodeSubTypeSite, NodeSubTypePlant,
	NodeSubTypeShip, NodeSubTypeSystem, NodeSubTypeFunctionalLocation,
	NodeSubTypeAsset, NodeSubTypeMeasurementPoint, NodeSubTypeInspectionPoint,
}

var nodeTypeClasses = map[NodeType][]NodeSubType{
	NodeTypeCompany:            {NodeSubTypeCompany},
	NodeTypeSite:               {NodeSubTypeSite},
	NodeTypePlant:              {NodeSubTypePlant, NodeSubTypeShip},
	NodeTypeSystem:             {NodeSubTypeSystem},
	NodeTypeFunctionalLocation: {NodeSubTypeFunctionalLocation},
	NodeTypeAsset:              {NodeSubTypeAsset},
	NodeTypeMeasurementPoint:   {NodeSubTypeMeasurementPoint},
	NodeTypeInspectionPoint:    {NodeSubTypeInspectionPoint},
}

func (nc NodeSubType) String() string {
	return string(nc)
}

func (nc NodeSubType) IsTypeOf(nt NodeType) bool {
	for nodeType, nodeClasses := range nodeTypeClasses {
		if nodeType == nt {
			for _, nodeClass := range nodeClasses {
				if nodeClass == nc {
					return true
				}
			}
		}
	}

	return false
}

func (nc NodeSubType) Validate() error {
	for _, nodeClass := range allNodeSubTypees {
		if nc == nodeClass {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid node class", nc)
}
