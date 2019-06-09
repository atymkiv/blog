package service

import (
	"github.com/atymkiv/echo_frame_learning/blog/pkg/utl/messages"
	"github.com/nats-io/go-nats"
	"log"
)

func Subscriber(connN *nats.Conn, subj string) {
	natsService := messages.Create(connN)
	outPath := "/home/" + subj + ".txt"

	err := natsService.Subscribe(subj, outPath)
	if err != nil {
		log.Println(err)
	}
	log.Printf("New subscriber, listening subject %s", subj)

}
