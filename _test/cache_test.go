package _test

import (
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"testing"
	"testing_w_L0/database"
	"testing_w_L0/my_cache"
	"time"
)

var c *cache.Cache

var testOrderData = database.Order{
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

func testCache() *cache.Cache {
	c = cache.New(15*time.Minute, 20*time.Minute)
	return c
}

func TestLoadCacheFromDB(t *testing.T) {
	orders := []database.Order{testOrderData}

	c := my_cache.LoadCacheFromDB(testCache(), orders)

	assert.NotEmpty(t, c.Items(), "Кэш не должен быть пустым после загрузки")
	assert.Equal(t, len(c.Items()), 1)

	c.Flush()
}

func TestAddCacheFromDB(t *testing.T) {
	c := my_cache.AddOrderToCache(testCache(), testOrderData)

	assert.NotEmpty(t, c.Items(), "Кэш не должен быть пустым после загрузки")
	assert.Equal(t, len(c.Items()), 1)

	c.Flush()
}

func TestGetOrderFromCache(t *testing.T) {
	c := testCache()

	my_cache.AddOrderToCache(c, testOrderData)

	cacheOrder := my_cache.GetOrderFromCache(c, testOrderData.OrderUid)

	assert.NotEmpty(t, cacheOrder, "Кэш не должен быть пустым после загрузки")
}
