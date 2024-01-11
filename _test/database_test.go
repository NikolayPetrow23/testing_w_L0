package _test

import (
	"github.com/ory/dockertest/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func createAndStartTestDb() (*dockertest.Pool, *dockertest.Resource) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=postgres",
			"POSTGRES_PASSWORD=postgres",
			"POSTGRES_DB=postgres",
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start PostgreSQL container: %s", err)
	}

	return pool, resource
}

func сonnectionDb() *gorm.DB {
	dsn := "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=postgres"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func сloseConnection(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
}

func teardown(pool *dockertest.Pool, resource *dockertest.Resource) {
	defer func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}()
}
