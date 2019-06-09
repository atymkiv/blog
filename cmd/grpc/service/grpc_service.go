package service

import (
	"context"
	"encoding/json"
	pb "github.com/atymkiv/echo_frame_learning/blog/cmd/grpc/routeguide"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/messages"
	"github.com/go-redis/redis"
	"log"
)

const QUEUE = "POST_"
const TOPIC = "posts"

func New(natsService *messages.Service, red *redis.Client) *GrpcRedisServer {
	return &GrpcRedisServer{redisConnection: red, natsService: natsService}
}

type GrpcRedisServer struct {
	redisConnection *redis.Client
	natsService     *messages.Service
}

func (s *GrpcRedisServer) CreatePost(ctx context.Context, post *pb.Post) (*pb.Result, error) {
	//push post into redis
	redisKey := QUEUE + post.Id
	dbRecord := pb.Post{post.Id, post.From, post.Body}
	marshalledValues, _ := json.Marshal(dbRecord)

	err := s.redisConnection.RPush(redisKey, string(marshalledValues)).Err()
	result := pb.Result{0}

	if err != nil {
		result.Code = -1
	}

	if err := s.natsService.PushMessage(post, TOPIC); err != nil {
		log.Fatalf("failed pushing post into nuts; err: %v", err)
		return nil, err
	}

	return &result, err
}
