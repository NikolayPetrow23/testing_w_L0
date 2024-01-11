package database

func Migrations() {
	db := ConnectionDb()
	db.AutoMigrate(&User{}, &Pay{}, &Product{})
	db.AutoMigrate(&Order{})
	CloseConnectionDb(db)
}
