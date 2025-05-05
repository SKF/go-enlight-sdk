package iam

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/SKF/proto/v2/common"

	"github.com/SKF/go-eventsource/eventsource"
	iam_grpcapi "github.com/SKF/proto/v2/iam"
)

func (c *client) CheckAuthentication(token, arn string) (claims *iam_grpcapi.UserClaims, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.CheckAuthenticationWithContext(ctx, token, arn)
}
func (c *client) CheckAuthenticationWithContext(ctx context.Context, token, arn string) (claims *iam_grpcapi.UserClaims, err error) {
	input := &iam_grpcapi.CheckAuthenticationInput{Token: token, MethodArn: arn}
	output, err := c.api.CheckAuthentication(ctx, input)
	if output != nil {
		claims = output
	} else if err == nil {
		err = errors.New("no output")
	}
	return
}

func (c *client) CheckAuthenticationByEndpoint(token, api, method, endpoint string) (claims *iam_grpcapi.UserClaims, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.CheckAuthenticationByEndpointWithContext(ctx, token, api, method, endpoint)
}
func (c *client) CheckAuthenticationByEndpointWithContext(ctx context.Context, token, api, method, endpoint string) (claims *iam_grpcapi.UserClaims, err error) {
	input := &iam_grpcapi.CheckAuthenticationByEndpointInput{
		Api:      api,
		Token:    token,
		Method:   method,
		Endpoint: endpoint,
	}
	output, err := c.api.CheckAuthenticationByEndpoint(ctx, input)
	if output != nil {
		claims = output
	} else if err == nil {
		err = errors.New("no output")
	}
	return
}

func (c *client) GetNodesByUser(userID string) (nodeIDs []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetNodesByUserWithContext(ctx, userID)
}
func (c *client) GetNodesByUserWithContext(ctx context.Context, userID string) (nodeIDs []string, err error) {
	input := &iam_grpcapi.GetHierarchyRelationsInput{UserId: userID}
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
	return c.GetEventRecordsWithContext(ctx, since, limit)
}
func (c *client) GetEventRecordsWithContext(ctx context.Context, since int, limit *int32) (records []eventsource.Record, err error) {
	input := iam_grpcapi.GetEventRecordsInput{Since: int64(since)}
	if limit != nil {
		input.Limit = &common.PrimitiveInt32{Value: *limit}
	}

	output, err := c.api.GetEventRecords(ctx, &input)
	if err != nil {
		return
	}

	err = json.Unmarshal(output.Records, &records)
	return
}
