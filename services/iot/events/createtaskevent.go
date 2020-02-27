package events

import (
	"github.com/SKF/go-enlight-sdk/services/iot/models"
	"github.com/SKF/go-eventsource/eventsource"
)

type CreateTaskEvent struct {
	*eventsource.BaseEvent
	models.Task
}

func (e CreateTaskEvent) GetAggregateID() string {
	return e.BaseEvent.AggregateID
}

func (e CreateTaskEvent) GetUserID() (userID string) {
	return e.BaseEvent.UserID
}
