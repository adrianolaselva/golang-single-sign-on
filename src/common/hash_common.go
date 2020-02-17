package common

import "golang.org/x/crypto/bcrypt"

type Hash struct {

}

func (h *Hash) BCryptGenerate(password string) string{
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (h *Hash) BCryptCompare(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}