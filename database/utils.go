package database

func CreateUser(
	name string, zip uint,
	city string, address string,
	region string, email string) (*User, error) {
	db := ConnectionDb()
	newUser := &User{
		Name:    name,
		Zip:     zip,
		City:    city,
		Address: address,
		Region:  region,
		Email:   email,
	}

	if err := db.Create(newUser).Error; err != nil {
		return nil, err
	}

	CloseConnection(db)

	return newUser, nil
}

func CreateOrder(
	trackNumber string, entry string,
	delivery uint, payment uint,
	items uint, locale string,
	internalSignature string, customerID string,
	deliveryService string, shardKey string,
	smId uint, oofShard string) (*Order, error) {
	db := ConnectionDb()
	newOrder := &Order{
		TrackNumber:       trackNumber,
		Entry:             entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            locale,
		InternalSignature: internalSignature,
		CustomerID:        customerID,
		DeliveryService:   deliveryService,
		ShardKey:          shardKey,
		SmId:              smId,
		OofShard:          oofShard,
	}

	if err := db.Create(newOrder).Error; err != nil {
		return nil, err
	}

	CloseConnection(db)

	return newOrder, nil
}

func CreatePay(
	transaction uint, requestId uint,
	currency string, provider string,
	amount uint, paymentDt uint,
	bank string, deliveryCost uint,
	goodsTotal uint, customFee uint) (*Pay, error) {
	db := ConnectionDb()
	newPay := &Pay{
		Transaction:  transaction,
		RequestId:    requestId,
		Currency:     currency,
		Provider:     provider,
		Amount:       amount,
		PaymentDt:    paymentDt,
		Bank:         bank,
		DeliveryCost: deliveryCost,
		GoodsTotal:   goodsTotal,
		CustomFee:    customFee,
	}

	if err := db.Create(newPay).Error; err != nil {
		return nil, err
	}

	CloseConnection(db)

	return newPay, nil
}

func CreateProduct(
	chrtID uint, trackNumber uint,
	price uint, rid string,
	name string, sale uint,
	size string, totalPrice uint,
	nmId uint, brand string,
	status uint) (*Product, error) {
	db := ConnectionDb()
	newProduct := &Product{
		ChrtID:      chrtID,
		TrackNumber: trackNumber,
		Price:       price,
		Rid:         rid,
		Name:        name,
		Sale:        sale,
		Size:        size,
		TotalPrice:  totalPrice,
		NmId:        nmId,
		Brand:       brand,
		Status:      status,
	}

	if err := db.Create(newProduct).Error; err != nil {
		return nil, err
	}

	CloseConnection(db)

	return newProduct, nil
}
