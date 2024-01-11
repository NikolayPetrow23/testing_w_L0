package my_cache

import (
	"testing_w_L0/database"
	"time"
)

type OrderData struct {
	ID                uint
	OrderUid          string
	TrackNumber       string
	Entry             string
	Delivery          database.User
	Payment           database.Pay
	Items             database.Product
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	ShardKey          string
	SmID              uint
	DateCreated       time.Time
	OofShard          string
}

type OrdersMap struct {
	Orders map[string]OrderData
}

func orderToMap(order database.Order) OrdersMap {
	// Функция преобразования данных из бд в структуру OrderData для добавления в кэш.

	orderMap := OrdersMap{
		Orders: make(map[string]OrderData),
	}

	orderMap.Orders[order.OrderUid] = OrderData{
		ID:                order.ID,
		OrderUid:          order.OrderUid,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Delivery:          order.User,
		Payment:           order.Pay,
		Items:             order.Product,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerID:        order.CustomerID,
		DeliveryService:   order.DeliveryService,
		ShardKey:          order.ShardKey,
		SmID:              order.SmId,
		DateCreated:       order.DateCreated,
		OofShard:          order.OofShard,
	}

	return orderMap
}
