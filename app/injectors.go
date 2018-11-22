//+build wireinject

package app

import (
	"github.com/google/go-cloud/wire"
)

//NewServer injects Server
func NewServer() *Server{
	wire.Build(AppProvider)
	return &Server{}
}