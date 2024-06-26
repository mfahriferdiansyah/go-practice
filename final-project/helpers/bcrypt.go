package helpers

import "golang.org/x/crypto/bcrypt"

func Hash(p string) string {
	cost := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, cost)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
