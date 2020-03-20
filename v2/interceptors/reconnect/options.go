package reconnect

import (
	"context"

	"google.golang.org/grpc"
)

var (
	defaultOptions = &options{
		newClientConn: func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
			return ctx, cc, opts, nil
		},
	}
)

type options struct {
	newClientConn NewConnectionFunc
}

type NewConnectionFunc func(context.Context, *grpc.ClientConn, ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error)
type CallOption func(opt *options)

func WithNewConnection(f NewConnectionFunc) CallOption {
	return func(o *options) {
		o.newClientConn = f
	}
}

func evaluateCallOptions(opts []CallOption) *options {
	copy := &options{}
	*copy = *defaultOptions
	for _, opt := range opts {
		opt(copy)
	}
	return copy
}
