package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQzMDA2MTgsIlVzZXIiOnsiaWQiOiI4ZDQyZWUzZS01NzE3LTRiNjUtYjBiNi0yMTgzNjFmOTgxYjMiLCJuYW1lIjoiQWRyaWFubyIsImxhc3RfbmFtZSI6Ik1vcmVpcmEgTGEgU2VsdmEiLCJlbWFpbCI6ImFkcmlhbm9sYXNlbHZhQGdtYWlsLmNvbSIsInVzZXJuYW1lIjoiYWRyaWFub2xhc2VsdmEiLCJhY3RpdmF0ZWQiOnRydWUsImJpcnRoZGF5IjoiMTk4Ny0wMi0xMSIsImNyZWF0ZWRfYXQiOiIyMDIwLTAzLTE0IDE2OjA1OjA0IiwidXBkYXRlZF9hdCI6IjIwMjAtMDMtMTQgMTY6NDk6MjAiLCJleHBpcmVzX2F0IjoiIiwiZGVsZXRlZF9hdCI6IiJ9LCJEYXRhIjp7fX0.PcD-CT9HVKvB73DmzQPvqE_KEWNAX1jPtTzR0JYIZLo"
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABA"), nil
	})

	fmt.Printf("token: %v\n\n", token)
	fmt.Printf("token raw: %s\n\n", token.Raw)

	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return
	}

	if err = claims.Valid(); err != nil {
		fmt.Printf("error: %v\n", err.Error())
		return
	}

	user := claims["User"].(map[string]interface{})
	fmt.Printf("user: %v\n\n\n", user)

	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
}
