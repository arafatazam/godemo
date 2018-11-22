package app

import (
	"net/http"
)

func (s *Server) handleHelloWorld() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		s.helloer.PrintHello(w)
	}
}