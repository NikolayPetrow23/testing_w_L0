package main

import (
	"log"
	"net/http"
	"testing_w_L0/cmd/web"
	"testing_w_L0/database"
	"testing_w_L0/natsStreaming"
)

func main() {
	database.Migrations()
	//fmt.Println(database.ConnectionDb().Select(database.User{Name: "Николай"}))

	//natsStreaming.ConnectionNats()
	natsStreaming.ConnectNats()
	//fmt.Println(database.FetchDataFromDB())

	//database.SelectUser()
	//database.SelectOrders()
	static := http.FileServer(http.Dir("static"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", static))
	mux.HandleFunc("/", web.Index)
	mux.HandleFunc("/orderID", web.SelectOrderID)
	log.Println("Запуск сервера на http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
