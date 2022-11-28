package helper

import "math/rand"

func CreatePin(length int) string {
	letters := []byte("0123456789")

	otp := make([]byte, length)
	for i := range otp {
		otp[i] = letters[rand.Intn(len(letters))]
	}

	return string(otp)
}
