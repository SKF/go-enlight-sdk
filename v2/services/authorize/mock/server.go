package mock

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"

	authorize "github.com/SKF/proto/v2/authorize"
	"github.com/SKF/proto/v2/common"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AuthorizeServer struct {
	mock.Mock
	grpc *grpc.Server
	done chan error
	sock *os.File
}

func createSockFile() (file *os.File, err error) {
	if file, err = ioutil.TempFile("", "authorize-grpc-server-*.sock"); err != nil {
		return
	}

	err = os.RemoveAll(file.Name())

	return
}

func NewServer() (server *AuthorizeServer, err error) {
	server = &AuthorizeServer{
		grpc: grpc.NewServer(),
	}

	authorize.RegisterAuthorizeServer(server.grpc, server)
	reflection.Register(server.grpc)

	if server.sock, err = createSockFile(); err != nil {
		return
	}

	lis, err := net.Listen("unix", server.sock.Name())
	if err != nil {
		return
	}

	server.done = make(chan error)
	go func() {
		server.done <- server.grpc.Serve(lis)
	}()

	return
}

func (s *AuthorizeServer) HostPort() (string, string) {
	return "unix", fmt.Sprintf("//%s", s.sock.Name())
}

func (s *AuthorizeServer) AssertExpectations(t *testing.T) {
	s.grpc.Stop()
	require.NoError(t, <-s.done)
	s.Mock.AssertExpectations(t)
}

func (s *AuthorizeServer) DeepPing(ctx context.Context, void *common.Void) (*common.PrimitiveString, error) {
	args := s.Called(ctx, void)
	return args.Get(0).(*common.PrimitiveString), args.Error(1)
}

func (s *AuthorizeServer) LogClientState(ctx context.Context, clientInfo *authorize.LogClientStateInput) (*common.Void, error) {
	args := s.Called(ctx, clientInfo)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) IsAuthorized(ctx context.Context, in *authorize.IsAuthorizedInput) (*authorize.IsAuthorizedOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.IsAuthorizedOutput), args.Error(1)
}

func (s *AuthorizeServer) IsAuthorizedBulk(ctx context.Context, in *authorize.IsAuthorizedBulkInput) (*authorize.IsAuthorizedBulkOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.IsAuthorizedBulkOutput), args.Error(1)
}

func (s *AuthorizeServer) IsAuthorizedByEndpoint(ctx context.Context, in *authorize.IsAuthorizedByEndpointInput) (*authorize.IsAuthorizedByEndpointOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.IsAuthorizedByEndpointOutput), args.Error(1)
}

func (s *AuthorizeServer) AddResource(ctx context.Context, in *authorize.AddResourceInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) RemoveResource(ctx context.Context, in *authorize.RemoveResourceInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) GetResource(ctx context.Context, in *authorize.GetResourceInput) (*authorize.GetResourceOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourceOutput), args.Error(1)
}

func (s *AuthorizeServer) AddResources(ctx context.Context, in *authorize.AddResourcesInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) RemoveResources(ctx context.Context, in *authorize.RemoveResourcesInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) GetResourcesByUserAction(ctx context.Context, in *authorize.GetResourcesByUserActionInput) (*authorize.GetResourcesByUserActionOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesByUserActionOutput), args.Error(1)
}

func (s *AuthorizeServer) GetResourcesByType(ctx context.Context, in *authorize.GetResourcesByTypeInput) (*authorize.GetResourcesByTypeOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesByTypeOutput), args.Error(1)
}

func (s *AuthorizeServer) GetResourceParents(ctx context.Context, in *authorize.GetResourceParentsInput) (*authorize.GetResourcesOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesOutput), args.Error(1)
}

func (s *AuthorizeServer) GetResourceChildren(ctx context.Context, in *authorize.GetResourceChildrenInput) (*authorize.GetResourcesOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesOutput), args.Error(1)
}

func (s *AuthorizeServer) GetActionsByUserRole(ctx context.Context, in *authorize.GetActionsByUserRoleInput) (*authorize.GetActionsByUserRoleOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetActionsByUserRoleOutput), args.Error(1)
}

func (s *AuthorizeServer) GetResourcesAndActionsByUser(ctx context.Context, in *authorize.GetResourcesAndActionsByUserInput) (*authorize.GetResourcesAndActionsByUserOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesAndActionsByUserOutput), args.Error(1)
}

func (s *AuthorizeServer) GetResourcesAndActionsByUserAndResource(ctx context.Context, in *authorize.GetResourcesAndActionsByUserAndResourceInput) (*authorize.GetResourcesAndActionsByUserAndResourceOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesAndActionsByUserAndResourceOutput), args.Error(1)
}

func (s *AuthorizeServer) GetResourcesByOriginAndType(ctx context.Context, in *authorize.GetResourcesByOriginAndTypeInput) (*authorize.GetResourcesByOriginAndTypeOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesByOriginAndTypeOutput), args.Error(1)
}

func (s *AuthorizeServer) GetResourcesWithActionsAccess(ctx context.Context, in *authorize.GetResourcesWithActionsAccessInput) (*authorize.GetResourcesWithActionsAccessOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetResourcesWithActionsAccessOutput), args.Error(1)
}

func (s *AuthorizeServer) GetUserIDsWithAccessToResource(ctx context.Context, in *authorize.GetUserIDsWithAccessToResourceInput) (*authorize.GetUserIDsWithAccessToResourceOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetUserIDsWithAccessToResourceOutput), args.Error(1)
}

func (s *AuthorizeServer) AddResourceRelation(ctx context.Context, in *authorize.AddResourceRelationInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) RemoveResourceRelation(ctx context.Context, in *authorize.RemoveResourceRelationInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) AddResourceRelations(ctx context.Context, in *authorize.AddResourceRelationsInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) RemoveResourceRelations(ctx context.Context, in *authorize.RemoveResourceRelationsInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) ApplyUserAction(ctx context.Context, in *authorize.ApplyUserActionInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) GetUserActions(ctx context.Context, in *authorize.GetUserActionsInput) (*authorize.GetUserActionsOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetUserActionsOutput), args.Error(1)
}

func (s *AuthorizeServer) RemoveUserAction(ctx context.Context, in *authorize.RemoveUserActionInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) AddUserRole(ctx context.Context, in *authorize.UserRole) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) GetUserRole(ctx context.Context, in *authorize.GetUserRoleInput) (*authorize.UserRole, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.UserRole), args.Error(1)
}

func (s *AuthorizeServer) RemoveUserRole(ctx context.Context, in *authorize.RemoveUserRoleInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) AddAction(ctx context.Context, in *authorize.AddActionInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) RemoveAction(ctx context.Context, in *authorize.RemoveActionInput) (*common.Void, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*common.Void), args.Error(1)
}

func (s *AuthorizeServer) GetAction(ctx context.Context, in *authorize.GetActionInput) (*authorize.GetActionOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetActionOutput), args.Error(1)
}

func (s *AuthorizeServer) GetAllActions(ctx context.Context, in *common.Void) (*authorize.GetAllActionsOutput, error) {
	args := s.Called(ctx, in)
	return args.Get(0).(*authorize.GetAllActionsOutput), args.Error(1)
}