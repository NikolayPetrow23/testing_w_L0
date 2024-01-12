package my_cache

import (
	"github.com/patrickmn/go-cache"
	"log"
	"testing_w_L0/database"
)

func LoadCacheFromDB(c *cache.Cache, orders []database.Order) *cache.Cache {
	// Функция загрузки заказов из бд в кэш.

	ordersMap := OrdersMap{
		Orders: make(map[string]OrderData),
	}
	for _, order := range orders {
		orderMap := OrderToMap(order)
		ordersMap.Orders[order.OrderUid] = orderMap.Orders[order.OrderUid]
		c.Set(order.OrderUid, orderMap.Orders[order.OrderUid], cache.DefaultExpiration)
	}

	if len(c.Items()) != 0 {
		log.Println("Данные были успешно загружены в кэш из базы данных!")
	}

	if len(c.Items()) == 0 {
		log.Println("Данные не были загружены из базы данных, произошла ошибка или данных нет!")
	}

	return c
}

func GetOrderFromCache(c *cache.Cache, OrderUid string) interface{} {
	// Функция поиска заказа в кэшэ по OrderUid.

	orderData, found := c.Get(OrderUid)

	if !found {
		return nil
	}

	return orderData
}

func AddOrderToCache(c *cache.Cache, order database.Order) *cache.Cache {
	// Функция добавления нового заказа в кэш.
	orderMap := OrderToMap(order)

	c.Set(order.OrderUid, orderMap.Orders[order.OrderUid], cache.DefaultExpiration)

	return c
}
