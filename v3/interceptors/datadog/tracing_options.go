package datadog

import (
	"context"
	"encoding/binary"
	"fmt"
	"strings"

	oc_trace "go.opencensus.io/trace"
	oc_propagation "go.opencensus.io/trace/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	dd_tracer "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// DatadogToOpenCensusUnaryInterceptor is used when your application is using Datadog for tracing,
// but the called client is using OpenCensus.
//
// The interceptor will create a Datadog span and inject the span as gRPC metadata according to:
// https://github.com/census-instrumentation/opencensus-go/tree/master/plugin/ocgrpc
func DatadogToOpenCensusUnaryInterceptor(serviceName string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		const traceContextKey = "grpc-trace-bin"

		resourceName := strings.TrimPrefix(method, "/")
		resourceName = strings.Replace(resourceName, "/", ".", -1)

		span, ctx := dd_tracer.StartSpanFromContext(
			ctx, "grpc.request",
			dd_tracer.ServiceName(serviceName),
			dd_tracer.ResourceName(resourceName),
		)

		defer span.Finish()

		ocSpanContext, err := ddToOCSpanContext(ctx)
		if err != nil {
			return fmt.Errorf("couldn't convert Datadog span to OpenCensus span-context: %w", err)
		}

		traceContextBinary := oc_propagation.Binary(ocSpanContext)
		ctx = metadata.AppendToOutgoingContext(ctx, traceContextKey, string(traceContextBinary))

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// ddToOCSpanContext takes a context including a Datadog span and
// returns an OpenCensus SpanContext
func ddToOCSpanContext(ctx context.Context) (ocSpan oc_trace.SpanContext, err error) {
	span, exists := dd_tracer.SpanFromContext(ctx)
	if !exists {
		err = fmt.Errorf("couldn't find span in context")
		return
	}

	binary.BigEndian.PutUint64(ocSpan.TraceID[8:16], span.Context().TraceID())
	binary.BigEndian.PutUint64(ocSpan.SpanID[0:8], span.Context().SpanID())

	return
}
