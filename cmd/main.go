package main

import (
	"log"
	"net/http"
	"testing_w_L0/database"
	"testing_w_L0/my_cache"
	"testing_w_L0/natsStreaming"
	"testing_w_L0/router"
	"time"
)

func main() {
	database.Migrations()

	my_cache.LoadCacheFromDB(database.SelectAllOrders())

	go natsStreaming.SubscribeNuts()

	go func() {
		time.Sleep(3 * time.Second)
		natsStreaming.OrderMessage()
	}()

	static := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", static))
	mux.HandleFunc("/", router.IndexRouter)
	mux.HandleFunc("/orderuid", router.OrderRouter)
	log.Println("Запуск сервера на http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
