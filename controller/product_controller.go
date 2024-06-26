package controller

import (
	"strconv"
	"strings"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/models"
)

func GetProductByID(id int) (*models.Product, error) {
	var product *models.Product
	err := database.Client.Model(&models.Product{}).
		Preload("Notes").
		First(&product, id).Error
	return product, err
}

func SearchProducts(query string) ([]*models.Product, error) {
	var products []*models.Product
	var idd int

	if id, err := strconv.Atoi(query); err != nil {
		id = 0
	} else {
		idd = id
	}

	err := database.Client.Model(&models.Product{}).Where("LOWER(name) LIKE ? OR id = ? OR LOWER(description) LIKE ?", "%"+strings.ToLower(query)+"%", idd, "%"+strings.ToLower(query)+"%").
		Preload("Notes").
		Find(&products).Error

	return products, err
}

func CreateProduct(name string, description string, price float32, notes []*models.Note) (*models.Product, error) {
	product := &models.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
	err := product.Save(database.Client)
	return product, err
}
