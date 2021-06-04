package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientAuthInterceptor struct {
	accessToken string
}

// Returns a new ClientAuthInterceptor to attach the authorization access token to outgoing
// requests.
func NewClientAuthInterceptor(accessToken string) *ClientAuthInterceptor {
	return &ClientAuthInterceptor{accessToken: accessToken}
}

// Unary interceptor to attach the access token to outgoing requests.
func (i *ClientAuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		return invoker(i.attachToken(ctx), method, req, reply, cc, opts...)
	}
}

// Attaches the authorization access token to the outgoing context.
func (i *ClientAuthInterceptor) attachToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", i.accessToken)
}
