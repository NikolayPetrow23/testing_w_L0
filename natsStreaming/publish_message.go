package natsStreaming

import (
	"log"
	"os"
)

func OrderMessage() {
	// Функция публикации сообщения (json) для создания заказа.

	sc, err := connectionNuts(os.Getenv("NUTS_SUBSCRIBER"))

	message := []byte(`{
		"order_uid": "b563feb7b2b84b6test",
		"track_number": "WBILMTESTTRACK",
		"entry": "WBIL",
		"delivery": 1,
		"payment": 1,
		"items": 1,
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shard_key": "9",
		"sm_id": 99,
		"oof_shard": "1"
	}`)

	err = sc.Publish("foo", message)

	if err != nil {
		log.Fatal(err)
	}

	defer sc.Close()

}
