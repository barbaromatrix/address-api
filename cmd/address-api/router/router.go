package router

import (
	"net"
	"time"

	"address-api/internal/api"
	"address-api/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lockp111/go-easyzap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	echoTracer "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

func NewMetaApp() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(echoTracer.Middleware(echoTracer.WithServiceName(config.GetConfig().AppName)))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST},
	}))
	e.GET("/prometheus", echo.WrapHandler(promhttp.Handler()))
	api.NewHealthCheck().Register(e)
	return e
}

func Start(e *echo.Echo, host string) {
	err := e.Start(host)
	easyzap.Errorf("Failed to start the server. Error: ", err)
}

func RunServer() error {
	serverEnforcement := keepalive.EnforcementPolicy{
		// By default, gRPC disconnects clients that send "too many" pings,
		// but we don't really care about that, so configure the server to be
		// as permissive as possible.
		MinTime:             time.Nanosecond,
		PermitWithoutStream: true,
	}
	serverKeepalive := keepalive.ServerParameters{
		// Send periodic pings on the connection when there is no other traffic.
		// PingInterval is the interval between network heartbeat pings. It is used
		// both for RPC heartbeat intervals and gRPC server keepalive pings. It is
		// set to 1 second in order to fail fast, but with large default timeouts
		// to tolerate high-latency multiregion clusters.
		Time: time.Second,
		// If the pings don't get a response within the timeout, we might be
		// experiencing a network partition. gRPC will close the transport-level
		// connection and all the pending RPCs (which may not have timeouts) will
		// fail eagerly.
		Timeout: 2 * time.Second,
		// MaxConnectionAge is a duration for the maximum amount of time a
		// connection may exist before it will be closed by sending a GoAway. A
		// random jitter of +/-10% will be added to MaxConnectionAge to spread out
		// connection storms.
		MaxConnectionAge: 1 * time.Minute,
		// MaxConnectionAgeGrace is an additive period after MaxConnectionAge after
		// which the connection will be forcibly closed.
		MaxConnectionAgeGrace: 1 * time.Minute,
	}

	grpcServer := grpc.NewServer(
		// The default number of concurrent streams/requests on a client connection
		// is 100, while the server is unlimited. The client setting can only be
		// controlled by adjusting the server value. Set a very large value for the
		// server value so that we have no fixed limit on the number of concurrent
		// streams/requests on either the client or server.
		grpc.MaxConcurrentStreams(400),
		grpc.KeepaliveParams(serverKeepalive),
		grpc.KeepaliveEnforcementPolicy(serverEnforcement),
	)
	reflection.Register(grpcServer)
	//cfg := config.GetConfig()
	//ipClient := client.NewIIpClient(cfg)
	//ipService := service.NewIpService(ipClient)
	listener, err := net.Listen("tcp", ":"+config.GetConfig().ServerHost)
	if err != nil {
		easyzap.Errorf("Error while initializing server: ", err)
	}
	easyzap.Info("Running gRPC server")

	return grpcServer.Serve(listener)
}
