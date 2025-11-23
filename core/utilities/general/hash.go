package general

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func MakeHash(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("❌ Error To Make Hash %v", err)
	}
	return string(hashed)
}

func CheckHash(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
