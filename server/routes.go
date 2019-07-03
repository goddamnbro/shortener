package server

func (s *Server) RegisterRoutes() {
	s.Router.HandleFunc("/", s.HandleIndex()).Methods("GET")
}
