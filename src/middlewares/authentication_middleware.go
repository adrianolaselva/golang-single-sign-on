package middlewares

import (
	"encoding/base64"
	"log"
	"net/http"
	"oauth2/src/service"
	"strings"
)

type AuthenticationMiddleware interface {
	ValidateClientIdAndSecret(next http.Handler) http.Handler
	ValidateUserAndPassword(next http.Handler) http.Handler
	ValidateJWTToken(next http.Handler, hasScopes []string) http.Handler
}

type authenticationMiddleware struct {
	userService service.UserService
}

func NewAuthenticationMiddleware(userService service.UserService) *authenticationMiddleware {
	return &authenticationMiddleware{userService}
}

func (b *authenticationMiddleware) ValidateClientIdAndSecret(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if user, err := b.userService.LoginByUserAndPassword(pair[0], pair[1]); err != nil {
			log.Println(user)
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (b *authenticationMiddleware) ValidateJWTToken(next http.Handler, hasScopes []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("hasScopes", hasScopes)
		log.Println("header", r.Header)
		next.ServeHTTP(w, r)
	})
}

func (b *authenticationMiddleware) ValidateUserAndPassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			http.Error(w, "authorization failed", http.StatusUnauthorized)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if user, err := b.userService.LoginByUserAndPassword(pair[0], pair[1]); err != nil {
			log.Println(user)
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}