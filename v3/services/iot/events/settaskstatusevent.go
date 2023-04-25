package events

import (
	"github.com/SKF/go-eventsource/eventsource"
)

type SetTaskStatusEvent struct {
	*eventsource.BaseEvent
	TaskStatus          string
	TaskStatusUpdatedAt int64
}

func (e SetTaskStatusEvent) GetAggregateID() string {
	return e.BaseEvent.AggregateID
}

func (e SetTaskStatusEvent) GetUserID() (userID string) {
	return e.BaseEvent.UserID
}
