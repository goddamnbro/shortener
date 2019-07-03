package server

import (
	"fmt"
	"log"
	"net/http"

	"shortener/shortener"

	"github.com/gorilla/mux"
)

type Server struct {
	Storage *shortener.Storage
	Router  *mux.Router
	Port    uint
}

func (s *Server) Run() {
	log.Printf("Server is starting http://localhost:%d \n", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Router))
}
