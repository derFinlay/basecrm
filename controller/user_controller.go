package controller

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/ent"
	"github.com/derfinlay/basecrm/ent/user"
	"github.com/derfinlay/basecrm/service"
)

var ErrUserNotFound = errors.New("User not found")
var ErrInvalidLogin = errors.New("Invalid login")

func CreateUser(name string, password string, ctx context.Context) (*ent.User, error) {
	username := GenerateUsernameFromName(name)
	passwordHash, err := service.CreateHash(password)

	if err != nil {
		return nil, err
	}

	return database.Client.User.Create().
		SetName(name).
		SetUsername(username).
		SetPassword(passwordHash).
		Save(ctx)
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

func Login(username string, password string, ctx context.Context) (*ent.User, error) {
	user, err := database.Client.User.Query().Where(user.Username(username)).First(ctx)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidLogin
	}
	if !service.CompareHash(password, user.Password) {
		return nil, ErrInvalidLogin
	}
	user.Update().SetLastLogin(time.Now()).Save(ctx)
	return user, nil
}

func GetUsers(ctx context.Context) ([]*ent.User, error) {
	return database.Client.User.Query().All(ctx)
}

func GetUserById(id int, ctx context.Context) (*ent.User, error) {
	return database.Client.User.Query().Where(user.ID(id)).First(ctx)
}

func SearchUserByName(name string, ctx context.Context) ([]*ent.User, error) {
	return database.Client.User.Query().Where(user.Or(user.UsernameHasPrefix(name), user.NameHasPrefix(name))).All(ctx)
}

func ResetUserPassword(user *ent.User, password string, ctx context.Context) (*ent.User, error) {
	hash, err := service.CreateHash(password)
	if err != nil {
		return nil, err
	}
	return user.Update().SetPassword(hash).Save(ctx)
}
