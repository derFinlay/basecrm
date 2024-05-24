package controller

import "github.com/derfinlay/basecrm/models"

func CountProductsInList(product *models.Product, products []*models.Product) int {
	count := 0
	for _, p := range products {
		if p.ID == product.ID {
			count++
		}
	}
	return count
}

func RemoveDuplicates[T string | int | *models.Product](input []T) []T {
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
