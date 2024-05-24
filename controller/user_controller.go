package controller

import (
	"errors"
	"strings"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/models"
	"github.com/derfinlay/basecrm/service"
)

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidLogin = errors.New("invalid login")

func CreateUser(name string, password string) (*models.User, error) {
	username := GenerateUsernameFromName(name)
	passwordHash, herr := service.CreateHash(password)

	if herr != nil {
		return nil, herr
	}

	newUser := &models.User{
		Name:     name,
		Username: username,
		Password: passwordHash,
	}

	err := newUser.Save(database.Client)

	return newUser, err
}

func GenerateUsernameFromName(name string) string {

	//Input: Finlay Sandy Gress
	//Output: FiSaGr

	prefixes := strings.Split(name, " ")
	for i, prefix := range prefixes {
		prefixes[i] = strings.ToUpper(prefix[0:1]) + strings.ToLower(prefix[1:2])
	}

	username := strings.Join(prefixes, "")

	return username
}

func GetUserSessionByToken(token string) (*models.UserSession, error) {
	var session *models.UserSession
	err := database.Client.Model(session).Where("token = ?", token).Preload("User").Error
	return session, err
}

func Login(username string, password string) (*models.UserSession, error) {
	var user *models.User
	err := database.Client.Find(user).Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidLogin
	}
	if !service.CompareHash(password, user.Password) {
		return nil, ErrInvalidLogin
	}

	session, err := user.CreateUserSession(database.Client)

	return session, err
}

func GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := database.Client.Find(users).Error
	return users, err
}

func GetUserById(id int) (*models.User, error) {
	var user *models.User
	err := database.Client.Model(&models.User{}).Where("id = ?", id).First(user).Error
	return user, err
}

func SearchUserByName(name string) ([]*models.User, error) {
	var users []*models.User
	err := database.Client.Model(&models.User{}).Where("name LIKE ?", name).First(&users).Error
	return users, err
}
