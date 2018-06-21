package models

import (
	"fmt"
)

type NodeType string

const (
	NodeTypeCompany            NodeType = "company"
	NodeTypeSite               NodeType = "site"
	NodeTypePlant              NodeType = "plant"
	NodeTypeSystem             NodeType = "system"
	NodeTypeFunctionalLocation NodeType = "functional_location"
	NodeTypeAsset              NodeType = "asset"
	NodeTypeMeasurementPoint   NodeType = "measurement_point"
	NodeTypeInspectionPoint    NodeType = "inspection_point"
)

var allNodeTypes = []NodeType{
	NodeTypeCompany, NodeTypeSite, NodeTypePlant, NodeTypeSystem,
	NodeTypeFunctionalLocation, NodeTypeAsset, NodeTypeMeasurementPoint,
	NodeTypeInspectionPoint,
}

var relations = map[NodeType][]NodeType{
	NodeTypeCompany:            {NodeTypeCompany},
	NodeTypeSite:               {NodeTypeCompany},
	NodeTypePlant:              {NodeTypeSite},
	NodeTypeSystem:             {NodeTypeSite, NodeTypePlant, NodeTypeSystem},
	NodeTypeFunctionalLocation: {NodeTypeSite, NodeTypePlant, NodeTypeSystem, NodeTypeFunctionalLocation},
	NodeTypeAsset:              {NodeTypeFunctionalLocation},
	NodeTypeMeasurementPoint:   {NodeTypeAsset},
	NodeTypeInspectionPoint:    {NodeTypeAsset},
}

func (nt NodeType) String() string {
	return string(nt)
}

func (nt NodeType) IsChildOf(parentType NodeType) bool {
	for key, relationTypes := range relations {
		if nt == key {
			for _, relationType := range relationTypes {
				if relationType == parentType {
					return true
				}
			}
		}
	}

	return false
}

func (nt NodeType) HasSubType(cls NodeSubType) bool {
	return cls.IsTypeOf(nt)
}

func (nt NodeType) Validate() error {
	for _, nodeType := range allNodeTypes {
		if nt == nodeType {
			return nil
		}
	}
	return fmt.Errorf("'%s' is not a valid node type", nt)
}
