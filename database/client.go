package database

import (
	"context"
	"fmt"
	"log"

	"github.com/derfinlay/basecrm/ent"
	"github.com/derfinlay/basecrm/service"
	_ "github.com/lib/pq"
)

var Client *ent.Client

func Create(host string, port int, user string, password string, ctx context.Context) (*ent.Client, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, "basecrm", password)
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	log.Printf("Connection to postgres successful")

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}
	log.Printf("Schema migration successful")
	Client = client

	if count, err := GetUserCount(ctx); count > 0 {
		return client, nil
	} else if err != nil {
		return nil, err
	}

	if _, err := CreateDefaultUser(ctx); err != nil {
		return nil, err
	}

	return client, nil
}

func Close() error {
	return Client.Close()
}

func GetUserCount(ctx context.Context) (int, error) {
	return Client.User.Query().Count(ctx)
}

func CreateDefaultUser(ctx context.Context) (*ent.User, error) {
	hash, err := service.CreateHash("admin")
	if err != nil {
		return nil, err
	}
	adminUser, err := Client.User.Create().SetName("Administrator").SetUsername("admin").SetPassword(hash).Save(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("Created admin user")
	return adminUser, nil
}
