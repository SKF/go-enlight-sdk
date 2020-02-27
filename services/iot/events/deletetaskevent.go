package events

import (
	"github.com/SKF/go-eventsource/eventsource"
)

type DeleteTaskEvent struct {
	*eventsource.BaseEvent
}

func (e DeleteTaskEvent) GetAggregateID() string {
	return e.BaseEvent.AggregateID
}

func (e DeleteTaskEvent) GetUserID() (userID string) {
	return e.BaseEvent.UserID
}
