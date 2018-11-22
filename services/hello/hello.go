package hello

import (
	"io"
)

//Hello implements helloer
type Hello struct{
	message string
}

//PrintHello prints hello world
func (h Hello)PrintHello(w io.Writer){
	w.Write([]byte(h.message))
}