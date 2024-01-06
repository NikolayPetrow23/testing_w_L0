package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID      uint `gorm:"primaryKey"`
	Name    string
	Zip     uint
	City    string
	Address string
	Region  string
	Email   string
}

type Pay struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	Transaction  uint
	RequestId    uint
	Currency     string
	Provider     string
	Amount       uint
	PaymentDt    uint
	Bank         string
	DeliveryCost uint
	GoodsTotal   uint
	CustomFee    uint
}

type Product struct {
	gorm.Model
	ID          uint `gorm:"primaryKey"`
	ChrtID      uint
	TrackNumber uint
	Price       uint
	Rid         string
	Name        string
	Sale        uint
	Size        string
	TotalPrice  uint
	NmId        uint
	Brand       string
	Status      uint
}

type Order struct {
	gorm.Model
	OrderUid          uint `gorm:"primaryKey;column:uid"`
	TrackNumber       string
	Entry             string
	Delivery          uint
	User              User `gorm:"foreignKey:Delivery"`
	Payment           uint
	Pay               Pay `gorm:"foreignKey:Payment"`
	Items             uint
	Product           Product `gorm:"foreignKey:Items"`
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	ShardKey          string
	SmId              uint
	DateCreated       time.Time
	OofShard          string
}

func Migrations() {
	db := ConnectionDb()
	db.AutoMigrate(&User{}, &Pay{}, &Product{}, &Order{})
	CloseConnection(db)
}

//func main() {
//	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
//		{
//			ID: "202201041200",
//			Migrate: func(tx *gorm.DB) error {
//				return tx.AutoMigrate(&User{})
//			},
//			Rollback: func(tx *gorm.DB) error {
//				return tx.Migrator().DropTable(&User{})
//			},
//		},
//		// Добавьте другие миграции для ваших фикстур
//	})
//
//	if err := m.Migrate(); err != nil {
//		log.Fatal(err)
//	}
//
//	// Теперь у вас есть база данных с миграциями, включая таблицу User
//	// Вы можете добавить код для создания фикстур данных в этой таблице
//}
