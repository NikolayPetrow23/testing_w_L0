package main

import (
	"fmt"

	"testing_w_L0/database"
	"testing_w_L0/natsStreaming"
)

func main() {
	database.Migrations()
	fmt.Println(database.ConnectionDb().Select(database.User{Name: "Николай"}))
	natsStreaming.ConnectionNats()
}
