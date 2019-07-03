package server

import (
	"fmt"
	"net/http"
)

func (s *Server) HandleIndex() http.HandlerFunc {
	// you can do something here
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, girls %s!", r.URL.Path[1:])
	}
}
