package models

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

var relations = map[NodeType][]NodeType{ // nolint
	NodeTypeCompany:            {NodeTypeCompany},
	NodeTypeSite:               {NodeTypeCompany},
	NodeTypePlant:              {NodeTypeSite},
	NodeTypeSystem:             {NodeTypeSite, NodeTypePlant, NodeTypeSystem},
	NodeTypeFunctionalLocation: {NodeTypeSite, NodeTypePlant, NodeTypeSystem, NodeTypeFunctionalLocation},
	NodeTypeAsset:              {NodeTypeFunctionalLocation},
	NodeTypeMeasurementPoint:   {NodeTypeAsset},
	NodeTypeInspectionPoint:    {NodeTypeAsset},
}
