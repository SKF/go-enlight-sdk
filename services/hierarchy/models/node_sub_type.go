package models

type NodeSubType string

const (
	NodeSubTypeRoot               NodeSubType = "root"
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

var nodeTypeClasses = map[NodeType][]NodeSubType{ // nolint
	NodeTypeRoot:               {NodeSubTypeRoot},
	NodeTypeCompany:            {NodeSubTypeCompany},
	NodeTypeSite:               {NodeSubTypeSite},
	NodeTypePlant:              {NodeSubTypePlant, NodeSubTypeShip},
	NodeTypeSystem:             {NodeSubTypeSystem},
	NodeTypeFunctionalLocation: {NodeSubTypeFunctionalLocation},
	NodeTypeAsset:              {NodeSubTypeAsset},
	NodeTypeMeasurementPoint:   {NodeSubTypeMeasurementPoint},
	NodeTypeInspectionPoint:    {NodeSubTypeInspectionPoint},
}
