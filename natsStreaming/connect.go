package natsStreaming

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func ConnectNats() {
	sc, _ := stan.Connect("test-cluster", "client-123", stan.NatsURL("nats://localhost:4222"))

	sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.DurableName("my-durable"))

	fmt.Println(sc)
}
