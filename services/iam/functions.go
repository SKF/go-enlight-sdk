package iam

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/SKF/proto/common"

	"github.com/SKF/go-eventsource/eventsource"
	iam_grpcapi "github.com/SKF/proto/iam"
)

func (c *client) CheckAuthentication(token, arn string) (claims iam_grpcapi.UserClaims, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.CheckAuthenticationWithContext(ctx, token, arn)
}
func (c *client) CheckAuthenticationWithContext(ctx context.Context, token, arn string) (claims iam_grpcapi.UserClaims, err error) {
	input := &iam_grpcapi.CheckAuthenticationInput{Token: token, MethodArn: arn}
	output, err := c.api.CheckAuthentication(ctx, input)
	if output != nil {
		claims = *output
	}
	if output == nil && err == nil {
		err = errors.New("No output")
	}
	return
}

func (c *client) CheckAuthenticationByEndpoint(token, api, method, endpoint string) (claims iam_grpcapi.UserClaims, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.CheckAuthenticationByEndpointWithContext(ctx, token, api, method, endpoint)
}
func (c *client) CheckAuthenticationByEndpointWithContext(ctx context.Context, token, api, method, endpoint string) (claims iam_grpcapi.UserClaims, err error) {
	input := &iam_grpcapi.CheckAuthenticationByEndpointInput{
		Api:      api,
		Token:    token,
		Method:   method,
		Endpoint: endpoint,
	}
	output, err := c.api.CheckAuthenticationByEndpoint(ctx, input)
	if output != nil {
		claims = *output
	}
	if output == nil && err == nil {
		err = errors.New("No output")
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

func (c *client) IsAuthorized(userID, action string, resource *common.Origin) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.IsAuthorizedWithContext(ctx, userID, action, resource)
}
func (c *client) IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error) {
	result, err := c.api.IsAuthorized(ctx, &iam_grpcapi.IsAuthorizedInput{
		UserId:   userID,
		Action:   action,
		Resource: resource,
	})
	if err != nil {
		return false, err
	}

	return result.Ok, err
}

func (c *client) AddAuthorizationResource(resource common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddAuthorizationResourceWithContext(ctx, resource)
}
func (c *client) AddAuthorizationResourceWithContext(ctx context.Context, resource common.Origin) error {
	_, err := c.api.AddAuthorizationResource(ctx, &iam_grpcapi.AddAuthorizationResourceInput{
		Resource: &resource,
	})
	return err
}

func (c *client) RemoveAuthorizationResource(resource common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveAuthorizationResourceWithContext(ctx, resource)
}
func (c *client) RemoveAuthorizationResourceWithContext(ctx context.Context, resource common.Origin) error {
	_, err := c.api.RemoveAuthorizationResource(ctx, &iam_grpcapi.RemoveAuthorizationResourceInput{
		Resource: &resource,
	})
	return err
}

func (c *client) GetAuthorizationResourcesByType(resourceType string) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetAuthorizationResourcesByTypeWithContext(ctx, resourceType)
}
func (c *client) GetAuthorizationResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error) {
	input := iam_grpcapi.GetAuthorizationResourcesByTypeInput{ResourceType: resourceType}
	output, err := c.api.GetAuthorizationResourcesByType(ctx, &input)
	if err != nil {
		return
	}
	if output != nil {
		for _, resource := range output.Resources {
			if resource != nil {
				resources = append(resources, *resource)
			}
		}
	}
	return
}

func (c *client) AddAuthorizationResourceRelation(resource common.Origin, parent common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddAuthorizationResourceRelationWithContext(ctx, resource, parent)
}
func (c *client) AddAuthorizationResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	_, err := c.api.AddAuthorizationResourceRelation(ctx, &iam_grpcapi.AddAuthorizationResourceRelationInput{
		Resource: &resource,
		Parent:   &parent,
	})
	return err
}

func (c *client) RemoveAuthorizationResourceRelation(resource common.Origin, parent common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveAuthorizationResourceRelationWithContext(ctx, resource, parent)
}
func (c *client) RemoveAuthorizationResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	_, err := c.api.RemoveAuthorizationResourceRelation(ctx, &iam_grpcapi.RemoveAuthorizationResourceRelationInput{
		Resource: &resource,
		Parent:   &parent,
	})
	return err
}

func (c *client) GetAuthorizationResourceRelations(resource common.Origin) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetAuthorizationResourceRelationsWithContext(ctx, resource)
}
func (c *client) GetAuthorizationResourceRelationsWithContext(ctx context.Context, resource common.Origin) (resources []common.Origin, err error) {
	input := iam_grpcapi.GetAuthorizationResourceRelationsInput{}
	output, err := c.api.GetAuthorizationResourceRelations(ctx, &input)
	if err != nil {
		return
	}
	if output != nil {
		for _, resource := range output.Resources {
			if resource != nil {
				resources = append(resources, *resource)
			}
		}
	}
	return
}

func (c *client) AddUserPermission(userID, role string, resource common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddUserPermissionWithContext(ctx, userID, role, resource)
}
func (c *client) AddUserPermissionWithContext(ctx context.Context, userID, role string, resource common.Origin) error {
	_, err := c.api.AddUserPermission(ctx, &iam_grpcapi.AddUserPermissionInput{
		UserId:   userID,
		Role:     role,
		Resource: &resource,
	})
	return err
}

func (c *client) RemoveUserPermission(userID, role string, resource common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveUserPermissionWithContext(ctx, userID, role, resource)
}
func (c *client) RemoveUserPermissionWithContext(ctx context.Context, userID, role string, resource common.Origin) error {
	_, err := c.api.RemoveUserPermission(ctx, &iam_grpcapi.RemoveUserPermissionInput{
		UserId:   userID,
		Role:     role,
		Resource: &resource,
	})
	return err
}
