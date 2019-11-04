package events

import (
	"github.com/SKF/go-enlight-sdk/services/hierarchy/models"
	"github.com/SKF/go-eventsource/eventsource"
	"github.com/SKF/go-utility/uuid"
)

// CreateNodeEvent ...
type CreateNodeEvent struct {
	*eventsource.BaseEvent
	models.BaseNode
	*models.AssetNode
	*models.MeasurementPoint
	*models.InspectionPoint
	*models.LubricationPoint
}

// UpdateNodeEvent ...
type UpdateNodeEvent struct {
	*eventsource.BaseEvent
	models.BaseNode
	*models.AssetNode
	*models.MeasurementPoint
	*models.InspectionPoint
	*models.LubricationPoint
}

// CopyNodeEvent ...
type CopyNodeEvent struct {
	*eventsource.BaseEvent
	models.BaseNode
	*models.AssetNode
	*models.MeasurementPoint
	*models.InspectionPoint
	*models.LubricationPoint
	SrcNodeID uuid.UUID `json:"srcNodeId"`
}

// DeleteNodeEvent ...
type DeleteNodeEvent struct {
	*eventsource.BaseEvent
}
