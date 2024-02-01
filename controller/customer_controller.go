package controller

import (
	"context"

	"github.com/derfinlay/basecrm/database"
	"github.com/derfinlay/basecrm/ent"
	"github.com/derfinlay/basecrm/ent/billingaddress"
	"github.com/derfinlay/basecrm/ent/customer"
	"github.com/derfinlay/basecrm/ent/deliveryaddress"
	"github.com/derfinlay/basecrm/ent/note"
)

func GetCustomerByID(id int, ctx context.Context) (*ent.Customer, error) {
	return database.Client.Customer.Query().Where(customer.ID(id)).First(ctx)
}

func CreateCustomer(name string, gender customer.Gender, user *ent.User, ctx context.Context) (*ent.Customer, error) {
	return database.Client.Customer.Create().
		SetName(name).
		SetGender(gender).
		SetCreatedBy(user).
		Save(ctx)
}

func SearchCustomer(search string, searchIn []string, ctx context.Context) ([]*ent.Customer, error) {
	customers := []*ent.Customer{}
	for i := 0; i < len(searchIn); i++ {
		toSearch := searchIn[i]
		switch toSearch {
		case "name":
			byName, _ := SearchCustomerByName(search, ctx)
			customers = append(customers, byName...)
		case "notes":
			byNotes, _ := SearchCustomerByNotes(search, ctx)
			customers = append(customers, byNotes...)
		case "billingaddress":
			byBillingAddress, _ := SearchCustomerByBillingAddress(search, ctx)
			customers = append(customers, byBillingAddress...)
		case "deliveryaddress":
			byDeliveryAddress, _ := SearchCustomerByDeliveryAddress(search, ctx)
			customers = append(customers, byDeliveryAddress...)

		}
	}
	return customers, nil
}

func SearchCustomerByName(name string, ctx context.Context) ([]*ent.Customer, error) {
	return database.Client.Customer.Query().
		Where(
			customer.NameContains(name)).
		All(ctx)
}

func SearchCustomerByNotes(content string, ctx context.Context) ([]*ent.Customer, error) {
	return database.Client.Customer.Query().
		Where(
			customer.HasNotesWith(
				note.ContentContains(content))).
		All(ctx)
}

func SearchCustomerByBillingAddress(address string, ctx context.Context) ([]*ent.Customer, error) {
	return database.Client.Customer.Query().
		Where(
			customer.HasBillingAddressesWith(
				billingaddress.Or(
					billingaddress.CityContains(address),
					billingaddress.StreetContains(address),
					billingaddress.ZipContains(address),
				))).
		All(ctx)

}

func SearchCustomerByDeliveryAddress(address string, ctx context.Context) ([]*ent.Customer, error) {
	return database.Client.Customer.Query().
		Where(
			customer.HasDeliveryAddressesWith(
				deliveryaddress.Or(
					deliveryaddress.CityContains(address),
					deliveryaddress.StreetContains(address),
					deliveryaddress.ZipContains(address),
				))).
		All(ctx)
}

func CreateDeliveryAddress(street string, city string, zipcode string, housenumber string, customer *ent.Customer, ctx context.Context) (*ent.DeliveryAddress, error) {
	return database.Client.DeliveryAddress.
		Create().
		SetCustomer(customer).
		SetCity(city).
		SetHousenumber(housenumber).
		SetStreet(street).
		SetZip(zipcode).
		Save(ctx)

}

func CreateBillingAddres(street string, city string, zipcode string, housenumber string, customer *ent.Customer, ctx context.Context) (*ent.BillingAddress, error) {
	return database.Client.BillingAddress.
		Create().
		SetCustomer(customer).
		SetCity(city).
		SetHousenumber(housenumber).
		SetStreet(street).
		SetZip(zipcode).
		Save(ctx)
}
