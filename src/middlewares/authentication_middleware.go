package middlewares

import (
	"log"
	"net/http"
)

type AuthenticationMiddleware interface {
	ValidateClientIdAndSecret(next http.Handler) http.Handler
	ValidateUserAndPassword(next http.Handler) http.Handler
	ValidateJWTToken(next http.Handler) http.Handler
}

type authenticationMiddleware struct {

}

func NewAuthenticationMiddleware() *authenticationMiddleware {
	return &authenticationMiddleware{}
}

func (b *authenticationMiddleware) ValidateClientIdAndSecret(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("header", r.Header)
		next.ServeHTTP(w, r)
	})
}

func (b *authenticationMiddleware) ValidateJWTToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("header", r.Header)
		next.ServeHTTP(w, r)
	})
}

func (b *authenticationMiddleware) ValidateUserAndPassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("header", r.Header)
		next.ServeHTTP(w, r)
	})
}