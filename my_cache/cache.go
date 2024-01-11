package my_cache

import (
	"github.com/patrickmn/go-cache"
	"log"
	"testing_w_L0/database"
	"time"
)

var c *cache.Cache

func LoadCacheFromDB(orders []database.Order) {
	// Функция загрузки заказов из бд в кэш.

	c = cache.New(15*time.Minute, 20*time.Minute)

	ordersMap := OrdersMap{
		Orders: make(map[string]OrderData),
	}
	for _, order := range orders {
		orderMap := orderToMap(order)
		ordersMap.Orders[order.OrderUid] = orderMap.Orders[order.OrderUid]
		c.Set(order.OrderUid, orderMap.Orders[order.OrderUid], cache.DefaultExpiration)
	}

	if len(c.Items()) != 0 {
		log.Println("Данные были успешно загружены в кэш из базы данных!")
	}

	if len(c.Items()) == 0 {
		log.Println("Данные не были загружены из базы данных, произошла ошибка или данных нет!")
	}
}

func GetOrderFromCache(OrderUid string) interface{} {
	// Функция поиска заказа в кэшэ по OrderUid.

	orderData, found := c.Get(OrderUid)

	if !found {
		return nil
	}

	return orderData
}

func AddOrderToCache(order database.Order) {
	// Функция добавления нового заказа в кэш.

	orderMap := orderToMap(order)
	c.Set(order.OrderUid, orderMap.Orders[order.OrderUid], cache.DefaultExpiration)
}
