package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	var urls strings.Builder

	for hash, url := range *s.Storage {
		urls.WriteString(fmt.Sprintf("http://localhost:%d/resolve/%s -> %s \n", s.Port, hash, url))
	}

	fmt.Fprintf(w, urls.String())
}

func (s *Server) HandleResolve(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]
	url, ok := s.Storage.Resolve(hash)
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func (s *Server) HandleShorten(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	hash, err := s.Storage.Shorten(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	shortURL := fmt.Sprintf("http://localhost:%d/resolve/%s", s.Port, hash)

	fmt.Fprintf(w, "Your short link: %s", shortURL)
}
