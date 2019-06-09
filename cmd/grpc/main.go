package main

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/grpc/service"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/config"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/grpc"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/messages"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/nats"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/redis"
)

func main() {
	cfg, err := config.Load("./cmd/grpc/config.json")
	checkErr(err)

	natsClient, err := nats.New(cfg.Nats)
	checkErr(err)

	redisClient, err := redis.New(cfg.Redis)
	checkErr(err)

	messageService := messages.Create(natsClient)
	grpcService := service.New(messageService, redisClient)
	grpc.Start(grpcService, &grpc.Config{cfg.GRPC.ListeningHost})
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
