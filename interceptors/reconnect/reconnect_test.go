package reconnect_test

import (
	"context"
	"testing"
	"time"

	"github.com/SKF/go-enlight-sdk/interceptors/reconnect"
	"github.com/SKF/go-enlight-sdk/tests/server"
	pb "github.com/SKF/go-enlight-sdk/tests/server/helloworld"
	"github.com/SKF/go-utility/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	bufSize     = 1024 * 1024
	timeout     = time.Millisecond * 100
	timeoutWait = time.Millisecond * 150
)

func Test_ReconnectInterceptor_HappyCase(t *testing.T) {
	ctx := context.Background()
	s := server.New(bufSize)
	defer s.Stop()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithNewConnection(func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
				conn, err := grpc.DialContext(ctx, "bufnet",
					grpc.WithContextDialer(s.Dialer()),
					grpc.WithTransportCredentials(insecure.NewCredentials()),
				)

				if err != nil {
					log.Infof("Failed to dial bufnet: %v", err)
					return ctx, cc, opts, err
				}
				return ctx, conn, opts, nil
			}),
		)),
		grpc.WithContextDialer(s.Dialer()),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err, "failed to dial bufnet")
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Lasse Kongo"})
	assert.NoError(t, err, "failed to call first SayHello")

	s.Restart()

	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err, "failed to call last SayHello")

	s.Restart()

	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err, "failed to call last SayHello")
}
