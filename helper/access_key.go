package helper

import (
	"math/rand"
)

type AccessKeyEmailData struct {
	AccessKey string
	FormURL   string
}

func CreateAccessKey(length int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	str := make([]byte, length)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	return string(str)
}
