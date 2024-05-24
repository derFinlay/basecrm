package service

import (
	"math/rand"
)

func GenerateRandomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.="
	b := ""
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(letters))
		b += string(letters[randomIndex])
	}
	return string(b)
}
