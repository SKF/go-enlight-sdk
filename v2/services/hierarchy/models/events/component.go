package events

import (
	"github.com/SKF/go-enlight-sdk/v2/services/hierarchy/models"
	"github.com/SKF/go-eventsource/eventsource"
	"github.com/SKF/go-utility/v2/uuid"
)

type CreateComponentEvent struct {
	*eventsource.BaseEvent
	AssetID uuid.UUID
	models.Component
}

type UpdateComponentEvent struct {
	*eventsource.BaseEvent
	AssetID uuid.UUID
	models.Component
}

type DeleteComponentEvent struct {
	*eventsource.BaseEvent
	AssetID uuid.UUID
}

type CopyComponentEvent struct {
	*eventsource.BaseEvent
	models.Component
	AssetID  uuid.UUID `json:"assetID"`
	SourceID uuid.UUID `json:"sourceID"`
}
