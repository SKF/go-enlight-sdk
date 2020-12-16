package authorize

import (
	"context"
	"errors"
	"fmt"
	"time"

	authorizeApi "github.com/SKF/proto/v2/authorize"
	"github.com/SKF/proto/v2/common"
)

const REQUEST_LENGTH_LIMIT = 1000

func toArrayOfPointers(resources []common.Origin) []*common.Origin {
	var result []*common.Origin
	for i := 0; i < len(resources); i++ {
		result = append(result, &resources[i])
	}

	return result
}

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
	result, err := c.api.IsAuthorized(ctx, &authorizeApi.IsAuthorizedInput{
		UserId:   userID,
		Action:   action,
		Resource: resource,
	})
	if err != nil {
		return false, err
	}

	return result.Ok, err
}

// Deprecated: This functions is deprecated in favor of
// IsAuthorizedBulkWithResources as this variant returns a list of resource IDs
// which aren't unique.
func (c *client) IsAuthorizedBulk(userID, action string, reqResources []common.Origin) ([]string, []bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.IsAuthorizedBulkWithContext(ctx, userID, action, reqResources)
}

// Deprecated: This functions is deprecated in favor of
// IsAuthorizedBulkWithResources as this variant returns a list of resource IDs
// which aren't unique.
func (c *client) IsAuthorizedBulkWithContext(ctx context.Context, userID, action string, reqResources []common.Origin) ([]string, []bool, error) {
	resources, oks, err := c.IsAuthorizedBulkWithResources(ctx, userID, action, reqResources)

	if err != nil {
		return nil, nil, err
	}

	ids := make([]string, len(resources))

	for i := range resources {
		ids[i] = resources[i].GetId()
	}

	return ids, oks, nil
}

func (c *client) IsAuthorizedBulkWithResources(ctx context.Context, userID, action string, reqResources []common.Origin) ([]common.Origin, []bool, error) {
	if err := requestLengthLimit(len(reqResources)); err != nil {
		return nil, nil, err
	}

	resourcesInput := toArrayOfPointers(reqResources)

	results, err := c.api.IsAuthorizedBulk(ctx, &authorizeApi.IsAuthorizedBulkInput{
		UserId:    userID,
		Action:    action,
		Resources: resourcesInput,
	})
	if err != nil {
		return nil, nil, err
	}

	responses := results.GetResponses()
	resources := make([]common.Origin, len(responses))
	oks := make([]bool, len(responses))

	for i := range responses {
		resource := responses[i].GetResource()
		// If running against an old server which doesn't set the resource
		if resource == nil {
			resource = &common.Origin{
				Id:       responses[i].GetResourceId(), //nolint: staticcheck
				Type:     "",
				Provider: "",
			}
		}
		resources[i] = *resource
		oks[i] = responses[i].GetOk()
	}

	return resources, oks, err
}

func (c *client) IsAuthorizedByEndpoint(api, method, endpoint, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.IsAuthorizedByEndpointWithContext(ctx, api, method, endpoint, userID)
}
func (c *client) IsAuthorizedByEndpointWithContext(ctx context.Context, api, method, endpoint, userID string) (bool, error) {
	result, err := c.api.IsAuthorizedByEndpoint(ctx, &authorizeApi.IsAuthorizedByEndpointInput{
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
	_, err := c.api.AddResource(ctx, &authorizeApi.AddResourceInput{
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
	input := authorizeApi.GetResourceInput{
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

	resourcesInput := toArrayOfPointers(resources)

	_, err := c.api.AddResources(ctx, &authorizeApi.AddResourcesInput{
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
	_, err := c.api.RemoveResource(ctx, &authorizeApi.RemoveResourceInput{
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

	resourcesInput := toArrayOfPointers(resources)

	_, err := c.api.RemoveResources(ctx, &authorizeApi.RemoveResourcesInput{
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
	input := authorizeApi.GetResourcesByUserActionInput{
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
	input := authorizeApi.GetResourcesWithActionsAccessInput{
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
	input := authorizeApi.GetResourcesByTypeInput{ResourceType: resourceType}
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
	_, err := c.api.AddResourceRelation(ctx, &authorizeApi.AddResourceRelationInput{
		Resource: &resource,
		Parent:   &parent,
	})
	return err
}

func (c *client) AddResourceRelations(resources authorizeApi.AddResourceRelationsInput) error {
	if err := requestLengthLimit(len(resources.Relation)); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.AddResourceRelationsWithContext(ctx, resources)
}

func (c *client) AddResourceRelationsWithContext(ctx context.Context, resources authorizeApi.AddResourceRelationsInput) error {
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
	_, err := c.api.RemoveResourceRelation(ctx, &authorizeApi.RemoveResourceRelationInput{
		Resource: &resource,
		Parent:   &parent,
	})
	return err
}

func (c *client) RemoveResourceRelations(resources authorizeApi.RemoveResourceRelationsInput) error {
	if err := requestLengthLimit(len(resources.Relation)); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveResourceRelationsWithContext(ctx, resources)
}

func (c *client) RemoveResourceRelationsWithContext(ctx context.Context, resources authorizeApi.RemoveResourceRelationsInput) error {
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
	_, err := c.api.ApplyUserAction(ctx, &authorizeApi.ApplyUserActionInput{
		UserId:   userID,
		Action:   action,
		Resource: resource,
	})
	return err
}

func (c *client) ApplyRolesForUserOnResources(userID string, roles []string, resources []common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.ApplyRolesForUserOnResourcesWithContext(ctx, userID, roles, resources)
}

func (c *client) ApplyRolesForUserOnResourcesWithContext(ctx context.Context, userID string, roles []string, resources []common.Origin) error {
	resourcesInput := toArrayOfPointers(resources)

	_, err := c.api.ApplyRolesForUserOnResources(ctx, &authorizeApi.ApplyRolesForUserOnResourcesInput{
		UserId:    userID,
		Roles:     roles,
		Resources: resourcesInput,
	})
	return err
}

func (c *client) RemoveUserAction(userID, action string, resource *common.Origin) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveUserActionWithContext(ctx, userID, action, resource)
}
func (c *client) RemoveUserActionWithContext(ctx context.Context, userID, action string, resource *common.Origin) error {
	_, err := c.api.RemoveUserAction(ctx, &authorizeApi.RemoveUserActionInput{
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
	input := authorizeApi.GetResourcesByOriginAndTypeInput{ResourceType: resourceType, Resource: &resource, Depth: depth}
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
	input := authorizeApi.GetResourceParentsInput{ParentOriginType: parentOriginType, Resource: &resource}
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
	input := authorizeApi.GetResourceChildrenInput{ChildOriginType: childOriginType, Resource: &resource}
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
	input := authorizeApi.GetUserIDsWithAccessToResourceInput{Resource: &resource}
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

func (c *client) GetActionsByUserRole(userRole string) ([]authorizeApi.Action, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetActionsByUserRoleWithContext(ctx, userRole)
}
func (c *client) GetActionsByUserRoleWithContext(ctx context.Context, userRole string) (actions []authorizeApi.Action, err error) {
	input := authorizeApi.GetActionsByUserRoleInput{UserRole: userRole}
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

func (c *client) GetResourcesAndActionsByUser(userID string) ([]authorizeApi.ActionResource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetResourcesAndActionsByUserWithContext(ctx, userID)
}
func (c *client) GetResourcesAndActionsByUserWithContext(ctx context.Context, userID string) (actionResources []authorizeApi.ActionResource, err error) {
	input := authorizeApi.GetResourcesAndActionsByUserInput{UserId: userID}
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

func (c *client) GetResourcesAndActionsByUserAndResource(userID string, resource *common.Origin) ([]authorizeApi.ActionResource, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetResourcesAndActionsByUserAndResourceWithContext(ctx, userID, resource)
}
func (c *client) GetResourcesAndActionsByUserAndResourceWithContext(ctx context.Context, userID string, resource *common.Origin) (actionResources []authorizeApi.ActionResource, err error) {
	input := authorizeApi.GetResourcesAndActionsByUserAndResourceInput{UserId: userID, Resource: resource}
	output, err := c.api.GetResourcesAndActionsByUserAndResource(ctx, &input)
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

func (c *client) AddAction(action authorizeApi.Action) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.AddActionWithContext(ctx, action)
}

func (c *client) AddActionWithContext(ctx context.Context, action authorizeApi.Action) error {
	_, err := c.api.AddAction(ctx, &authorizeApi.AddActionInput{Action: &action})
	return err
}

func (c *client) RemoveAction(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.RemoveActionWithContext(ctx, name)
}

func (c *client) RemoveActionWithContext(ctx context.Context, name string) error {
	_, err := c.api.RemoveAction(ctx, &authorizeApi.RemoveActionInput{Name: name})
	return err
}

func (c *client) GetAction(name string) (authorizeApi.Action, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetActionWithContext(ctx, name)
}

func (c *client) GetActionWithContext(ctx context.Context, name string) (actions authorizeApi.Action, err error) {
	input := authorizeApi.GetActionInput{Name: name}
	action, err := c.api.GetAction(ctx, &input)
	if err != nil {
		return
	}
	return *action.Action, err
}

func (c *client) GetAllActions() ([]authorizeApi.Action, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetAllActionsWithContext(ctx)
}

func (c *client) GetAllActionsWithContext(ctx context.Context) (actions []authorizeApi.Action, err error) {
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

func (c *client) GetUserActions(userID string) ([]authorizeApi.Action, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetUserActionsWithContext(ctx, userID)
}

func (c *client) GetUserActionsWithContext(ctx context.Context, userID string) (actions []authorizeApi.Action, err error) {
	result, err := c.api.GetUserActions(ctx, &authorizeApi.GetUserActionsInput{
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

func (c *client) AddUserRole(role authorizeApi.UserRole) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.AddUserRoleWithContext(ctx, role)
}

func (c *client) AddUserRoleWithContext(ctx context.Context, role authorizeApi.UserRole) error {
	_, err := c.api.AddUserRole(ctx, &role)
	return err
}

func (c *client) GetUserRole(roleName string) (authorizeApi.UserRole, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.requestTimeout)
	defer cancel()
	return c.GetUserRoleWithContext(ctx, roleName)
}

func (c *client) GetUserRoleWithContext(ctx context.Context, roleName string) (role authorizeApi.UserRole, err error) {
	result, err := c.api.GetUserRole(ctx, &authorizeApi.GetUserRoleInput{
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
	_, err := c.api.RemoveUserRole(ctx, &authorizeApi.RemoveUserRoleInput{
		RoleName: roleName,
	})
	return err
}

func (c *client) IsAuthorizedWithReason(ctx context.Context, userID, action string, resource *common.Origin) (bool, string, error) {
	result, err := c.api.IsAuthorizedWithReason(ctx, &authorizeApi.IsAuthorizedInput{
		UserId:   userID,
		Action:   action,
		Resource: resource,
	})
	if err != nil {
		return false, result.Reason, err
	}

	return result.Ok, result.Reason, err
}
