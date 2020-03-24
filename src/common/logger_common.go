package common

import (
    "github.com/subchen/go-log"
    "net/http"
    "time"
)

func Logger(inner http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        inner.ServeHTTP(w, r)

        log.Printf("[%s %s %s]", r.Method, r.RequestURI, time.Since(start))
    })
}

