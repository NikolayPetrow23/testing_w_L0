package natsStreaming

import (
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func ConnectionNats() {
	natsURL := "nats://localhost:4222"

	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subject := "example_subject"

	message := []byte("Hello, NATS Streaming!")

	err = nc.Publish(subject, message)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message published successfully")

	time.Sleep(1 * time.Second)
}
