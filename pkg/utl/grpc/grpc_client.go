package grpc

import (
	pb "github.com/atymkiv/echo_frame_learning/blog/cmd/grpc/routeguide"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/config"
	"google.golang.org/grpc"
	"log"
)

func New(cfg *config.GRPC) (pb.RouteGuideClient, error) {
	conn, err := grpc.Dial(cfg.ListeningHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldn't connect to %v, err: %v", cfg.ListeningHost, err)
		return nil, err
	}

	client := pb.NewRouteGuideClient(conn)

	return client, nil

}
