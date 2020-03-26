package authorize_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"

	"github.com/SKF/go-enlight-sdk/v2/services/authorize"
	authMock "github.com/SKF/go-enlight-sdk/v2/services/authorize/mock"
	authAPI "github.com/SKF/proto/authorize"
	"github.com/SKF/proto/common"
	"google.golang.org/grpc"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func newSocket(t *testing.T) (net.Listener, *os.File) {
	file, err := ioutil.TempFile("", "grpc-server-*.sock")

	require.NoError(t, err)

	err = os.RemoveAll(file.Name())

	require.NoError(t, err)

	lis, err := net.Listen("unix", file.Name())

	require.NoError(t, err)

	return lis, file
}

func createServer(t *testing.T) (*authMock.MockAuthorizeServer, *grpc.Server, string, <-chan error) {
	listener, sock := newSocket(t)

	mockServer := authMock.NewMockServer()
	grpcServer := mockServer.MakeGRPCServer()

	done := make(chan error)

	go func() {
		done <- grpcServer.Serve(listener)
	}()

	return mockServer, grpcServer, fmt.Sprintf("unix://%s", sock.Name()), done
}

func Test_DeepPing(t *testing.T) {
	mockServer, grpcServer, host, done := createServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	client := authorize.CreateClient()

	err := client.DialTarget(ctx, host, grpc.WithInsecure())

	require.NoError(t, err)

	mockServer.On("DeepPing", mock.Anything, mock.Anything).Return(&common.PrimitiveString{Value: ""}, nil)

	err = client.DeepPingWithContext(ctx)

	assert.NoError(t, err)

	grpcServer.Stop()

	assert.NoError(t, <-done)

	mockServer.AssertExpectations(t)
}

func Test_IsAuthorizedBulkWithResources(t *testing.T) {
	mockServer, grpcServer, host, done := createServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	client := authorize.CreateClient()

	err := client.DialTarget(ctx, host, grpc.WithInsecure())

	require.NoError(t, err)

	mockServer.On("IsAuthorizedBulk", mock.Anything, &authAPI.IsAuthorizedBulkInput{
		UserId: "testUser",
		Action: "testAction",
		Resources: []*common.Origin{
			&common.Origin{
				Id:       "0",
				Type:     "node",
				Provider: "1",
			},
		},
	}).
		Return(&authAPI.IsAuthorizedBulkOutput{
			Responses: []*authAPI.IsAuthorizedOutItem{
				&authAPI.IsAuthorizedOutItem{
					ResourceId: "0",
					Ok:         true,
					Resource: &common.Origin{
						Id:       "0",
						Type:     "node",
						Provider: "1",
					},
				},
			},
		}, nil)

	res, oks, err := client.IsAuthorizedBulkWithResources(context.Background(), "testUser", "testAction", []common.Origin{
		{
			Id:       "0",
			Type:     "node",
			Provider: "1",
		},
	})

	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Len(t, oks, 1)

	assert.Equal(t, "0", res[0].GetId())
	assert.Equal(t, "node", res[0].GetType())
	assert.Equal(t, "1", res[0].GetProvider())
	assert.True(t, oks[0])

	grpcServer.Stop()

	assert.NoError(t, <-done)

	mockServer.AssertExpectations(t)
}

// Test that if the server does *not* include the resource in the reply the
// client doesn't crash
func Test_IsAuthorizedBulkWithResourcesNoResourceInResonse(t *testing.T) {
	mockServer, grpcServer, host, done := createServer(t)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	client := authorize.CreateClient()

	err := client.DialTarget(ctx, host, grpc.WithInsecure())

	require.NoError(t, err)

	mockServer.On("IsAuthorizedBulk", mock.Anything, &authAPI.IsAuthorizedBulkInput{
		UserId: "testUser",
		Action: "testAction",
		Resources: []*common.Origin{
			&common.Origin{
				Id:       "0",
				Type:     "node",
				Provider: "1",
			},
		},
	}).
		Return(&authAPI.IsAuthorizedBulkOutput{
			Responses: []*authAPI.IsAuthorizedOutItem{
				&authAPI.IsAuthorizedOutItem{
					ResourceId: "0",
					Ok:         true,
				},
			},
		}, nil)

	res, oks, err := client.IsAuthorizedBulkWithResources(context.Background(), "testUser", "testAction", []common.Origin{
		{
			Id:       "0",
			Type:     "node",
			Provider: "1",
		},
	})

	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Len(t, oks, 1)

	assert.Equal(t, "0", res[0].GetId())
	assert.Equal(t, "", res[0].GetType())
	assert.Equal(t, "", res[0].GetProvider())
	assert.True(t, oks[0])

	grpcServer.Stop()

	assert.NoError(t, <-done)

	mockServer.AssertExpectations(t)
}
