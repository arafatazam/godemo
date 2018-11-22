package app

import (
	"github.com/google/go-cloud/wire"
	"github.com/go-chi/chi"
)

//ProvideChi provides go chi router
func ProvideChi() *chi.Mux{
	return chi.NewRouter()
}

//AppProvider provides application dependencies
var AppProvider = wire.NewSet(
	ProvideChi,
	Server{},
)