package authorize_test

import (
	"context"
	"testing"
	"time"

	"github.com/SKF/go-enlight-sdk/v2/services/authorize"
	authMock "github.com/SKF/go-enlight-sdk/v2/services/authorize/mock"
	grpcapi "github.com/SKF/proto/v2/authorize"
	"github.com/SKF/proto/v2/common"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func clientFor(t *testing.T, server *authMock.AuthorizeServer) authorize.AuthorizeClient {
	host, port := server.HostPort()

	client := authorize.CreateClient()
	server.On("LogClientState", mock.Anything, mock.Anything).Return(&common.Void{}, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	require.NoError(t, client.DialWithContext(ctx, host, port, grpc.WithInsecure()))

	return client
}

func Test_DeepPing(t *testing.T) {
	server, err := authMock.NewServer()
	require.NoError(t, err)

	client := clientFor(t, server)

	server.On("DeepPing", mock.Anything, mock.Anything).Return(&common.PrimitiveString{Value: ""}, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	err = client.DeepPingWithContext(ctx)
	assert.NoError(t, err)

	server.AssertExpectations(t)
}

func Test_IsAuthorizedBulkWithResources(t *testing.T) {
	server, err := authMock.NewServer()
	require.NoError(t, err)

	client := clientFor(t, server)

	server.On("IsAuthorizedBulk", mock.Anything, &grpcapi.IsAuthorizedBulkInput{
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
		Return(&grpcapi.IsAuthorizedBulkOutput{
			Responses: []*grpcapi.IsAuthorizedOutItem{
				&grpcapi.IsAuthorizedOutItem{
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	res, oks, err := client.IsAuthorizedBulkWithResources(ctx, "testUser", "testAction", []common.Origin{
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

	server.AssertExpectations(t)
}

// Test that if the server does *not* include the resource in the reply the
// client doesn't crash
func Test_IsAuthorizedBulkWithResourcesNoResourceInResonse(t *testing.T) {
	server, err := authMock.NewServer()
	require.NoError(t, err)

	client := clientFor(t, server)

	server.On("IsAuthorizedBulk", mock.Anything, &grpcapi.IsAuthorizedBulkInput{
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
		Return(&grpcapi.IsAuthorizedBulkOutput{
			Responses: []*grpcapi.IsAuthorizedOutItem{
				&grpcapi.IsAuthorizedOutItem{
					ResourceId: "0",
					Ok:         true,
				},
			},
		}, nil)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	res, oks, err := client.IsAuthorizedBulkWithResources(ctx, "testUser", "testAction", []common.Origin{
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

	server.AssertExpectations(t)
}
