package middleware

import (
	"net/http"

	"github.com/sqoopdata/madoc/pkg/application"
)

type Middleware func(http.HandlerFunc, *application.Application) http.HandlerFunc

func Chain(h http.HandlerFunc, a *application.Application, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}

	wrapped := h

	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped, a)
	}

	return wrapped
}
