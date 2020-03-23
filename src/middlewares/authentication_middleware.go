package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"oauth2/src/dto"
	"oauth2/src/service"
	"oauth2/src/service/oauth"
	"strings"
)

type AuthenticationMiddleware interface {
	//Filter(next http.Handler) http.Handler
	//ValidateClientIdAndSecret(next http.Handler) http.Handler
	//ValidateUserAndPassword(next http.Handler) http.Handler
	ValidateToken(next http.Handler, hasScopes []string) http.Handler
}

type authenticationMiddleware struct {
	userService service.UserService
	authFlow oauth.AuthFlow
}

func NewAuthenticationMiddleware(userService service.UserService, authFlow oauth.AuthFlow) *authenticationMiddleware {
	return &authenticationMiddleware{userService, authFlow}
}

func (b *authenticationMiddleware) ValidateToken(next http.Handler, hasScopes []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) < 2 || auth[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
				Message: "invalid token",
			})
			return
		}

		_, claims, err := b.authFlow.ValidateAccessToken(auth[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(&dto.DefaultResponseDto{
				Message: err.Error(),
			})
			return
		}

		//if hasScopes != nil {
		//	for _, hasScope := range hasScopes {
		//		for _, scope := range claims.Scope {
		//			if hasScope == scope {
		//
		//			}
		//		}
		//	}
		//}

		log.Println(claims.ClientID)
		log.Println(claims.Scope)
		log.Println(claims.ExpiresAt)
		log.Println(claims.Profile.ID)
		log.Println(claims.Profile.Username)
		log.Println("hasScopes", hasScopes)

		r.Header.Add("user_id", claims.Profile.ID)
		r.Header.Add("client_id", claims.ClientID)

		log.Println("header", r.Header)
		next.ServeHTTP(w, r)
	})
}


func (b *authenticationMiddleware) Filter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		//
		//if len(auth) != 2 || auth[0] != "Basic" {
		//	http.Error(w, "authorization failed", http.StatusUnauthorized)
		//	return
		//}
		//
		//payload, _ := base64.StdEncoding.DecodeString(auth[1])
		//pair := strings.SplitN(string(payload), ":", 2)
		//
		//if user, err := b.userService.LoginByUserAndPassword(pair[0], pair[1]); err != nil {
		//	log.Println(user)
		//	next.ServeHTTP(w, r)
		//	return
		//}

		next.ServeHTTP(w, r)
	})
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