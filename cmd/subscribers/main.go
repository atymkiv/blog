package main

import (
	"github.com/atymkiv/echo_frame_learning/blog/cmd/subscribers/service"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/config"
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/nats"
	"sync"
)

func main() {
	cfg, err := config.Load("./cmd/subscribers/config.json")
	checkErr(err)

	natsClient, err := nats.New(cfg.Nats)
	checkErr(err)

	service.Subscriber(natsClient, "users")

	service.Subscriber(natsClient, "posts")

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
