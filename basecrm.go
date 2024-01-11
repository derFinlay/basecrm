package basecrm

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/derfinlay/basecrm/ent"
)

func Test() bool {
	fmt.Println("baseCRM >> Test successful")
	return true
}

func Init() {
	log.SetPrefix("baseCRM >> ")
	return
}

func CreateClient(host string, port int, user string, password string) (*ent.Client, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, "basecrm", password)
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client, err
}
