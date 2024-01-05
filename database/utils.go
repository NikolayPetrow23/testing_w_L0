package database

func CreateUser(name string, zip uint, city string, address string, region string, email string) (*User, error) {
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
