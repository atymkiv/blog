package mock

import (
	"context"
	pb "github.com/atymkiv/echo_frame_learning/blog/cmd/grpc/routeguide"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	CreatePostFn func(ctx context.Context, in *pb.Post, opts ...grpc.CallOption) (*pb.Result, error)
}

func (client *GrpcClient) CreatePost(ctx context.Context, in *pb.Post, opts ...grpc.CallOption) (*pb.Result, error) {
	return client.CreatePostFn(ctx, in, nil)
}
