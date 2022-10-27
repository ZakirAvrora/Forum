package services

import (
	"Forum/utilities/e"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePasswordHash(password []byte) ([]byte, error) {
	hashPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, e.Wrap("can't generate hash passwd", err)
	}

	return hashPassword, nil
}

func ComparePassword(hashedPassword []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		err = e.Wrap("hashed password is not same as password", err)
	}
	return err
}
