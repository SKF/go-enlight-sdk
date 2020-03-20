package reconnect

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	defaultOptions = &options{
		codes: []codes.Code{},
		newClientConn: func(ctx context.Context, cc *grpc.ClientConn, opts ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error) {
			return ctx, cc, opts, nil
		},
	}
)

type options struct {
	codes         []codes.Code
	newClientConn NewConnectionFunc
}

type NewConnectionFunc func(context.Context, *grpc.ClientConn, ...grpc.CallOption) (context.Context, *grpc.ClientConn, []grpc.CallOption, error)
type CallOption func(opt *options)

func WithCodes(errorCodes ...codes.Code) CallOption {
	return func(o *options) {
		o.codes = errorCodes
	}
}

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
