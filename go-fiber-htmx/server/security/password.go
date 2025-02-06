package security

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(password string, value string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(value)) == nil
}

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
