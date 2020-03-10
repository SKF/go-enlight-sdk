package models

import (
	"fmt"
)

type NodeType string

const (
	NodeTypeRoot               NodeType = "root"
	NodeTypeCompany            NodeType = "company"
	NodeTypeSite               NodeType = "site"
	NodeTypePlant              NodeType = "plant"
	NodeTypeSystem             NodeType = "system"
	NodeTypeFunctionalLocation NodeType = "functional_location"
	NodeTypeAsset              NodeType = "asset"
	NodeTypeMeasurementPoint   NodeType = "measurement_point"
	NodeTypeInspectionPoint    NodeType = "inspection_point"
	NodeTypeLubricationPoint   NodeType = "lubrication_point"
	NodeTypeUnknown            NodeType = "unknown"
)

var allNodeTypes = []NodeType{
	NodeTypeRoot, NodeTypeCompany, NodeTypeSite, NodeTypePlant, NodeTypeSystem,
	NodeTypeFunctionalLocation, NodeTypeAsset, NodeTypeMeasurementPoint,
	NodeTypeInspectionPoint, NodeTypeLubricationPoint,
}

var relations = map[NodeType][]NodeType{
	NodeTypeCompany:            {NodeTypeRoot},
	NodeTypeSite:               {NodeTypeCompany},
	NodeTypePlant:              {NodeTypeSite},
	NodeTypeSystem:             {NodeTypeSite, NodeTypePlant, NodeTypeSystem},
	NodeTypeFunctionalLocation: {NodeTypeSite, NodeTypePlant, NodeTypeSystem, NodeTypeFunctionalLocation},
	NodeTypeAsset:              {NodeTypeFunctionalLocation},
	NodeTypeMeasurementPoint:   {NodeTypeAsset},
	NodeTypeInspectionPoint:    {NodeTypeAsset},
	NodeTypeLubricationPoint:   {NodeTypeAsset},
}

func (nt NodeType) String() string {
	return string(nt)
}

func (nt NodeType) IsChildOf(parentType NodeType) bool {
	relationTypes := relations[nt]
	for _, relationType := range relationTypes {
		if relationType == parentType {
			return true
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
