package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) []byte {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		panic("Error to hash the password")
	}

	return passwordHash
}

func Verify(password []byte, passwordHash []byte) error {
	return bcrypt.CompareHashAndPassword(passwordHash, password)
}
