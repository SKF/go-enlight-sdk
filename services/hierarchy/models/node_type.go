package models

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
)

var relations = map[NodeType][]NodeType{ // nolint
	NodeTypeCompany:            {NodeTypeRoot},
	NodeTypeSite:               {NodeTypeCompany},
	NodeTypePlant:              {NodeTypeSite},
	NodeTypeSystem:             {NodeTypeSite, NodeTypePlant, NodeTypeSystem},
	NodeTypeFunctionalLocation: {NodeTypeSite, NodeTypePlant, NodeTypeSystem, NodeTypeFunctionalLocation},
	NodeTypeAsset:              {NodeTypeFunctionalLocation},
	NodeTypeMeasurementPoint:   {NodeTypeAsset},
	NodeTypeInspectionPoint:    {NodeTypeAsset},
}
