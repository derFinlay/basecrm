package service

import (
	"fmt"
	"net/smtp"
)

type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type EmailClient struct {
	Authentication smtp.Auth
	Address        string
	From           string
}

var Client *EmailClient

func Create(config *EmailConfig) *EmailClient {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)

	address := config.Host + ":" + fmt.Sprint(config.Port)

	Client = &EmailClient{
		auth,
		address,
		config.From,
	}

	return Client
}

func (config *EmailClient) SendEmail(to []string, msg string) error {
	messageBytes := []byte(msg)

	return smtp.SendMail(config.Address, config.Authentication, config.From, to, messageBytes)
}
