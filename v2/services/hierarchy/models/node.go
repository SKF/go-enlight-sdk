package models

import (
	"errors"

	"github.com/SKF/go-utility/v2/uuid"
	common_proto "github.com/SKF/proto/v2/common"
	hierarchy_proto "github.com/SKF/proto/v2/hierarchy"
)

type Node struct {
	BaseNode
	*AssetNode
	*MeasurementPoint
	*InspectionPoint
	*LubricationPoint
}

func (node Node) Validate() (err error) {

	if err = node.BaseNode.Validate(); err != nil {
		return err
	}

	switch node.Type {
	case NodeTypeAsset:
		if node.AssetNode == nil {
			return errors.New("Node type is Asset, but data is nil")
		}
		if err = node.AssetNode.Validate(); err != nil {
			return err
		}
	case NodeTypeMeasurementPoint:
		if node.MeasurementPoint == nil {
			return errors.New("Node type is MeasurementPoint, but data is nil")
		}
		if err = node.MeasurementPoint.Validate(); err != nil {
			return err
		}
	case NodeTypeInspectionPoint:
		if node.InspectionPoint == nil {
			return errors.New("Node type is InspectionPoint, but data is nil")
		}
		if err = node.InspectionPoint.Validate(); err != nil {
			return err
		}
	case NodeTypeLubricationPoint:
		if node.LubricationPoint == nil {
			return errors.New("Node type is LubricationPoint, but data is nil")
		}
		if err = node.LubricationPoint.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func ParseNodeType(nodetypeInput string) NodeType {

	switch nodetypeInput {
	case "root":
		return NodeTypeRoot
	case "company":
		return NodeTypeCompany
	case "site":
		return NodeTypeSite
	case "plant":
		return NodeTypePlant
	case "system":
		return NodeTypeSystem
	case "functional_location":
		return NodeTypeFunctionalLocation
	case "asset":
		return NodeTypeAsset
	case "measurement_point":
		return NodeTypeMeasurementPoint
	case "inspection_point":
		return NodeTypeInspectionPoint
	case "lubrication_point":
		return NodeTypeLubricationPoint
	default:
		return NodeTypeUnknown
	}
}

func (n Node) ToGRPC() (node *hierarchy_proto.Node) {

	node = &hierarchy_proto.Node{
		Id:          n.ID.String(),
		Type:        n.Type.String(),
		SubType:     n.SubType.String(),
		Description: n.Description,
		Label:       n.Label,
		ParentId:    n.ParentID.String(),
	}

	if n.Origin != nil {
		node.Origin = n.Origin.ToGRPC()
	}

	if n.Industry != nil {
		node.IndustrySegment = &common_proto.PrimitiveString{Value: n.Industry.String()}
	}

	if n.MeasurementPoint != nil {
		node.MeasurementPoint = n.MeasurementPoint.ToGRPC()
	}

	if n.InspectionPoint != nil {
		node.InspectionPoint = n.InspectionPoint.ToGRPC()
	}

	if n.LubricationPoint != nil {
		node.LubricationPoint = n.LubricationPoint.ToGRPC()
	}

	if n.AssetNode != nil {
		node.AssetNode = n.AssetNode.ToGRPC()
	}

	if n.Position != nil {
		node.Position = &common_proto.PrimitiveInt64{Value: *n.Position}
	}

	if n.Tags != nil {
		node.Tags = &common_proto.PrimitiveString{Value: *n.Tags}
	}

	return node
}

func (n *Node) FromGRPC(node hierarchy_proto.Node) {
	n.ID = uuid.UUID(node.Id)
	n.Type = NodeType(node.Type)
	n.SubType = NodeSubType(node.SubType)
	n.Description = node.Description
	n.Label = node.Label
	n.ParentID = uuid.UUID(node.ParentId)
	if node.IndustrySegment != nil {
		segment := IndustrySegment(node.IndustrySegment.Value)
		n.Industry = &segment
	}
	if node.Position != nil {
		n.Position = &node.Position.Value
	}

	if node.MeasurementPoint != nil {
		n.MeasurementPoint = &MeasurementPoint{}
		n.MeasurementPoint.FromGRPC(*node.MeasurementPoint)
	}

	if node.InspectionPoint != nil {
		n.InspectionPoint = &InspectionPoint{}
		n.InspectionPoint.FromGRPC(*node.InspectionPoint)
	}

	if node.LubricationPoint != nil {
		n.LubricationPoint = &LubricationPoint{}
		n.LubricationPoint.FromGRPC(*node.LubricationPoint)
	}

	if node.AssetNode != nil {
		n.AssetNode = &AssetNode{}
		n.AssetNode.FromGRPC(*node.AssetNode)
	}
	if node.Origin != nil {
		n.Origin = &Origin{}
		n.Origin.FromGRPC(*node.Origin)
	}

	if node.Tags != nil {
		n.Tags = &node.Tags.Value
	}
}

func (n Node) ToGRPCAncestorNode() (node hierarchy_proto.AncestorNode) {
	node.Id = n.ID.String()
	node.Type = n.Type.String()
	node.SubType = n.SubType.String()
	node.Description = n.Description
	node.Label = n.Label
	node.ParentId = n.ParentID.String()

	if n.Origin != nil {
		node.Origin = n.Origin.ToGRPC()
	}

	return node
}
