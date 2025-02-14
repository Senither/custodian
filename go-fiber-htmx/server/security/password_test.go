package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyPassword(t *testing.T) {
	t.Run("it verifies the password successfully", func(t *testing.T) {
		password := "mysecretpassword"
		encryptedPassword, _ := EncryptPassword(password)

		isValid := VerifyPassword(encryptedPassword, password)
		assert.True(t, isValid)
	})

	t.Run("it fails to verify an incorrect password", func(t *testing.T) {
		password := "mysecretpassword"
		encryptedPassword, _ := EncryptPassword(password)

		isValid := VerifyPassword(encryptedPassword, "wrongpassword")
		assert.False(t, isValid)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("it encrypts the password successfully", func(t *testing.T) {
		password := "mysecretpassword"
		encryptedPassword, err := EncryptPassword(password)

		assert.Nil(t, err)
		assert.NotEqual(t, password, encryptedPassword)
		assert.True(t, len(encryptedPassword) > 0)
	})

	t.Run("it returns an error pass is too long", func(t *testing.T) {
		_, err := EncryptPassword("12345678901234567890123456789012345678901234567890123456789012345678901234567890")
		assert.NotNil(t, err)
		assert.Equal(t, "bcrypt: password length exceeds 72 bytes", err.Error())
	})
}
