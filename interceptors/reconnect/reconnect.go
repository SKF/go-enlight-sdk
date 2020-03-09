package reconnect

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/SKF/go-utility/v2/log"
)

func UnaryInterceptor(opts ...CallOption) func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	options := evaluateCallOptions(opts)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			for _, code := range options.codes {
				if code == status.Code(err) {
					log.WithTracing(ctx).WithError(err).WithField("code", code).Debug("Calling reconnect function")
					newCtx, newCC, newOpts, errConn := options.newClientConn(ctx, cc, opts...)
					if errConn != nil {
						return errors.Wrapf(err, "failed to re-connect: %s", errConn.Error())
					}

					return invoker(newCtx, method, req, reply, newCC, newOpts...)
				}
			}
		}

		return err
	}
}
