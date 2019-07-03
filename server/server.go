package server

import (
	"log"
	"net/http"
	"shortener/shortener"

	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

	"github.com/gorilla/mux"
)

type Server struct {
	Storage *shortener.Storage
	Router  *mux.Router
	Port    uint
}

func (s *Server) Run() {
	log.Printf("Server is starting http://localhost:%d \n", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Printf(":%d", s.Port), s.Router))
}
