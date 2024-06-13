package events

import (
	"github.com/SKF/go-eventsource/eventsource"
)

type AppendNodeDataEvent struct {
	*eventsource.BaseEvent
	NodeDataCreatedAt   int64 `json:"nodeDataCreatedAt"`
	NodeDataContentType int32 `json:"nodeDataContentType"`
}

func (e AppendNodeDataEvent) GetAggregateID() string {
	return e.BaseEvent.AggregateID
}

func (e AppendNodeDataEvent) GetUserID() (userID string) {
	return e.BaseEvent.UserID
}
