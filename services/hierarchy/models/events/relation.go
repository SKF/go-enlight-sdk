package events

import (
	"github.com/SKF/go-eventsource/eventsource"
	"github.com/SKF/go-utility/v2/uuid"
)

// CreateRelationEvent this struct used to create node relation-event
type CreateRelationEvent struct {
	*eventsource.BaseEvent
	FromID uuid.UUID
	ToID   uuid.UUID
}

// DeleteRelationEvent this struct used to delate node relation-event
type DeleteRelationEvent struct {
	*eventsource.BaseEvent
}
