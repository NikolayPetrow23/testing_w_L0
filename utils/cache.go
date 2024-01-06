//package database
//
//import (
//	"time"
//
//	"github.com/patrickmn/go-cache"
//)
//
//var c *cache.Cache
//
//func init() {
//	c = cache.New(5*time.Minute, 10*time.Minute)
//}
//
//func FetchDataFromDB() ([]User, error) {
//	db := ConnectionDb()
//
//	var result []User
//
//	if err := db.Find(&result).Error; err != nil {
//		return nil, err
//	}
//	return result, nil
//}

//func cacheDataFromDB() {
//	dataFromDB, err := FetchDataFromDB()
//	if err != nil {
//		// Обработка ошибки при получении данных из базы данных
//		fmt.Println("Ошибка при получении данных из базы данных:", err)
//		return
//	}
//
//	// Сохраните данные в кэше
//	for _, data := range dataFromDB {
//		c.Set(
//			data.Key, data.Value, cache.DefaultExpiration,
//		)
//	}
//
//	fmt.Println("Данные успешно закэшированы.")
//}
