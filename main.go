package main

import (
	"shortener/server"
	"shortener/shortener"

	"github.com/gorilla/mux"
)

func main() {
	server := server.Server{
		Storage: shortener.NewStorage(),
		Router:  mux.NewRouter(),
		Port:    5300,
	}
	server.RegisterRoutes()
	server.Run()
}
