package app

import(
	"github.com/go-chi/chi"
	"net/http"
)

//Server holds dependencies
type Server struct{
	Mux *chi.Mux
	helloer helloer
}

func (s *Server)ServeHTTP(w http.ResponseWriter, r *http.Request){
	s.Mux.ServeHTTP(w, r)
}