package events

import (
	"github.com/SKF/go-enlight-sdk/v2/services/hierarchy/models"

	"github.com/SKF/go-eventsource/eventsource"
	"github.com/SKF/go-utility/v2/uuid"
)

type CreateNodeEvent struct {
	*eventsource.BaseEvent
	models.BaseNode
	*models.AssetNode
	*models.MeasurementPoint
	*models.InspectionPoint
	*models.LubricationPoint
}

type UpdateNodeEvent struct {
	*eventsource.BaseEvent
	models.BaseNode
	*models.AssetNode
	*models.MeasurementPoint
	*models.InspectionPoint
	*models.LubricationPoint
}

type DeleteOriginEvent struct {
	*eventsource.BaseEvent
}

type CopyNodeEvent struct {
	*eventsource.BaseEvent
	models.BaseNode
	*models.AssetNode
	*models.MeasurementPoint
	*models.InspectionPoint
	*models.LubricationPoint
	SrcNodeID uuid.UUID `json:"srcNodeId"`
}

type DeleteNodeEvent struct {
	*eventsource.BaseEvent
}
