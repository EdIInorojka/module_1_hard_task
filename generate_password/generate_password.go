package generate_password

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// letters - список допустимых символов в пароле
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GeneratePassword(n int) (string, error) {
	if n <= 0 {
		return "", errors.New("Длина пароля не может быть отрицательной или равна нулю")
	}

	password := make([]byte, n)
	length := big.NewInt(int64(len(letters)))

	for i := range password {
		num, err := rand.Int(rand.Reader, length)
		if err != nil {
			return "", err
		}
		password[i] = letters[num.Int64()]
	}

	return string(password), nil
}
