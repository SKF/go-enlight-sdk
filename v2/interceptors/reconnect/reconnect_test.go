package reconnect_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/SKF/go-utility/v2/log"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	"github.com/SKF/go-enlight-sdk/v2/interceptors/reconnect"
	"github.com/SKF/go-enlight-sdk/v2/tests/server"
	pb "github.com/SKF/go-enlight-sdk/v2/tests/server/helloworld"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func Test_ReconnectInterceptor_RepeatedReconnects(t *testing.T) {
	ctx := context.Background()
	s := server.New(bufSize)
	defer s.Stop()

	interceptorCalled := 0
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithNewConnection(func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
				interceptorCalled++

				conn, err := grpc.DialContext(ctx, "bufnet",
					grpc.WithContextDialer(s.Dialer()),
					grpc.WithInsecure(),
					grpc.WithBlock(),
				)

				if err != nil {
					err = errors.Wrap(err, "inside")
					return ctx, cc, opts, err
				}

				return ctx, conn, opts, nil
			}),
		)),
		grpc.WithContextDialer(s.Dialer()),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	require.NoError(t, err, "failed to dial bufnet")
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	childCtx, _ := context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Lasse Kongo"})
	assert.NoError(t, err, "failed to call first SayHello")

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("Loop no %d", i)
		s.Stop()

		childCtx, _ = context.WithTimeout(ctx, timeout)
		_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
		assert.True(t, strings.HasPrefix(err.Error(), `failed to reconnect: inside: context deadline exceeded: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error:`), msg, err.Error())

		time.Sleep(timeoutWait)

		childCtx, _ = context.WithTimeout(ctx, timeout)
		_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
		assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = "transport: Error while dialing closed"`, msg)

		s.Start()

		childCtx, _ = context.WithTimeout(ctx, timeout)
		_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
		assert.NoError(t, err, "failed to call last SayHello", msg)
	}

	assert.Equal(t, 15, interceptorCalled)
}

func Test_ReconnectInterceptor_RepeatedReconnectsWithClose(t *testing.T) {
	ctx := context.Background()
	s := server.New(bufSize)
	defer s.Stop()

	interceptorCalled := 0
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithNewConnection(func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
				interceptorCalled++

				conn, err := grpc.DialContext(ctx, "bufnet",
					grpc.WithContextDialer(s.Dialer()),
					grpc.WithInsecure(),
					grpc.WithBlock(),
				)

				if err != nil {
					err = errors.Wrap(err, "inside")
					return ctx, cc, opts, err
				}

				_ = cc.Close()
				return ctx, conn, opts, nil
			}),
		)),
		grpc.WithContextDialer(s.Dialer()),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	require.NoError(t, err, "failed to dial bufnet")
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	childCtx, _ := context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Lasse Kongo"})
	assert.NoError(t, err, "failed to call first SayHello")

	s.Stop()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.True(t, strings.HasPrefix(err.Error(), `failed to reconnect: inside: context deadline exceeded: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error:`), err.Error())

	time.Sleep(timeoutWait)

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error: desc = "transport: Error while dialing closed"`)

	s.Start()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err)

	s.Stop()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Canceled desc = grpc: the client connection is closing`)

	time.Sleep(timeoutWait)

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Canceled desc = grpc: the client connection is closing`)

	s.Start()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err)

	s.Stop()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Canceled desc = grpc: the client connection is closing`)

	time.Sleep(timeoutWait)

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Canceled desc = grpc: the client connection is closing`)

	s.Start()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err)

	time.Sleep(timeoutWait)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `rpc error: code = DeadlineExceeded desc = context deadline exceeded`)

	assert.Equal(t, 9, interceptorCalled)
}

func Test_ReconnectInterceptor_RepeatedReconnectsWithFirstClose(t *testing.T) {
	ctx := context.Background()
	s := server.New(bufSize)
	defer s.Stop()

	interceptorCalled := 0
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithUnaryInterceptor(reconnect.UnaryInterceptor(
			reconnect.WithCodes(codes.Unavailable, codes.Canceled),
			reconnect.WithNewConnection(func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
				interceptorCalled++
				_ = cc.Close()

				conn, err := grpc.DialContext(ctx, "bufnet",
					grpc.WithContextDialer(s.Dialer()),
					grpc.WithInsecure(),
					grpc.WithBlock(),
				)

				if err != nil {
					err = errors.Wrap(err, "inside")
					return ctx, cc, opts, err
				}

				return ctx, conn, opts, nil
			}),
		)),
		grpc.WithContextDialer(s.Dialer()),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	require.NoError(t, err, "failed to dial bufnet")
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	childCtx, _ := context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Lasse Kongo"})
	assert.NoError(t, err, "failed to call first SayHello")

	s.Stop()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.True(t, strings.HasPrefix(err.Error(), `failed to reconnect: inside: context deadline exceeded: rpc error: code = Unavailable desc = all SubConns are in TransientFailure, latest connection error: connection error:`), err.Error())

	time.Sleep(timeoutWait)

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Canceled desc = grpc: the client connection is closing`)

	s.Start()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err)

	s.Stop()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Canceled desc = grpc: the client connection is closing`)

	time.Sleep(timeoutWait)

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.EqualError(t, err, `failed to reconnect: inside: context deadline exceeded: rpc error: code = Canceled desc = grpc: the client connection is closing`)

	s.Start()

	childCtx, _ = context.WithTimeout(ctx, timeout)
	_, err = client.SayHello(childCtx, &pb.HelloRequest{Name: "Kalle Anka"})
	assert.NoError(t, err)

	assert.Equal(t, 6, interceptorCalled)
}
