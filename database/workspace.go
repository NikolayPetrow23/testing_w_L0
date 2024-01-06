package database

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

//func SelectUser() {
//	c = cache.New(5*time.Minute, 10*time.Minute)
//
//	db := ConnectionDb()
//
//	var users []User
//	if err := db.Select("*").Find(&users).Error; err != nil {
//		panic("Failed to fetch data from database")
//	}
//
//	CloseConnection(db)
//
//	userMap := make(map[uint]map[string]interface{})
//	for _, user := range users {
//		userMap[user.ID] = map[string]interface{}{
//			"ID":      user.ID,
//			"Name":    user.Name,
//			"Zip":     user.Zip,
//			"City":    user.City,
//			"Address": user.Address,
//			"Region":  user.Region,
//			"Email":   user.Email,
//		}
//		c.Set(string(user.ID), userMap[user.ID], cache.DefaultExpiration)
//	}
//
//	fmt.Println(c.Items())
//
//}

func SelectOrders() {
	c = cache.New(5*time.Minute, 10*time.Minute)

	db := ConnectionDb()

	var orders []Order
	if err := db.Select("*").Preload("User").Preload("Pay").Preload("Product").Find(&orders).Error; err != nil {
		panic("Failed to fetch data from database")
	}
	//.Preload("User").Preload("Pay").Preload("Product")
	//if err := db.Preload("User").Find(&orders).Error; err != nil {
	//	panic("Failed to fetch data from database")
	//}

	CloseConnection(db)

	userMap := make(map[uint]map[string]interface{})
	for _, order := range orders {
		userMap[order.ID] = map[string]interface{}{
			"ID":          order.ID,
			"TrackNumber": order.TrackNumber,
			"Entry":       order.Entry,
			"Delivery":    order.User,
			//"User":              order.User,
			"Payment": order.Pay,
			//"Pay":               order.Pay,
			"Items": order.Product,
			//"Product":           order.Product,
			"Locale":            order.Locale,
			"InternalSignature": order.InternalSignature,
			"CustomerID":        order.CustomerID,
			"DeliveryService":   order.DeliveryService,
			"ShardKey":          order.ShardKey,
			"SmId":              order.SmId,
			"DateCreated":       order.DateCreated,
			"OofShard":          order.OofShard,
		}
		c.Set(string(order.ID), userMap[order.ID], cache.DefaultExpiration)
	}

	fmt.Println(c.Items())
}
