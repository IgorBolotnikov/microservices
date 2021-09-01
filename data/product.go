package data

import "time"

// Product defined a structure for an AI product
type Product struct {
	ID int
	Name string
	Description string
	Price float32
	SKU string
	CreatedOn string
	UpdatedOn string
	DeletedOn string
}

var productList = []*Product{
	&Product{
		ID: 1,
		Name: "Espresso",
		Description: "Short and strong coffee",
		Price: 1.99,
		SKU: "aa00",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
	&Product{
		ID: 1,
		Name: "Cappuccino",
		Description: "Larger coffee with frothed milk",
		Price: 2.49,
		SKU: "aa01",
		CreatedOn: time.Now().UTC().String(),
		UpdatedOn: time.Now().UTC().String(),
	},
}