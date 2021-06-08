package server

import (
	"context"
	"net"
	"strings"

	pb "github.com/ovotech/kiss/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
)

// Logs requests for audit trail.
func auditLog(
	ctx context.Context,
	authorized bool,
	request string,
	user string,
	namespace string,
) {
	clientIP := getRequestIPBestEffort(ctx)
	if clientIP == nil {
		clientIP = &net.IP{}
	}

	var msg string
	if authorized {
		msg = "authorized"
	} else {
		msg = "blocked"
	}

	log.Info().Msgf(
		"%s: %s user '%s' (from host '%s') on namespace '%s'",
		request,
		msg,
		user,
		clientIP,
		namespace,
	)
}

func getRequestIPBestEffort(ctx context.Context) *net.IP {
	// We first check if we can retrieve client IP from the load balancer forward header
	if forwardedIP, ok := getForwardedIP(ctx); ok {
		clientIP := net.ParseIP(forwardedIP)
		return &clientIP
	}

	// If client IP not forwarded via load balancer, use the source IP in gRPC
	p, ok := peer.FromContext(ctx)
	if ok {
		clientIP := net.ParseIP(strings.Split(p.Addr.String(), ":")[0])
		return &clientIP
	}

	return nil
}

func getForwardedIP(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	if forwardedIP, ok := getFirstValueForRequestHeader(md, pb.XForwardedForHeader); ok {
		return forwardedIP, true
	}

	return "", false
}

func getFirstValueForRequestHeader(md metadata.MD, key string) (string, bool) {
	headers := md.Get(key)
	if len(headers) == 0 || len(headers[0]) == 0 {
		return "", false
	}

	return headers[0], true
}
