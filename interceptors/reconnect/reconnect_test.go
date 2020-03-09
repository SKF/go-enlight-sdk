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
	stop           bool
	listener       *bufconn.Listener
	restartChannel chan time.Duration

	Server      *grpc.Server
	TestService *service
}

type service struct {
	calls int
}

func (s *service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	s.calls++
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func newServer(t *testing.T, bufferSize int) *server {
	s := &server{
		stop:           false,
		restartChannel: make(chan time.Duration),
	}

	// serverRunning := make(chan bool)
	go func() {
		for !s.stop {
			s.Server = grpc.NewServer()
			s.TestService = &service{}
			pb.RegisterGreeterServer(s.Server, s.TestService)
			go func() {
				s.listener = bufconn.Listen(bufSize)
				s.Server.Serve(s.listener)
				// require.NoError(t, err, "server exited")
			}()

			// serverRunning <- true

			d := <-s.restartChannel
			s.Server.Stop()
			time.Sleep(d)
		}
	}()

	// <-serverRunning

	time.Sleep(time.Second * 1)
	return s
}

func (s *server) Dialer() func(context.Context, string) (net.Conn, error) {
	return func(context.Context, string) (net.Conn, error) {
		return s.listener.Dial()
	}
}

func (s *server) NumberOfClientCalls() int {
	return s.TestService.calls
}

func (s *server) Restart(d time.Duration) {
	s.restartChannel <- d
}

func (s *server) RestartWithBlocking(d time.Duration) {
	s.Restart(d)
	time.Sleep(d)
}

func (s *server) Close() {
	s.stop = true
	s.Restart(0 * time.Second)
}

// func

// func newServerDialer(t *testing.T, bufSize int) func(context.Context, string) (net.Conn, error) {
// 	lis := bufconn.Listen(bufSize)
// 	s := grpc.NewServer()
// 	pb.RegisterGreeterServer(s, &server{
// 		grpcServer: s,
// 	})
// 	go func() {
// 		err := s.Serve(lis)
// 		require.NoError(t, err, "server exited")
// 	}()

// 	return func(context.Context, string) (net.Conn, error) {
// 		return lis.Dial()
// 	}
// }

func Test_ReconnectInterceptor_HappyCase(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	s := newServer(t, bufSize)
	defer s.Close()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithCodes(codes.Unavailable),
			reconnect.WithNewConnection(func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
				conn, err := grpc.DialContext(ctx, "bufnet",
					grpc.WithContextDialer(s.Dialer()),
					grpc.WithInsecure(),
				)

				if err != nil {
					log.Printf("Failed to dial bufnet: %v", err)
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

	s.RestartWithBlocking(time.Millisecond * 0)
	time.Sleep(time.Millisecond * 100)

	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err, "failed to call last SayHello")

	s.RestartWithBlocking(time.Millisecond * 0)

	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err, "failed to call last SayHello")

	// s.Close()

	// _, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Anders And"})
	// assert.Error(t, err, "context deadline exceeded")
}

// func Test_ReconnectInterceptor_ConnectionClosed(t *testing.T) {
// 	ctx := context.Background()
// 	conn, err := grpc.DialContext(ctx, "bufnet",
// 		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
// 			reconnect.WithCodes(codes.Unavailable),
// 		)),
// 		grpc.WithContextDialer(newServerDialer(t, bufSize)),
// 		grpc.WithInsecure(),
// 	)
// 	require.NoError(t, err, "failed to dial bufnet")
// 	defer conn.Close()

// 	client := pb.NewGreeterClient(conn)
// 	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Lasse Kongo"})
// 	assert.NoError(t, err, "failed to call first SayHello")

// 	time.Sleep(time.Millisecond * 1000)
// 	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
// 	assert.EqualError(t, err, `rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = "transport: Error while dialing closed"`)
// }

// func Test_ReconnectInterceptor_ConnectionClosedRetry(t *testing.T) {
// 	ctx := context.Background()
// 	conn, err := grpc.DialContext(ctx, "bufnet",
// 		grpc.WithUnaryInterceptor(
// 			reconnect.UnaryInterceptor(
// 				reconnect.WithCodes(codes.Unavailable),
// 			),
// 			retry.UnaryClientInterceptor(
// 				retry.WithCodes(codes.Unavailable),
// 				retry.WithMax(3),
// 				retry.WithBackoff(func(attempts int) time.Duration {
// 					log.Printf("attempts %d", attempts)
// 					return 50 * time.Millisecond
// 				}),
// 			),
// 		),
// 		grpc.WithContextDialer(newServerDialer(t, bufSize)),
// 		grpc.WithInsecure(),
// 	)
// 	require.NoError(t, err, "failed to dial bufnet")
// 	defer conn.Close()

// 	client := pb.NewGreeterClient(conn)
// 	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Lasse Kongo"})
// 	assert.NoError(t, err, "failed to call first SayHello")

// 	time.Sleep(time.Millisecond * 1000)
// 	_, err = client.SayHello(ctx, &pb.HelloRequest{Name: "Kalle Anka"})
// 	assert.EqualError(t, err, `rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = "transport: Error while dialing closed"`)
// }
