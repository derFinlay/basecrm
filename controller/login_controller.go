package controller

import (
	"errors"
	"strings"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/models"
	"github.com/derfinlay/basecrm/service"
)

var ErrTokenNotFound = errors.New("invalid reset token")
var ErrTokenAlreadyUsed = errors.New("reset token was already used")

func GetLoginByCustomerId(customerId int) (*models.Login, error) {
	customer, err := GetCustomerByID(customerId)
	if err != nil {
		return nil, err
	}
	return customer.Login, nil
}

func GetLoginByEmail(email string) (*models.Login, error) {
	var login *models.Login
	err := database.Client.Model(&models.Login{}).Where("email = ?", strings.ToLower(email)).First(login).Error
	return login, err
}

func GetLoginResetByToken(token string) (*models.LoginReset, error) {
	var loginReset *models.LoginReset
	err := database.Client.Model(&models.LoginReset{}).Where("token = ?", token).First(loginReset).Error
	return loginReset, err
}

func GetActiveLoginResetByLogin(login *models.Login) (*models.LoginReset, error) {
	reset := &models.LoginReset{}
	err := database.Client.Model(reset).Where("id = ?", login.ID).Preload("LoginResets").First(login).Error
	return reset, err
}

func CreateLoginPasswordReset(login *models.Login) (*models.LoginReset, error) {
	if activeReset, err := GetActiveLoginResetByLogin(login); err == nil {
		return activeReset, nil
	}

	token := service.GenerateRandomString(32)
	reset := &models.LoginReset{
		Active: true,
		Token:  token,
	}
	err := database.Client.Model(login).Association("LoginResets").Append(reset)
	return reset, err
}

func ResetLoginPassword(resetToken string, newPassword string) (*models.Login, error) {
	reset, err := GetLoginResetByToken(resetToken)
	if err != nil {
		return nil, err
	}

	if err := CheckLoginReset(reset); err != nil {
		return nil, err
	}

	login, err := GetLoginById(int(reset.LoginID))
	if err != nil {
		return nil, err
	}

	return nil, login.ResetPassword(newPassword, database.Client)
}

func GetLoginById(id int) (*models.Login, error) {
	var login *models.Login
	err := database.Client.First(login).Where("id = ?", id).Error
	return login, err
}

func CheckLoginReset(reset *models.LoginReset) error {
	if reset == nil {
		return ErrTokenNotFound
	}

	if !reset.Active {
		return ErrTokenAlreadyUsed
	}
	return nil
}
