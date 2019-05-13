package authorize

import (
	"context"
	"fmt"
	"time"

	authorize_grpcapi "github.com/SKF/proto/authorize"
	"github.com/SKF/proto/common"
)

const REQUEST_LENGTH_LIMIT = 1000

func requestLengthLimit(requestLength int) error {
	if requestLength > REQUEST_LENGTH_LIMIT {
		return fmt.Errorf("request length limit exceeded. max: %d actual: %d", REQUEST_LENGTH_LIMIT, requestLength)
	}
	return nil
}

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

func (c *client) IsAuthorizedBulk(userID, action string, resources []common.Origin) ([]string, []bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.IsAuthorizedBulkWithContext(ctx, userID, action, resources)
}

func (c *client) IsAuthorizedBulkWithContext(ctx context.Context, userID, action string, resources []common.Origin) ([]string, []bool, error) {
	if err := requestLengthLimit(len(resources)); err != nil {
		return nil, nil, err
	}

	var resourcesInput []*common.Origin
	for i := 0; i < len(resources); i++ {
		resourcesInput = append(resourcesInput, &resources[i])
	}

	results, err := c.api.IsAuthorizedBulk(ctx, &authorize_grpcapi.IsAuthorizedBulkInput{
		UserId:    userID,
		Action:    action,
		Resources: resourcesInput,
	})
	if err != nil {
		return nil, nil, err
	}

	items := results.Responses
	var ids []string
	var oks []bool
	for _, item := range items {
		ids = append(ids, item.ResourceId)
		oks = append(oks, item.Ok)
	}

	return ids, oks, err
}

func (c *client) IsAuthorizedByEndpoint(api, method, endpoint, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.IsAuthorizedByEndpointWithContext(ctx, api, method, endpoint, userID)
}
func (c *client) IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string) (bool, error) {
	result, err := c.api.IsAuthorizedByEndpoint(ctx, &authorize_grpcapi.IsAuthorizedByEndpointInput{
		Api:      api,
		Method:   method,
		Endpoint: endpoint,
		UserId:   userID,
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
	if err := requestLengthLimit(len(resources)); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddResourcesWithContext(ctx, resources)
}

func (c *client) AddResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	if err := requestLengthLimit(len(resources)); err != nil {
		return err
	}

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
	if err := requestLengthLimit(len(resources)); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveResourcesWithContext(ctx, resources)
}

func (c *client) RemoveResourcesWithContext(ctx context.Context, resources []common.Origin) error {
	if err := requestLengthLimit(len(resources)); err != nil {
		return err
	}

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
	if err := requestLengthLimit(len(resources.Relation)); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddResourceRelationsWithContext(ctx, resources)
}

func (c *client) AddResourceRelationsWithContext(ctx context.Context, resources authorize_grpcapi.AddResourceRelationsInput) error {
	if err := requestLengthLimit(len(resources.Relation)); err != nil {
		return err
	}

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
	if err := requestLengthLimit(len(resources.Relation)); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveResourceRelationsWithContext(ctx, resources)
}

func (c *client) RemoveResourceRelationsWithContext(ctx context.Context, resources authorize_grpcapi.RemoveResourceRelationsInput) error {
	if err := requestLengthLimit(len(resources.Relation)); err != nil {
		return err
	}

	_, err := c.api.RemoveResourceRelations(ctx, &resources)
	return err
}

func (c *client) AddUserPermission(userID, action string, resource *common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.AddUserPermissionWithContext(ctx, userID, action, resource)
}
func (c *client) AddUserPermissionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	_, err := c.api.AddUserPermission(ctx, &authorize_grpcapi.AddUserPermissionInput{
		UserId:   userID,
		Action:   action,
		Resource: resource,
	})
	return err
}

func (c *client) RemoveUserPermission(userID, action string, resource *common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.RemoveUserPermissionWithContext(ctx, userID, action, resource)
}
func (c *client) RemoveUserPermissionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	_, err := c.api.RemoveUserPermission(ctx, &authorize_grpcapi.RemoveUserPermissionInput{
		UserId:   userID,
		Action:   action,
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

func (c *client) GetActionsByUserRole(userRole string) ([]authorize_grpcapi.Action, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetActionsByUserRoleWithContext(ctx, userRole)
}
func (c *client) GetActionsByUserRoleWithContext(ctx context.Context, userRole string) (actions []authorize_grpcapi.Action, err error) {
	input := authorize_grpcapi.GetActionsByUserRoleInput{UserRole: userRole}
	output, err := c.api.GetActionsByUserRole(ctx, &input)
	if err != nil {
		return
	}

	if output != nil {
		for _, action := range output.Actions {
			if action != nil {
				actions = append(actions, *action)
			}
		}
	}

	return
}

func (c *client) GetResourcesAndActionsByUser(userID string) ([]authorize_grpcapi.ActionResource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	return c.GetResourcesAndActionsByUserWithContext(ctx, userID)
}
func (c *client) GetResourcesAndActionsByUserWithContext(ctx context.Context, userID string) (actionResources []authorize_grpcapi.ActionResource, err error) {
	input := authorize_grpcapi.GetResourcesAndActionsByUserInput{UserId: userID}
	output, err := c.api.GetResourcesAndActionsByUser(ctx, &input)
	if err != nil {
		return
	}

	if output != nil {
		for _, elem := range output.Data {
			if elem != nil {
				actionResources = append(actionResources, *elem)
			}
		}
	}

	return
}
