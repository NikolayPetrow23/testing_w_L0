package database

import (
	"time"
)

type User struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Phone   string
	Zip     string
	City    string
	Address string
	Region  string
	Email   string
}

type Pay struct {
	ID           uint `gorm:"primaryKey"`
	Transaction  string
	RequestID    string
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
	ID          uint `gorm:"primaryKey"`
	ChrtID      uint
	TrackNumber string
	Price       uint
	Rid         string
	Name        string
	Sale        uint
	Size        string
	TotalPrice  uint
	NmID        uint
	Brand       string
	Status      uint
}

type Order struct {
	ID                uint   `gorm:"primaryKey"`
	OrderUid          string `gorm:"type:varchar(64);uniqueIndex"`
	TrackNumber       string `gorm:"type:varchar(64);uniqueIndex"`
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
