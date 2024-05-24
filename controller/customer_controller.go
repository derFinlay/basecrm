package controller

import (
	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/models"
)

func GetCustomerByID(id int) (*models.Customer, error) {
	var customer *models.Customer
	err := database.Client.Model(&models.Customer{}).First(customer, id).Preload("Orders").Preload("BillingAddresses").Preload("DevliveryAddresses").Preload("Login").Preload("CreatedBy").Error
	return customer, err
}

func CreateCustomer(name string, gender string, user *models.User) (*models.Customer, error) {
	customer := &models.Customer{
		Name:      name,
		Gender:    gender,
		CreatedBy: user,
	}
	err := customer.Save(database.Client)
	return customer, err
}

func SearchCustomer(query string) ([]*models.Customer, error) {
	var customers []*models.Customer

	err := database.Client.Model(&models.Customer{}).Where("name = ? OR id = ?", query, query).Find(customers).Error

	return customers, err
}

func CreateDeliveryAddress(street string, city string, zipcode string, housenumber string, customer *models.Customer, user *models.User) (*models.DeliveryAddress, error) {
	deliveryAddress := &models.DeliveryAddress{
		City:      city,
		Street:    street,
		ZIP:       zipcode,
		Number:    housenumber,
		CreatedBy: user,
	}

	err := database.Client.Model(customer).Association("DeliveryAddresses").Append(deliveryAddress)
	return deliveryAddress, err
}

func CreateBillingAddres(street string, city string, zipcode string, housenumber string, customer *models.Customer, user *models.User) (*models.BillingAddress, error) {
	billingAddress := &models.BillingAddress{
		City:      city,
		Street:    street,
		ZIP:       zipcode,
		Number:    housenumber,
		CreatedBy: user,
	}

	err := database.Client.Model(customer).Association("BillingAddresses").Append(billingAddress)
	return billingAddress, err
}
