package controller

import (
	"strconv"
	"strings"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/models"
)

func GetCustomerByID(id int) (*models.Customer, error) {
	var customer *models.Customer
	err := database.Client.Model(&models.Customer{}).
		Preload("Notes").
		Preload("BillingAddress").
		Preload("Orders").
		Preload("DeliveryAddresses").
		Preload("DeliveryAddresses.CreatedBy").
		Preload("DeliveryAddresses.Notes").
		Preload("DeliveryAddresses.Tels").
		Preload("Login").
		Preload("CreatedBy").
		First(&customer, id).Error
	return customer, err
}

func CreateCustomer(name string, gender string, billingAddress *models.BillingAddress, deliveryAddresses []*models.DeliveryAddress, tels []*models.Tel, notes []*models.Note, user *models.User) (*models.Customer, error) {
	billingAddress.CreatedByID = user.ID
	customer := &models.Customer{
		Name:              name,
		Gender:            gender,
		CreatedBy:         user,
		Tels:              tels,
		Notes:             notes,
		DeliveryAddresses: deliveryAddresses,
		BillingAddress:    billingAddress,
	}
	err := customer.Save(database.Client)
	return customer, err
}

func SearchCustomer(query string) ([]*models.Customer, error) {
	var customers []*models.Customer
	var idd int

	if id, err := strconv.Atoi(query); err != nil {
		id = 0
	} else {
		idd = id
	}

	err := database.Client.Model(&models.Customer{}).Where("LOWER(name) LIKE ? OR id = ?", "%"+strings.ToLower(query)+"%", idd).
		Preload("BillingAddress").
		Preload("Notes").
		Preload("CreatedBy").
		Preload("Tels").
		Find(&customers).Error

	return customers, err
}

func CreateDeliveryAddress(street string, city string, zipcode string, housenumber string, customer *models.Customer, user *models.User) (*models.DeliveryAddress, error) {
	deliveryAddress := &models.DeliveryAddress{
		City:       city,
		Street:     street,
		ZIP:        zipcode,
		Number:     housenumber,
		CustomerID: customer.ID,
		CreatedBy:  user,
	}
	err := deliveryAddress.Save(database.Client)
	return deliveryAddress, err
}
