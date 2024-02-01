package controller

import (
	"context"
	"time"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/ent"
	"github.com/derfinlay/basecrm/service"
)

func CreatePosition(name string, description string, amount int, price float64, ctx context.Context) (*ent.Position, error) {
	return database.Client.Position.Create().
		SetName(name).
		SetDescription(description).
		SetAmount(amount).
		SetUnitPrice(price).
		Save(ctx)
}

func CreateOrder(customer *ent.Customer, billingAddress *ent.BillingAddress, deliveryAddress *ent.DeliveryAddress, products []*ent.Product, taxMultiplier float64, due time.Time, user *ent.User, ctx context.Context) (*ent.Order, error) {
	positions := make([]*ent.Position, 0)

	p := service.RemoveDuplicates(products)
	for _, product := range p {
		amount := CountProductsInList(product, products)
		position, err := CreatePosition(product.Name, "", amount, product.Price, ctx)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	order, err := database.Client.Order.Create().
		SetCustomer(customer).
		SetBillingAddress(billingAddress).
		SetDeliveryAddress(deliveryAddress).
		SetTax(taxMultiplier).
		SetDue(due).
		SetCreatedBy(user).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func CountProductsInList(product *ent.Product, products []*ent.Product) int {
	count := 0
	for _, p := range products {
		if p.ID == product.ID {
			count++
		}
	}
	return count
}
