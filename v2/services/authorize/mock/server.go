package mock

import (
	"context"

	authorize "github.com/SKF/proto/authorize"
	"github.com/SKF/proto/common"
	"github.com/stretchr/testify/mock"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type MockAuthorizeServer struct {
	mock.Mock
}

func NewMockServer() *MockAuthorizeServer {
	return &MockAuthorizeServer{}
}

func (s *MockAuthorizeServer) MakeGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	authorize.RegisterAuthorizeServer(grpcServer, s)
	reflection.Register(grpcServer)

	return grpcServer
}

func (s *MockAuthorizeServer) DeepPing(ctx context.Context, void *common.Void) (*common.PrimitiveString, error) {
	args := s.Called(ctx, void)
	return args.Get(0).(*common.PrimitiveString), args.Error(1)
}

func (s *MockAuthorizeServer) LogClientState(ctx context.Context, clientInfo *authorize.LogClientStateInput) (*common.Void, error) {
	args := s.Called(ctx, clientInfo)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) IsAuthorized(ctx context.Context, in *authorize.IsAuthorizedInput) (*authorize.IsAuthorizedOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.IsAuthorizedOutput), args.Error(1)
}

func (s *MockAuthorizeServer) IsAuthorizedBulk(ctx context.Context, in *authorize.IsAuthorizedBulkInput) (*authorize.IsAuthorizedBulkOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.IsAuthorizedBulkOutput), args.Error(1)
}

func (s *MockAuthorizeServer) IsAuthorizedByEndpoint(ctx context.Context, in *authorize.IsAuthorizedByEndpointInput) (*authorize.IsAuthorizedByEndpointOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.IsAuthorizedByEndpointOutput), args.Error(1)
}

func (s *MockAuthorizeServer) AddResource(ctx context.Context, in *authorize.AddResourceInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) RemoveResource(ctx context.Context, in *authorize.RemoveResourceInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) GetResource(ctx context.Context, in *authorize.GetResourceInput) (*authorize.GetResourceOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourceOutput), args.Error(1)
}

func (s *MockAuthorizeServer) AddResources(ctx context.Context, in *authorize.AddResourcesInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) RemoveResources(ctx context.Context, in *authorize.RemoveResourcesInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) GetResourcesByUserAction(ctx context.Context, in *authorize.GetResourcesByUserActionInput) (*authorize.GetResourcesByUserActionOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesByUserActionOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetResourcesByType(ctx context.Context, in *authorize.GetResourcesByTypeInput) (*authorize.GetResourcesByTypeOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesByTypeOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetResourceParents(ctx context.Context, in *authorize.GetResourceParentsInput) (*authorize.GetResourcesOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetResourceChildren(ctx context.Context, in *authorize.GetResourceChildrenInput) (*authorize.GetResourcesOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetActionsByUserRole(ctx context.Context, in *authorize.GetActionsByUserRoleInput) (*authorize.GetActionsByUserRoleOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetActionsByUserRoleOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetResourcesAndActionsByUser(ctx context.Context, in *authorize.GetResourcesAndActionsByUserInput) (*authorize.GetResourcesAndActionsByUserOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesAndActionsByUserOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetResourcesAndActionsByUserAndResource(ctx context.Context, in *authorize.GetResourcesAndActionsByUserAndResourceInput) (*authorize.GetResourcesAndActionsByUserAndResourceOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesAndActionsByUserAndResourceOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetResourcesByOriginAndType(ctx context.Context, in *authorize.GetResourcesByOriginAndTypeInput) (*authorize.GetResourcesByOriginAndTypeOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesByOriginAndTypeOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetResourcesWithActionsAccess(ctx context.Context, in *authorize.GetResourcesWithActionsAccessInput) (*authorize.GetResourcesWithActionsAccessOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesWithActionsAccessOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetUserIDsWithAccessToResource(ctx context.Context, in *authorize.GetUserIDsWithAccessToResourceInput) (*authorize.GetUserIDsWithAccessToResourceOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetUserIDsWithAccessToResourceOutput), args.Error(1)
}

func (s *MockAuthorizeServer) AddResourceRelation(ctx context.Context, in *authorize.AddResourceRelationInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) RemoveResourceRelation(ctx context.Context, in *authorize.RemoveResourceRelationInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) AddResourceRelations(ctx context.Context, in *authorize.AddResourceRelationsInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) RemoveResourceRelations(ctx context.Context, in *authorize.RemoveResourceRelationsInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) ApplyUserAction(ctx context.Context, in *authorize.ApplyUserActionInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) GetUserActions(ctx context.Context, in *authorize.GetUserActionsInput) (*authorize.GetUserActionsOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetUserActionsOutput), args.Error(1)
}

func (s *MockAuthorizeServer) RemoveUserAction(ctx context.Context, in *authorize.RemoveUserActionInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) AddUserRole(ctx context.Context, in *authorize.UserRole) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) GetUserRole(ctx context.Context, in *authorize.GetUserRoleInput) (*authorize.UserRole, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.UserRole), args.Error(1)
}

func (s *MockAuthorizeServer) RemoveUserRole(ctx context.Context, in *authorize.RemoveUserRoleInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) AddAction(ctx context.Context, in *authorize.AddActionInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) RemoveAction(ctx context.Context, in *authorize.RemoveActionInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *MockAuthorizeServer) GetAction(ctx context.Context, in *authorize.GetActionInput) (*authorize.GetActionOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetActionOutput), args.Error(1)
}

func (s *MockAuthorizeServer) GetAllActions(ctx context.Context, in *common.Void) (*authorize.GetAllActionsOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetAllActionsOutput), args.Error(1)
}
