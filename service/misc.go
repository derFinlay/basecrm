package service

import (
	"math/rand"

	"github.com/derfinlay/basecrm/ent"
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

func RemoveDuplicates[T string | int | *ent.Product](input []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range input {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
