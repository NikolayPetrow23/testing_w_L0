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
