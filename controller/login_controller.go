package controller

import (
	"context"
	"errors"
	"strings"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/ent"
	"github.com/derfinlay/basecrm/ent/login"
	"github.com/derfinlay/basecrm/ent/loginreset"
	"github.com/derfinlay/basecrm/service"
)

var ErrTokenNotFound = errors.New("invalid reset token")
var ErrTokenAlreadyUsed = errors.New("reset token was already used")

func CreateLogin(c *ent.Customer, email string, ctx context.Context) (*ent.Customer, error) {
	err := service.ValidateEmail(email)
	if err != nil {
		return nil, err
	}
	defaultPassword := service.GenerateRandomString(12)
	login, err := database.Client.Login.Create().SetEmail(email).Save(ctx)

	if err != nil {
		return nil, err
	}

	service.SendDefaultPasswordEmail(login, defaultPassword)

	return c.Update().
		SetLogin(login).
		Save(ctx)
}

func GetLoginByCustomerId(customerId int, ctx context.Context) (*ent.Login, error) {
	customer, err := GetCustomerByID(customerId, ctx)
	if err != nil {
		return nil, err
	}
	return customer.QueryLogin().First(ctx)
}

func GetLoginByEmail(email string, ctx context.Context) (*ent.Login, error) {
	return database.Client.Login.Query().Where(login.Email(strings.ToLower(email))).First(ctx)
}

func GetLoginResetByToken(token string, ctx context.Context) (*ent.LoginReset, error) {
	return database.Client.LoginReset.Query().Where(loginreset.Token(token)).First(ctx)
}

func GetActiveLoginResetByLogin(login *ent.Login, ctx context.Context) (*ent.LoginReset, error) {
	return login.QueryLoginResets().Where(loginreset.Active(true)).First(ctx)
}

func CreateLoginPasswordReset(login *ent.Login, ctx context.Context) (*ent.LoginReset, error) {
	if activeReset, err := GetActiveLoginResetByLogin(login, ctx); err == nil {
		return activeReset, nil
	}

	token := service.GenerateRandomString(32)
	return database.Client.LoginReset.Create().SetActive(true).SetToken(token).SetLogin(login).Save(ctx)
}

func ResetLoginPassword(resetToken string, newPassword string, ctx context.Context) (*ent.Login, error) {
	reset, err := GetLoginResetByToken(resetToken, ctx)
	if err != nil {
		return nil, err
	}

	if err := CheckLoginReset(reset); err != nil {
		return nil, err
	}

	hash, hashErr := service.CreateHash(newPassword)
	if hashErr != nil {
		return nil, hashErr
	}

	login := reset.Edges.Login
	return login.Update().SetPassword(hash).Save(ctx)
}

func CheckLoginReset(reset *ent.LoginReset) error {
	if reset == nil {
		return ErrTokenNotFound
	}

	if !reset.Active {
		return ErrTokenAlreadyUsed
	}
	return nil
}
