package reconnect

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func ReconnectUnaryInterceptor(opts ...CallOption) func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	options := evaluateCallOptions(opts)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			for _, code := range options.codes {
				if code == status.Code(err) {
					cc, opts = options.newClientConn(ctx, cc, opts...)
					return invoker(ctx, method, req, reply, cc, opts...)
				}
			}
		}

		return err
	}
}
