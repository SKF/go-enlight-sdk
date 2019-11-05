package events

import (
	"github.com/SKF/go-eventsource/eventsource"
	"github.com/SKF/go-eventsource/eventsource/serializers/json"
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateHierarchyEventSerializer() eventsource.Serializer {
	return json.NewSerializer(
		UpdateNodeEvent{},
		CreateNodeEvent{},
		DeleteNodeEvent{},
		CreateRelationEvent{},
		DeleteRelationEvent{},
		CopyNodeEvent{},
	)
}

func CreateHierarchyEventRepository(sess *session.Session, store eventsource.Store) eventsource.Repository {
	return eventsource.NewRepository(
		store,
		CreateHierarchyEventSerializer(),
	)
}
