package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:9002",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
