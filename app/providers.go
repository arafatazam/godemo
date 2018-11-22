package app

import (
	"github.com/arafatazam/godemo/services/hello"
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
	hello.ProvideHello,
	wire.Bind(new(helloer), new(hello.Hello)),
	Server{},
)