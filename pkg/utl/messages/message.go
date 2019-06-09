package messages

import (
	"encoding/json"
	"github.com/nats-io/go-nats"
	"log"
	"os"
)

type Service struct {
	connectionN *nats.Conn
}

func Create(connN *nats.Conn) *Service {
	return &Service{
		connectionN: connN,
	}
}

func (s *Service) PushMessage(message interface{}, subject string) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = s.connectionN.Publish(subject, msg)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Subscribe(subject string, outPath string) error {
	_, err := s.connectionN.Subscribe(subject, func(m *nats.Msg) {
		msg := m.Data
		log.Printf("Recieved message for %s subscriber", subject)
		if err := writeToFile(outPath, msg); err != nil {
			log.Printf("failed writing msg into file from nats; err: %v", err)
		}

	})
	if err != nil {
		return err
	}

	return nil
}

func writeToFile(outPath string, data []byte) error {
	var _, err = os.Stat(outPath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(outPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}

	f, err := os.OpenFile(outPath, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(string(data) + "\n")
	if err != nil {
		return err
	}
	return nil
}
