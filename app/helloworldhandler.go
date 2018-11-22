package app

import (
	"fmt"
	"net/http"
)

func (s *Server) handleHelloWorld() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World")
	}
}