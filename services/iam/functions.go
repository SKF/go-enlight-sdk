package iam

import (
	"context"
	"encoding/json"
	"time"

	"github.com/SKF/go-eventsource/eventsource"

	"github.com/SKF/go-enlight-sdk/services/iam/grpcapi"
)

func (c *client) CheckAuthentication(token, method string) (user grpcapi.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	input := &grpcapi.CheckAuthenticationInput{Token: token, MethodArn: method}
	output, err := c.api.CheckAuthentication(ctx, input)
	if output != nil {
		user = *output
	}
	return

}

func (c *client) GetNodesByUser(userID string) (nodeIDs []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	input := &grpcapi.GetHierarchyRelationsInput{UserId: userID}
	output, err := c.api.GetHierarchyRelations(ctx, input)
	if err != nil {
		return
	}
	nodeIDs = output.NodeIds
	return
}

func (c *client) GetEventRecords(since int, limit *int32) (records []eventsource.Record, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	input := grpcapi.GetEventRecordsInput{Since: int64(since)}
	if limit != nil {
		input.Limit = &grpcapi.PrimitiveInt32{Value: *limit}
	}

	output, err := c.api.GetEventRecords(ctx, &input)
	if err != nil {
		return
	}

	err = json.Unmarshal(output.Records, &records)
	return
}
