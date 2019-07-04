package server

func (s *Server) RegisterRoutes() {
	s.Router.HandleFunc("/", s.HandleIndex).Methods("GET")
	s.Router.HandleFunc("/resolve/{hash}", s.HandleResolve).Methods("GET")
	s.Router.HandleFunc("/shorten", s.HandleShorten).Methods("POST")
}
