package common

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Route struct {
	Method string
	Pattern string
	Handler http.Handler
}

func (r *Route) Initialize(router *mux.Router, routes []*Route) {
	for _, route := range routes {
		router.
			Handle(route.Pattern, Logger(route.Handler)).
			Methods(route.Method)
	}
}

func (r *Route) Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}