package _test

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"testing_w_L0/database"
	"testing_w_L0/my_cache"
	"time"
)

var TestOrderData = database.Order{
	ID:          1,
	OrderUid:    "test_order_uid",
	TrackNumber: "test_track_number",
	Entry:       "test_entry",
	Delivery:    1,
	User: database.User{
		ID:      1,
		Name:    "test_user",
		Phone:   "+99999999999",
		Zip:     "test_zip",
		City:    "test_city",
		Address: "test_address",
		Region:  "test_region",
		Email:   "test_email",
	},
	Payment: 1,
	Pay: database.Pay{
		ID:           1,
		Transaction:  "test_transaction",
		RequestID:    "test_request",
		Currency:     "USD",
		Provider:     "test_provider",
		Amount:       123,
		PaymentDt:    123,
		Bank:         "test_bank",
		DeliveryCost: 123,
		GoodsTotal:   123,
		CustomFee:    123,
	},
	Items: 1,
	Product: database.Product{
		ID:          1,
		ChrtID:      123,
		TrackNumber: "test_track_number",
		Price:       123,
		Rid:         "test_rid",
		Name:        "test_name",
		Sale:        123,
		Size:        "test_size",
		TotalPrice:  123,
		NmID:        123,
		Brand:       "",
		Status:      123,
	},
	Locale:            "en",
	InternalSignature: "test_internal_signature",
	CustomerID:        "test_customer_id",
	DeliveryService:   "test_delivery_service",
	ShardKey:          "test_shard_key",
	SmId:              123,
	DateCreated:       time.Now(),
	OofShard:          "test_oof_shard",
}

func TestLoadCacheFromDB(t *testing.T) {
	var c *cache.Cache

	orders := []database.Order{TestOrderData}

	my_cache.LoadCacheFromDB(orders)

	fmt.Println(c.Items())

	assert.NotEmpty(t, c.Items(), "Кэш не должен быть пустым после загрузки")

	c.Flush()
}
