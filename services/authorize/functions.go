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

func (c *client) IsAuthorizedByEndpoint(api, method, endpoint, userID string, resource *common.Origin) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.IsAuthorizedByEndpointWithContext(ctx, api, method, endpoint, userID, resource)
}
func (c *client) IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string, resource *common.Origin) (bool, error) {
	result, err := c.api.IsAuthorizedByEndpoint(ctx, &authorize_grpcapi.IsAuthorizedByEndpointInput{
		Api:      api,
		Method:   method,
		Endpoint: endpoint,
		UserId:   userID,
		Resource: resource,
	})
	if err != nil {
		return false, err
	}

	return result.Ok, err
}

func (c *client) AddResource(resource common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddResourceWithContext(ctx, resource)
}

func (c *client) AddResourceWithContext(ctx context.Context, resource common.Origin) error {
	_, err := c.api.AddResource(ctx, &authorize_grpcapi.AddResourceInput{
		Resource: &resource,
	})
	return err
}

func (c *client) AddResources(resources []common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddResourcesWithContext(ctx, resources)
}

func (c *client) AddResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	var resourcesInput []*common.Origin
	for i := 0; i < len(resources); i++ {
		resourcesInput = append(resourcesInput, &resources[i])
	}

	_, err := c.api.AddResources(ctx, &authorize_grpcapi.AddResourcesInput{
		Resource: resourcesInput,
	})
	return err
}

func (c *client) RemoveResource(resource common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveResourceWithContext(ctx, resource)
}
func (c *client) RemoveResourceWithContext(ctx context.Context, resource common.Origin) error {
	_, err := c.api.RemoveResource(ctx, &authorize_grpcapi.RemoveResourceInput{
		Resource: &resource,
	})
	return err
}

func (c *client) RemoveResources(resources []common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveResourcesWithContext(ctx, resources)
}

func (c *client) RemoveResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	var resourcesInput []*common.Origin
	for i := 0; i < len(resources); i++ {
		resourcesInput = append(resourcesInput, &resources[i])
	}

	_, err := c.api.RemoveResources(ctx, &authorize_grpcapi.RemoveResourcesInput{
		Resource: resourcesInput,
	})
	return err
}

func (c *client) GetResourcesByType(resourceType string) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetResourcesByTypeWithContext(ctx, resourceType)
}
func (c *client) GetResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error) {
	input := authorize_grpcapi.GetResourcesByTypeInput{ResourceType: resourceType}
	output, err := c.api.GetResourcesByType(ctx, &input)
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

func (c *client) AddResourceRelation(resource common.Origin, parent common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddResourceRelationWithContext(ctx, resource, parent)
}
func (c *client) AddResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	_, err := c.api.AddResourceRelation(ctx, &authorize_grpcapi.AddResourceRelationInput{
		Resource: &resource,
		Parent:   &parent,
	})
	return err
}

func (c *client) AddResourceRelations(resources authorize_grpcapi.AddResourceRelationsInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddResourceRelationsWithContext(ctx, resources)
}

func (c *client) AddResourceRelationsWithContext(ctx context.Context, resources authorize_grpcapi.AddResourceRelationsInput) error {
	_, err := c.api.AddResourceRelations(ctx, &resources)
	return err
}

func (c *client) RemoveResourceRelation(resource common.Origin, parent common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveResourceRelationWithContext(ctx, resource, parent)
}
func (c *client) RemoveResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error {
	_, err := c.api.RemoveResourceRelation(ctx, &authorize_grpcapi.RemoveResourceRelationInput{
		Resource: &resource,
		Parent:   &parent,
	})
	return err
}

func (c *client) RemoveResourceRelations(resources authorize_grpcapi.RemoveResourceRelationsInput) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveResourceRelationsWithContext(ctx, resources)
}

func (c *client) RemoveResourceRelationsWithContext(ctx context.Context, resources authorize_grpcapi.RemoveResourceRelationsInput) error {
	_, err := c.api.RemoveResourceRelations(ctx, &resources)
	return err
}

func (c *client) GetResourceRelations(resource common.Origin) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetResourceRelationsWithContext(ctx, resource)
}
func (c *client) GetResourceRelationsWithContext(ctx context.Context, resource common.Origin) (resources []common.Origin, err error) {
	input := authorize_grpcapi.GetResourceRelationsInput{}
	output, err := c.api.GetResourceRelations(ctx, &input)
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

func (c *client) AddUserPermission(userID, role string, resource *common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddUserPermissionWithContext(ctx, userID, role, resource)
}
func (c *client) AddUserPermissionWithContext(ctx context.Context, userID, role string, resource *common.Origin) error {
	_, err := c.api.AddUserPermission(ctx, &authorize_grpcapi.AddUserPermissionInput{
		UserId:   userID,
		Role:     role,
		Resource: resource,
	})
	return err
}

func (c *client) RemoveUserPermission(userID, role string, resource *common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveUserPermissionWithContext(ctx, userID, role, resource)
}
func (c *client) RemoveUserPermissionWithContext(ctx context.Context, userID, role string, resource *common.Origin) error {
	_, err := c.api.RemoveUserPermission(ctx, &authorize_grpcapi.RemoveUserPermissionInput{
		UserId:   userID,
		Role:     role,
		Resource: resource,
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
