package product

import (
	"fmt"
	"slices"
)

type ProductID string
type Currency string
type Name string
type Availability bool

type Price struct {
	Value    float64
	Currency Currency
}

type Product struct {
	ID            ProductID
	OzonCardPrice Price
	Price         Price
	Name          Name
	Available     Availability
}

const (
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

const (
	AvailableProduct    Availability = true
	NotAvailableProduct Availability = false
)

func GetCurrencies() []Currency {
	return []Currency{USD, USD, EUR}
}

func (currency Currency) IsValid() bool {
	return slices.Contains(GetCurrencies(), currency)
}

func (price Price) IsValid() bool {
	return price.Value > 0 && price.Currency.IsValid()
}

func NewProduct(ID ProductID, ozonCardPrice Price, price Price, name Name, available Availability) (*Product, error) {

	if ozonCardPrice.IsValid() {
		return nil, fmt.Errorf("Invalid ozon card price")
	}

	if !price.IsValid() {
		return nil, fmt.Errorf("Invalid price")
	}

	return &Product{
		ID:            ID,
		OzonCardPrice: ozonCardPrice,
		Price:         price,
		Name:          name,
		Available:     available,
	}, nil
}
