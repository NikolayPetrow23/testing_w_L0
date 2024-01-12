package nats_streaming

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"testing_w_L0/database"
	"testing_w_L0/my_cache"
)

func connectionNuts(clientID string) (stan.Conn, error) {
	godotenv.Load()

	sc, err := stan.Connect(
		os.Getenv("NUTS_CLUSTER"),
		clientID,
		stan.NatsURL(
			fmt.Sprintf("nats://%s:%s", os.Getenv("NUTS_HOST"), os.Getenv("NUTS_PORT")),
		),
	)
	if err != nil {
		log.Fatal("Biba: ", err, os.Getenv("NUTS_PORT"), os.Getenv("NUTS_HOST"))
	}

	return sc, err
}

func SubscribeNuts() {
	// Функция подписки на канал.

	sc, err := connectionNuts(os.Getenv("NUTS_SUBSCRIBE"))

	subscription, err := sc.Subscribe(
		"foo",
		func(msg *stan.Msg) {
			messageProcessing(msg)

			fmt.Printf("Сообщение: %s\n", string(msg.Data))

			if err != nil {
				log.Println("Неверные данные в сообщении!", err)
				return
			}
			msg.Ack()
		},
		stan.DurableName("durable-name"),
		stan.StartWithLastReceived(),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer subscription.Close()

	fmt.Println("Subscriber is waiting for messages. Press Ctrl+C to exit.")

	select {}

}

func messageProcessing(msg *stan.Msg) {
	// Функция обработки сообщения из строки в json.

	var order database.OrderDataJSON

	err := json.Unmarshal(msg.Data, &order)

	if err != nil {
		log.Println("Неверные данные в сообщении!", err)
		return
	}

	if database.ValidateOrderDataJSON(order) != true {
		log.Println("Данные не прошли валидацию!")
		return
	}

	newOrderUID, err := database.CreateOrder(order)
	selectOrder := database.SelectOrderUID(newOrderUID)
	my_cache.AddOrderToCache(my_cache.Cache(), selectOrder)
}
