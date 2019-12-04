package reconnect_test

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/test/bufconn"

	"github.com/SKF/go-enlight-sdk/interceptors/reconnect"
	pb "github.com/SKF/go-enlight-sdk/interceptors/reconnect/helloworld"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const bufSize = 1024 * 1024

type server struct {
	grpcServer *grpc.Server
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	go func() {
		time.Sleep(time.Millisecond * 500)
		s.grpcServer.Stop()
	}()
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func newServerDialer(t *testing.T, bufSize int) func(context.Context, string) (net.Conn, error) {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{
		grpcServer: s,
	})
	go func() {
		err := s.Serve(lis)
		require.NoError(t, err, "server exited")
	}()

	return func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}
}

func Test_ReconnectInterceptor_HappyCase(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithCodes(codes.Unavailable),
			reconnect.WithNewConnection(func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (*grpc.ClientConn, []grpc.CallOption, error) {
				conn, err := grpc.DialContext(ctx, "bufnet",
					grpc.WithContextDialer(newServerDialer(t, bufSize)),
					grpc.WithInsecure(),
				)

				if err != nil {
					log.Printf("Failed to dial bufnet: %v", err)
					return cc, opts, err
				}
				return conn, opts, nil
			}),
		)),
		grpc.WithContextDialer(newServerDialer(t, bufSize)),
		grpc.WithInsecure(),
	)
	require.NoError(t, err, "failed to dial bufnet")
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Lasse Kongo"})
	assert.NoError(t, err, "failed to call first SayHello")

	time.Sleep(time.Millisecond * 1000)
	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err, "failed to call last SayHello")

	time.Sleep(time.Millisecond * 1000)
	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err, "failed to call last SayHello")
}

func Test_ReconnectInterceptor_ConnectionClosed(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithCodes(codes.Unavailable),
		)),
		grpc.WithContextDialer(newServerDialer(t, bufSize)),
		grpc.WithInsecure(),
	)
	require.NoError(t, err, "failed to dial bufnet")
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Lasse Kongo"})
	assert.NoError(t, err, "failed to call first SayHello")

	time.Sleep(time.Millisecond * 1000)
	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = "transport: Error while dialing closed"`)
}
