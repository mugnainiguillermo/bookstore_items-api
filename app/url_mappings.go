package app

import (
	"github.com/mugnainiguillermo/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingControllers.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsControllers.Create).Methods(http.MethodPost)
}
