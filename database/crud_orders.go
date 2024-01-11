package database

import (
	"log"
	"time"
)

type OrderDataJSON struct {
	OrderUid          string `json:"order_uid"`
	TrackNumber       string `json:"track_number"`
	Entry             string `json:"entry"`
	Delivery          uint   `json:"delivery"`
	Payment           uint   `json:"payment"`
	Items             uint   `json:"items"`
	Locale            string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerID        string `json:"customer_id"`
	DeliveryService   string `json:"delivery_service"`
	ShardKey          string `json:"shard_key"`
	SmId              uint   `json:"sm_id"`
	OofShard          string `json:"oof_shard"`
}

func SelectAllOrders() []Order {
	// Функция взятия всех заказов из бд.

	var orders []Order

	db := ConnectionDb()

	if err := db.Select("*").Preload("User").Preload(
		"Pay").Preload("Product").Find(
		&orders).Error; err != nil {
		panic("Failed to fetch data from database")
	}

	CloseConnectionDb(db)

	return orders
}

func SelectOrderUID(orderUid string) Order {
	// Функция поиска заказа по OrderUid в бд.

	var order Order

	db := ConnectionDb()

	if err := db.Select("*").Preload("User").Preload(
		"Pay").Preload("Product").Where(
		"order_uid = ?", orderUid).Find(
		&order).Error; err != nil {
		panic("Failed to fetch data from database")
	}

	CloseConnectionDb(db)

	return order
}

func CreateOrder(orderData OrderDataJSON) (string, error) {
	// Функция создания заказа в бд.

	db := ConnectionDb()

	newOrder := &Order{
		OrderUid:          orderData.OrderUid,
		TrackNumber:       orderData.TrackNumber,
		Entry:             orderData.Entry,
		Delivery:          orderData.Delivery,
		Payment:           orderData.Payment,
		Items:             orderData.Items,
		Locale:            orderData.Locale,
		InternalSignature: orderData.InternalSignature,
		CustomerID:        orderData.CustomerID,
		DeliveryService:   orderData.DeliveryService,
		ShardKey:          orderData.ShardKey,
		SmId:              orderData.SmId,
		DateCreated:       time.Now(),
		OofShard:          orderData.OofShard,
	}

	if err := db.Create(newOrder).Error; err != nil {
		return "nil", err
		log.Fatalln("Заказ не был создан, произошла ошибка!", err)
	}

	CloseConnectionDb(db)

	return newOrder.OrderUid, nil
}

func ValidateOrderDataJSON(orderData OrderDataJSON) bool {
	// Функция валидации важных полей json'a приходящего для создания заказа.

	if len(orderData.OrderUid) < 10 || len(orderData.TrackNumber) < 10 {
		return false
	}

	if orderData.Delivery < 1 || orderData.Payment < 1 || orderData.Items < 1 {
		return false
	}

	if orderData.Locale != "en" && orderData.Locale != "rus" && orderData.Locale != "kz" {
		return false
	}

	return true
}
