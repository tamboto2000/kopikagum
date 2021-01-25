package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc(`/api/{path:[a-zA-Z0-9=\-\_\/]+}`, gateway.API)

	srv := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
