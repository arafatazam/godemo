package app

import (
	"io"
)

type helloer interface{
	PrintHello(w io.Writer)
}