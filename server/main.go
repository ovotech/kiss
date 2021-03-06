package server

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"

	serverauthz "github.com/ovotech/kiss/pkg/authz/server"
	"github.com/ovotech/kiss/pkg/aws"
	pb "github.com/ovotech/kiss/proto"
)

var (
	AWSManager *aws.Manager
)

func init() {
	// Set up zerolog
	initLogging()

}

type kissServer struct {
	pb.UnimplementedKISSServer
}

func newServer() *kissServer {
	s := &kissServer{}
	return s
}

// Run starts the server, which will execute asychronously until one of the channels have been
// notified. Error is returned immediately if server cannot bootstrap.
func Run(
	host string,
	port int,
	errChan chan error,
	sigChan chan os.Signal,
	jwksURL *string,
	namespacesKey *string,
	namespacesRegex *string,
	identifierKey *string,
	adminNamespace *string,
	roleBindingPrefix *string,
	kubeconfigPath *string,
	enableTracing bool,
) (*grpc.Server, error) {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}
	log.Info().Msgf("listening on %s:%d", host, port)

	authInterceptor := serverauthz.NewServerAuthzInterceptor(
		*jwksURL,
		*namespacesKey,
		*namespacesRegex,
		*identifierKey,
		*adminNamespace,
		*roleBindingPrefix,
		*kubeconfigPath,
	)
	var grpcServer *grpc.Server
	if enableTracing {

		log.Info().Msg("Starting tracing...")

		grpcServer = grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				grpctrace.UnaryServerInterceptor(grpctrace.WithAnalytics(true)),
				authInterceptor.Unary(),
			),
		)
	} else {

		grpcServer = grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				authInterceptor.Unary(),
			),
		)
	}

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	pb.RegisterKISSServer(grpcServer, newServer())

	// Prepare channel for graceful termination of requests when killed
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// Start server asychronously, based on
	// https://stackoverflow.com/a/55800690
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	return grpcServer, nil
}

func setLogLevel() {
	// Set log level (trace, debug, info, warn, error, fatal, panic) Default info
	value, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		return
	}
	logLevel, err := zerolog.ParseLevel(value)
	if err != nil {
		// Defaults to Info level
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Error().Msgf("Error parsing \"LOG_LEVEL\" with value %s", value)
		return
	}
	zerolog.SetGlobalLevel(logLevel)
}

func initLogging() {
	// Better performance (datadog is be able to process it)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// Add contextual fields to the global logger
	// Caller adds the file and line number to log
	log.Logger = log.With().Str("application", "kiss").Caller().Logger()
	// Set log level (trace, debug, info, warn, error, fatal, panic) Default info
	setLogLevel()
	log.Info().Msgf("Logging level set to %s", zerolog.GlobalLevel())
}
