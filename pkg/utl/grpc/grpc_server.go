package grpc

import (
	pb "github.com/atymkiv/echo_frame_learning/blog/cmd/grpc/routeguide"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Config represents server specific config
type Config struct {
	ListeningHost string
}

func Start(svr pb.RouteGuideServer, cfg *Config) {
	lis, err := net.Listen("tcp", cfg.ListeningHost)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, svr)
	log.Printf("Server listening on endpoint '%s'...", cfg.ListeningHost)
	grpcServer.Serve(lis)
}
