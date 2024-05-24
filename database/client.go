package database

import (
	"fmt"
	"log"

	"github.com/derfinlay/basecrm/models"
	"github.com/derfinlay/basecrm/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Create(host string, port int, user string, password string, dbname string) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, dbname, password)
	client, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Printf("Connection to postgres successful")
	Client = client

	// Run the auto migration tool.
	autoMigrate()
	log.Printf("Schema migration successful")

	if count, err := getUserCount(); count > 0 {
		return client, nil
	} else if err != nil {
		return nil, err
	}

	if _, err := createDefaultUser(); err != nil {
		return nil, err
	}

	return client, nil
}

func autoMigrate() {
	Client.AutoMigrate(&models.BillingAddress{})
	Client.AutoMigrate(&models.Customer{})
	Client.AutoMigrate(&models.DeliveryAddress{})
	Client.AutoMigrate(&models.Invoice{})
	Client.AutoMigrate(&models.LoginReset{})
	Client.AutoMigrate(&models.Login{})
	Client.AutoMigrate(&models.Note{})
	Client.AutoMigrate(&models.Order{})
	Client.AutoMigrate(&models.Position{})
	Client.AutoMigrate(&models.Product{})
	Client.AutoMigrate(&models.Role{})
	Client.AutoMigrate(&models.Tel{})
	Client.AutoMigrate(&models.UserSession{})
	Client.AutoMigrate(&models.User{})
}

func getUserCount() (int64, error) {
	var count int64
	err := Client.Model(&models.User{}).Count(&count).Error
	return count, err
}

func createDefaultUser() (*models.User, error) {
	hash, err := service.CreateHash("admin")
	if err != nil {
		return nil, err
	}

	adminUser := &models.User{
		Name: "Administrator", Username: "admin", Password: hash,
	}

	cerr := Client.Create(adminUser).Error
	if err != nil {
		return nil, cerr
	}
	log.Printf("Created admin user")
	return adminUser, nil
}
