package app

func (s *Server) Routes(){
	s.Mux.Handle("/", s.handleHelloWorld())
}