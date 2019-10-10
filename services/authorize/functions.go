package authorize

import (
	"context"
	"errors"
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

func (c *client) SetRequestTimeout(d time.Duration) {
	c.requestTimeout = d
}

func (c *client) IsAuthorized(userID, action string, resource *common.Origin) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.AddResourceWithContext(ctx, resource)
}

func (c *client) AddResourceWithContext(ctx context.Context, resource common.Origin) error {
	_, err := c.api.AddResource(ctx, &authorize_grpcapi.AddResourceInput{
		Resource: &resource,
	})
	return err
}

func (c *client) GetResource(id string, originType string) (common.Origin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetResourceWithContext(ctx, id, originType)
}
func (c *client) GetResourceWithContext(ctx context.Context, id string, originType string) (common.Origin, error) {
	input := authorize_grpcapi.GetResourceInput{
		Id:         id,
		OriginType: originType,
	}
	resource, err := c.api.GetResource(ctx, &input)

	if err != nil {
		return common.Origin{}, err
	}

	return *resource.Resource, err
}

func (c *client) AddResources(resources []common.Origin) error {
	if err := requestLengthLimit(len(resources)); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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

	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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

func (c *client) GetResourcesByUserAction(userID, actionName, resourceType string) ([]common.Origin, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetResourcesByUserActionWithContext(ctx, userID, actionName, resourceType)
}
func (c *client) GetResourcesByUserActionWithContext(ctx context.Context, userID, actionName, resourceType string) (resources []common.Origin, err error) {
	input := authorize_grpcapi.GetResourcesByUserActionInput{
		UserId:       userID,
		Action:       actionName,
		ResourceType: resourceType,
	}
	output, err := c.api.GetResourcesByUserAction(ctx, &input)
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

func (c *client) GetResourcesWithActionsAccess(actions []string, resourceType string, resource *common.Origin) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetResourcesWithActionsAccessWithContext(ctx, actions, resourceType, resource)
}
func (c *client) GetResourcesWithActionsAccessWithContext(ctx context.Context, actions []string, resourceType string, resource *common.Origin) (resources []common.Origin, err error) {
	input := authorize_grpcapi.GetResourcesWithActionsAccessInput{
		Actions:      actions,
		ResourceType: resourceType,
		Resource:     resource,
	}
	output, err := c.api.GetResourcesWithActionsAccess(ctx, &input)
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

func (c *client) GetResourcesByType(resourceType string) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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

	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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

	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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

func (c *client) ApplyUserAction(userID, action string, resource *common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.ApplyUserActionWithContext(ctx, userID, action, resource)
}
func (c *client) ApplyUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	_, err := c.api.ApplyUserAction(ctx, &authorize_grpcapi.ApplyUserActionInput{
		UserId:   userID,
		Action:   action,
		Resource: resource,
	})
	return err
}

func (c *client) RemoveUserAction(userID, action string, resource *common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveUserActionWithContext(ctx, userID, action, resource)
}
func (c *client) RemoveUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	_, err := c.api.RemoveUserAction(ctx, &authorize_grpcapi.RemoveUserActionInput{
		UserId:   userID,
		Action:   action,
		Resource: resource,
	})
	return err
}

func (c *client) GetResourcesByOriginAndType(resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetResourcesByOriginAndTypeWithContext(ctx, resource, resourceType, depth)
}
func (c *client) GetResourcesByOriginAndTypeWithContext(ctx context.Context, resource common.Origin, resourceType string, depth int32) (resources []common.Origin, err error) {
	input := authorize_grpcapi.GetResourcesByOriginAndTypeInput{ResourceType: resourceType, Resource: &resource, Depth: depth}
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

func (c *client) GetResourceParents(resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetResourceParentsWithContext(ctx, resource, parentOriginType)
}
func (c *client) GetResourceParentsWithContext(ctx context.Context, resource common.Origin, parentOriginType string) (resources []common.Origin, err error) {	
	input := authorize_grpcapi.GetResourceParentsInput{ParentOriginType: parentOriginType, Resource: &resource}
	output, err := c.api.GetResourceParents(ctx, &input)
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

func (c *client) GetResourceChildren(resource common.Origin, childOriginType string) (resources []common.Origin, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetResourceChildrenWithContext(ctx, resource, childOriginType)
}
func (c *client) GetResourceChildrenWithContext(ctx context.Context, resource common.Origin, childOriginType string) (resources []common.Origin, err error) {	
	input := authorize_grpcapi.GetResourceChildrenInput{ChildOriginType: childOriginType, Resource: &resource}
	output, err := c.api.GetResourceChildren(ctx, &input)
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




func (c *client) GetUserIDsWithAccessToResource(resource common.Origin) (resources []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetUserIDsWithAccessToResourceWithContext(ctx, resource)
}
func (c *client) GetUserIDsWithAccessToResourceWithContext(ctx context.Context, resource common.Origin) (userIds []string, err error) {
	input := authorize_grpcapi.GetUserIDsWithAccessToResourceInput{Resource: &resource}
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
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
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

func (c *client) AddAction(action authorize_grpcapi.Action) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.AddActionWithContext(ctx, action)
}

func (c *client) AddActionWithContext(ctx context.Context, action authorize_grpcapi.Action) error {
	_, err := c.api.AddAction(ctx, &authorize_grpcapi.AddActionInput{Action: &action})
	return err
}

func (c *client) RemoveAction(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveActionWithContext(ctx, name)
}

func (c *client) RemoveActionWithContext(ctx context.Context, name string) error {
	_, err := c.api.RemoveAction(ctx, &authorize_grpcapi.RemoveActionInput{Name: name})
	return err
}

func (c *client) GetAction(name string) (authorize_grpcapi.Action, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetActionWithContext(ctx, name)
}

func (c *client) GetActionWithContext(ctx context.Context, name string) (actions authorize_grpcapi.Action, err error) {
	input := authorize_grpcapi.GetActionInput{Name: name}
	action, err := c.api.GetAction(ctx, &input)
	if err != nil {
		return
	}
	return *action.Action, err
}

func (c *client) GetAllActions() ([]authorize_grpcapi.Action, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetAllActionsWithContext(ctx)
}

func (c *client) GetAllActionsWithContext(ctx context.Context) (actions []authorize_grpcapi.Action, err error) {
	allActions, err := c.api.GetAllActions(ctx, &common.Void{})
	if err != nil {
		return
	}
	if allActions != nil {
		for _, action := range allActions.Actions {
			if action != nil {
				actions = append(actions, *action)
			}
		}
	}

	return
}

func (c *client) GetUserActions(userID string) ([]authorize_grpcapi.Action, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetUserActionsWithContext(ctx, userID)
}

func (c *client) GetUserActionsWithContext(ctx context.Context, userID string) (actions []authorize_grpcapi.Action, err error) {
	result, err := c.api.GetUserActions(ctx, &authorize_grpcapi.GetUserActionsInput{
		UserId: userID,
	})
	if err != nil {
		return
	}
	if result != nil {
		for _, action := range result.Actions {
			if action != nil {
				actions = append(actions, *action)
			}
		}
	}

	return
}

func (c *client) AddUserRole(role authorize_grpcapi.UserRole) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.AddUserRoleWithContext(ctx, role)
}

func (c *client) AddUserRoleWithContext(ctx context.Context, role authorize_grpcapi.UserRole) error {
	_, err := c.api.AddUserRole(ctx, &role)
	return err
}

func (c *client) GetUserRole(roleName string) (authorize_grpcapi.UserRole, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetUserRoleWithContext(ctx, roleName)
}

func (c *client) GetUserRoleWithContext(ctx context.Context, roleName string) (role authorize_grpcapi.UserRole, err error) {
	result, err := c.api.GetUserRole(ctx, &authorize_grpcapi.GetUserRoleInput{
		RoleName: roleName,
	})
	if err != nil {
		return
	}

	if result != nil {
		role = *result
	} else {
		err = errors.New("No result")
	}
	return
}

func (c *client) RemoveUserRole(roleName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveUserRoleWithContext(ctx, roleName)
}

func (c *client) RemoveUserRoleWithContext(ctx context.Context, roleName string) error {
	_, err := c.api.RemoveUserRole(ctx, &authorize_grpcapi.RemoveUserRoleInput{
		RoleName: roleName,
	})
	return err
}
