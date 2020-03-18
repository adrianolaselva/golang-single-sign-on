package common

import (
	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	BCryptGenerate(password string) (string, error)
	BCryptCompare(hashedPassword string, password string) (bool, error)
}

type hashImpl struct {

}

func NewHash() *hashImpl {
	return &hashImpl{}
}

func (h *hashImpl) BCryptGenerate(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (h *hashImpl) BCryptCompare(hashedPassword string,  password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, err
	}

	return true, err
}