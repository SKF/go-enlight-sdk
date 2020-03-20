package reconnect

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/SKF/go-utility/v2/log"
)

func UnaryInterceptor(opts ...CallOption) func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	options := evaluateCallOptions(opts)
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			switch cc.GetState() {
			case connectivity.Idle, connectivity.Connecting, connectivity.Ready:
			default:
				var triggerReconnect = len(options.codes) == 0
				for _, code := range options.codes {
					triggerReconnect = triggerReconnect || code == status.Code(err)
				}

				if triggerReconnect {
					log.WithTracing(ctx).
						WithError(err).
						WithField("code", status.Code(err)).
						WithField("state", cc.GetState().String()).
						Debug("Calling reconnect function")

					if ctx.Err() == nil {
						// Can/should we extend the context ??
						// THe invoker can eat all of the time and we have no time left to reconnect
						// Can use ctx.Deadline to be able to reconnect

						newCtx, newCC, newOpts, errConn := options.newClientConn(ctx, cc, opts...)
						if errConn != nil {
							return errors.Wrapf(err, "failed to reconnect: %s", errConn.Error())
						}
						return invoker(newCtx, method, req, reply, newCC, newOpts...)
					}
				}
			}
		}

		return err
	}
}
