package authorize

import (
	"context"
	"time"

	"github.com/SKF/proto/common"

	authorize_grpcapi "github.com/SKF/proto/authorize"
)

func (c *client) IsAuthorized(userID, action string, resource *common.Origin) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.IsAuthorizedWithContext(ctx, userID, action, resource)
}
func (c *client) IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error) {
	result, err := c.api.IsAuthorized(ctx, &authorize_grpcapi.IsAuthorizedInput{
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
	_, err := c.api.AddAuthorizationResource(ctx, &authorize_grpcapi.AddAuthorizationResourceInput{
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
	_, err := c.api.RemoveAuthorizationResource(ctx, &authorize_grpcapi.RemoveAuthorizationResourceInput{
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
	input := authorize_grpcapi.GetAuthorizationResourcesByTypeInput{ResourceType: resourceType}
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
	_, err := c.api.AddAuthorizationResourceRelation(ctx, &authorize_grpcapi.AddAuthorizationResourceRelationInput{
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
	_, err := c.api.RemoveAuthorizationResourceRelation(ctx, &authorize_grpcapi.RemoveAuthorizationResourceRelationInput{
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
	input := authorize_grpcapi.GetAuthorizationResourceRelationsInput{}
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
	_, err := c.api.AddUserPermission(ctx, &authorize_grpcapi.AddUserPermissionInput{
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
	_, err := c.api.RemoveUserPermission(ctx, &authorize_grpcapi.RemoveUserPermissionInput{
		UserId:   userID,
		Role:     role,
		Resource: &resource,
	})
	return err
}

func (c *client) GetResourcesByOriginAndType(originID string, resourceType string) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetResourcesByOriginAndTypeWithContext(ctx, originID, resourceType)
}
func (c *client) GetResourcesByOriginAndTypeWithContext(ctx context.Context, originID string, resourceType string) (resources []common.Origin, err error) {
	input := authorize_grpcapi.GetResourcesByOriginAndTypeInput{ResourceType: resourceType, OriginId: originID}
	output, err := c.api.GetResourcesByOriginAndType(ctx, &input)
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

func (c *client) GetUserIDsWithAccessToResource(originID string) (resources []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetUserIDsWithAccessToResourceWithContext(ctx, originID)
}
func (c *client) GetUserIDsWithAccessToResourceWithContext(ctx context.Context, originID string) (userIds []string, err error) {
	input := authorize_grpcapi.GetUserIDsWithAccessToResourceInput{OriginId: originID}
	output, err := c.api.GetUserIDsWithAccessToResource(ctx, &input)
	if err != nil {
		return
	}
	if output != nil {
		for _, userID := range output.UserIds {
			if userID != "" {
				userIds = append(userIds, userID)
			}
		}
	}
	return
}
