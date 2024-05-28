package controller

import (
	"strconv"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/models"
)

func GetCustomerByID(id int) (*models.Customer, error) {
	var customer *models.Customer
	err := database.Client.Model(&models.Customer{}).Preload("BillingAddress").First(&customer, id).Preload("Orders").Preload("DevliveryAddresses").Preload("Login").Preload("CreatedBy").Error
	return customer, err
}

func CreateCustomer(name string, gender string, billingAddress *models.BillingAddress, deliveryAddresses []*models.DeliveryAddress, user *models.User) (*models.Customer, error) {
	billingAddress.CreatedByID = user.ID
	customer := &models.Customer{
		Name:              name,
		Gender:            gender,
		CreatedBy:         user,
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

	err := database.Client.Model(&models.Customer{}).Where("name = ? OR id = ?", query, idd).Find(&customers).Error

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
