package helper

import "golang.org/x/crypto/bcrypt"

// HashString encrypt given text
func HashString(text string) (string, error) {
	bt, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}