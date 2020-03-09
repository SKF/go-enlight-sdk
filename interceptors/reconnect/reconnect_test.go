package reconnect_test

import (
	"context"
	"testing"

	"github.com/SKF/go-utility/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/SKF/go-enlight-sdk/interceptors/reconnect"
	"github.com/SKF/go-enlight-sdk/tests/server"
	pb "github.com/SKF/go-enlight-sdk/tests/server/helloworld"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const bufSize = 1024 * 1024

func Test_ReconnectInterceptor_HappyCase(t *testing.T) {
	ctx := context.Background()
	s := server.New(bufSize)
	defer s.Stop()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithCodes(codes.Unavailable),
			reconnect.WithNewConnection(func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
				conn, err := grpc.DialContext(ctx, "bufnet",
					grpc.WithContextDialer(s.Dialer()),
					grpc.WithInsecure(),
				)

				if err != nil {
					log.Infof("Failed to dial bufnet: %v", err)
					return ctx, cc, opts, err
				}
				return ctx, conn, opts, nil
			}),
		)),
		grpc.WithContextDialer(s.Dialer()),
		grpc.WithInsecure(),
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

func Test_ReconnectInterceptor_ConnectionClosed(t *testing.T) {
	ctx := context.Background()
	s := server.New(bufSize)
	defer s.Stop()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithCodes(codes.Unavailable),
		)),
		grpc.WithContextDialer(s.Dialer()),
		grpc.WithInsecure(),
	)
	require.NoError(t, err, "failed to dial bufnet")
	defer conn.Close()

	s.Stop()

	client := pb.NewGreeterClient(conn)
	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = "transport: Error while dialing closed"`)
}
